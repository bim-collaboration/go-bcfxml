package bcfxml

import (
	"encoding/xml"
	"errors"
	"github.com/gogf/gf/v2/os/gfile"
)

type Version struct {
	XMLName   xml.Name `xml:"Version"`
	VersionId string   `xml:"VersionId,attr" json:"version_id"`
}

type sVersion struct {
}

func InsVersion() *sVersion {
	return &sVersion{}
}




func (s *sVersion) ReadFileVersion(path string) (out *Version, err error) {
	out = &Version{}
	if !gfile.IsFile(path) {
		err = errors.New("path is error")
		return
	}
	err = xml.Unmarshal(gfile.GetBytes(path), out)
	return
}

func (s *sVersion) ReadVersion(data []byte) (out *Version, err error) {
	out = &Version{}
	err = xml.Unmarshal(data, out)
	return
}

func (s *sVersion) WriteVersion(in *Version) (out []byte, err error) {
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

func (s *sVersion) WriteFileVersion(path string, in *Version) (err error) {
	if in == nil {
		err=errors.New("null data")
		return
	}
	fileHandle, err := gfile.Create(path)
	defer fileHandle.Close()
	if err != nil {
		return
	}
	if da, err := s.WriteVersion(in); err == nil {
		_, err = fileHandle.Write(da)
	}
	return
}
