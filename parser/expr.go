package parser

import (
	"strings"
)

type strkey string

func (sk strkey) keyword() string { return string(sk) + ".keyword" }
func (sk strkey) string() string  { return string(sk) }

func parseExpr(token string) (key strkey, pattern, value string) {
	kpv := strings.Split(token, "=")
	kp := strings.Split(kpv[0], "__")
	key, pattern, value = strkey(kp[0]), kp[1], kpv[1]
	return key, pattern, value
}
