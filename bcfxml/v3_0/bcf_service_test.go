package v3_0

import (
	"github.com/gogf/gf/v2/os/gfile"
	"testing"
)

func TestWriteReadFold(t *testing.T) {
	fold := gfile.Dir(gfile.Dir(".")) + "\\testdata\\v3_0\\read"
	data,err := InsBcfService().ReadFoldData(fold)
	t.Log("WriteFileProject err=", err)
	t.Log("data =", data)
}

func TestWriteFold(t *testing.T) {
	fold := gfile.Dir(gfile.Dir(".")) + "\\testdata\\v3_0\\write"
	in := &BcfData{
		ProjectInfo: &ProjectInfo{
			Project: Project{
				ProjectId: "de894a86-3a08-4ea0-b2d1-6c222b5602d1",
				Name:      "BCF 3.0 test cases",
			},
		},
		DocumentInfo: &DocumentInfo{
			Documents: []Document{{
				Guid:     "b1d1b7f0-60b9-457d-ad12-16e0fb997bc5",
				Filename: "ThisIsADocument.txt",
			}},
		},
		Extensions: &Extensions{
			TopicTypes:    []string{"ERROR", "WARNING", "INFORMATION", "CLASH", "OTHER"},
			TopicStatuses: []string{"OPEN", "IN_PROGRESS", "SOLVED", "CLOSED"},
			Priorities:    []string{"LOW", "MEDIUM", "HIGH", "CRITICAL"},
			Users:         []string{"Architect@example.com", "Engineer@example.com", "MEPDesigner@example.com"},
		},
		Markups: []*Markup{
			&Markup{
				Header: []File{{
					IfcProject:                 "0M6o7Znnv7hxsbWgeu7oQq",
					IfcSpatialStructureElement: "23B$bNeGHFQuMYJzvUX0FD",
					IsExternal:                 false,
					Filename:                   "BCF-ARK",
					Date:                       "2021-01-04T09:37:45.000Z",
					Reference:                  "../Architectural.ifc",
				}},
				Topic: Topic{
					Guid:"8ac9822a-761a-4deb-9f39-f61286acbf6a",
					ServerAssignedId:"6",
					TopicStatus:"OPEN",
					TopicType:"ERROR",
					Title:"Document Reference Internal",
					CreationDate:"2021-02-17T09:16:04.160Z",
					CreationAuthor:"Architect@example.com",
					ModifiedDate:"2021-02-17T09:16:04.160Z",
					AssignedTo:"OtherUser@doe.com",
					DocumentReferences:[]DocumentReference{{
						Guid: "048e898f-555f-47c3-a273-6c664fb2ef69",
						DocumentGuid: "b1d1b7f0-60b9-457d-ad12-16e0fb997bc5",
						Description: "ThisIsADocument.txt",
					}},
					Comments: []Comment{{
						Guid: "c045ef04-324d-4c36-9ac8-831f2a15c6d6",
						Date:"2021-02-17T09:16:04.160Z",
						Author: "Architect@example.com",
						Comment: "A comment",
						Viewpoint:CommentViewpoint{
							Guid: "20ad2ff7-ceac-4d4b-b288-ad92a0f65182",
						},
					}},
					ViewPoints: []ViewPoint{{
						Guid: "20ad2ff7-ceac-4d4b-b288-ad92a0f65182",
						Viewpoint:"Viewpoint_20ad2ff7-ceac-4d4b-b288-ad92a0f65182.bcfv",
						Snapshot: "Snapshot_20ad2ff7-ceac-4d4b-b288-ad92a0f65182.png",
					}},
				},
			},
		},
		Visinfos: map[string][]*VisualizationInfo{"8ac9822a-761a-4deb-9f39-f61286acbf6a":{&VisualizationInfo{
			Guid: "20ad2ff7-ceac-4d4b-b288-ad92a0f65182",
			Components: Components{
				Visibility: ComponentVisibility{
					DefaultVisibility: true,
					ViewSetupHints: ViewSetupHints{
						SpacesVisible: false,
						SpaceBoundariesVisible: false,
						OpeningsVisible: false,
					},
					Exceptions: []Component{{
						OriginatingSystem:"ARCHICAD",
						IfcGuid: "1m5wAJelDFdhn6qBdOGjos",
					}},
				},
			},
			PerspectiveCamera: PerspectiveCamera{
				CameraViewPoint: Point{
					X:23.112083480037754,
					Y:-25.45574897560043,
					Z:18.52519828443092,
				},
				CameraDirection: Direction{
					X:-0.6289237666876837,
					Y:0.5933487270610425,
					Z:-0.5023864884632315,
				},
				CameraUpVector: Direction{
					X:-0.36542566067664123,
					Y:0.3447553774917179,
					Z:0.8646431727652648,
				},
				FieldOfView: 60.0,
				AspectRatio: 1.0,
			},
		}}},
	}
	err := InsBcfService().WriteFoldData(fold, in)
	t.Log("WriteFileProject err:", err)
}
