package httpd

import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
//import "fmt"
import "log"
import "strings"
import "encoding/json"


func ReadJsonObj(inStream Lnsservlet_inStream) LnsAny {
    txt, _ := inStream.ReadAll()

    jsonMap := map[string] LnsAny{}
    if txt != nil {
        decoder := json.NewDecoder( strings.NewReader( txt.(string) ) ) 
        if err := decoder.Decode( &jsonMap ); err != nil {
            log.Print( err )
            return nil
        }
    }
    ret := Lns_mapFromGo( jsonMap )
    return ret
}
func WriteJsonObj(outStream Lnsservlet_outStream, jsonMap *LnsMap ) {
    goMap := Lns_valToGo( jsonMap, true )
    
    bytes, _ := json.Marshal( goMap )
    outStream.Write( string(bytes) )
}
