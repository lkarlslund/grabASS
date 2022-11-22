package main

//go:generate easyjson types.go

//easyjson:json
type Result struct {
	Site    string              `json:"site,omitempty"`
	Body    string              `json:"body,omitempty"`
	Headers map[string][]string `json:"headers,omitempty"`
	Code    int                 `json:"resultcode,omitempty"`
}

type logger struct {
}

func (l logger) Errorf(format string, v ...interface{}) {}
func (l logger) Warnf(format string, v ...interface{})  {}
func (l logger) Debugf(format string, v ...interface{}) {}