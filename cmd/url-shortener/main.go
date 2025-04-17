package main

import (
	"REST_API/internal/config"
	"fmt"
)

func main() {

	cfg := config.MustLoad()
	fmt.Println(cfg)

}
ыы