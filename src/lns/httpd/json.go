package httpd

import (
	"encoding/json"
	"log"
	"strings"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

//import "fmt"

func ReadJsonObj(_env *LnsEnv, inStream Lnsservlet_inStream) LnsAny {
	txt, _ := inStream.ReadAll(_env)

	jsonMap := map[string]LnsAny{}
	if txt != nil {
		decoder := json.NewDecoder(strings.NewReader(txt.(string)))
		if err := decoder.Decode(&jsonMap); err != nil {
			log.Print(err)
			return nil
		}
	}
	ret := Lns_mapFromGo(jsonMap)
	return ret
}
func WriteJsonObj(_env *LnsEnv, outStream Lnsservlet_outStream, jsonMap *LnsMap) {
	goMap := Lns_valToGo(jsonMap, true)

	bytes, _ := json.Marshal(goMap)
	outStream.Write(_env, string(bytes))
}
