package sms

type Config struct {
	// Encoding sets whether we are working with GSM.7 or UCS.2.
	// By default, this is EncodingGSM7.
	Encoding Encoding
	// Language configures the character set to use.
	// By default, this is LangEnglish.
	Language Language
	// CharReplace will attempt to replace illegal characters with an equivalent
	// that is legal within the configured encoding space.
	CharReplace bool
	// CharRemove will strip characters that aren't within the encoding space.
	CharRemove bool
}
