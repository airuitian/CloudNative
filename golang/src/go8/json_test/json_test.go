package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info""`
	JobInfo   JobInfo   `json:"job_info""`
}

type BasicInfo struct {
	Name string `json:"name""`
	Age  int    `json:"age""`
}

type JobInfo struct {
	Skills []string `json:"skills""`
}

var jsonStr = `{
	"basic_info":{
      "name":"Mike",
      "age":30
    },
    "job_info":{
      "skills":["java","go","c"]
    }
}`

func TestJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)

	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}

}
