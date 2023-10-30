package sms

import (
	"strings"
	"testing"
)

const defaultBaseSet = "*@jw"
const defaultExtendedSet = defaultBaseSet + "}"

func TestLen(t *testing.T) {
	type args struct {
		encoding Encoding
		lang     Language
		msg      string
	}
	tests := []struct {
		name         string
		args         args
		wantNumChars int
		wantNumSMS   int
	}{
		{
			name: "should return 0, 0 for an empty byte array with default encoding and default language",
			args: args{
				encoding: EncodingDefault,
				lang:     LangDefault,
				msg:      "",
			},
			wantNumChars: 0,
			wantNumSMS:   0,
		},
		{
			name: "should single default message",
			args: args{
				encoding: EncodingDefault,
				lang:     LangDefault,
				msg:      "hello world",
			},
			wantNumChars: 11,
			wantNumSMS:   1,
		},
		{
			name: "test single default message with extended charset",
			args: args{
				encoding: EncodingDefault,
				lang:     LangDefault,
				msg:      "{hello world}",
			},
			wantNumChars: 15,
			wantNumSMS:   1,
		},
		{
			name: "test single default message with extended charset boundary",
			args: args{
				encoding: EncodingDefault,
				lang:     LangDefault,
				msg:      strings.Repeat("{", 40) + strings.Repeat("}", 40),
			},
			wantNumChars: 160,
			wantNumSMS:   1,
		},
		{
			name: "test multi default message with base charset",
			args: args{
				encoding: EncodingDefault,
				lang:     LangDefault,
				msg:      "Hello world, I hope you are having a good day today. I'm looking forward to being able to message you in a way that will end up causing two SMS's to be sent. Bye",
			},
			wantNumChars: 161,
			wantNumSMS:   2,
		},
		{
			name: "test multi default message with extended charset",
			args: args{
				encoding: EncodingDefault,
				lang:     LangDefault,
				msg:      strings.Repeat("{", 40) + "0" + strings.Repeat("}", 40),
			},
			wantNumChars: 161,
			wantNumSMS:   2,
		},
		{
			name: "test UTF8 default message with charset",
			args: args{
				encoding: EncodingDefault,
				lang:     LangDefault,
				msg:      "🫥",
			},
			wantNumChars: 0,
			wantNumSMS:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNumChars, gotNumSMS := Len(tt.args.encoding, tt.args.lang, tt.args.msg)
			if gotNumChars != tt.wantNumChars {
				t.Errorf("Len() gotNumChars = %v, want %v", gotNumChars, tt.wantNumChars)
			}
			if gotNumSMS != tt.wantNumSMS {
				t.Errorf("Len() gotNumSMS = %v, want %v", gotNumSMS, tt.wantNumSMS)
			}
		})
	}
}

func BenchmarkLen(b *testing.B) {
	var x, y int
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		msg := strings.Repeat(defaultExtendedSet, b.N)
		b.StartTimer()

		x, y = Len(EncodingDefault, LangDefault, msg)
	}

	b.StopTimer()
	y, x = x, y
	b.StartTimer()
}

func Benchmark_segmentLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = segmentLen(EncodingDefault, b.N)
	}
}

func Benchmark_strLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		msg := strings.Repeat(defaultExtendedSet, b.N)
		b.StartTimer()

		_ = strLen(EncodingDefault, LangDefault, msg)
	}
}
