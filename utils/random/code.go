package random

import (
	"fmt"
	"math/rand"
	"time"
)

var stringCode = "abcdefg"

func Code(length int) string {
	rand.Seed(time.Now().UnixNano())
	code := fmt.Sprintf("%v", rand.Intn(9999999))
	return code[0:length]
}
