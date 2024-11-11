package server

import (
	"log"
	"sandbox-service/configs"
)

func ExternalDeps(cfg configs.Config) {
	log.Println(cfg)
}
