package stable

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdMessageId{}

func TestNewUserIdMailFolderIdChildFolderIdMessageID(t *testing.T) {
	id := NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.MailFolderId != "mailFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderId")
	}

	if id.MailFolderId1 != "mailFolderId1" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId1'", id.MailFolderId1, "mailFolderId1")
	}

	if id.MessageId != "messageId" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageId")
	}
}

func TestFormatUserIdMailFolderIdChildFolderIdMessageID(t *testing.T) {
	actual := NewUserIdMailFolderIdChildFolderIdMessageID("userId", "mailFolderId", "mailFolderId1", "messageId").ID()
	expected := "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/messages/messageId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMailFolderIdChildFolderIdMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdChildFolderIdMessageId
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
			Input: "/users/userId/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/messages",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/messages/messageId",
			Expected: &UserIdMailFolderIdChildFolderIdMessageId{
				UserId:        "userId",
				MailFolderId:  "mailFolderId",
				MailFolderId1: "mailFolderId1",
				MessageId:     "messageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/messages/messageId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdChildFolderIdMessageID(v.Input)
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

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

	}
}

func TestParseUserIdMailFolderIdChildFolderIdMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdChildFolderIdMessageId
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
			Input: "/users/userId/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/mEsSaGeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/messages/messageId",
			Expected: &UserIdMailFolderIdChildFolderIdMessageId{
				UserId:        "userId",
				MailFolderId:  "mailFolderId",
				MailFolderId1: "mailFolderId1",
				MessageId:     "messageId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/childFolders/mailFolderId1/messages/messageId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/mEsSaGeS/mEsSaGeId",
			Expected: &UserIdMailFolderIdChildFolderIdMessageId{
				UserId:        "uSeRiD",
				MailFolderId:  "mAiLfOlDeRiD",
				MailFolderId1: "mAiLfOlDeRiD1",
				MessageId:     "mEsSaGeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/cHiLdFoLdErS/mAiLfOlDeRiD1/mEsSaGeS/mEsSaGeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdChildFolderIdMessageIDInsensitively(v.Input)
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

		if actual.MailFolderId != v.Expected.MailFolderId {
			t.Fatalf("Expected %q but got %q for MailFolderId", v.Expected.MailFolderId, actual.MailFolderId)
		}

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

	}
}

func TestSegmentsForUserIdMailFolderIdChildFolderIdMessageId(t *testing.T) {
	segments := UserIdMailFolderIdChildFolderIdMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMailFolderIdChildFolderIdMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
