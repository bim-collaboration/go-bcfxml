package v3_0

import (
	"errors"
	"github.com/gogf/gf/v2/os/gfile"
)

type BcfData struct {
	ProjectInfo *ProjectInfo
	DocumentInfo *DocumentInfo
	Extensions *Extensions
	Markups []*Markup
	Visinfos map[string][]*VisualizationInfo
}


type sBcfService struct {
	InsProjectInfo *sProjectInfo
	InsDocumentInfo *sDocumentInfo
	InsExtensions *sExtensions
	InsMarkup *sMarkup
	InsVisualizationInfo *sVisualizationInfo
}

var (
	insBcfService=sBcfService{
		InsProjectInfo:newInsProjectInfo(),
		InsDocumentInfo:newInsDocumentInfo(),
		InsExtensions:newInsExtensions(),
		InsMarkup:newInsMarkup(),
		InsVisualizationInfo:newInsVisualizationInfo(),
	}
)

func InsBcfService() *sBcfService {
	return &insBcfService
}

func (s *sBcfService) ReadFoldData(fold string) (out *BcfData, err error) {
	if !gfile.IsDir(fold) {
		err = errors.New("fold not found")
		return
	}
	out=&BcfData{}

	//ProjectInfo
	if da,err:=s.InsProjectInfo.ReadFileProjectInfo(fold+"\\project.bcfp");err==nil {
		out.ProjectInfo=da
	}
	//DocumentInfo
	if da,err:=s.InsDocumentInfo.ReadFileDocumentInfo(fold+"\\documents.xml");err==nil {
		out.DocumentInfo=da
	}
	//Extensions
	if da,err:=s.InsExtensions.ReadFileExtensions(fold+"\\extensions.xml");err==nil {
		out.Extensions=da
	}
	//loop markup.bcf
	dirNames,_:=gfile.DirNames(fold)
	for _,pa:=range dirNames{
		pathmarkup:=fold+"\\"+pa
		if !gfile.IsDir(pathmarkup) {
			continue
		}
		pathmarkup+="\\markup.bcf"
		if !gfile.IsFile(pathmarkup) {
			continue
		}
		//markup
		if da,err:=s.InsMarkup.ReadFileMarkup(pathmarkup);err==nil {
			if da.Topic.Guid == pa {
				out.Markups=append(out.Markups,da)
			}
		}
	}
	//loop VisualizationInfo
	out.Visinfos=make(map[string][]*VisualizationInfo)
	for _,m:=range out.Markups{
		vis:=make([]*VisualizationInfo,0)
		for _,vi:=range m.Topic.ViewPoints{
			//Viewpoint
			pathVisinfo:=fold+"\\"+m.Topic.Guid+"\\"+vi.Viewpoint
			if da,err:=s.InsVisualizationInfo.ReadFileVisualizationInfo(pathVisinfo);err==nil {
				vis=append(vis,da)
			}
		}
		out.Visinfos[m.Topic.Guid]=vis
	}
	err=nil
	return
}

func (s *sBcfService) WriteFoldData(fold string,in *BcfData) (err error) {
	if in == nil {
		err = errors.New("no data")
		return
	}
	if !gfile.IsDir(fold) {
		err = errors.New("fold not found")
		return
	}
	//ProjectInfo
	s.InsProjectInfo.WriteFileProjectInfo(fold+"\\project.bcfp",in.ProjectInfo)
	//DocumentInfo
	s.InsDocumentInfo.WriteFileDocumentInfo(fold+"\\documents.xml",in.DocumentInfo)
	//Extensions
	s.InsExtensions.WriteFileExtensions(fold+"\\extensions.xml",in.Extensions)
	//loop markup.bcf
	for _,m:=range in.Markups{
		err=gfile.Mkdir(fold+"\\"+m.Topic.Guid)
		if err != nil {
			continue
		}
		s.InsMarkup.WriteFileMarkup(fold+"\\"+m.Topic.Guid+"\\markup.bcf",m)
	}
	//loop viewpoint
	for k,vis:=range in.Visinfos{
		for _,vi:=range vis{
			s.InsVisualizationInfo.WriteFileVisualizationInfo(fold+"\\"+k+"\\Viewpoint_"+vi.Guid+".bcfv",vi)
		}
	}
	err=nil
	return
}


