package memailfoldermessagerule

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderMessageRuleId{}

func TestNewMeMailFolderMessageRuleID(t *testing.T) {
	id := NewMeMailFolderMessageRuleID("mailFolderIdValue", "messageRuleIdValue")

	if id.MailFolderId != "mailFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderIdValue")
	}

	if id.MessageRuleId != "messageRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageRuleId'", id.MessageRuleId, "messageRuleIdValue")
	}
}

func TestFormatMeMailFolderMessageRuleID(t *testing.T) {
	actual := NewMeMailFolderMessageRuleID("mailFolderIdValue", "messageRuleIdValue").ID()
	expected := "/me/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMailFolderMessageRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderMessageRuleId
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
			Input: "/me/mailFolders/mailFolderIdValue/messageRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue",
			Expected: &MeMailFolderMessageRuleId{
				MailFolderId:  "mailFolderIdValue",
				MessageRuleId: "messageRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderMessageRuleID(v.Input)
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

		if actual.MessageRuleId != v.Expected.MessageRuleId {
			t.Fatalf("Expected %q but got %q for MessageRuleId", v.Expected.MessageRuleId, actual.MessageRuleId)
		}

	}
}

func TestParseMeMailFolderMessageRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderMessageRuleId
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
			Input: "/me/mailFolders/mailFolderIdValue/messageRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue",
			Expected: &MeMailFolderMessageRuleId{
				MailFolderId:  "mailFolderIdValue",
				MessageRuleId: "messageRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderIdValue/messageRules/messageRuleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeRuLeS/mEsSaGeRuLeIdVaLuE",
			Expected: &MeMailFolderMessageRuleId{
				MailFolderId:  "mAiLfOlDeRiDvAlUe",
				MessageRuleId: "mEsSaGeRuLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/mEsSaGeRuLeS/mEsSaGeRuLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderMessageRuleIDInsensitively(v.Input)
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

		if actual.MessageRuleId != v.Expected.MessageRuleId {
			t.Fatalf("Expected %q but got %q for MessageRuleId", v.Expected.MessageRuleId, actual.MessageRuleId)
		}

	}
}

func TestSegmentsForMeMailFolderMessageRuleId(t *testing.T) {
	segments := MeMailFolderMessageRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMailFolderMessageRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
