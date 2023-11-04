package sms

const (
	LenGSM7       = 160
	LenGSM7Concat = 153
	LenUCS2       = 70
	LenUCS2Concat = 70
)

func New(message string, options ...OptionFunc) *SMS {
	sms := &SMS{
		encoding: EncodingDefault,
		language: LangDefault,
		Message:  message,
	}

	for _, opt := range options {
		sms = opt(sms)
	}

	return sms
}

type SMS struct {
	encoding Encoding
	language Language
	Message  string
}

func (s SMS) Encode() (string, error) {
	return Encode(s.encoding, s.language, s.Message)
}

func (s SMS) Valid() bool {
	return Validate(s.encoding, s.language, s.Message)
}

func (s SMS) Len() int {
	return strLen(s.encoding, s.language, s.Message)
}

func (s SMS) Segments() int {
	return segmentLen(s.encoding, strLen(s.encoding, s.language, s.Message))
}
