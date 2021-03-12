package utils

import (
	"encoding/json"
	"net/http"
)

func Success(w http.ResponseWriter, value interface{}) error {
	res, err := json.Marshal(&value)
	if err != nil {
		return err
	}
	w.Write(res)
	return nil
}
