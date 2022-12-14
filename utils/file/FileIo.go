package file

import (
	"fmt"
	"io/ioutil"
	"log"
)

func ReadOffset() string {
	file := "./resources/offset.txt"
	buf, _ := ioutil.ReadFile(file)
	log.Println("开始读取offset:")
	fmt.Printf("起始offset为"+"%s\n", buf)
	return string(buf)
}

func WriteOffset(str string) {
	file := "./resources/offset.txt"
	d := []byte(str)
	err := ioutil.WriteFile(file, d, 0666)
	if err != nil {
		fmt.Println("更新offset失败")
	}
	fmt.Println("更新offset成功")
}
