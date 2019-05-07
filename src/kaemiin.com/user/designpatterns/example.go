package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"kaemiin.com/user/singleton"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	isDebug := os.Getenv("IS_DEBUG")
	fmt.Println(isDebug)
	fmt.Println("USE SINGLETON...")
	s := singleton.New()
	s["this"] = "that"
	s2 := singleton.New()
	fmt.Println("This is ", s2["this"])
	fmt.Println("USE SINGLETON...")
	// TODO
}
