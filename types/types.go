package types

type TemplateData struct {
	Count    int
	Contacts []Contact
}

type Contact struct {
	Name  string
	Email string
}
