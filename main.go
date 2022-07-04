package main

import (
	"api-2/src"
	"sync"

	"github.com/joho/godotenv"
)

func main() {
	var loadOnce sync.Once
	godotenv.Load()
	srv := src.InitServer()

	loadOnce.Do(func() {
		srv.Run()
	})

}
