package helpers

import (
	"encoding/json"
	"log"
)

func ToJson(data any) string {
	out, err := json.Marshal(data)
	if err != nil {
		log.Fatal("unamble to unmarshall")
	}

	return string(out)
}
