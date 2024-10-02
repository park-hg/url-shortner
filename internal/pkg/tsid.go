package pkg

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/redis/go-redis/v9"
)

type IDGenerator interface {
	GenerateTSID() (int64, error)
	ToString(int64) (string, error)
	ToID(string) (int64, error)
}

const (
	maxServerID          = 1024
	baseUnixMilliSeconds = 1727868113 // 2024-10-02 11:21
)

var (
	ErrGenerateTSID = errors.New("fail to generate tsid")
)

// tsid: 64bit
// serverID: 10bit
// timestamp: 40bit (34.8 yrs)
// serialNumber: 14bit (can handler almost 10,000 req/s)

type TSIDGenerator struct {
	rdb    redis.UniversalClient
	ticker *time.Ticker
	enc    Encoder

	serverID int
	serialNo int32
}

func NewTSIDGenerator(rdb redis.UniversalClient, enc Encoder) *TSIDGenerator {
	ctx := context.Background()

	for i := 0; i < maxServerID; i++ {
		res, err := rdb.SetNX(ctx, fmt.Sprintf("server_id:%d", i), 1, redis.KeepTTL).Result()
		if err != nil {
			panic(err)
		}
		if res == true {
			ticker := time.NewTicker(1 * time.Second)
			var serialNo int32
			go func() {
				for range ticker.C {
					atomic.StoreInt32(&serialNo, 0)
				}
			}()

			return &TSIDGenerator{rdb: rdb, ticker: ticker, enc: enc, serverID: i, serialNo: serialNo}
		}
	}

	panic("no spare server id")
}

// GenerateTSID creates integer tsID as following components
// tsID structure:
//
//		---------------------------------------------
//		|  timestamp  |  server_id  | serial_number |
//		---------------------------------------------
//		| <- 40bit -> | <- 10bit -> |  <- 14bit ->  |
//		---------------------------------------------
//		|  ~34.8 yrs  |  1024 srvs  |  ~10,000req/s |
//	    ---------------------------------------------
func (g *TSIDGenerator) GenerateTSID() (int64, error) {
	relativeTimestamp := time.Now().UnixMilli() - baseUnixMilliSeconds

	for i := 0; i < 100; i++ {
		oldValue := atomic.LoadInt32(&g.serialNo)
		newValue := oldValue + 1

		if atomic.CompareAndSwapInt32(&g.serialNo, oldValue, newValue) {
			return relativeTimestamp<<24 | int64(g.serverID)<<14 | int64(g.serialNo), nil
		}
	}

	return 0, ErrGenerateTSID
}

func (g *TSIDGenerator) ToString(id int64) (string, error) {
	return g.enc.Encode(id), nil
}

func (g *TSIDGenerator) ToID(strID string) (int64, error) {
	return g.enc.Decode(strID)
}

func (g *TSIDGenerator) Close() error {
	ctx := context.Background()
	g.ticker.Stop()
	return g.rdb.Del(ctx, fmt.Sprintf("server_id:%d", g.serverID)).Err()
}
