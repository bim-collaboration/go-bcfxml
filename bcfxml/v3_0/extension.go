package v3_0

import (
	"encoding/xml"
	"errors"
	"github.com/gogf/gf/v2/os/gfile"
)

type Extensions struct {
	XMLName     xml.Name `xml:"Extensions"`
	TopicTypes   []string `xml:"TopicTypes>TopicType,omitempty"`
	TopicStatuses []string `xml:"TopicStatuses>TopicStatus,omitempty"`
	Priorities    []string `xml:"Priorities>Priority,omitempty"`
	TopicLabels  []string `xml:"TopicLabels>TopicLabel,omitempty"`
	Users        []string `xml:"Users>User,omitempty"`
	SnippetTypes []string `xml:"SnippetTypes>SnippetType,omitempty"`
	Stages       []string `xml:"Stages>Stage,omitempty"`
}

type sExtensions struct {
}


func newInsExtensions() *sExtensions {
	return &sExtensions{}
}

func (s *sExtensions) ReadFileExtensions(path string) (out *Extensions, err error) {

	if !gfile.IsFile(path) {
		err = errors.New("path is error")
		return
	}
	out = &Extensions{}
	err = xml.Unmarshal(gfile.GetBytes(path), out)
	return
}

func (s *sExtensions) ReadExtensions(data []byte) (out *Extensions, err error) {
	out = &Extensions{}
	err = xml.Unmarshal(data, out)
	return
}

func (s *sExtensions) WriteExtensions(in *Extensions) (out []byte, err error) {
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

func (s *sExtensions) WriteFileExtensions(path string, in *Extensions) (err error) {
	if in == nil {
		err=errors.New("null data")
		return
	}
	fileHandle, err := gfile.Create(path)
	defer fileHandle.Close()
	if err != nil {
		return
	}
	if da, err := s.WriteExtensions(in); err == nil {
		_, err = fileHandle.Write(da)
	}
	return
}
