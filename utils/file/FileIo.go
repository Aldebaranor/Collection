package file

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadOffset() string {
	file := "./resources/offset.txt"
	buf, _ := ioutil.ReadFile(file)
	log.Println("Collection Program offset:")
	fmt.Printf("%s\n", buf)
	return string(buf)
}

func WriteOffset(str string) {
	file := "./resources/offset.txt"
	d := []byte(str)
	err := ioutil.WriteFile(file, d, 0666)
	if err != nil {
		fmt.Println("write fail")
	}
	fmt.Println("write success")
}
