package meoutlooktaskattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskId{}

func TestNewMeOutlookTaskID(t *testing.T) {
	id := NewMeOutlookTaskID("outlookTaskIdValue")

	if id.OutlookTaskId != "outlookTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskIdValue")
	}
}

func TestFormatMeOutlookTaskID(t *testing.T) {
	actual := NewMeOutlookTaskID("outlookTaskIdValue").ID()
	expected := "/me/outlook/tasks/outlookTaskIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOutlookTaskID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/tasks",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/tasks/outlookTaskIdValue",
			Expected: &MeOutlookTaskId{
				OutlookTaskId: "outlookTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/tasks/outlookTaskIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OutlookTaskId != v.Expected.OutlookTaskId {
			t.Fatalf("Expected %q but got %q for OutlookTaskId", v.Expected.OutlookTaskId, actual.OutlookTaskId)
		}

	}
}

func TestParseMeOutlookTaskIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/tasks/outlookTaskIdValue",
			Expected: &MeOutlookTaskId{
				OutlookTaskId: "outlookTaskIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/tasks/outlookTaskIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKs/oUtLoOkTaSkIdVaLuE",
			Expected: &MeOutlookTaskId{
				OutlookTaskId: "oUtLoOkTaSkIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKs/oUtLoOkTaSkIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.OutlookTaskId != v.Expected.OutlookTaskId {
			t.Fatalf("Expected %q but got %q for OutlookTaskId", v.Expected.OutlookTaskId, actual.OutlookTaskId)
		}

	}
}

func TestSegmentsForMeOutlookTaskId(t *testing.T) {
	segments := MeOutlookTaskId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOutlookTaskId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
