package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile() {
	file, _ := os.Open("./resources/test.txt")
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	fmt.Println(string(bytes))
}
