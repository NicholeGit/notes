package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	Test int32
	M    int32
}

func main() {
	x := `{
			"test":"1",
			"m":1
		}`

	t := &T{}
	err := json.Unmarshal([]byte(x), t)
	if err != nil {
		fmt.Println(err, t)
	} else {
		fmt.Println(t)
	}

}
