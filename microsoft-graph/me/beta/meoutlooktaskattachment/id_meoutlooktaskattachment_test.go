package meoutlooktaskattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeOutlookTaskAttachmentId{}

func TestNewMeOutlookTaskAttachmentID(t *testing.T) {
	id := NewMeOutlookTaskAttachmentID("outlookTaskIdValue", "attachmentIdValue")

	if id.OutlookTaskId != "outlookTaskIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskIdValue")
	}

	if id.AttachmentId != "attachmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentIdValue")
	}
}

func TestFormatMeOutlookTaskAttachmentID(t *testing.T) {
	actual := NewMeOutlookTaskAttachmentID("outlookTaskIdValue", "attachmentIdValue").ID()
	expected := "/me/outlook/tasks/outlookTaskIdValue/attachments/attachmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeOutlookTaskAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskAttachmentId
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
			// Incomplete URI
			Input: "/me/outlook/tasks/outlookTaskIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/tasks/outlookTaskIdValue/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/tasks/outlookTaskIdValue/attachments/attachmentIdValue",
			Expected: &MeOutlookTaskAttachmentId{
				OutlookTaskId: "outlookTaskIdValue",
				AttachmentId:  "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/tasks/outlookTaskIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskAttachmentID(v.Input)
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseMeOutlookTaskAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeOutlookTaskAttachmentId
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
			// Incomplete URI
			Input: "/me/outlook/tasks/outlookTaskIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKs/oUtLoOkTaSkIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/outlook/tasks/outlookTaskIdValue/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKs/oUtLoOkTaSkIdVaLuE/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/outlook/tasks/outlookTaskIdValue/attachments/attachmentIdValue",
			Expected: &MeOutlookTaskAttachmentId{
				OutlookTaskId: "outlookTaskIdValue",
				AttachmentId:  "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/outlook/tasks/outlookTaskIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKs/oUtLoOkTaSkIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe",
			Expected: &MeOutlookTaskAttachmentId{
				OutlookTaskId: "oUtLoOkTaSkIdVaLuE",
				AttachmentId:  "aTtAcHmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/oUtLoOk/tAsKs/oUtLoOkTaSkIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeOutlookTaskAttachmentIDInsensitively(v.Input)
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

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForMeOutlookTaskAttachmentId(t *testing.T) {
	segments := MeOutlookTaskAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeOutlookTaskAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
