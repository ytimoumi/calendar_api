package server

import (
	"log"
	"os"
)

func Init() {
	r := NewRouter()
	if err := r.Run(":8080"); err != nil {
		log.Println("Listening failed")
		os.Exit(1)
	}
}
