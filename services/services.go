package services

import (
	"fmt"
	"os"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type Personalizations struct {
	To          []string
	DynamicData map[string]interface{}
	TemplateID  string
}

func ComposeDynamicTemplateEmail(personalizations Personalizations) {

	m := mail.NewV3Mail()

	//from
	e := mail.NewEmail(os.Getenv("SENDER_NAME"), os.Getenv("SENDER_EMAIL"))
	m.SetFrom(e)

	//template id
	m.SetTemplateID(personalizations.TemplateID)

	p := mail.NewPersonalization()

	for _, to := range personalizations.To {
		//mail.NewEmail("", to)
		p.AddTos(mail.NewEmail("", to))
	}

	// tos := []*mail.Email{
	// 	mail.NewEmail("Example User", "test1@example.com"),
	// 	mail.NewEmail("Example User", "test2@example.com"),
	// }
	// p.AddTos(tos...)

	//p.SetDynamicTemplateData("receipt", "true")
	p.DynamicTemplateData = personalizations.DynamicData

	m.AddPersonalizations(p)

	sendDynamicTemplateEmail(mail.GetRequestBody(m))

}

func sendDynamicTemplateEmail(body []byte) {
	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	var Body = body
	request.Body = Body
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
