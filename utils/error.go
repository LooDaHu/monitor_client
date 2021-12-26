package utils

import "fmt"

var (
	MarshalError    = fmt.Errorf("json marshal error")
	HTTPError       = fmt.Errorf("http error")
	ConfigReadError = fmt.Errorf("read config error")
)
