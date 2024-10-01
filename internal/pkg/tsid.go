package pkg

type IDGenerator interface {
	GenerateTSID() (uint64, error)
	ToString(uint64) (string, error)
	ToID(string) (uint64, error)
}

type TSIDGenerator struct {
}

func NewTSIDGenerator() *TSIDGenerator {
	return &TSIDGenerator{}
}

func (g *TSIDGenerator) GenerateTSID() (uint64, error) {

}

func (g *TSIDGenerator) ToString(id uint64) (string, error) {

}

func (g *TSIDGenerator) ToID(strID string) (uint64, error) {

}
