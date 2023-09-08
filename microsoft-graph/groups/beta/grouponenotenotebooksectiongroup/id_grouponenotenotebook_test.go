package grouponenotenotebooksectiongroup

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupOnenoteNotebookId{}

func TestNewGroupOnenoteNotebookID(t *testing.T) {
	id := NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue")

	if id.GroupId != "groupIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'GroupId'", id.GroupId, "groupIdValue")
	}

	if id.NotebookId != "notebookIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'NotebookId'", id.NotebookId, "notebookIdValue")
	}
}

func TestFormatGroupOnenoteNotebookID(t *testing.T) {
	actual := NewGroupOnenoteNotebookID("groupIdValue", "notebookIdValue").ID()
	expected := "/groups/groupIdValue/onenote/notebooks/notebookIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseGroupOnenoteNotebookID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupOnenoteNotebookId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/notebooks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/onenote/notebooks/notebookIdValue",
			Expected: &GroupOnenoteNotebookId{
				GroupId:    "groupIdValue",
				NotebookId: "notebookIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/onenote/notebooks/notebookIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupOnenoteNotebookID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

	}
}

func TestParseGroupOnenoteNotebookIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *GroupOnenoteNotebookId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/groups/groupIdValue/onenote/notebooks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/nOtEbOoKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/groups/groupIdValue/onenote/notebooks/notebookIdValue",
			Expected: &GroupOnenoteNotebookId{
				GroupId:    "groupIdValue",
				NotebookId: "notebookIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/groups/groupIdValue/onenote/notebooks/notebookIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe",
			Expected: &GroupOnenoteNotebookId{
				GroupId:    "gRoUpIdVaLuE",
				NotebookId: "nOtEbOoKiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/gRoUpS/gRoUpIdVaLuE/oNeNoTe/nOtEbOoKs/nOtEbOoKiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseGroupOnenoteNotebookIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.GroupId != v.Expected.GroupId {
			t.Fatalf("Expected %q but got %q for GroupId", v.Expected.GroupId, actual.GroupId)
		}

		if actual.NotebookId != v.Expected.NotebookId {
			t.Fatalf("Expected %q but got %q for NotebookId", v.Expected.NotebookId, actual.NotebookId)
		}

	}
}

func TestSegmentsForGroupOnenoteNotebookId(t *testing.T) {
	segments := GroupOnenoteNotebookId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("GroupOnenoteNotebookId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
