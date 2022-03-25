package app

import (
	"github.com/swanden/rentateam/pkg/config"
	"log"
)

func Run(configFile string) {
	_, err := config.New(configFile)
	if err != nil {
		log.Fatalf("Config error: %s\n", err)
	}
}
