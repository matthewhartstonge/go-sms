package sms

type Language string

func (l Language) String() string {
	return string(l)
}

const (
	LangDefault Language = "Default"
	LangEnglish Language = "English"
)
