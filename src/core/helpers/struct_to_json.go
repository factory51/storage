package helpers

import (
	"encoding/json"
	"fmt"
)

// StructToJson implements generics to accept any structure and unmarshall it to JSON

func StructToJson[T any](structure T) (decoded string, err error) {

	bytes, err := json.Marshal(structure)

	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	decoded = string(bytes)

	return

}
