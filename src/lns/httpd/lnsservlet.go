// This code is transcompiled by LuneScript.
package httpd
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

// declaration Class -- luaInStream
type Lnsservlet_luaInStreamMtd interface {
    Read(arg1 LnsInt)(LnsAny, string)
    ReadAll()(LnsAny, string)
    readStream(arg1 LnsAny)(LnsAny, string)
}
type Lnsservlet_luaInStream struct {
    stream Lns_iStream
    FP Lnsservlet_luaInStreamMtd
}
func Lnsservlet_luaInStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Lnsservlet_luaInStream).FP
}
type Lnsservlet_luaInStreamDownCast interface {
    ToLnsservlet_luaInStream() *Lnsservlet_luaInStream
}
func Lnsservlet_luaInStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Lnsservlet_luaInStreamDownCast)
    if ok { return work.ToLnsservlet_luaInStream() }
    return nil
}
func (obj *Lnsservlet_luaInStream) ToLnsservlet_luaInStream() *Lnsservlet_luaInStream {
    return obj
}
func NewLnsservlet_luaInStream(arg1 Lns_iStream) *Lnsservlet_luaInStream {
    obj := &Lnsservlet_luaInStream{}
    obj.FP = obj
    obj.InitLnsservlet_luaInStream(arg1)
    return obj
}
func (self *Lnsservlet_luaInStream) InitLnsservlet_luaInStream(arg1 Lns_iStream) {
    self.stream = arg1
}
// 12: decl @lns.@httpd.@lnsservlet.luaInStream.readStream
func (self *Lnsservlet_luaInStream) readStream(mode LnsAny)(LnsAny, string) {
    {
        _bin := self.stream.Read(mode)
        if _bin != nil {
            bin := _bin.(string)
            return bin, ""
        }
    }
    return nil, "err"
}

// 18: decl @lns.@httpd.@lnsservlet.luaInStream.read
func (self *Lnsservlet_luaInStream) Read(size LnsInt)(LnsAny, string) {
    return self.FP.readStream(size)
}

// 21: decl @lns.@httpd.@lnsservlet.luaInStream.readAll
func (self *Lnsservlet_luaInStream) ReadAll()(LnsAny, string) {
    return self.FP.readStream("*a")
}


// declaration Class -- luaOutStream
type Lnsservlet_luaOutStreamMtd interface {
    Write(arg1 string) string
}
type Lnsservlet_luaOutStream struct {
    stream Lns_oStream
    FP Lnsservlet_luaOutStreamMtd
}
func Lnsservlet_luaOutStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Lnsservlet_luaOutStream).FP
}
type Lnsservlet_luaOutStreamDownCast interface {
    ToLnsservlet_luaOutStream() *Lnsservlet_luaOutStream
}
func Lnsservlet_luaOutStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Lnsservlet_luaOutStreamDownCast)
    if ok { return work.ToLnsservlet_luaOutStream() }
    return nil
}
func (obj *Lnsservlet_luaOutStream) ToLnsservlet_luaOutStream() *Lnsservlet_luaOutStream {
    return obj
}
func NewLnsservlet_luaOutStream(arg1 Lns_oStream) *Lnsservlet_luaOutStream {
    obj := &Lnsservlet_luaOutStream{}
    obj.FP = obj
    obj.InitLnsservlet_luaOutStream(arg1)
    return obj
}
func (self *Lnsservlet_luaOutStream) InitLnsservlet_luaOutStream(arg1 Lns_oStream) {
    self.stream = arg1
}
// 28: decl @lns.@httpd.@lnsservlet.luaOutStream.write
func (self *Lnsservlet_luaOutStream) Write(bin string) string {
    var err LnsAny
    _,err = self.stream.Write(bin)
    if err != nil{
        err_46 := err.(string)
        return err_46
    }
    return ""
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
// 45: DeclConstr
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

// declaration Class -- HostingInfo
type Lnsservlet_HostingInfoMtd interface {
}
type Lnsservlet_HostingInfo struct {
    LocalPath string
    UrlPath string
    FP Lnsservlet_HostingInfoMtd
}
func Lnsservlet_HostingInfo2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*Lnsservlet_HostingInfo).FP
}
type Lnsservlet_HostingInfoDownCast interface {
    ToLnsservlet_HostingInfo() *Lnsservlet_HostingInfo
}
func Lnsservlet_HostingInfoDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(Lnsservlet_HostingInfoDownCast)
    if ok { return work.ToLnsservlet_HostingInfo() }
    return nil
}
func (obj *Lnsservlet_HostingInfo) ToLnsservlet_HostingInfo() *Lnsservlet_HostingInfo {
    return obj
}
func NewLnsservlet_HostingInfo(arg1 string, arg2 string) *Lnsservlet_HostingInfo {
    obj := &Lnsservlet_HostingInfo{}
    obj.FP = obj
    obj.InitLnsservlet_HostingInfo(arg1, arg2)
    return obj
}
func (self *Lnsservlet_HostingInfo) InitLnsservlet_HostingInfo(arg1 string, arg2 string) {
    self.LocalPath = arg1
    self.UrlPath = arg2
}

func Lns_lnsservlet_init() {
    if init_lnsservlet { return }
    init_lnsservlet = true
    lnsservlet__mod__ = "@lns.@httpd.@lnsservlet"
    Lns_InitMod()
}
func init() {
    init_lnsservlet = false
}
