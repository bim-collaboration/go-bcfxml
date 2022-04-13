package bcfxml

import (
	"errors"
	"github.com/bim-collaboration/go-bcfxml/bcfxml/v2_1"
	"github.com/bim-collaboration/go-bcfxml/bcfxml/v3_0"
	"github.com/gogf/gf/v2/encoding/gcompress"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/text/gstr"
	"os"
)

type BcfXml struct {
	Version *Version
	V3_0_Data *v3_0.BcfData
	V2_1_Data *v2_1.BcfData
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
	if out.Version,err=InsVersion().ReadFileVersion(desfold+"/bcf.version");err != nil {
		return
	}
	switch out.Version.VersionId {
	case "3.0":
		out.V3_0_Data,err=v3_0.InsBcfService().ReadFoldData(desfold)
	case "2.1":
		out.V2_1_Data,err=v2_1.InsBcfService().ReadFoldData(desfold)
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
	case "2.1":
		err=v2_1.InsBcfService().WriteFoldData(fold,in.V2_1_Data)
	}
	InsVersion().WriteFileVersion(fold+"/bcf.version",in.Version)
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

func (s *sBcfXml) CopyFile(srcFold,desFold string,in *BcfXml) (err error) {
	if in == nil {
		err = errors.New("no data")
		return
	}
	if !gfile.IsDir(srcFold) {
		err = errors.New("srcFold not found")
		return
	}
	if !gfile.IsDir(desFold) {
		err = errors.New("desFold not found")
		return
	}
	switch in.Version.VersionId {
	case "3.0":
		for _,m:=range in.V3_0_Data.Markups{
			for _,vis:=range m.Topic.ViewPoints{
				gfile.CopyFile(srcFold+"/"+m.Topic.Guid+"/"+vis.Snapshot,desFold+"/"+m.Topic.Guid+"/"+vis.Snapshot)
			}
		}
	case "2.1":
		for _,m:=range in.V2_1_Data.Markups{
			for _,vis:=range m.Viewpoints{
				gfile.CopyFile(srcFold+"/"+m.Topic.Guid+"/"+vis.Snapshot,desFold+"/"+m.Topic.Guid+"/"+vis.Snapshot)
			}
			for _,doc:=range m.Topic.DocumentReference{
				gfile.CopyFile(srcFold+"/"+gstr.TrimLeft(doc.ReferencedDocument,"../"),desFold+"/"+gstr.TrimLeft(doc.ReferencedDocument,"../"))
			}
			if  m.Topic.BimSnippet!= nil {
				gfile.CopyFile(srcFold+"/"+m.Topic.Guid+"/"+m.Topic.BimSnippet.Reference,desFold+"/"+m.Topic.Guid+"/"+m.Topic.BimSnippet.Reference)
			}
			for _,fi:=range m.Header{
				gfile.CopyFile(srcFold+"/"+gstr.TrimLeft(fi.Reference,"../"),desFold+"/"+gstr.TrimLeft(fi.Reference,"../"))
			}
		}
		for k,visinfos:=range in.V2_1_Data.Visinfos{
			for _,vis:=range visinfos{
				for _,bi:=range vis.Bitmap{
					gfile.CopyFile(srcFold+"/"+k+"/"+bi.Reference,desFold+"/"+k+"/"+bi.Reference)
				}
			}
		}
	}
	return
}



