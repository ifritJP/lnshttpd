// bind.lns から生成して、処理を追加している。



package lnshttpd
import . "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
import "io"
import "io/ioutil"
var init_bind bool
var bind__mod__ string
// declaration Class -- goInStream
type bind_goInStreamMtd interface {
    Read(arg1 LnsInt)(LnsAny, string)
    ReadAll()(LnsAny, string)
}
type bind_goInStream struct {
    reader io.Reader
    FP bind_goInStreamMtd
}
func bind_goInStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*bind_goInStream).FP
}
type bind_goInStreamDownCast interface {
    Tobind_goInStream() *bind_goInStream
}
func bind_goInStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(bind_goInStreamDownCast)
    if ok { return work.Tobind_goInStream() }
    return nil
}
func (obj *bind_goInStream) Tobind_goInStream() *bind_goInStream {
    return obj
}
func Newbind_goInStream(arg1 io.Reader) *bind_goInStream {
    obj := &bind_goInStream{}
    obj.FP = obj
    obj.Initbind_goInStream(arg1)
    return obj
}
func (self *bind_goInStream) Initbind_goInStream(arg1 LnsAny) {
    self.reader = arg1.(io.Reader)
}
// 5: decl @lnshttpd.@bind.goInStream.read
func (self *bind_goInStream) Read(size LnsInt)(LnsAny, string) {
    buf := make( []byte, size )
    size, err := self.reader.Read( buf )
    if err == io.EOF {
        return nil, "eof"
    }
    if err != nil {
        return nil, "err"
    }
    return string( buf[ :size] ), ""
}

// 8: decl @lnshttpd.@bind.goInStream.readAll
func (self *bind_goInStream) ReadAll()(LnsAny, string) {
    buf, err := ioutil.ReadAll( self.reader )
    if err == io.EOF {
        return nil, "eof"
    }
    if err != nil {
        return nil, "err"
    }
    return string( buf ), ""
}


// declaration Class -- goOutStream
type bind_goOutStreamMtd interface {
    Write(arg1 string) string
}
type bind_goOutStream struct {
    writer io.Writer
    FP bind_goOutStreamMtd
}
func bind_goOutStream2Stem( obj LnsAny ) LnsAny {
    if obj == nil {
        return nil
    }
    return obj.(*bind_goOutStream).FP
}
type bind_goOutStreamDownCast interface {
    Tobind_goOutStream() *bind_goOutStream
}
func bind_goOutStreamDownCastF( multi ...LnsAny ) LnsAny {
    if len( multi ) == 0 { return nil }
    obj := multi[ 0 ]
    if ddd, ok := multi[ 0 ].([]LnsAny); ok { obj = ddd[0] }
    work, ok := obj.(bind_goOutStreamDownCast)
    if ok { return work.Tobind_goOutStream() }
    return nil
}
func (obj *bind_goOutStream) Tobind_goOutStream() *bind_goOutStream {
    return obj
}
func Newbind_goOutStream(arg1 io.Writer) *bind_goOutStream {
    obj := &bind_goOutStream{}
    obj.FP = obj
    obj.Initbind_goOutStream(arg1)
    return obj
}
func (self *bind_goOutStream) Initbind_goOutStream(arg1 LnsAny) {
    self.writer = arg1.(io.Writer)
}
// 15: decl @lnshttpd.@bind.goOutStream.write
func (self *bind_goOutStream) Write(bin string) string {
    self.writer.Write( []byte(bin) )
    return ""
}


func Lns_bind_init() {
    if init_bind { return }
    init_bind = true
    bind__mod__ = "@lnshttpd.@bind"
    Lns_InitMod()
    Lns_lnsservlet_init()
    // {
    //     _ = Newbind_goInStream(0)
    //     _ = Newbind_goOutStream(0)
    // }
}
func init() {
    init_bind = false
}
