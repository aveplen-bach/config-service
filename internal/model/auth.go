package model

import (
	"gorm.io/gorm"
)

type (
	AuthenticationDatabaseConfig struct {
		gorm.Model
		Host     string
		Port     int32
		Username string
		Password string
		Database string
	}

	AuthenticationConfig struct {
		DatabaseConfig        AuthenticationDatabaseConfig
		FacerecRegistration   FaceRecognitionRegistration
		S3GatewayRegistration S3GatewayRegistration
	}

	AuthenticationRegistration struct {
		gorm.Model
		Host string `bson:"host"`
		Port int32  `bson:"port"`
	}
)
