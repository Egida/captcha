package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albrazodiac/captcha"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/verify", verifyHandler)

	log.Println("Server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	captchaValue, captchaImageBase64, err := captcha.GenerateCaptcha()
	if err != nil {
		http.Error(w, "Failed to generate CAPTCHA", http.StatusInternalServerError)
		return
	}

	// Display the CAPTCHA image and provide an input field for the user
	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body>
			<img src="data:image/png;base64,%s" alt="CAPTCHA Image"><br>
			<form action="/verify" method="POST">
				<input type="text" name="captcha" required>
				<input type="hidden" name="captchaValue" value="%s">
				<input type="submit" value="Verify">
			</form>
		</body>
		</html>
	`, captchaImageBase64, captchaValue)

	fmt.Fprint(w, html)
}

func verifyHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	inputValue := r.Form.Get("captcha")
	captchaValue := r.Form.Get("captchaValue")

	err = captcha.VerifyCaptcha(inputValue)
	if err != nil {
		http.Error(w, "CAPTCHA verification failed", http.StatusForbidden)
		return
	}

	// CAPTCHA verified successfully
	fmt.Fprint(w, "CAPTCHA verified!")
}
