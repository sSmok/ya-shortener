// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package model

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson3df1d826DecodeGithubComSSmokYaShortenerInternalModel(in *jlexer.Lexer, out *OriginalLink) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "url":
			out.URL = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson3df1d826EncodeGithubComSSmokYaShortenerInternalModel(out *jwriter.Writer, in OriginalLink) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"url\":"
		out.RawString(prefix[1:])
		out.String(string(in.URL))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v OriginalLink) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson3df1d826EncodeGithubComSSmokYaShortenerInternalModel(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v OriginalLink) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson3df1d826EncodeGithubComSSmokYaShortenerInternalModel(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *OriginalLink) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson3df1d826DecodeGithubComSSmokYaShortenerInternalModel(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *OriginalLink) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson3df1d826DecodeGithubComSSmokYaShortenerInternalModel(l, v)
}