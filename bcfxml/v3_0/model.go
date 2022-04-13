package v3_0

import "encoding/xml"

//************************BCF3.0 common struct**********************//
//**************************************documents*********************************//
type Document struct {
	Guid        string `xml:"Guid,attr"`
	Description string `xml:"Description,omitempty"`
	Filename    string `xml:"Filename"`
}

//**************************************extensions*********************************//

//**************************************project*********************************//
type Project struct {
	ProjectId string   `xml:"ProjectId,attr"`
	Name      string   `xml:"Name,omitempty"`
}
//**************************************markup*********************************//

type ViewPoint struct {
	Viewpoint string `xml:"Viewpoint,omitempty"`
	Snapshot  string `xml:"Snapshot,omitempty"`
	Index     int    `xml:"Index,omitempty"`
	Guid      string `xml:"Guid,attr"`
}
type BimSnippet struct {
	Reference       string `xml:"Reference"`
	ReferenceSchema string `xml:"ReferenceSchema"`
	SnippetType     string `xml:"SnippetType,attr"`
	isExternal      bool   `xml:"isExternal,attr"`
}

type Topic struct {
	XMLName            xml.Name            `xml:"Topic"`
	ReferenceLinks     []string            `xml:"ReferenceLinks>ReferenceLink,omitempty"`
	Title              string              `xml:"Title"`
	Priority           string              `xml:"Priority,omitempty"`
	Index              string              `xml:"Index,omitempty"`
	Labels             []string            `xml:"Labels>Label,omitempty"`
	CreationDate       string              `xml:"CreationDate"`
	CreationAuthor     string              `xml:"CreationAuthor"`
	ModifiedDate       string              `xml:"ModifiedDate,omitempty"`
	ModifiedAuthor     string              `xml:"ModifiedAuthor,omitempty"`
	DueDate            string              `xml:"DueDate,omitempty"`
	AssignedTo         string              `xml:"AssignedTo,omitempty"`
	Stage              string              `xml:"Stage,omitempty"`
	Description        string              `xml:"Description,omitempty"`
	BimSnippet         *BimSnippet          `xml:"BimSnippet,omitempty"`
	DocumentReferences []*DocumentReference `xml:"DocumentReferences>DocumentReference,omitempty"`
	RelatedTopics      []*RelatedTopic      `xml:"RelatedTopics>RelatedTopic,omitempty"`
	Comments           []*Comment           `xml:"Comments>Comment,omitempty"`
	ViewPoints         []*ViewPoint         `xml:"Viewpoints>ViewPoint,omitempty"`
	Guid               string              `xml:"Guid,attr"`
	ServerAssignedId   string              `xml:"ServerAssignedId,attr,omitempty"`
	TopicType          string              `xml:"TopicType,attr"`
	TopicStatus        string              `xml:"TopicStatus,attr"`
}

type File struct {
	Filename                   string `xml:"Filename,omitempty"`
	Date                       string `xml:"Date,omitempty"`
	Reference                  string `xml:"Reference,omitempty"`
	IfcProject                 string `xml:"IfcProject,attr,omitempty"`
	IfcSpatialStructureElement string `xml:"IfcSpatialStructureElement,attr,omitempty"`
	IsExternal                 bool `xml:"IsExternal,attr"`
}

type RelatedTopic struct {
	Guid string `xml:"Guid,attr"`
}

type DocumentReference struct {
	DocumentGuid string `xml:"DocumentGuid,omitempty"`
	Url          string `xml:"Url,omitempty"`
	Description  string `xml:"Description,omitempty"`
	Guid         string `xml:"Guid,attr"`
}

type Comment struct {
	Date                  string           `xml:"Date"`
	Author                string           `xml:"Author"`
	Comment               string           `xml:"Comment,omitempty"`
	Viewpoint             *CommentViewpoint `xml:"Viewpoint,omitempty"`
	ModifiedDate          string           `xml:"ModifiedDate,omitempty"`
	ModifiedAuthor        string           `xml:"ModifiedAuthor,omitempty"`
	Guid                  string           `xml:"Guid,attr"`
}

type CommentViewpoint struct {
	Guid string `xml:"Guid,attr"`
}

//**************************************visinfo*********************************//
type OrthogonalCamera struct {
	CameraViewPoint  Point     `xml:"CameraViewPoint"`
	CameraDirection  Direction `xml:"CameraDirection"`
	CameraUpVector   Direction `xml:"CameraUpVector"`
	ViewToWorldScale float64   `xml:"ViewToWorldScale"`
	AspectRatio      float64   `xml:"AspectRatio"`
}

type PerspectiveCamera struct {
	CameraViewPoint Point     `xml:"CameraViewPoint"`
	CameraDirection Direction `xml:"CameraDirection"`
	CameraUpVector  Direction `xml:"CameraUpVector"`
	FieldOfView     float64   `xml:"FieldOfView"`
	AspectRatio     float64   `xml:"AspectRatio"`
}

type Point struct {
	X float64  `xml:"X"`
	Y float64  `xml:"Y"`
	Z float64  `xml:"Z"`
}

type Direction struct {
	X float64  `xml:"X"`
	Y float64  `xml:"Y"`
	Z float64  `xml:"Z"`
}

type Components struct {
	Selection  *ComponentSelection  `xml:"Selection,omitempty"`
	Visibility *ComponentVisibility `xml:"Visibility,omitempty"`
	Coloring   *ComponentColoring   `xml:"Coloring>Color,omitempty"`
}

type ComponentSelection struct {
	Component []*Component `xml:"Component,omitempty"`
}
type ComponentVisibility struct {
	ViewSetupHints *ViewSetupHints `xml:"ViewSetupHints,omitempty"`
	Exceptions      []*Component      `xml:"Exceptions>Component,omitempty"`
	DefaultVisibility      bool      `xml:"DefaultVisibility"`
}

type ViewSetupHints struct {
	SpacesVisible          bool `xml:"SpacesVisible,attr"`
	SpaceBoundariesVisible bool `xml:"SpaceBoundariesVisible,attr"`
	OpeningsVisible        bool `xml:"OpeningsVisible,attr"`
}
type ComponentColoring struct {
	Components []*Component `xml:"Components>Component"`
	Color      string      `xml:"Color,attr"`
}

type Component struct {
	OriginatingSystem string `xml:"OriginatingSystem,omitempty"`
	AuthoringToolId   string `xml:"AuthoringToolId,omitempty"`
	IfcGuid   string `xml:"IfcGuid,omitempty"`
}

type Line struct {
	StartPoint Point `xml:"StartPoint"`
	EndPoint   Point `xml:"EndPoint"`
}

type ClippingPlane struct {
	Location  Point     `xml:"Location"`
	Direction Direction `xml:"Direction"`
}

type Bitmap struct {
	Format    string    `xml:"Format"`
	Reference string    `xml:"Reference"`
	Location  Point     `xml:"Location"`
	Normal    Direction `xml:"Normal"`
	Up        Direction `xml:"Up"`
	Height    string    `xml:"Height"`
}
