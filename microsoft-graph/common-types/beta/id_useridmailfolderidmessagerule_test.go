package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMailFolderIdMessageRuleId{}

func TestNewUserIdMailFolderIdMessageRuleID(t *testing.T) {
	id := NewUserIdMailFolderIdMessageRuleID("userId", "mailFolderId", "messageRuleId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.MailFolderId != "mailFolderId" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderId")
	}

	if id.MessageRuleId != "messageRuleId" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageRuleId'", id.MessageRuleId, "messageRuleId")
	}
}

func TestFormatUserIdMailFolderIdMessageRuleID(t *testing.T) {
	actual := NewUserIdMailFolderIdMessageRuleID("userId", "mailFolderId", "messageRuleId").ID()
	expected := "/users/userId/mailFolders/mailFolderId/messageRules/messageRuleId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMailFolderIdMessageRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdMessageRuleId
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
			Input: "/users/userId/mailFolders/mailFolderId/messageRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/messageRules/messageRuleId",
			Expected: &UserIdMailFolderIdMessageRuleId{
				UserId:        "userId",
				MailFolderId:  "mailFolderId",
				MessageRuleId: "messageRuleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/messageRules/messageRuleId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdMessageRuleID(v.Input)
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

func TestParseUserIdMailFolderIdMessageRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMailFolderIdMessageRuleId
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
			Input: "/users/userId/mailFolders/mailFolderId/messageRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/mailFolders/mailFolderId/messageRules/messageRuleId",
			Expected: &UserIdMailFolderIdMessageRuleId{
				UserId:        "userId",
				MailFolderId:  "mailFolderId",
				MessageRuleId: "messageRuleId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/mailFolders/mailFolderId/messageRules/messageRuleId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeRuLeS/mEsSaGeRuLeId",
			Expected: &UserIdMailFolderIdMessageRuleId{
				UserId:        "uSeRiD",
				MailFolderId:  "mAiLfOlDeRiD",
				MessageRuleId: "mEsSaGeRuLeId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mAiLfOlDeRs/mAiLfOlDeRiD/mEsSaGeRuLeS/mEsSaGeRuLeId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMailFolderIdMessageRuleIDInsensitively(v.Input)
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

func TestSegmentsForUserIdMailFolderIdMessageRuleId(t *testing.T) {
	segments := UserIdMailFolderIdMessageRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMailFolderIdMessageRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
