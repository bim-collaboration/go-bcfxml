package bcfxml

import (
	"github.com/gogf/gf/v2/os/gfile"
	"testing"
)

//func TestReadFile(t *testing.T) {
//	//bcffile := gfile.Dir(".") + "/testdata/bcfxmltest/read/test1.bcf"
//	//fold := gfile.Dir(".") + "/testdata/bcfxmltest/read/test1"
//	fold := gfile.Dir(".") + "/testdata/bcfxmltest/read/unzipped"
//	//err:=InsBcfXml().UnCompressFile(bcffile,fold)
//	//if err != nil {
//	//	t.Log("UnCompressFile err=", err)
//	//}
//	data,err:=InsBcfXml().ReadBcfFold(fold)
//	if err != nil {
//		t.Log("ReadBcfFold err=", err)
//	}
//	t.Log("data =", data)
//}
//
//func TestWriteFile(t *testing.T) {
//	bcffile := gfile.Dir(".") + "/testdata/bcfxmltest/read/test1.bcf"
//	fold := gfile.Dir(".") + "/testdata/bcfxmltest/read/test1"
//	writefold := gfile.Dir(".") + "/testdata/bcfxmltest/write/test1"
//	writefile := gfile.Dir(".") + "/testdata/bcfxmltest/write/test1.bcf"
//	err:=InsBcfXml().UnCompressFile(bcffile,fold)
//	if err != nil {
//		t.Log("UnCompressFile err=", err)
//	}
//	data,err:=InsBcfXml().ReadBcfFold(fold)
//	if err != nil {
//		t.Log("ReadBcfFold err=", err)
//	}
//	gfile.Mkdir(writefold)
//	err=InsBcfXml().WriteBcfFold(writefold,data)
//	if err != nil {
//		t.Log("WriteBcfFold err=", err)
//	}
//	err=InsBcfXml().CompressFile(writefold,writefile)
//	if err != nil {
//		t.Log("CompressFile err=", err)
//	}
//}


//func TestReadFold(t *testing.T) {
//	fold := "D:/gotest/Test Cases/v2.1/Markup/ExternalBIMSnippet/unzipped"
//	data,err:=InsBcfXml().ReadBcfFold(fold)
//	if err != nil {
//		t.Log("ReadBcfFold err=", err)
//	}
//	t.Log("data =", data)
//}

func TestWriteFole(t *testing.T) {
	fi:="single invisible wall"
	fold := "D:/gotest/Test Cases/v2.1/Visualization/"+fi+"/unzipped"
	data,err:=InsBcfXml().ReadBcfFold(fold)
	if err != nil {
		t.Log("ReadBcfFold err=", err)
	}
	writefold :="D:/gotest/Test Cases/v2.1/Visualization/"+fi+"/"+fi
	writefile := "D:/gotest/Test Cases/v2.1/Visualization/"+fi+"/"+fi+".bcfzip"
	gfile.Mkdir(writefold)
	err=InsBcfXml().WriteBcfFold(writefold,data)
	if err != nil {
		t.Log("WriteBcfFold err=", err)
	}
	//copy file
	err=InsBcfXml().CopyFile(fold,writefold,data)
	if err != nil {
		t.Log("copy file err=", err)
	}

	err=InsBcfXml().CompressFile(writefold,writefile)
	if err != nil {
		t.Log("CompressFile err=", err)
	}
}