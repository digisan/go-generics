package v2

import (
	"bytes"
	"io"
	"log"
	"strings"
)

func StreamToBytes(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(stream); err != nil {
		log.Fatalln(err)
	}
	return buf.Bytes()
}

func BytesToStream(data []byte) io.Reader {
	return bytes.NewReader(data)
}

func StreamToString(stream io.Reader) string {
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(stream); err != nil {
		log.Fatalln(err)
	}
	return buf.String()
}

func StringToStream(str string) io.Reader {
	return strings.NewReader(str)
}
