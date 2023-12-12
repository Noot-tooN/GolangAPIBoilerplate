package models

import basemodels "golangapi/models/base"

type UserInfo struct {
	basemodels.BaseUuidModelSoftDelete
	Email         string `gorm:"unique;index;not null;default:null"`
	EmailVerified bool   `gorm:"index;not null;default:false"`
	Password 	  string `gorm:"not null; default: null"`
}
