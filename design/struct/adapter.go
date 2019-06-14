package _struct

import "os"

//class Adapter

type IPs2 interface {
	transform(int) int
}

type IUsb interface {
	transform(int,int) int
}

type Usb struct {

}

func (u *Usb)transform(i1, i2 int) int  {
	return i1 + i2
}

//Ps2 -> Usb
type Ps2Adapter struct {
	usb IUsb //对象适配器
}
//类适配方式不存在
func (adapter *Ps2Adapter)transform(i int) int  {
	return adapter.usb.transform(i, 0)
}

//接口适配
type DbEngine interface {
	Read([]byte) error
}

type FileEngine interface {
	Read(os.File, []byte) error
}

type DbAdapter struct {
	FileEngine
}
//将数据库读取适配到从文件读取
func (adapter *DbAdapter)Read(byt []byte) error  {
	adapter.FileEngine.Read(os.File{}, byt)
	return nil
}