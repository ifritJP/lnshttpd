// This code is transcompiled by LuneScript.
package main
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import lnsservlet "github.com/ifritJP/lnshttpd/src/lns/httpd"
import servlet "github.com/ifritJP/lnshttpd/src/lns/httpd"
var init_test bool
var test__mod__ string
// 3: decl @test.handle
func Test_handle(_env *LnsEnv, req *lnsservlet.Lnsservlet_RequestInfo) *lnsservlet.Lnsservlet_ResponseInfo {
    __func__ := "@test.handle"
    Lns_print([]LnsAny{__func__, req.Url})
    var resp *lnsservlet.Lnsservlet_ResponseInfo
    resp = lnsservlet.NewLnsservlet_ResponseInfo(_env)
    resp.Header.Set("HogeHeader",NewLnsList([]LnsAny{"HogeVal"}))
    resp.StatusCode = 200
    resp.Writer = lnsservlet.Lnsservlet_writerFunc(handle___anonymous_0_)
    return resp
}

// 24: decl @test.start
func Test_start(_env *LnsEnv, port LnsInt) {
    servlet.Start(_env, port, NewLnsList([]LnsAny{lnsservlet.Lnsservlet_HandlerInfo2Stem(lnsservlet.NewLnsservlet_HandlerInfo(_env, "/code", lnsservlet.Lnsservlet_handlerFunc(Test_handle)))}), NewLnsList([]LnsAny{lnsservlet.Lnsservlet_HostingInfo2Stem(lnsservlet.NewLnsservlet_HostingInfo(_env, "./indexableContents", "/contents"))}))
}

// 35: decl @test.__main
func Test___main(_env *LnsEnv, argList *LnsList) LnsInt {
    Lns_test_init( _env )
    Test_start(_env, 8080)
    return 0
}

func handle___anonymous_0_(_env *LnsEnv, stream lnsservlet.Lnsservlet_outStream) {
    stream.Write(_env, "resp1")
    stream.Write(_env, "resp2")
}

func Lns_test_init(_env *LnsEnv) {
    if init_test { return }
    init_test = true
    test__mod__ = "@test"
    Lns_InitMod()
    lnsservlet.Lns_lnsservlet_init(_env)
}
func init() {
    init_test = false
}
