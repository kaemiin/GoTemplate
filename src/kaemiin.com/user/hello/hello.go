package main

/*
	Commentary here
*/

// import "fmt"
import (
	"fmt"
	"github.com/joho/godotenv"
	"kaemiin.com/user/stringutil"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	isDebug := os.Getenv("IS_DEBUG")
	fmt.Println(stringutil.Reverse("Hello, world."))
	fmt.Println(isDebug)
}
