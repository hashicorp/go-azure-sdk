package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMessageIdAttachmentId{}

func TestNewMeMessageIdAttachmentID(t *testing.T) {
	id := NewMeMessageIdAttachmentID("messageId", "attachmentId")

	if id.MessageId != "messageId" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageId")
	}

	if id.AttachmentId != "attachmentId" {
		t.Fatalf("Expected %q but got %q for Segment 'AttachmentId'", id.AttachmentId, "attachmentId")
	}
}

func TestFormatMeMessageIdAttachmentID(t *testing.T) {
	actual := NewMeMessageIdAttachmentID("messageId", "attachmentId").ID()
	expected := "/me/messages/messageId/attachments/attachmentId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMessageIdAttachmentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMessageIdAttachmentId
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
			Input: "/me/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/messages/messageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/messages/messageId/attachments",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/messages/messageId/attachments/attachmentId",
			Expected: &MeMessageIdAttachmentId{
				MessageId:    "messageId",
				AttachmentId: "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/messages/messageId/attachments/attachmentId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMessageIdAttachmentID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestParseMeMessageIdAttachmentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMessageIdAttachmentId
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
			Input: "/me/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/messages/messageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/messages/messageId/attachments",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeId/aTtAcHmEnTs",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/messages/messageId/attachments/attachmentId",
			Expected: &MeMessageIdAttachmentId{
				MessageId:    "messageId",
				AttachmentId: "attachmentId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/messages/messageId/attachments/attachmentId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeId/aTtAcHmEnTs/aTtAcHmEnTiD",
			Expected: &MeMessageIdAttachmentId{
				MessageId:    "mEsSaGeId",
				AttachmentId: "aTtAcHmEnTiD",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeId/aTtAcHmEnTs/aTtAcHmEnTiD/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMessageIdAttachmentIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.AttachmentId != v.Expected.AttachmentId {
			t.Fatalf("Expected %q but got %q for AttachmentId", v.Expected.AttachmentId, actual.AttachmentId)
		}

	}
}

func TestSegmentsForMeMessageIdAttachmentId(t *testing.T) {
	segments := MeMessageIdAttachmentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMessageIdAttachmentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
