package util

type DetailedError interface {
	Error() string
}

type detailedError struct {
	err      string
	detail   string
	location string
}
