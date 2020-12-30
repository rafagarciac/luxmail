package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readEnv() (string, int, string, string) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Host := os.Getenv("SMTP_HOST")
	Port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	Username := os.Getenv("SMTP_USERNAME")
	Password := os.Getenv("SMTP_PASSWORD")

	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	return Host, Port, Username, Password
}

func readHTML() string {
	dat, err := ioutil.ReadFile("./html/index.html")
	check(err)
	return string(dat)
}

func main() {
	Host, Port, Username, Password := readEnv()

	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", readHTML())
	// m.Attach("/home/rafa/Pictures/arteon.png")

	d := gomail.NewDialer(Host, Port, Username, Password)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
