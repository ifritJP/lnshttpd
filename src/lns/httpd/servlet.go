package httpd

import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import "net/http"
import "fmt"
import "log"
import "sync"

var EnvMutex sync.Mutex


func handleMain(
    resp http.ResponseWriter, req *http.Request,
    handler func ( req *Lnsservlet_RequestInfo ) *Lnsservlet_ResponseInfo ) {

    EnvMutex.Lock()
    
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

    EnvMutex.Unlock()
}

func Start( port int, infoList *LnsList, hostingList *LnsList ) {
    log.Print( "start -- ", port )

    
    for _, info := range( infoList.Items ) {
        handlerInfo := info.(*Lnsservlet_HandlerInfo)
        http.HandleFunc(
            handlerInfo.Path,
            func ( resp http.ResponseWriter, req *http.Request ) {
                handleMain( resp, req, handlerInfo.Handler )
            } )
    }
    for _, info := range( hostingList.Items ) {
        hostingInfo := info.(*Lnsservlet_HostingInfo)

        localPath := hostingInfo.LocalPath
        if localPath[ len(localPath) - 1 ] != '/' {
            localPath += "/"
        }
        urlPath := hostingInfo.UrlPath
        if urlPath[ len( urlPath ) - 1 ] != '/' {
            urlPath += "/"
        }
        
        log.Printf( "hosting: %s: %s", localPath, urlPath );
        http.Handle(
            urlPath,
            http.StripPrefix( urlPath, http.FileServer(http.Dir( localPath ))))
    }
    
    log.Fatal(http.ListenAndServe( fmt.Sprintf( ":%d", port ), nil))
}
