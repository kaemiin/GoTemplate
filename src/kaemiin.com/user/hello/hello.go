package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
    "os"
	"kaemiin.com/user/stringutil"
)

func main()  {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	isDebug := os.Getenv("IS_DEBUG")
	fmt.Println(stringutil.Reverse("Hello, world."))
	fmt.Println(isDebug)
}
