package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	now := time.Now()
	version := now.Format("20060102150405")

	file, err := os.OpenFile("./app/config/migrations.go", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Println(err)
		os.Exit(7)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("%v", version))
	if err != nil {
		logger.Println(err)
		os.Exit(7)
	}

	logger.Println("Migration Created")
}
