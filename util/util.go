package util

import (
	"encoding/json"
	"fmt"
)

func PP(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	fmt.Println(" ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
	fmt.Println(" ")
}
