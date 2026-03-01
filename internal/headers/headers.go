package headers

import (
	"fmt"
	"strings"
)

type Headers map[string]string

func NewHeaders() Headers {
	return make(Headers)
}
func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := strings.Index(string(data), "\r\n")
	if idx == -1 {
		return 0, false, nil
	}
	if idx == 0 {
		return 2, true, nil
	}
	line := string(data[:idx])
	colon := strings.IndexByte(line, ':')
	if colon <= 0 {
		return 0, false, fmt.Errorf("invalid header line")
	}
	rawKey := line[:colon]
	// Field-name must not contain leading/trailing/interior whitespace.
	if rawKey == "" || rawKey != strings.TrimSpace(rawKey) || strings.ContainsAny(rawKey, " \t") {
		return 0, false, fmt.Errorf("invalid header line")
	}
	value := strings.Trim(line[colon+1:], " \t")
	if strings.Contains(value, " : ") {
		return 0, false, fmt.Errorf("invalid header line")
	}
	if strings.Contains(value, " :") {
		return 0, false, fmt.Errorf("invalid header line")
	}
	if strings.Contains(value, ": ") {
		return 0, false, fmt.Errorf("invalid header line")
	}
	h[rawKey] = value
	return idx + 2, false, nil

}
