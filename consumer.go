package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func eamilWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {
	// signal done when the function exits
	defer wg.Done()
	// consume from the channel
	for recipient := range ch {
		smtpHost := "localhost"
		smtpPort := "1025"

		// formatedMsg := fmt.Sprintf("To: %s\r\nSubject: Test Emai\r\n\r\n%s\r\n", recipient.Email, "Just testing our email comapin.")
		// msg := []byte(formatedMsg)

		msg, err := executeTemplate(recipient)
		if err != nil {
			fmt.Printf("Worker :%d Error parsing template for %s", id, recipient.Email)
			continue
		}

		fmt.Printf("Worker %d: Sending email to %s \n", id, recipient.Email)

		err = smtp.SendMail(smtpHost+":"+smtpPort, nil, "rakibutsho1920@gmail.com", []string{recipient.Email}, []byte(msg))

		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(50 * time.Millisecond)

		fmt.Printf("Worker %d: sent email to %s \n", id, recipient.Email)

	}
}
