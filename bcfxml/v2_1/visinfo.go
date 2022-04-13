package v2_1

import (
	"encoding/xml"
	"errors"
	"github.com/gogf/gf/v2/os/gfile"
)

type VisualizationInfo struct {
	XMLName           xml.Name          `xml:"VisualizationInfo"`
	Components        *Components        `xml:"Components,omitempty"`
	OrthogonalCamera  *OrthogonalCamera  `xml:"OrthogonalCamera,omitempty"`
	PerspectiveCamera *PerspectiveCamera `xml:"PerspectiveCamera,omitempty"`
	Lines             *Line              `xml:"Lines>Line,omitempty"`
	ClippingPlanes    *ClippingPlane     `xml:"ClippingPlanes>ClippingPlane,omitempty"`
	Bitmap           []*Bitmap            `xml:"Bitmap,omitempty"`
	Guid           string            `xml:"Guid,attr"`
}

type sVisualizationInfo struct {
}



func newInsVisualizationInfo() *sVisualizationInfo {
	return &sVisualizationInfo{}
}

func (s *sVisualizationInfo) ReadFileVisualizationInfo(path string) (out *VisualizationInfo, err error) {

	if !gfile.IsFile(path) {
		err = errors.New("path is error")
		return
	}
	out = &VisualizationInfo{}
	err = xml.Unmarshal(gfile.GetBytes(path), out)
	return
}

func (s *sVisualizationInfo) ReadVisualizationInfo(data []byte) (out *VisualizationInfo, err error) {
	out = &VisualizationInfo{}
	err = xml.Unmarshal(data, out)
	return
}

func (s *sVisualizationInfo) WriteVisualizationInfo(in *VisualizationInfo) (out []byte, err error) {
	if in == nil {
		err=errors.New("null data")
		return
	}
	out = make([]byte, 0)
	if da, err := xml.Marshal(in); err == nil {
		out = append(out, []byte(xml.Header)...)
		out = append(out, da...)
	}
	return
}

func (s *sVisualizationInfo) WriteFileVisualizationInfo(path string, in *VisualizationInfo) (err error) {
	if in == nil {
		err=errors.New("null data")
		return
	}
	fileHandle, err := gfile.Create(path)
	defer fileHandle.Close()
	if err != nil {
		return
	}
	if da, err := s.WriteVisualizationInfo(in); err == nil {
		_, err = fileHandle.Write(da)
	}
	return
}
