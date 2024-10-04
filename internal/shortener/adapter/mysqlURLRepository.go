package adapter

import (
	"context"
	"time"

	"gorm.io/gorm"

	"traffic-reporter/internal/pkg"
)

type URLMappingTable struct {
	ID          int64     `gorm:"type:bigint unsigned;not null;primaryKey;autoIncrement:false"`
	OriginalURL string    `gorm:"type:varchar(255);not null;uniqueIndex:UNIQ_originalUrl"`
	CreatedAt   time.Time `gorm:"type:timestamp;not null;default:CURRENT_TIMESTAMP"`
}

type MySQLURLRepository struct {
	db          *gorm.DB
	idGenerator pkg.IDGenerator
}

func NewMySQLURLRepository(db *gorm.DB, idGenerator pkg.IDGenerator) *MySQLURLRepository {
	return &MySQLURLRepository{db: db, idGenerator: idGenerator}
}

func (m *MySQLURLRepository) Shorten(ctx context.Context, original string) (string, error) {
	id, err := m.idGenerator.GenerateTSID()
	if err != nil {
		return "", err
	}
	strID, err := m.idGenerator.ToString(id)
	if err != nil {
		return "", err
	}

	target := URLMappingTable{
		ID:          id,
		OriginalURL: original,
		CreatedAt:   time.Now(),
	}

	err = m.db.WithContext(ctx).Create(&target).Error
	return strID, err
}

func (m *MySQLURLRepository) GetShortened(ctx context.Context, original string) (string, error) {
	dest := URLMappingTable{}
	err := m.db.WithContext(ctx).Where("original_url = ?", original).First(&dest).Error
	if err != nil {
		return "", err
	}
	return m.idGenerator.ToString(dest.ID)
}

func (m *MySQLURLRepository) RetrieveOriginal(ctx context.Context, shortened string) (string, error) {
	id, err := m.idGenerator.ToID(shortened)
	if err != nil {
		return "", err
	}

	var dest URLMappingTable
	err = m.db.WithContext(ctx).Where("id = ?", id).First(&dest).Error
	return dest.OriginalURL, err
}
