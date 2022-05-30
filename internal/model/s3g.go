package model

import (
	"gorm.io/gorm"
)

type (
	S3GatewayConfig struct {
		gorm.Model
		Host string
		Port int32
	}

	S3GatewayRegistration struct {
		gorm.Model
		Host string
		Port int32
	}
)
