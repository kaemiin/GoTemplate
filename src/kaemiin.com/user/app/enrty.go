package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jasonlvhit/gocron"
	"gopkg.in/natefinch/lumberjack.v2"
	"kaemiin.com/user/greet"
	"kaemiin.com/user/greet/zhTW"
)

func init() {
	fmt.Println("app/entry.go ==> init()")
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	LOG_FILE_LOCATION := os.Getenv("LOG_FILE_LOCATION")
	if LOG_FILE_LOCATION != "" {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LOG_FILE_LOCATION,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		})
	}
}

func task() {
	fmt.Println("app/entry.go ==> task()")
	isDebug := os.Getenv("IS_DEBUG")
	log.Println(isDebug)
	log.Println(greet.Morning)
	log.Println(zhTW.Morning)
	_, time := gocron.NextRun()
	fmt.Println(time)
}

func main() {
	fmt.Println("app/entry.go ==> main()")
	gocron.Every(3).Minutes().Do(task)
	<-gocron.Start()
}
