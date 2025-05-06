package utils

import (
	"math/rand"
	"strings"

	constants "github.com/guncv/tech-exam-software-engineering/constant"
)

func RandomString(n int) string {
	var sb strings.Builder
	k := len(constants.Alphabet)

	for i := 0; i < n; i++ {
		c := constants.Alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
