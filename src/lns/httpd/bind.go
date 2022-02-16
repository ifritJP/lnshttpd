// This code is transcompiled by LuneScript.
package httpd

import (
	"io"
	"io/ioutil"

	. "github.com/ifritJP/LuneScript/src/lune/base/runtime_go"
)

var init_bind bool
var bind__mod__ string

// 5: decl @lns.@httpd.@bind.goInStream.read
func (self *bind_goInStream) Read(_env *LnsEnv, size LnsInt) (LnsAny, string) {
	buf := make([]byte, size)
	size, err := self.reader.Read(buf)
	if err == io.EOF {
		return nil, "eof"
	}
	if err != nil {
		return nil, "err"
	}
	return string(buf[:size]), ""
}

// 8: decl @lns.@httpd.@bind.goInStream.readAll
func (self *bind_goInStream) ReadAll(_env *LnsEnv) (LnsAny, string) {
	buf, err := ioutil.ReadAll(self.reader)
	if err == io.EOF {
		return nil, "eof"
	}
	if err != nil {
		return nil, "err"
	}
	return buf, ""
}

// 15: decl @lns.@httpd.@bind.goOutStream.write
func (self *bind_goOutStream) Write(_env *LnsEnv, bin string) string {
	(self.writer.(io.Writer)).Write([]byte(bin))
	return bin
}

// declaration Class -- goInStream
type bind_goInStreamMtd interface {
	Read(_env *LnsEnv, arg1 LnsInt) (LnsAny, string)
	ReadAll(_env *LnsEnv) (LnsAny, string)
}
type bind_goInStream struct {
	reader io.Reader
	FP     bind_goInStreamMtd
}

func bind_goInStream2Stem(obj LnsAny) LnsAny {
	if obj == nil {
		return nil
	}
	return obj.(*bind_goInStream).FP
}

type bind_goInStreamDownCast interface {
	Tobind_goInStream() *bind_goInStream
}

func bind_goInStreamDownCastF(multi ...LnsAny) LnsAny {
	if len(multi) == 0 {
		return nil
	}
	obj := multi[0]
	if ddd, ok := multi[0].([]LnsAny); ok {
		obj = ddd[0]
	}
	work, ok := obj.(bind_goInStreamDownCast)
	if ok {
		return work.Tobind_goInStream()
	}
	return nil
}
func (obj *bind_goInStream) Tobind_goInStream() *bind_goInStream {
	return obj
}
func Newbind_goInStream(_env *LnsEnv, arg1 io.Reader) *bind_goInStream {
	obj := &bind_goInStream{}
	obj.FP = obj
	obj.Initbind_goInStream(_env, arg1)
	return obj
}
func (self *bind_goInStream) Initbind_goInStream(_env *LnsEnv, arg1 io.Reader) {
	self.reader = arg1
}

// declaration Class -- goOutStream
type bind_goOutStreamMtd interface {
	Write(_env *LnsEnv, arg1 string) string
}
type bind_goOutStream struct {
	writer io.Writer
	FP     bind_goOutStreamMtd
}

func bind_goOutStream2Stem(obj LnsAny) LnsAny {
	if obj == nil {
		return nil
	}
	return obj.(*bind_goOutStream).FP
}

type bind_goOutStreamDownCast interface {
	Tobind_goOutStream() *bind_goOutStream
}

func bind_goOutStreamDownCastF(multi ...LnsAny) LnsAny {
	if len(multi) == 0 {
		return nil
	}
	obj := multi[0]
	if ddd, ok := multi[0].([]LnsAny); ok {
		obj = ddd[0]
	}
	work, ok := obj.(bind_goOutStreamDownCast)
	if ok {
		return work.Tobind_goOutStream()
	}
	return nil
}
func (obj *bind_goOutStream) Tobind_goOutStream() *bind_goOutStream {
	return obj
}
func Newbind_goOutStream(_env *LnsEnv, arg1 io.Writer) *bind_goOutStream {
	obj := &bind_goOutStream{}
	obj.FP = obj
	obj.Initbind_goOutStream(_env, arg1)
	return obj
}
func (self *bind_goOutStream) Initbind_goOutStream(_env *LnsEnv, arg1 io.Writer) {
	self.writer = arg1
}

func Lns_bind_init(_env *LnsEnv) {
	if init_bind {
		return
	}
	init_bind = true
	bind__mod__ = "@lns.@httpd.@bind"
	Lns_InitMod()
	Lns_lnsservlet_init(_env)
	{
		_ = Newbind_goInStream(_env, nil)
		_ = Newbind_goOutStream(_env, nil)
	}
}
func init() {
	init_bind = false
}
