package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"gopkg.in/natefinch/lumberjack.v2"
	"net"
	"net/mail"
	"net/smtp"
	"crypto/tls"
)

func init() {
	fmt.Println("email/sender.go ==> init()")
	err := godotenv.Load()
	if err != nil {
		panic(errors.New("Error loading .env file"))
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

func send(subject string, body string, email string) (err error) {
	defer func() {
		if r := recover(); r != nil {
		   err = r.(error)
		}
	}()
	from := mail.Address{"", os.Getenv("EMAIL_FROM")}
	to   := mail.Address{"", email}
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
	headers["Subject"] = subject
	message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body
	servername := fmt.Sprintf("%s:%s", os.Getenv("EMAIL_HOST"), os.Getenv("EMAIL_PORT"))
	host, _, _ := net.SplitHostPort(servername)
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_USERNAME"), os.Getenv("EMAIL_PASSWORD"), host)
	c, err := smtp.Dial(servername)
    if err != nil {
        panic(err)
	}
	isStarttls, err := strconv.ParseBool(os.Getenv("EMIAL_STARTTLS.REQUIRED"))
	if err != nil {
        panic(err)
	}
	tlsconfig := &tls.Config {
        InsecureSkipVerify: !isStarttls,
        ServerName: host,
	}
    c.StartTLS(tlsconfig)
    if err = c.Auth(auth); err != nil {
        panic(err)
	}
	if err = c.Mail(from.Address); err != nil {
        panic(err)
    }
    if err = c.Rcpt(to.Address); err != nil {
        panic(err)
	}
	w, err := c.Data()
    if err != nil {
        panic(err)
    }
	_, err = w.Write([]byte(message))
    if err != nil {
        panic(err)
	}
	err = w.Close()
    c.Quit()
	return
}

func main() {
	fmt.Println("email/sender.go ==> main()")
	isDebug := os.Getenv("IS_DEBUG")
	log.Println(isDebug)
	subject := "THIS IS A TEST EMAIL 測試信件"
	body := "TEST 這是測試"
	email := "YYY@gmail.com"
	err := send(subject, body, email)
	if err != nil {
		log.Panic(err)
	}
}
