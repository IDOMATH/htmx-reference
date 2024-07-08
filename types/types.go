package types

type TemplateData struct {
	Count    int
	Contacts []Contact
}

type Contact struct {
	Name  string
	Email string
}

func (d *TemplateData) HasEmail(email string) bool {
	for _, contact := range d.Contacts {
		if contact.Email == email {
			return true
		}
	}
	return false
}

type FormData struct {
	Values map[string]string
	Errors map[string]string
}

func NewFormData() *FormData {
	return &FormData{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}
