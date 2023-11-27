package pkg

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"math/big"
	"net/smtp"
)

func GenerateOTP(email string) (string, error) {
	// Sender's email address and password
	from := "varun.sugur03@gmail.com"
	password := "lhtj mbsq ztiy iann"

	// Recipient's email address
	to := email

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	otp := generateOTP()

	// Message content
	message := []byte("Subject: Verify Email\n\nThe Verification code for verifying your email is:" + otp)

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err := smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		log.Error().Err(err).Msg("error in sending mail")
		fmt.Println("Error sending email:", err)

		return "", errors.New("mail is not sent")
	}
	fmt.Println("Email sent successfully!")
	return otp, nil

}

func generateOTP() string {
	const digits = 6
	otp := ""

	for i := 0; i < digits; i++ {
		// Generate a random number between 0 and 9
		randomNumber, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			log.Error().Err(err).Msg("error in generating random number")
			fmt.Println("Error generating random number:", err)
			return ""
		}

		// Append the random number to the OTP
		otp += randomNumber.String()
	}

	return otp
}
