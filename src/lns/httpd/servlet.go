package httpd

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

var EnvMutex sync.Mutex

func handleMain(
	_env *LnsEnv, resp http.ResponseWriter, req *http.Request,
	handler func(_env *LnsEnv, req *Lnsservlet_RequestInfo) *Lnsservlet_ResponseInfo) {

	EnvMutex.Lock()

	// ヘッダとステータスコードをセットする
	reqInfo := NewLnsservlet_RequestInfo(
		_env, req.Method, req.URL.RequestURI(),
		Lns_mapFromGo(req.Header), Newbind_goInStream(_env, req.Body))
	info := handler(_env, reqInfo)

	for key, list := range info.Header.Items {
		lnsList := list.(*LnsList)
		for _, val := range lnsList.Items {
			resp.Header().Add(key.(string), val.(string))
		}
	}

	resp.WriteHeader(info.StatusCode)

	if info.Writer != nil {
		info.Writer.(Lnsservlet_writerFunc)(_env, Newbind_goOutStream(_env, resp))
	} else {
		if txt := info.Txt; txt != nil {
			resp.Write([]byte(txt.(string)))
		}
	}

	EnvMutex.Unlock()
}

func Start(_env *LnsEnv, port int, infoList *LnsList, hostingList *LnsList) {
	log.Print("start -- ", port)

	for _, info := range infoList.Items {
		handlerInfo := info.(*Lnsservlet_HandlerInfo)
		http.HandleFunc(
			handlerInfo.Path,
			func(resp http.ResponseWriter, req *http.Request) {
				handleMain(_env, resp, req, handlerInfo.Handler)
			})
	}
	for _, info := range hostingList.Items {
		hostingInfo := info.(*Lnsservlet_HostingInfo)

		localPath := hostingInfo.LocalPath
		if localPath[len(localPath)-1] != '/' {
			localPath += "/"
		}
		urlPath := hostingInfo.UrlPath
		if urlPath[len(urlPath)-1] != '/' {
			urlPath += "/"
		}

		log.Printf("hosting: %s: %s", localPath, urlPath)
		http.Handle(
			urlPath,
			http.StripPrefix(urlPath, http.FileServer(http.Dir(localPath))))
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
