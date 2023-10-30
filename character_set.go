package sms

import "github.com/matthewhartstonge/sms/gsm7"

// NewCharacterSet returns a character encoding set, with reversed maps for
// quick look up from either side.
func NewCharacterSet(
	language Language,
	runeBase map[rune]empty,
	runeExtended map[rune]empty,
	runeReplacements map[rune]string,
) *CharacterSet {
	return &CharacterSet{
		language: language,

		runeBase:         runeBase,
		runeExtended:     runeExtended,
		runeReplacements: runeReplacements,
	}
}

type empty = struct{}

// CharacterSet defines the character set that an SMS message is
type CharacterSet struct {
	language         Language
	runeBase         map[rune]empty
	runeExtended     map[rune]empty
	runeReplacements map[rune]string
}

var charSetsGsm7 = map[Language]*CharacterSet{
	LangDefault: NewCharacterSet(LangDefault, gsm7.DefaultBasic, gsm7.DefaultExtended, gsm7.DefaultReplacements),
	LangEnglish: NewCharacterSet(LangEnglish, gsm7.EnglishBasic, gsm7.EnglishExtended, gsm7.EnglishReplacements),
}
