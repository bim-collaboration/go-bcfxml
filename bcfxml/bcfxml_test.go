package bcfxml

import (
	"github.com/gogf/gf/v2/os/gfile"
	"testing"
)

func TestReadFile1(t *testing.T) {
	bcffile := gfile.Dir(".") + "\\testdata\\bcfxmltest\\read\\test1.bcf"
	fold := gfile.Dir(".") + "\\testdata\\bcfxmltest\\read\\test1"
	err:=InsBcfXml().UnCompressFile(bcffile,fold)
	if err != nil {
		t.Log("UnCompressFile err=", err)
	}
	data,err:=InsBcfXml().ReadBcfFold(fold)
	if err != nil {
		t.Log("ReadBcfFold err=", err)
	}
	t.Log("data =", data)
}

func TestWriteFile1(t *testing.T) {
	bcffile := gfile.Dir(".") + "\\testdata\\bcfxmltest\\read\\test1.bcf"
	fold := gfile.Dir(".") + "\\testdata\\bcfxmltest\\read\\test1"
	writefold := gfile.Dir(".") + "\\testdata\\bcfxmltest\\write\\test1"
	writefile := gfile.Dir(".") + "\\testdata\\bcfxmltest\\write\\test1.bcf"
	err:=InsBcfXml().UnCompressFile(bcffile,fold)
	if err != nil {
		t.Log("UnCompressFile err=", err)
	}
	data,err:=InsBcfXml().ReadBcfFold(fold)
	if err != nil {
		t.Log("ReadBcfFold err=", err)
	}
	gfile.Mkdir(writefold)
	err=InsBcfXml().WriteBcfFold(writefold,data)
	if err != nil {
		t.Log("WriteBcfFold err=", err)
	}
	err=InsBcfXml().CompressFile(writefold,writefile)
	if err != nil {
		t.Log("CompressFile err=", err)
	}
}