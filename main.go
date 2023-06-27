package main

import (
	"fmt"
	"golang-yaml/validation"

	"github.com/hashicorp/go-multierror"
)

func main() {
	//even if one file fails the overall validation should fail.
	arr := make([]string, 0)
	arr = append(arr, "../fold2/exmp.yaml")
	arr = append(arr, "../fold2/exmp2.yaml")
	listOfErrors := validation.ValidateYamlFile(arr).(*multierror.Error)
	if len(listOfErrors.Errors) > 0 {
		fmt.Println(validation.ValidateYamlFile(arr))
	} else {
		fmt.Println("Validation passed!")
	}

	fmt.Println()
}
