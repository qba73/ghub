package main

import (
	"fmt"

	basicpb "github.com/qba73/ghub/src/basic"
)

func main() {
	doBasic()

}

func doBasic() {
	bm := basicpb.BasicMessage{
		Id:         1234,
		IsSimple:   true,
		Name:       "First basic message",
		SampleList: []int32{1, 2, 3, 4},
	}

	fmt.Println(bm)
}
