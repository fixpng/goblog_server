package pwd

import (
	"fmt"
	"testing"
)

func TestHashPwd(t *testing.T) {
	fmt.Println(HashPwd("1234"))
}

func TestCheckPwd(t *testing.T) {
	fmt.Println(CheckPwd("$2a$04$JAqrLh3RnDIQQF6fjiRpeua6GT3oIchzFCKadJhCtnaZsPr1WLgfe", "1234"))
}
