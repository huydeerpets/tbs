package mail

import (
	"bytes"
	"github.com/huydeerpets/tbs/utils"
	"html/template"
	"net/smtp"

	"github.com/astaxie/beego"
)

// Body 
type Body struct {
	From    string
	To      string
	Subject string
	Message string
}

// Send 
func Send(email string, body []byte) error {
	if utils.IsTest() {
		return nil
	}

	auth := smtp.PlainAuth(
		"",
		beego.AppConfig.String("email"),
		beego.AppConfig.String("emailpass"),
		"smtp.gmail.com",
	)

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		beego.AppConfig.String("email"),
		[]string{email},
		body,
	)

	return err
}

// GetBody 
func GetBody(b Body) []byte {
	buffer := new(bytes.Buffer)
	template := template.Must(template.New("emailTemplate").Parse(getBodyTemplate()))
	template.Execute(buffer, &b)

	return buffer.Bytes()
}

// getBodyTemplate 
func getBodyTemplate() string {
	return "To: {{.To}}\r\n" +
		"Subject: {{.Subject}}\r\n" +
		"\r\n" +
		"{{.Message}}"
}

// ForgetpasswordTemplate
type ForgetpasswordTemplate struct {
	URL   string
	Host  string
	Email string
}

// GetForgetpasswordBody 
func GetForgetpasswordBody(f ForgetpasswordTemplate) []byte {
	t: = "Request for password change accepted. \r\n" +
"\r\n" +
"Please change your password from the following URL \r\n" +
"{{.URL}} \r\n" +
"\r\n" +
"※ This URL is valid for 1 hour from issuance. \r\n" +
"* If you have applied for password change multiple times within 1 hour, please note that only the most recently issued URL is valid. \r\n" +
"\r\n" +
"-------------------------------------------- \r\n" +
"dotstamp: {{.Host}} \r\n" +
"Contact Us: {{.Email}}"

	buffer := new(bytes.Buffer)
	template := template.Must(template.New("forgetPassword").Parse(t))
	template.Execute(buffer, &f)

	return buffer.Bytes()
}

// GetForgetpasswordURL 
func GetForgetpasswordURL(email string, keyword string) (string, error) {
	e, err := utils.Encrypter([]byte(email))
	if err != nil {
		return "", err
	}
	k, err := utils.Encrypter([]byte(keyword))
	if err != nil {
		return "", err
	}

	return utils.Urlencode(e) + "/" + utils.Urlencode(k), nil
}
