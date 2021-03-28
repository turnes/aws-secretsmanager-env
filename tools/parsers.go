package tools

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func JsonToEnv(j string) (env string, err error) {
	var objmap map[string]interface{}
	err = json.Unmarshal([]byte(j), &objmap)
	if err != nil {
		log.Println("It could not parse the json: ", err)
	}
	var fileString strings.Builder
	for k, v := range objmap {
		fmt.Fprintf(&fileString, "%s=%s\n", k, v)
	}
	env = fileString.String()
	return
}

//func EnvToJson(env file)  (json string, err error){
//	json = "ss"
//	return
//}
