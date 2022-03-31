package main

import (
	"errorsPractice/config"
	"fmt"
	"log"
)

func main() {
	confData, err := config.Load()
	if err != nil {
		log.Fatalf("Impossible to laad application config because: %s\n", err)
	}
	fmt.Println(confData)
}
