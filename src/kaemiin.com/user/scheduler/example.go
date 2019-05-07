package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jasonlvhit/gocron"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func task() {
	isDebug := os.Getenv("IS_DEBUG")
	fmt.Println(isDebug)
	_, time := gocron.NextRun()
	fmt.Println(time)
}

func main() {
	fmt.Println("START SCHEDULER...")
	gocron.Every(30).Minutes().Do(task)
	<-gocron.Start()
}
