package entity

type Mail struct {
	To      string
	Subject string
	Body    string
}

func (Mail) TableName() string {
	return "mails"
}