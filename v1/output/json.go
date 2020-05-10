package output

import (
	"encoding/json"
	"net/http"
)

type Json struct {

}

var jsonType = []string{"application/json; charset=utf-8"}

func (j Json) Content(rw http.ResponseWriter, Value interface{}) error {
	header := rw.Header()


	//fmt.Printf("%v %v", header, value)
	// 如果没有找到 就去设置
	if val := header["Content-Type"]; len(val)==0 {
		header["Content-Type"] = jsonType
	}

	jsonBytes, err := json.Marshal(Value)
	if err != nil {
		return err
	}

	//rw.WriteHeader(200)
	rw.Write(jsonBytes)
	return nil
}