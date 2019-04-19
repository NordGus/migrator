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
	filename := fmt.Sprintf("./migrations/%v_.sql", now.Format("20060102150405"))

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logger.Println(err)
		os.Exit(7)
	}
	file.Close()
	logger.Println("Migration Created")
}
