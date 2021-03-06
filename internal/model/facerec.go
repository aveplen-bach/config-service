package model

import (
	"gorm.io/gorm"
)

type (
	FaceRecognitionConfig struct {
		gorm.Model
		Threshold float64
	}

	FaceRecognitionRegistration struct {
		gorm.Model
		Host string
		Port int32
	}
)
