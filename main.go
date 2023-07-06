package main

import (
	"fmt"
	"golang-yaml/validation"

	m "github.com/hashicorp/go-multierror"
)

func main() {
	// even if one file fails the overall validation should fail.
	arr := make([]string, 0)

	arr = append(arr, "scripts/AnotherConcordPlaybook.yaml")
	arr = append(arr, "scripts/SimpleConcordPlaybook.yaml")
	listOfErrors := validation.ValidateYamlFile(arr).(*m.Error)
	if len(listOfErrors.Errors) > 0 {
		fmt.Println(validation.ValidateYamlFile(arr))
	} else {
		fmt.Println("Validation passed!")
	}

	fmt.Println()
}
