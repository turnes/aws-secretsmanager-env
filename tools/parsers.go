package tools

import (
	"encoding/json"
	"fmt"
)

func JsonToEnv(j string) (env string){
	var objmap map[string]interface{}
	json.Unmarshal([]byte(j), &objmap)
	for k, v := range objmap {
		fmt.Printf("%s=%s\n", k,v)
	}
	env = "ok"
	return 
}
