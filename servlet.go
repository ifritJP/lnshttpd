package lnshttpd

import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import "net/http"
import "fmt"
import "log"


func handleMain(
    resp http.ResponseWriter, req *http.Request,
    handler func ( req *Lnsservlet_RequestInfo ) *Lnsservlet_ResponseInfo ) {
    
    // ヘッダとステータスコードをセットする
    reqInfo := NewLnsservlet_RequestInfo(
        req.Method, req.URL.RequestURI(),
        Lns_mapFromGo( req.Header ), Newbind_goInStream( req.Body ) )
    info := handler( reqInfo )

    resp.WriteHeader( info.StatusCode )

    if info.Writer != nil {
        info.Writer.(Lnsservlet_writerFunc)( Newbind_goOutStream( resp ) )
    } else {
        fmt.Fprintf( resp, "%s", info.Txt )
    }
}

func Start( port int, infoList *LnsList ) {
    log.Print( "start -- ", port )

    
    for _, info := range( infoList.Items ) {
        handlerInfo := info.(*Lnsservlet_HandlerInfo)
        http.HandleFunc(
            handlerInfo.Path,
            func ( resp http.ResponseWriter, req *http.Request ) {
                handleMain( resp, req, handlerInfo.Handler )
            } )
    }
    
    log.Fatal(http.ListenAndServe( fmt.Sprintf( ":%d", port ), nil))
    
}
