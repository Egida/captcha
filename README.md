# How it's work
1. When a user visits the home page ("/"), the homeHandler function is called. It generates a new CAPTCHA using the GenerateCaptcha function from the captcha package.
2. The CAPTCHA value and the base64-encoded CAPTCHA image are sent to the user's browser as part of the HTML response. The CAPTCHA image is displayed on the page along with an input field for the user to enter the CAPTCHA value.
3. When the user submits the form by clicking the "Verify" button, the form data is sent to the server with a POST request to the /verify route.
4. The verifyHandler function is called to handle the verification process. It retrieves the user's input value and the stored CAPTCHA value from the form data.
5. The VerifyCaptcha function from the captcha package is called to compare the user's input value with the stored CAPTCHA value. If the values match, the CAPTCHA is considered verified. Otherwise, an error is returned.
6. If the CAPTCHA is verified successfully, a message indicating the successful verification is sent back to the user's browser as the response.
   
The captcha package contains the core logic for generating and verifying CAPTCHAs. The GenerateCaptcha function generates a random CAPTCHA value and a CAPTCHA image. The CAPTCHA value is stored using the storeCaptchaValue function (which should be implemented based on your chosen storage mechanism). The VerifyCaptcha function retrieves the stored CAPTCHA value, compares it with the user's input, and clears the stored CAPTCHA value afterward.
Please note that the actual implementation of generating CAPTCHA images and storing/retrieving CAPTCHA values depends on your chosen libraries and mechanisms. You'll need to customize those parts according to your requirements.
In summary, this CAPTCHA system generates a CAPTCHA image and value, stores the value securely, presents the image to the user, and verifies the user's input against the stored value to prevent botnet requests.

# How to run
1. Set up your Go environment by installing Go from the official website and configuring your workspace.
2. Create a directory for your project and navigate to it using the terminal or command prompt.
3. Create the following files: captcha.go, main.go, and go.mod. Copy the code provided in the respective files.
4. Replace "github.com/your_username/captcha" in the main.go file with your actual GitHub username and repository name.
5. Run the following command to initialize your Go module:
- bash:
```bash
go mod init github.com/your_username/captcha
```  
6. Install any necessary packages or libraries required for generating CAPTCHA images and interacting with your chosen storage mechanism. Refer to the documentation of the specific libraries you're using to determine the installation steps.
7. Build and run the program using the following command:
- go:
```bash
go run main.go
```
8. The server will start running, and you should see a message like: "Server listening on port 8080..."
9. Open your web browser and navigate to http://localhost:8080 to access the CAPTCHA page.
10. You should see a CAPTCHA image and an input field. Enter the CAPTCHA value displayed in the image and click the "Verify" button.
11. If the CAPTCHA value is verified successfully, you will see a message indicating the successful verification.
---
You can now test the CAPTCHA system by refreshing the page to generate a new CAPTCHA and verifying it again.
Please note that this is a basic example, and you may need to modify and extend the code to fit your specific requirements, such as adding proper error handling, implementing the storage mechanism, and integrating a CAPTCHA image generation library.
