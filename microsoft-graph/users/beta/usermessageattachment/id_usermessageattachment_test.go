package usermessageattachment

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMessageAttachmentId{}

func TestNewUserMessageAttachmentID(t *testing.T) {
	id := NewUserMessageAttachmentID("userIdValue", "messageIdValue", "attachmentIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.MessageId != "messageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageIdValue")
	}

	if id.AttachmentId != "attachmentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentIdValue")
	}
}

func TestFormatUserMessageAttachmentID(t *testing.T) {
	actual := NewUserMessageAttachmentID("userIdValue", "messageIdValue", "attachmentIdValue").ID()
	expected := "/users/userIdValue/messages/messageIdValue/attachments/attachmentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserMessageAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMessageAttachmentId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/messages/messageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/messages/messageIdValue/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/messages/messageIdValue/attachments/attachmentIdValue",
			Expected: &UserMessageAttachmentId{
				UserId:       "userIdValue",
				MessageId:    "messageIdValue",
				AttachmentId: "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/messages/messageIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMessageAttachmentID(v.Input)
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

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseUserMessageAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMessageAttachmentId
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
			Input: "/users/userIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/messages/messageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mEsSaGeS/mEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/messages/messageIdValue/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mEsSaGeS/mEsSaGeIdVaLuE/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/messages/messageIdValue/attachments/attachmentIdValue",
			Expected: &UserMessageAttachmentId{
				UserId:       "userIdValue",
				MessageId:    "messageIdValue",
				AttachmentId: "attachmentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/messages/messageIdValue/attachments/attachmentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mEsSaGeS/mEsSaGeIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe",
			Expected: &UserMessageAttachmentId{
				UserId:       "uSeRiDvAlUe",
				MessageId:    "mEsSaGeIdVaLuE",
				AttachmentId: "aTtAcHmEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mEsSaGeS/mEsSaGeIdVaLuE/aTtAcHmEnTs/aTtAcHmEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMessageAttachmentIDInsensitively(v.Input)
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

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForUserMessageAttachmentId(t *testing.T) {
	segments := UserMessageAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserMessageAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
