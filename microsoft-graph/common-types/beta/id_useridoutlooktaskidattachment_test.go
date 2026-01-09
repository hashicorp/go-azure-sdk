package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOutlookTaskIdAttachmentId{}

func TestNewUserIdOutlookTaskIdAttachmentID(t *testing.T) {
	id := NewUserIdOutlookTaskIdAttachmentID("userId", "outlookTaskId", "attachmentId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.OutlookTaskId != "outlookTaskId" {
		t.Fatalf("Expected %q but got %q for Segment 'OutlookTaskId'", id.OutlookTaskId, "outlookTaskId")
	}

	if id.AttachmentId != "attachmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentId")
	}
}

func TestFormatUserIdOutlookTaskIdAttachmentID(t *testing.T) {
	actual := NewUserIdOutlookTaskIdAttachmentID("userId", "outlookTaskId", "attachmentId").ID()
	expected := "/users/userId/outlook/tasks/outlookTaskId/attachments/attachmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdOutlookTaskIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookTaskIdAttachmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/tasks",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/tasks/outlookTaskId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/tasks/outlookTaskId/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/outlook/tasks/outlookTaskId/attachments/attachmentId",
			Expected: &UserIdOutlookTaskIdAttachmentId{
				UserId:        "userId",
				OutlookTaskId: "outlookTaskId",
				AttachmentId:  "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/outlook/tasks/outlookTaskId/attachments/attachmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookTaskIdAttachmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.OutlookTaskId != v.Expected.OutlookTaskId {
			t.Fatalf("Expected %q but got %q for OutlookTaskId", v.Expected.OutlookTaskId, actual.OutlookTaskId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseUserIdOutlookTaskIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdOutlookTaskIdAttachmentId
	}{
		{
			// Incomplete URI
			Input: "",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/tasks",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/tasks/outlookTaskId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKs/oUtLoOkTaSkId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/outlook/tasks/outlookTaskId/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/outlook/tasks/outlookTaskId/attachments/attachmentId",
			Expected: &UserIdOutlookTaskIdAttachmentId{
				UserId:        "userId",
				OutlookTaskId: "outlookTaskId",
				AttachmentId:  "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/outlook/tasks/outlookTaskId/attachments/attachmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs/aTtAcHmEnTiD",
			Expected: &UserIdOutlookTaskIdAttachmentId{
				UserId:        "uSeRiD",
				OutlookTaskId: "oUtLoOkTaSkId",
				AttachmentId:  "aTtAcHmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/oUtLoOk/tAsKs/oUtLoOkTaSkId/aTtAcHmEnTs/aTtAcHmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdOutlookTaskIdAttachmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.UserId != v.Expected.UserId {
			t.Fatalf("Expected %q but got %q for UserId", v.Expected.UserId, actual.UserId)
		}

		if actual.OutlookTaskId != v.Expected.OutlookTaskId {
			t.Fatalf("Expected %q but got %q for OutlookTaskId", v.Expected.OutlookTaskId, actual.OutlookTaskId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForUserIdOutlookTaskIdAttachmentId(t *testing.T) {
	segments := UserIdOutlookTaskIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdOutlookTaskIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
