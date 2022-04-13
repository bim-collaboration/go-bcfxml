package v2_1

import (
	"encoding/xml"
	"errors"
	"github.com/gogf/gf/v2/os/gfile"
)

type ProjectInfo struct {
	XMLName   xml.Name `xml:"ProjectExtension"`
	Project *Project `xml:"Project,omitempty"`
	ExtensionSchema string `xml:"ExtensionSchema,omitempty"`
}

type sProjectInfo struct {
}

func newInsProjectInfo() *sProjectInfo {
	return &sProjectInfo{}
}

func (s *sProjectInfo) ReadFileProjectInfo(path string) (out *ProjectInfo, err error) {

	if !gfile.IsFile(path) {
		err = errors.New("path is error")
		return
	}
	out = &ProjectInfo{}
	err = xml.Unmarshal(gfile.GetBytes(path), out)
	return
}

func (s *sProjectInfo) ReadProjectInfo(data []byte) (out *ProjectInfo, err error) {
	out = &ProjectInfo{}
	err = xml.Unmarshal(data, out)
	return
}

func (s *sProjectInfo) WriteProjectInfo(in *ProjectInfo) (out []byte, err error) {
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

func (s *sProjectInfo) WriteFileProjectInfo(path string, in *ProjectInfo) (err error) {
	if in == nil {
		err=errors.New("null data")
		return
	}
	fileHandle, err := gfile.Create(path)
	defer fileHandle.Close()
	if err != nil {
		return
	}
	if da, err := s.WriteProjectInfo(in); err == nil {
		_, err = fileHandle.Write(da)
	}
	return
}
