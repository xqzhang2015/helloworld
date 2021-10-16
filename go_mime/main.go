package main

import (
	"fmt"
	"io/ioutil"

	"github.com/h2non/filetype"
)

func main() {
	buf, _ := ioutil.ReadFile("sample.mp4")

	kind, _ := filetype.Match(buf)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
		return
	}

	fmt.Printf("File type: %s. MIME: %s, MIME: %#v\n", kind.Extension, kind.MIME.Value, kind.MIME)
}

// for sample.mp4
// File type: mp4. MIME: video/mp4, MIME: types.MIME{Type:"video", Subtype:"mp4", Value:"video/mp4"}

// for sample.jpg
// File type: jpg. MIME: image/jpeg, MIME: types.MIME{Type:"image", Subtype:"jpeg", Value:"image/jpeg"}

// more doc info
// https://dev.to/sistoi/golang-mime-type-handling-3fnd
