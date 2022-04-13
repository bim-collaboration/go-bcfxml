package v3_0

import (
	"encoding/xml"
	"errors"
	"github.com/gogf/gf/v2/os/gfile"
)

type DocumentInfo struct {
	XMLName   xml.Name   `xml:"DocumentInfo"`
	Documents []*Document `xml:"Documents>Document,omitempty" json:"Documents"`
}

type sDocumentInfo struct {
}


func newInsDocumentInfo() *sDocumentInfo {
	return &sDocumentInfo{}
}

func (s *sDocumentInfo) ReadFileDocumentInfo(path string) (out *DocumentInfo, err error) {

	if !gfile.IsFile(path) {
		err = errors.New("path is error")
		return
	}
	out = &DocumentInfo{}
	err = xml.Unmarshal(gfile.GetBytes(path), out)
	return
}

func (s *sDocumentInfo) ReadDocumentInfo(data []byte) (out *DocumentInfo, err error) {
	out = &DocumentInfo{}
	err = xml.Unmarshal(data, out)
	return
}

func (s *sDocumentInfo) WriteDocumentInfo(in *DocumentInfo) (out []byte, err error) {
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

func (s *sDocumentInfo) WriteFileDocumentInfo(path string, in *DocumentInfo) (err error) {
	if in == nil {
		err=errors.New("null data")
		return
	}
	fileHandle, err := gfile.Create(path)
	defer fileHandle.Close()
	if err != nil {
		return
	}
	if da, err := s.WriteDocumentInfo(in); err == nil {
		_, err = fileHandle.Write(da)
	}
	return
}
