package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/sirupsen/logrus"
)

// How to generate interface stub
// c+p e *Email EmailSlicer

type Email struct {
	Mail     string
	Username string
	Domain   string
}

type EmailSlicer interface {
	IsValidUsername() bool
	IsValidDomain() bool
	IsEmailSlicer() map[string]string
}

func (e *Email) IsValidUsername(mail string) bool {
	usernameRegex := `^[a-zA-Z0-9_]{3,20}$`
	// Compile the regular expression
	re := regexp.MustCompile(usernameRegex)
	// Use the regular expression to check the validity of the username
	return re.MatchString(mail)
}

func (e *Email) IsValidDomain(mail string) bool {
	_, err := net.LookupIP(mail)
	return err == nil
}

func (e *Email) IsEmailSlicer() *Email {
	result := strings.Split(strings.TrimSpace(string(e.Mail)), "@")
	if len(result) == 2 {
		if e.IsValidUsername(result[0]) && e.IsValidDomain(result[1]) {
			log.WithFields(logrus.Fields{
				"mail":     e.Mail,
				"username": result[0],
				"domain":   result[1],
			}).Info("Valid Mail address")
			return &Email{Domain: result[1], Username: result[0]}
		} else {
			log.WithFields(logrus.Fields{
				"mail":     e.Mail,
				"domain":   result[1],
				"username": result[0],
				"err":      "invalid mail",
			}).Error("mail is not a valid")
			return nil
		}
	} else {
		log.WithFields(logrus.Fields{
			"mail": e.Mail,
			"err":  "invalid mail",
		}).Error("mail is not a valid")
		return nil
	}
}

var log = logrus.New()

func main() {

	b := Email{}
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		// Clear existing content by truncating the file
		file.Truncate(0)
		file.Seek(0, 0)

		log.SetOutput(file)
		defer file.Close()

	}

	reader := bufio.NewReader(os.Stdin)

	res, _, err := reader.ReadLine()
	if err != nil {
		log.WithFields(logrus.Fields{
			"err": "nil",
		}).Info("This is a log message with fields")
	}
	b.Mail = string(res)
	fmt.Println(b.IsEmailSlicer())

}
