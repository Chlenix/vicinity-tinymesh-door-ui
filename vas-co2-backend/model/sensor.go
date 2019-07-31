package model

import (
	_ "database/sql"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Sensor struct {
	Oid      uuid.UUID `gorm:"primary_key"`
	Eid      string    `gorm:"type:varchar(100);unique_index"`
	Readings []Reading `gorm:"foreignkey:SensorOid"`
}

type Reading struct {
	gorm.Model
	Value     int8 `gorm:"type:tinyint(1) unsigned;not null"`
	Time      time.Time
	SensorOid uuid.UUID `gorm:"index"`
}
