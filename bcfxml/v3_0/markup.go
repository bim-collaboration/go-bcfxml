package v3_0

import (
	"encoding/xml"
	"errors"
	"github.com/gogf/gf/v2/os/gfile"
)

type Markup struct {
	XMLName    xml.Name    `xml:"Markup"`
	Header     []*File      `xml:"Header>Files>File,omitempty"`
	Topic      *Topic       `xml:"Topic"`
}

type sMarkup struct {
}


func newInsMarkup() *sMarkup {
	return &sMarkup{}
}

func (s *sMarkup) ReadFileMarkup(path string) (out *Markup, err error) {

	if !gfile.IsFile(path) {
		err = errors.New("path is error")
		return
	}
	out = &Markup{}
	err = xml.Unmarshal(gfile.GetBytes(path), out)
	return
}

func (s *sMarkup) ReadMarkup(data []byte) (out *Markup, err error) {
	out = &Markup{}
	err = xml.Unmarshal(data, out)
	return
}

func (s *sMarkup) WriteMarkup(in *Markup) (out []byte, err error) {
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

func (s *sMarkup) WriteFileMarkup(path string, in *Markup) (err error) {
	if in == nil {
		err=errors.New("null data")
		return
	}
	fileHandle, err := gfile.Create(path)
	defer fileHandle.Close()
	if err != nil {
		return
	}
	if da, err := s.WriteMarkup(in); err == nil {
		_, err = fileHandle.Write(da)
	}
	return
}
