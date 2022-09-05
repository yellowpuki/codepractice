package main

import "os"

func main() {
	os.OpenFile("file.txt", os.O_CREATE, 0777) // check umask!!!! default 0002cd
}
