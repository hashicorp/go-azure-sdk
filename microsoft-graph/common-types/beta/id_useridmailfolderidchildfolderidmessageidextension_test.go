package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdChildFolderIdMessageIdExtensionId{}

func TestNewUserIdMailFolderIdChildFolderIdMessageIdExtensionID(t *testing.T) {
	id := NewUserIdMailFolderIdChildFolderIdMessageIdExtensionID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue", "extensionIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.MailFolderId != "mailFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderIdValue")
	}

	if id.MailFolderId1 != "mailFolderId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId1'", id.MailFolderId1, "mailFolderId1Value")
	}

	if id.MessageId != "messageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageIdValue")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatUserIdMailFolderIdChildFolderIdMessageIdExtensionID(t *testing.T) {
	actual := NewUserIdMailFolderIdChildFolderIdMessageIdExtensionID("userIdValue", "mailFolderIdValue", "mailFolderId1Value", "messageIdValue", "extensionIdValue").ID()
	expected := "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMailFolderIdChildFolderIdMessageIdExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdChildFolderIdMessageIdExtensionId
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
			Input: "/users/userIdValue/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue/extensions/extensionIdValue",
			Expected: &UserIdMailFolderIdChildFolderIdMessageIdExtensionId{
				UserId:        "userIdValue",
				MailFolderId:  "mailFolderIdValue",
				MailFolderId1: "mailFolderId1Value",
				MessageId:     "messageIdValue",
				ExtensionId:   "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdChildFolderIdMessageIdExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseUserIdMailFolderIdChildFolderIdMessageIdExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdChildFolderIdMessageIdExtensionId
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
			Input: "/users/userIdValue/mailFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/mEsSaGeS/mEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/mEsSaGeS/mEsSaGeIdVaLuE/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue/extensions/extensionIdValue",
			Expected: &UserIdMailFolderIdChildFolderIdMessageIdExtensionId{
				UserId:        "userIdValue",
				MailFolderId:  "mailFolderIdValue",
				MailFolderId1: "mailFolderId1Value",
				MessageId:     "messageIdValue",
				ExtensionId:   "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messages/messageIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/mEsSaGeS/mEsSaGeIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &UserIdMailFolderIdChildFolderIdMessageIdExtensionId{
				UserId:        "uSeRiDvAlUe",
				MailFolderId:  "mAiLfOlDeRiDvAlUe",
				MailFolderId1: "mAiLfOlDeRiD1VaLuE",
				MessageId:     "mEsSaGeIdVaLuE",
				ExtensionId:   "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/mEsSaGeS/mEsSaGeIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdChildFolderIdMessageIdExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForUserIdMailFolderIdChildFolderIdMessageIdExtensionId(t *testing.T) {
	segments := UserIdMailFolderIdChildFolderIdMessageIdExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMailFolderIdChildFolderIdMessageIdExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
