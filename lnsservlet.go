// This code is transcompiled by LuneScript.
package lnshttpd
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
var init_lnsservlet bool
var lnsservlet__mod__ string
type Lnsservlet_writerFunc func (arg1 Lnsservlet_outStream)
type Lnsservlet_handlerFunc func (arg1 *Lnsservlet_RequestInfo) *Lnsservlet_ResponseInfo
type Lnsservlet_inStream interface {
        Read(arg1 LnsInt)(LnsAny, string)
        ReadAll()(LnsAny, string)
}
func Lns_cast2Lnsservlet_inStream( obj LnsAny ) LnsAny {
    if _, ok := obj.(Lnsservlet_inStream); ok { 
        return obj
    }
    return nil
}

type Lnsservlet_outStream interface {
        Write(arg1 string) string
}
func Lns_cast2Lnsservlet_outStream( obj LnsAny ) LnsAny {
    if _, ok := obj.(Lnsservlet_outStream); ok { 
        return obj
    }
    return nil
}

// declaration Class -- ResponseInfo
type Lnsservlet_ResponseInfoMtd interface {
}
type Lnsservlet_ResponseInfo struct {
    StatusCode LnsInt
    Header *LnsMap
    Writer LnsAny
    Txt LnsAny
    FP Lnsservlet_ResponseInfoMtd
}
func Lnsservlet_ResponseInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Lnsservlet_ResponseInfo).FP
}
type Lnsservlet_ResponseInfoDownCast interface {
    ToLnsservlet_ResponseInfo() *Lnsservlet_ResponseInfo
}
func Lnsservlet_ResponseInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Lnsservlet_ResponseInfoDownCast)
    if ok { return work.ToLnsservlet_ResponseInfo() }
    return nil
}
func (obj *Lnsservlet_ResponseInfo) ToLnsservlet_ResponseInfo() *Lnsservlet_ResponseInfo {
    return obj
}
func NewLnsservlet_ResponseInfo() *Lnsservlet_ResponseInfo {
    obj := &Lnsservlet_ResponseInfo{}
    obj.FP = obj
    obj.InitLnsservlet_ResponseInfo()
    return obj
}
// 18: DeclConstr
func (self *Lnsservlet_ResponseInfo) InitLnsservlet_ResponseInfo() {
    self.StatusCode = 200
    
    self.Header = NewLnsMap( map[LnsAny]LnsAny{})
    
    self.Writer = nil
    
    self.Txt = ""
    
}


// declaration Class -- RequestInfo
type Lnsservlet_RequestInfoMtd interface {
}
type Lnsservlet_RequestInfo struct {
    Method string
    Url string
    Header *LnsMap
    Reader Lnsservlet_inStream
    FP Lnsservlet_RequestInfoMtd
}
func Lnsservlet_RequestInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Lnsservlet_RequestInfo).FP
}
type Lnsservlet_RequestInfoDownCast interface {
    ToLnsservlet_RequestInfo() *Lnsservlet_RequestInfo
}
func Lnsservlet_RequestInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Lnsservlet_RequestInfoDownCast)
    if ok { return work.ToLnsservlet_RequestInfo() }
    return nil
}
func (obj *Lnsservlet_RequestInfo) ToLnsservlet_RequestInfo() *Lnsservlet_RequestInfo {
    return obj
}
func NewLnsservlet_RequestInfo(arg1 string, arg2 string, arg3 *LnsMap, arg4 Lnsservlet_inStream) *Lnsservlet_RequestInfo {
    obj := &Lnsservlet_RequestInfo{}
    obj.FP = obj
    obj.InitLnsservlet_RequestInfo(arg1, arg2, arg3, arg4)
    return obj
}
func (self *Lnsservlet_RequestInfo) InitLnsservlet_RequestInfo(arg1 string, arg2 string, arg3 *LnsMap, arg4 Lnsservlet_inStream) {
    self.Method = arg1
    self.Url = arg2
    self.Header = arg3
    self.Reader = arg4
}

// declaration Class -- HandlerInfo
type Lnsservlet_HandlerInfoMtd interface {
}
type Lnsservlet_HandlerInfo struct {
    Path string
    Handler Lnsservlet_handlerFunc
    FP Lnsservlet_HandlerInfoMtd
}
func Lnsservlet_HandlerInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Lnsservlet_HandlerInfo).FP
}
type Lnsservlet_HandlerInfoDownCast interface {
    ToLnsservlet_HandlerInfo() *Lnsservlet_HandlerInfo
}
func Lnsservlet_HandlerInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Lnsservlet_HandlerInfoDownCast)
    if ok { return work.ToLnsservlet_HandlerInfo() }
    return nil
}
func (obj *Lnsservlet_HandlerInfo) ToLnsservlet_HandlerInfo() *Lnsservlet_HandlerInfo {
    return obj
}
func NewLnsservlet_HandlerInfo(arg1 string, arg2 Lnsservlet_handlerFunc) *Lnsservlet_HandlerInfo {
    obj := &Lnsservlet_HandlerInfo{}
    obj.FP = obj
    obj.InitLnsservlet_HandlerInfo(arg1, arg2)
    return obj
}
func (self *Lnsservlet_HandlerInfo) InitLnsservlet_HandlerInfo(arg1 string, arg2 Lnsservlet_handlerFunc) {
    self.Path = arg1
    self.Handler = arg2
}

func Lns_lnsservlet_init() {
    if init_lnsservlet { return }
    init_lnsservlet = true
    lnsservlet__mod__ = "@lnsservlet"
    Lns_InitMod()
}
func init() {
    init_lnsservlet = false
}
