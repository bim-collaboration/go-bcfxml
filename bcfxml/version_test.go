package bcfxml

import (
	"github.com/gogf/gf/v2/os/gfile"
	"testing"
)

func TestReadFileVersion(t *testing.T) {
	bcffile := gfile.Dir(".") + "/testdata/bcfxmltest/read/test1/bcf.version"
	data,err:=InsVersion().ReadFileVersion(bcffile)
	t.Log("CompressFile err=", data)
	t.Log("CompressFile err=", err)
}

func TestWriteFileVersion(t *testing.T) {
	//DetailedVersion:="dfsfdsfsf"
	ver:=&Version{
		VersionId: "2.1",
		//DetailedVersion: &DetailedVersion,
	}
	bcffile := gfile.Dir(".") + "/testdata/bcfxmltest/write/test1/bcf.version"
	err:=InsVersion().WriteFileVersion(bcffile,ver)
	t.Log("CompressFile err=", err)
}
