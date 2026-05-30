package main

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

type ReadWriter interface {
	Reader
	Writer
}

type File struct {
	content string
}

func (f *File) Read() string {
	return f.content
}

func (f *File) Write(data string) {
	f.content = data
}

func copyData(rw ReadWriter) {
	data := rw.Read()

	rw.Write(data + " [processed]")
}

func compositionExample() {
	file := &File{
		content: "hello world",
	}

	copyData(file)

	fmt.Println(file.Read())
}