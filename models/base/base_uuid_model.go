package basemodels

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BaseUuidModelSoftDelete struct {
	Uuid      uuid.UUID `gorm:"type:uuid;primary_key;not null;default:null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (base *BaseUuidModelSoftDelete) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Uuid == uuid.Nil {
		base.Uuid = uuid.NewV4()
	}
	return
}

type BaseUuidModelHardDelete struct {
	Uuid      uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (base *BaseUuidModelHardDelete) BeforeCreate(tx *gorm.DB) (err error) {
	if base.Uuid == uuid.Nil {
		base.Uuid = uuid.NewV4()
	}
	return
}
