package basemodels

import (
	"time"

	"gorm.io/gorm"
)

type BaseIdModelSoftDelete struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type BaseIdModelHardDelete struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
