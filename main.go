package main

import "gopkg.in/gomail.v2"

// Host SMTP
const Host = "smtp.mailtrap.io"

// Port SMTP
const Port = 587

// Username SMTP
const Username = "0c93259d9cb1b3"

// Password SMTP
const Password = "8015ee435b41eb"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "alex@example.com")
	m.SetHeader("To", "bob@example.com", "cora@example.com")
	m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")
	// m.Attach("/home/rafa/Pictures/arteon.png")

	d := gomail.NewDialer(Host, Port, Username, Password)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
