package usermailfoldermessagerule

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMailFolderMessageRuleId{}

func TestNewUserMailFolderMessageRuleID(t *testing.T) {
	id := NewUserMailFolderMessageRuleID("userIdValue", "mailFolderIdValue", "messageRuleIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.MailFolderId != "mailFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderIdValue")
	}

	if id.MessageRuleId != "messageRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageRuleId'", id.MessageRuleId, "messageRuleIdValue")
	}
}

func TestFormatUserMailFolderMessageRuleID(t *testing.T) {
	actual := NewUserMailFolderMessageRuleID("userIdValue", "mailFolderIdValue", "messageRuleIdValue").ID()
	expected := "/users/userIdValue/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserMailFolderMessageRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMailFolderMessageRuleId
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
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/messageRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue",
			Expected: &UserMailFolderMessageRuleId{
				UserId:        "userIdValue",
				MailFolderId:  "mailFolderIdValue",
				MessageRuleId: "messageRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMailFolderMessageRuleID(v.Input)
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

		if actual.MessageRuleId != v.Expected.MessageRuleId {
			t.Fatalf("Expected %q but got %q for MessageRuleId", v.Expected.MessageRuleId, actual.MessageRuleId)
		}

	}
}

func TestParseUserMailFolderMessageRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserMailFolderMessageRuleId
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
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/messageRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue",
			Expected: &UserMailFolderMessageRuleId{
				UserId:        "userIdValue",
				MailFolderId:  "mailFolderIdValue",
				MessageRuleId: "messageRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeRuLeS/mEsSaGeRuLeIdVaLuE",
			Expected: &UserMailFolderMessageRuleId{
				UserId:        "uSeRiDvAlUe",
				MailFolderId:  "mAiLfOlDeRiDvAlUe",
				MessageRuleId: "mEsSaGeRuLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeRuLeS/mEsSaGeRuLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserMailFolderMessageRuleIDInsensitively(v.Input)
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

		if actual.MessageRuleId != v.Expected.MessageRuleId {
			t.Fatalf("Expected %q but got %q for MessageRuleId", v.Expected.MessageRuleId, actual.MessageRuleId)
		}

	}
}

func TestSegmentsForUserMailFolderMessageRuleId(t *testing.T) {
	segments := UserMailFolderMessageRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserMailFolderMessageRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
