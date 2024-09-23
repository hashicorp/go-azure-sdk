package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdMessageIdExtensionId{}

func TestNewUserIdMailFolderIdMessageIdExtensionID(t *testing.T) {
	id := NewUserIdMailFolderIdMessageIdExtensionID("userId", "mailFolderId", "messageId", "extensionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.MailFolderId != "mailFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderId")
	}

	if id.MessageId != "messageId" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageId")
	}

	if id.ExtensionId != "extensionId" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionId")
	}
}

func TestFormatUserIdMailFolderIdMessageIdExtensionID(t *testing.T) {
	actual := NewUserIdMailFolderIdMessageIdExtensionID("userId", "mailFolderId", "messageId", "extensionId").ID()
	expected := "/users/userId/mailFolders/mailFolderId/messages/messageId/extensions/extensionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMailFolderIdMessageIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdMessageIdExtensionId
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
			Input: "/users/userId/mailFolders/mailFolderId/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/messages/messageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/messages/messageId/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/messages/messageId/extensions/extensionId",
			Expected: &UserIdMailFolderIdMessageIdExtensionId{
				UserId:       "userId",
				MailFolderId: "mailFolderId",
				MessageId:    "messageId",
				ExtensionId:  "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/messages/messageId/extensions/extensionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdMessageIdExtensionID(v.Input)
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

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseUserIdMailFolderIdMessageIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdMessageIdExtensionId
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
			Input: "/users/userId/mailFolders/mailFolderId/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/messages/messageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS/mEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/mailFolders/mailFolderId/messages/messageId/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS/mEsSaGeId/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/messages/messageId/extensions/extensionId",
			Expected: &UserIdMailFolderIdMessageIdExtensionId{
				UserId:       "userId",
				MailFolderId: "mailFolderId",
				MessageId:    "messageId",
				ExtensionId:  "extensionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/messages/messageId/extensions/extensionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS/mEsSaGeId/eXtEnSiOnS/eXtEnSiOnId",
			Expected: &UserIdMailFolderIdMessageIdExtensionId{
				UserId:       "uSeRiD",
				MailFolderId: "mAiLfOlDeRiD",
				MessageId:    "mEsSaGeId",
				ExtensionId:  "eXtEnSiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeS/mEsSaGeId/eXtEnSiOnS/eXtEnSiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdMessageIdExtensionIDInsensitively(v.Input)
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

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForUserIdMailFolderIdMessageIdExtensionId(t *testing.T) {
	segments := UserIdMailFolderIdMessageIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMailFolderIdMessageIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
