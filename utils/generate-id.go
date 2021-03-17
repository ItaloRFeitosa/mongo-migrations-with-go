package utils

import (guuid "github.com/google/uuid" )

func GenerateId() string {
	return guuid.New().String()
}