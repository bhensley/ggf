package model

import "gorm.io/gorm"

type RedirectUTM struct {
	gorm.Model

	UTMSource   string `gorm:"not null"`
	UTMMedium   string `gorm:"not null"`
	UTMCampaign string `gorm:"not null"`
	UTMContent  string `gorm:"not null"`
	UTMTerm     string `gorm:"not null"`

	ShortURL string   `gorm:"not null"`
	Redirect Redirect `gorm:"foreignKey:ShortURL;references:ShortURL"`
}
