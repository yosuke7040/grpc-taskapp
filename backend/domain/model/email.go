package model

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	return &Email{value: value}, nil
}

func (email Email) Value() string {
	return email.value
}

func (email Email) Equals(other *Email) bool {
	return email.value == other.value
}
