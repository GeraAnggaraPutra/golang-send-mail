package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

const (
	CONFIG_SMTP_HOST      = "smtp.gmail.com"
	CONFIG_SMTP_PORT      = 587
	CONFIG_SENDER_NAME    = "PT. Makmur Subur Jaya <anggaragera@gmail.com>"
	CONFIG_AUTH_EMAIL     = "anggaragera@gmail.com"
	CONFIG_AUTH_PASSWORD  = "xdsoxgooroerestq"
	NUM_EMAILS            = 10
)

func main() {
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", CONFIG_SENDER_NAME)
	mailer.SetHeader("To", "gerdyoung1234@gmail.com")
	mailer.SetAddressHeader("Cc", "gera_098@smkassalaambandung.sch.id", "Tra Lala La")
	mailer.SetHeader("Subject", "Test mail")
	mailer.SetBody("text/html", "Hello, <b>have a nice day</b>")
	// mailer.Attach("./sample.png")

	dialer := gomail.NewDialer(
		CONFIG_SMTP_HOST,
		CONFIG_SMTP_PORT,
		CONFIG_AUTH_EMAIL,
		CONFIG_AUTH_PASSWORD,
	)

	done := make(chan int, NUM_EMAILS)

	for i := 1; i <= NUM_EMAILS; i++ {
		go func(iter int) {
			err := dialer.DialAndSend(mailer)
			if err != nil {
				log.Printf("Perulangan ke-%d: %s\n", iter, err.Error())
			} else {
				done <- iter
			}
		}(i)
	}

	for i := 1; i <= NUM_EMAILS; i++ {
		result := <-done
		log.Printf("Perulangan ke-%d selesai.\n", result)
	}

	log.Println("Semua email terkirim!")
}
