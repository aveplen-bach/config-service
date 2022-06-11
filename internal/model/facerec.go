package model

import (
	"gorm.io/gorm"
)

type FacerecConfig struct {
	gorm.Model
	Threshold float64
}
