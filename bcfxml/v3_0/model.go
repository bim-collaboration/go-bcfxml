package v3_0

import "encoding/xml"

//************************BCF3.0 common struct**********************//
//**************************************documents*********************************//
type Document struct {
	Guid        string `xml:"Guid,attr" json:"Guid"`
	Description string `xml:"Description,omitempty" json:"Description"`
	Filename    string `xml:"Filename,omitempty" json:"Filename"`
}

//**************************************extensions*********************************//

//**************************************project*********************************//
type Project struct {
	ProjectId string   `xml:"ProjectId,attr" json:"project_id"`
	Name      string   `xml:"Name,omitempty" json:"Name"`
}
//**************************************markup*********************************//

type ViewPoint struct {
	Viewpoint string `xml:"Viewpoint,omitempty" json:"Viewpoint"`
	Snapshot  string `xml:"Snapshot,omitempty" json:"Snapshot"`
	Index     int    `xml:"Index,omitempty" json:"Index"`
	Guid      string `xml:"Guid,attr" json:"Guid"`
}
type BimSnippet struct {
	Reference       string `xml:"Reference" json:"Reference"`
	ReferenceSchema string `xml:"ReferenceSchema" json:"ReferenceSchema"`
	SnippetType     string `xml:"SnippetType,attr" json:"SnippetType"`
	isExternal      bool   `xml:"isExternal,attr" json:"isExternal"`
}

type Topic struct {
	XMLName            xml.Name            `xml:"Topic"`
	ReferenceLinks     []string            `xml:"ReferenceLinks>ReferenceLink,omitempty" json:"ReferenceLinks"`
	Title              string              `xml:"Title" json:"Title"`
	Priority           string              `xml:"Priority,omitempty" json:"Priority"`
	Index              string              `xml:"Index,omitempty" json:"Index"`
	Labels             []string            `xml:"Labels>Label,omitempty" json:"Labels"`
	CreationDate       string              `xml:"CreationDate" json:"CreationDate"`
	CreationAuthor     string              `xml:"CreationAuthor" json:"CreationAuthor"`
	ModifiedDate       string              `xml:"ModifiedDate,omitempty" json:"ModifiedDate"`
	ModifiedAuthor     string              `xml:"ModifiedAuthor,omitempty" json:"ModifiedAuthor"`
	DueDate            string              `xml:"DueDate,omitempty" json:"DueDate"`
	AssignedTo         string              `xml:"AssignedTo,omitempty" json:"AssignedTo"`
	Stage              string              `xml:"Stage,omitempty" json:"Stage"`
	Description        string              `xml:"Description,omitempty" json:"Description"`
	BimSnippet         BimSnippet          `xml:"BimSnippet,omitempty" json:"BimSnippet"`
	DocumentReferences []DocumentReference `xml:"DocumentReferences>DocumentReference,omitempty" json:"DocumentReferences"`
	RelatedTopics      []RelatedTopic      `xml:"RelatedTopics>RelatedTopic,omitempty" json:"RelatedTopics"`
	Comments           []Comment           `xml:"Comments>Comment,omitempty" json:"Comments"`
	ViewPoints         []ViewPoint         `xml:"Viewpoints>ViewPoint,omitempty" json:"ViewPoints"`
	Guid               string              `xml:"Guid,attr" json:"Guid"`
	ServerAssignedId   string              `xml:"ServerAssignedId,attr,omitempty" json:"ServerAssignedId"`
	TopicType          string              `xml:"TopicType,attr" json:"TopicType"`
	TopicStatus        string              `xml:"TopicStatus,attr" json:"TopicStatus"`
}

type File struct {
	Filename                   string `xml:"Filename,omitempty" json:"Filename"`
	Date                       string `xml:"Date,omitempty" json:"Date"`
	Reference                  string `xml:"Reference" json:"Reference"`
	IfcProject                 string `xml:"IfcProject,attr,omitempty" json:"IfcProject"`
	IfcSpatialStructureElement string `xml:"IfcSpatialStructureElement,attr,omitempty" json:"IfcSpatialStructureElement"`
	IsExternal                 bool `xml:"IsExternal,attr,omitempty" json:"IsExternal"`
}

type RelatedTopic struct {
	Guid string `xml:"Guid,attr" json:"Guid"`
}

type DocumentReference struct {
	DocumentGuid string `xml:"DocumentGuid,omitempty" json:"DocumentGuid"`
	Url          string `xml:"Url,omitempty" json:"Url"`
	Description  string `xml:"Description,omitempty" json:"description"`
	Guid         string `xml:"Guid,attr" json:"guid"`
}

type Comment struct {
	Date                  string           `xml:"Date" json:"Date"`
	Author                string           `xml:"Author" json:"Author"`
	Comment               string           `xml:"Comment,omitempty" json:"Comment"`
	Viewpoint             CommentViewpoint `xml:"Viewpoint,omitempty" json:"Viewpoint"`
	ModifiedDate          string           `xml:"ModifiedDate" json:"ModifiedDate"`
	ModifiedDateSpecified bool             `xml:"ModifiedDateSpecified" json:"ModifiedDateSpecified"`
	ModifiedAuthor        string           `xml:"ModifiedAuthor" json:"ModifiedAuthor"`
	Guid                  string           `xml:"Guid,attr" json:"Guid"`
}

type CommentViewpoint struct {
	Guid string `xml:"Guid,attr" json:"Guid"`
}

//**************************************visinfo*********************************//
type OrthogonalCamera struct {
	CameraViewPoint  Point     `xml:"CameraViewPoint" json:"CameraViewPoint"`
	CameraDirection  Direction `xml:"CameraDirection" json:"CameraDirection"`
	CameraUpVector   Direction `xml:"CameraUpVector" json:"CameraUpVector"`
	ViewToWorldScale float64   `xml:"ViewToWorldScale" json:"ViewToWorldScale"`
	AspectRatio      float64   `xml:"AspectRatio" json:"AspectRatio"`
}

type PerspectiveCamera struct {
	CameraViewPoint Point     `xml:"CameraViewPoint" json:"CameraViewPoint"`
	CameraDirection Direction `xml:"CameraDirection" json:"CameraDirection"`
	CameraUpVector  Direction `xml:"CameraUpVector" json:"CameraUpVector"`
	FieldOfView     float64   `xml:"FieldOfView" json:"FieldOfView"`
	AspectRatio     float64   `xml:"AspectRatio" json:"AspectRatio"`
}

type Point struct {
	X float64 `json:"X" xml:"X"`
	Y float64 `json:"Y" xml:"Y"`
	Z float64 `json:"Z" xml:"Z"`
}

type Direction struct {
	X float64 `json:"X" xml:"X"`
	Y float64 `json:"Y" xml:"Y"`
	Z float64 `json:"Z" xml:"Z"`
}

type Components struct {
	Selection  ComponentSelection  `xml:"Selection,omitempty" json:"Selection"`
	Visibility ComponentVisibility `xml:"Visibility,omitempty" json:"Visibility"`
	Coloring   ComponentColoring   `xml:"Coloring>Color,omitempty" json:"Coloring"`
}

type ComponentSelection struct {
	Component []Component `xml:"Component,omitempty" json:"Component"`
}
type ComponentVisibility struct {
	ViewSetupHints ViewSetupHints `xml:"ViewSetupHints,omitempty" json:"ViewSetupHints"`
	Exceptions      []Component      `xml:"Exceptions>Component,omitempty" json:"Exceptions"`
	DefaultVisibility      bool      `xml:"DefaultVisibility" json:"DefaultVisibility"`
}

type ViewSetupHints struct {
	SpacesVisible          bool `xml:"SpacesVisible,attr" json:"SpacesVisible"`
	SpaceBoundariesVisible bool `xml:"SpaceBoundariesVisible,attr" json:"SpaceBoundariesVisible"`
	OpeningsVisible        bool `xml:"OpeningsVisible,attr" json:"OpeningsVisible"`
}
type ComponentColoring struct {
	Components []Component `xml:"Components>Component" json:"Components"`
	Color      string      `xml:"Color,attr" json:"Color"`
}

type Component struct {
	OriginatingSystem string `xml:"OriginatingSystem,omitempty" json:"OriginatingSystem"`
	AuthoringToolId   string `xml:"AuthoringToolId,omitempty" json:"AuthoringToolId"`
	IfcGuid   string `xml:"IfcGuid" json:"IfcGuid"`
}

type Line struct {
	StartPoint Point `xml:"StartPoint" json:"StartPoint"`
	EndPoint   Point `xml:"EndPoint" json:"EndPoint"`
}

type ClippingPlane struct {
	Location  Point     `xml:"Location" json:"Location"`
	Direction Direction `xml:"Direction" json:"Direction"`
}

type Bitmap struct {
	Format    string    `xml:"Format" json:"Format"`
	Reference string    `xml:"Reference" json:"Reference"`
	Location  Point     `xml:"Location" json:"Location"`
	Normal    Direction `xml:"Normal" json:"Normal"`
	Up        Direction `xml:"Up" json:"Up"`
	Height    string    `xml:"Height" json:"Height"`
}
