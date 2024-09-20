package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMailFolderIdMessageIdMentionId{}

func TestNewMeMailFolderIdMessageIdMentionID(t *testing.T) {
	id := NewMeMailFolderIdMessageIdMentionID("mailFolderId", "messageId", "mentionId")

	if id.MailFolderId != "mailFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderId")
	}

	if id.MessageId != "messageId" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageId")
	}

	if id.MentionId != "mentionId" {
		t.Fatalf("Expected %q but got %q for Segment 'MentionId'", id.MentionId, "mentionId")
	}
}

func TestFormatMeMailFolderIdMessageIdMentionID(t *testing.T) {
	actual := NewMeMailFolderIdMessageIdMentionID("mailFolderId", "messageId", "mentionId").ID()
	expected := "/me/mailFolders/mailFolderId/messages/messageId/mentions/mentionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMailFolderIdMessageIdMentionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderIdMessageIdMentionId
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
			Input: "/me/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/messages/messageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/messages/messageId/mentions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderId/messages/messageId/mentions/mentionId",
			Expected: &MeMailFolderIdMessageIdMentionId{
				MailFolderId: "mailFolderId",
				MessageId:    "messageId",
				MentionId:    "mentionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderId/messages/messageId/mentions/mentionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderIdMessageIdMentionID(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestParseMeMailFolderIdMessageIdMentionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderIdMessageIdMentionId
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
			Input: "/me/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/messages/messageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS/mEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderId/messages/messageId/mentions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS/mEsSaGeId/mEnTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderId/messages/messageId/mentions/mentionId",
			Expected: &MeMailFolderIdMessageIdMentionId{
				MailFolderId: "mailFolderId",
				MessageId:    "messageId",
				MentionId:    "mentionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderId/messages/messageId/mentions/mentionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS/mEsSaGeId/mEnTiOnS/mEnTiOnId",
			Expected: &MeMailFolderIdMessageIdMentionId{
				MailFolderId: "mAiLfOlDeRiD",
				MessageId:    "mEsSaGeId",
				MentionId:    "mEnTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS/mEsSaGeId/mEnTiOnS/mEnTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderIdMessageIdMentionIDInsensitively(v.Input)
		if err != nil {
			if v.Error {
				continue
			}

			t.Fatalf("Expect a value but got an error: %+v", err)
		}
		if v.Error {
			t.Fatal("Expect an error but didn't get one")
		}

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestSegmentsForMeMailFolderIdMessageIdMentionId(t *testing.T) {
	segments := MeMailFolderIdMessageIdMentionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMailFolderIdMessageIdMentionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
