package sms

import (
	"bytes"
	"math"
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

func strLen(encoding Encoding, lang Language, msg string) (strLen int) {
	if len(msg) == 0 {
		return 0
	}

	charSet := charSetsGsm7[lang]
	for _, r := range bytes.Runes([]byte(msg)) {
		if _, ok := charSet.runeBase[r]; ok {
			strLen += 1
			continue
		}
		if _, ok := charSet.runeExtended[r]; ok {
			strLen += 2
			continue
		}
		if s, ok := charSet.runeReplacements[r]; ok {
			strLen += len(s)
		}
	}

	return strLen
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

func Validate(encoding Encoding, lang Language, msg string) bool {
	if encoding == EncodingUCS2 {
		return true
	}

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
