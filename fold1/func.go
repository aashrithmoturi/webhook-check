package fold1

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ValidateYamlFile(file string) error {
	yamlData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	var data interface{}
	err = yaml.Unmarshal(yamlData, &data)
	if err != nil {
		return err
	}
	//check branches
	_, ok := data.(map[interface{}]interface{})
	if ok == false {
		return fmt.Errorf("It is not covertable to a map Structure!")
	}
	return nil
}

func Reverse(st string) string {
	s := []rune(st)
	for i := 0; i < len(s)/2; i++ {
		c := s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = c
	}
	return string(s)
}
