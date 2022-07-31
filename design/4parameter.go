package design

import "fmt"

type email struct {
	From, To, Subject, Body string
}

type Builder func(*email)

// not create instance
func Parameter() {
	sendEmail(func(e *email) {
		e.from("I'm 'from' part")
		e.to("I'm 'to' part")
		e.subject("I'm 'subject' part")
		e.body("I'm 'body' part")
	})
}

func sendEmail(act Builder) {
	e := email{}
	act(&e)
	send(&e)
}

func send(email *email) {
	fmt.Println(email)
}

func (e *email) from(from string) *email {
	e.From = from
	return e
}

func (e *email) to(to string) *email {
	e.To = to
	return e
}

func (e *email) subject(subject string) *email {
	e.Subject = subject
	return e
}

func (e *email) body(body string) *email {
	e.Body = body
	return e
}


