package idiomok

import (
	"encoding/json"
	"log"
	"strconv"
)

func UseCaseIdiomOk() {

	jsonBytes := []byte(`{"name":"John","age":"17","car":null}`)
	changeAge(jsonBytes)

}

func changeAge(jsonBytes []byte) {
	//Key is a string and value is an interface
	var jsonMap map[string]interface{}

	err := json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		log.Fatalln(err)
	}

	age := jsonMap["age"]
	if age == nil {
		log.Println("age is nil")
		return

	}

	//We're going to see what type is age

	//*Float64
	ageAsFloat64, ok := age.(float64)
	if ok {
		if ageAsFloat64 > 18 {
			log.Println("age is float64 and is greater than 18")
		} else {
			log.Println("age is float64 and is less than 18")
		}
		return
	}

	//*String
	if ageAsString, ok := age.(string); ok {
		//Convert ageAsString to float64
		ageAsFloat64, err := strconv.ParseFloat(ageAsString, 64)
		if err != nil {
			log.Println("age is string but can't be converted to float64 because it's not a number")
			return
		}
		if ageAsFloat64 > 18 {
			log.Println("age is string and is greater than 18")
		} else {
			log.Println("age is string and is less than 18")
		}
	} else {
		log.Println("age is not float64 or string")
	}
}
