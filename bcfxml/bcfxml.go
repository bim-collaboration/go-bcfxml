package bcfxml

import (
	"errors"
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/os/gfile"
	"go-bcfxml/bcfxml/v3_0"
	"os"
)

type BcfXml struct {
	Version *Version
	V3_0_Data *v3_0.BcfData
}

type sBcfXml struct {
}

var (
	insBcfXml=sBcfXml{}
)

func InsBcfXml() *sBcfXml {
	return &insBcfXml
}

func (s *sBcfXml) ReadBcfFold(desfold string) (out *BcfXml, err error) {
	if !gfile.IsDir(desfold) {
		err = errors.New("fold not found")
		return
	}
	out=&BcfXml{}
	if out.Version,err=InsVersion().ReadFileVersion(desfold+"\\bcf.version");err != nil {
		return
	}
	switch out.Version.VersionId {
	case "3.0":
		out.V3_0_Data,err=v3_0.InsBcfService().ReadFoldData(desfold)
	}
	return
}

func (s *sBcfXml) WriteBcfFold(fold string,in *BcfXml) (err error) {
	if in == nil {
		err = errors.New("no data")
		return
	}
	if !gfile.IsDir(fold) {
		err = errors.New("fold not found")
		return
	}
	switch in.Version.VersionId {
	case "3.0":
		err=v3_0.InsBcfService().WriteFoldData(fold,in.V3_0_Data)
	}
	return
}

func (s *sBcfXml)UnCompressFile(srcFileName ,desFold string) error {
	ok:=gfile.Exists(desFold)
	if !ok {
		gfile.Mkdir(desFold)
	}else {
		files, _ := gfile.ScanDir(desFold, "*", false)
		if len(files) != 0 {
			for _,fname:=range files{
				os.RemoveAll(fname)
			}
		}
	}
	err := gcompress.UnZipFile(srcFileName, desFold)
	return err
}

func (s *sBcfXml)CompressFile(srcFold ,desFileName string) error {
	files, err := gfile.ScanDir(srcFold, "*", false)
	if (err != nil)||(len(files) == 0) {
		return errors.New("no file")
	}
	filenames:=files[0]
	for i := 1; i <len(files) ; i++ {
		filenames+=","+files[i]
	}
	err = gcompress.ZipPath(filenames, desFileName)
	return err
}




