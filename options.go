package sms

// OptionFunc enables passing in variadic functions to dictate how a given SMS
// is encoded.
type OptionFunc func(sms *SMS) *SMS

func WithEncoding(encoding Encoding) OptionFunc {
	return func(sms *SMS) *SMS {
		sms.encoding = encoding
		return sms
	}
}

func WithLanguage(language Language) OptionFunc {
	return func(sms *SMS) *SMS {
		sms.language = language
		return sms
	}
}
