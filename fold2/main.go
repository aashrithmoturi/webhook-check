package main

import (
	"fmt"
	"golang-yaml/fold1"
)

func main() {

	arr := make([]string, 0)
	arr = append(arr, "exmp.yaml")
	arr = append(arr, "exmp2.yaml")

	for _, file := range arr {
		err := fold1.ValidateYamlFile(file)

		fmt.Println("This is for ", file)
		if err != nil {
			fmt.Println("Failed to validate:", err)
		} else {
			fmt.Println("Passed the validaion")
		}
		fmt.Println("")
	}

}
