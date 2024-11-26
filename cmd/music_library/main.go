package main

import (
	"fmt"

	"github.com/romanpovol/music_library/internal/config"
)

func main() {
	cfg := config.LoadBackendConfig()

	fmt.Println(*cfg)

	// TODO: init logger

	// TODO: init storage

	// TODO: init router

	// TODO: run server
}
