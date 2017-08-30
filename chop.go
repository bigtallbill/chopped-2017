package main

import (
	"crypto/md5"
	"crypto/rand"
	"io/ioutil"

	"io"
	"fmt"
)
const HASH_LENGTH = 64

func PrintLn(s string) {
	text := []byte(s)
	ioutil.WriteFile("/dev/stdout", text, 0777)

}

func RandomMd5() string {
	b := make([]byte, HASH_LENGTH)

	_, err := rand.Read(b)
	if err != nil {
		PrintLn("error:")
		PrintLn(err.Error())
	}

	h := md5.New()
	s := string(b)
	io.WriteString(h, s)
	s2 := h.Sum(nil)

	return fmt.Sprintf("%x", s2)
}



func main() {

	PrintLn(RandomMd5())
}
