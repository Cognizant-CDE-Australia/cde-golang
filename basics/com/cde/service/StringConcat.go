package service

import (
	"bytes"
	"container/list"
	"fmt"
)

func StringConcat() {
	var buffer bytes.Buffer
	buffer.WriteString("Bhagavane")
	buffer.WriteString(" Vasudeva")
	fmt.Println(buffer.String())

	values := list.New()
	values.PushBack("bird")
	values.PushBack("cat")

}
