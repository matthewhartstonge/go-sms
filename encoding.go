package sms

import (
	"bytes"
	"math"
	"unicode/utf8"
)

const EncodingDefault = EncodingGSM7

type Encoding int

const (
	EncodingGSM7 Encoding = iota
	EncodingUCS2
)

func Encode(encoding Encoding, lang Language, msg string) (string, error) {
	// TODO: implement me!
	return "", nil
}

// Len given a language and encoding, will inform the count of SMS messages
// that will need to be sent to transmit the entire the message.
// O(n)
func Len(encoding Encoding, lang Language, msg string) (chars int, segments int) {
	chars = strLen(encoding, lang, msg)
	segments = segmentLen(encoding, chars)
	return
}

func StrLen(encoding Encoding, lang Language, msg string) int {
	return strLen(encoding, lang, msg)
}

func SegmentLen(encoding Encoding, lang Language, msg string) int {
	return segmentLen(encoding, strLen(encoding, lang, msg))
}

// strLen returns how many characters are in the message.
func strLen(encoding Encoding, lang Language, msg string) (strLen int) {
	if len(msg) == 0 {
		return 0
	}

	switch encoding {
	case EncodingGSM7:
		return gsmStrLen(charSetsGsm7[lang], msg)

	case EncodingUCS2:
		// TODO: implement UCS char counting.
		return -1
	}

	return 0
}

// gsmStrLen returns the number of characters in the provided message based on
// conversion to the characters in the base, extended and replacement character
// sets.
func gsmStrLen(charSet *CharacterSet, msg string) int {
	msgLen := 0
	for _, r := range bytes.Runes([]byte(msg)) {
		if n := gsmRuneLen(charSet, r); n > 0 {
			msgLen += n
			continue
		}
		if s, ok := charSet.runeReplacements[r]; ok {
			// replacements can be multi-character, but also be from the
			// extended set, so the counting gets... harder.
			for _, r := range bytes.Runes([]byte(s)) {
				msgLen += gsmRuneLen(charSet, r)
			}
		}
	}

	return msgLen
}

// gsmRuneLen returns the number of characters a given rune counts as, returns 0
// if not in the base or extended gsm7 character sets.
func gsmRuneLen(charSet *CharacterSet, r rune) int {
	if _, ok := charSet.runeBase[r]; ok {
		return 1
	}
	if _, ok := charSet.runeExtended[r]; ok {
		return 2
	}

	return 0
}

// Benchmark_segmentLen-10    	1000000000	         0.4253 ns/op	       0 B/op	       0 allocs/op
func segmentLen(encoding Encoding, strlen int) int {
	if strlen == 0 {
		return 0
	}

	segmentLength := LenGSM7
	switch encoding {
	case EncodingGSM7:
		if strlen > 160 {
			segmentLength = LenGSM7Concat
		}

	case EncodingUCS2:
		segmentLength = LenUCS2
		if strlen > 160 {
			segmentLength = LenUCS2Concat
		}
	}

	return int(math.Ceil(float64(strlen) / float64(segmentLength)))
}

// Validate returns whether the provided message can be successfully encoded by
// the given encoding and language.
func Validate(encoding Encoding, lang Language, msg string) bool {
	switch encoding {
	case EncodingGSM7:
		return gsm7Valid(lang, msg)

	case EncodingUCS2:
		return ucsValid(msg)

	default:
		return false
	}
}

func gsm7Valid(lang Language, msg string) bool {
	charSet := charSetsGsm7[lang]
	for _, r := range bytes.Runes([]byte(msg)) {
		if _, ok := charSet.runeBase[r]; ok {
			continue
		}
		if _, ok := charSet.runeExtended[r]; ok {
			continue
		}
		if _, ok := charSet.runeReplacements[r]; ok {
			continue
		}

		return false
	}

	return true
}

func ucsValid(msg string) bool {
	for _, r := range bytes.Runes([]byte(msg)) {
		if l := utf8.RuneLen(r); l == -1 {
			// If a rune returns -1, it's not valid UTF-8, therefore by the
			// transitive properties of UCS == UTF-8 we can verify that a
			// message cannot be valid if it can't be encoded as UTF-8.
			return false
		}
	}

	return true
}
