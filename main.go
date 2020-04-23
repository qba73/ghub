package main

import (
	"fmt"

	example_basic "github.com/qba73/ghub/src/basic"
)

func main() {
	doBasic()

}

func doBasic() {
	bm := example_basic.BasicMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "First basic message",
		SampleList: []int32{1, 2, 3, 4},
	}

	fmt.Println(bm)
}
