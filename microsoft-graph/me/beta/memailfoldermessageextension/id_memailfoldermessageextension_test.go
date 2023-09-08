package memailfoldermessageextension

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderMessageExtensionId{}

func TestNewMeMailFolderMessageExtensionID(t *testing.T) {
	id := NewMeMailFolderMessageExtensionID("mailFolderIdValue", "messageIdValue", "extensionIdValue")

	if id.MailFolderId != "mailFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderIdValue")
	}

	if id.MessageId != "messageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageIdValue")
	}

	if id.ExtensionId != "extensionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ExtensionId'", id.ExtensionId, "extensionIdValue")
	}
}

func TestFormatMeMailFolderMessageExtensionID(t *testing.T) {
	actual := NewMeMailFolderMessageExtensionID("mailFolderIdValue", "messageIdValue", "extensionIdValue").ID()
	expected := "/me/mailFolders/mailFolderIdValue/messages/messageIdValue/extensions/extensionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMailFolderMessageExtensionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderMessageExtensionId
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
			Input: "/me/mailFolders/mailFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/messages/messageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/messages/messageIdValue/extensions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderIdValue/messages/messageIdValue/extensions/extensionIdValue",
			Expected: &MeMailFolderMessageExtensionId{
				MailFolderId: "mailFolderIdValue",
				MessageId:    "messageIdValue",
				ExtensionId:  "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderIdValue/messages/messageIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderMessageExtensionID(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestParseMeMailFolderMessageExtensionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderMessageExtensionId
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
			Input: "/me/mailFolders/mailFolderIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/messages/messageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeS/mEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/messages/messageIdValue/extensions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeS/mEsSaGeIdVaLuE/eXtEnSiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderIdValue/messages/messageIdValue/extensions/extensionIdValue",
			Expected: &MeMailFolderMessageExtensionId{
				MailFolderId: "mailFolderIdValue",
				MessageId:    "messageIdValue",
				ExtensionId:  "extensionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderIdValue/messages/messageIdValue/extensions/extensionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeS/mEsSaGeIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE",
			Expected: &MeMailFolderMessageExtensionId{
				MailFolderId: "mAiLfOlDeRiDvAlUe",
				MessageId:    "mEsSaGeIdVaLuE",
				ExtensionId:  "eXtEnSiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeS/mEsSaGeIdVaLuE/eXtEnSiOnS/eXtEnSiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderMessageExtensionIDInsensitively(v.Input)
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

		if actual.ExtensionId != v.Expected.ExtensionId {
			t.Fatalf("Expected %q but got %q for ExtensionId", v.Expected.ExtensionId, actual.ExtensionId)
		}

	}
}

func TestSegmentsForMeMailFolderMessageExtensionId(t *testing.T) {
	segments := MeMailFolderMessageExtensionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMailFolderMessageExtensionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
