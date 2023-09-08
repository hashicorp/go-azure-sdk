package memailfolderchildfoldermessagerule

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMailFolderChildFolderMessageRuleId{}

func TestNewMeMailFolderChildFolderMessageRuleID(t *testing.T) {
	id := NewMeMailFolderChildFolderMessageRuleID("mailFolderIdValue", "mailFolderId1Value", "messageRuleIdValue")

	if id.MailFolderId != "mailFolderIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId'", id.MailFolderId, "mailFolderIdValue")
	}

	if id.MailFolderId1 != "mailFolderId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'MailFolderId1'", id.MailFolderId1, "mailFolderId1Value")
	}

	if id.MessageRuleId != "messageRuleIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageRuleId'", id.MessageRuleId, "messageRuleIdValue")
	}
}

func TestFormatMeMailFolderChildFolderMessageRuleID(t *testing.T) {
	actual := NewMeMailFolderChildFolderMessageRuleID("mailFolderIdValue", "mailFolderId1Value", "messageRuleIdValue").ID()
	expected := "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messageRules/messageRuleIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMailFolderChildFolderMessageRuleID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderChildFolderMessageRuleId
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
			Input: "/me/mailFolders/mailFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messageRules",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messageRules/messageRuleIdValue",
			Expected: &MeMailFolderChildFolderMessageRuleId{
				MailFolderId:  "mailFolderIdValue",
				MailFolderId1: "mailFolderId1Value",
				MessageRuleId: "messageRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messageRules/messageRuleIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderChildFolderMessageRuleID(v.Input)
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

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.MessageRuleId != v.Expected.MessageRuleId {
			t.Fatalf("Expected %q but got %q for MessageRuleId", v.Expected.MessageRuleId, actual.MessageRuleId)
		}

	}
}

func TestParseMeMailFolderChildFolderMessageRuleIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMailFolderChildFolderMessageRuleId
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
			Input: "/me/mailFolders/mailFolderIdValue/childFolders",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messageRules",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/mEsSaGeRuLeS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messageRules/messageRuleIdValue",
			Expected: &MeMailFolderChildFolderMessageRuleId{
				MailFolderId:  "mailFolderIdValue",
				MailFolderId1: "mailFolderId1Value",
				MessageRuleId: "messageRuleIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/mailFolders/mailFolderIdValue/childFolders/mailFolderId1Value/messageRules/messageRuleIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/mEsSaGeRuLeS/mEsSaGeRuLeIdVaLuE",
			Expected: &MeMailFolderChildFolderMessageRuleId{
				MailFolderId:  "mAiLfOlDeRiDvAlUe",
				MailFolderId1: "mAiLfOlDeRiD1VaLuE",
				MessageRuleId: "mEsSaGeRuLeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mAiLfOlDeRs/mAiLfOlDeRiDvAlUe/cHiLdFoLdErS/mAiLfOlDeRiD1VaLuE/mEsSaGeRuLeS/mEsSaGeRuLeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMailFolderChildFolderMessageRuleIDInsensitively(v.Input)
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

		if actual.MailFolderId1 != v.Expected.MailFolderId1 {
			t.Fatalf("Expected %q but got %q for MailFolderId1", v.Expected.MailFolderId1, actual.MailFolderId1)
		}

		if actual.MessageRuleId != v.Expected.MessageRuleId {
			t.Fatalf("Expected %q but got %q for MessageRuleId", v.Expected.MessageRuleId, actual.MessageRuleId)
		}

	}
}

func TestSegmentsForMeMailFolderChildFolderMessageRuleId(t *testing.T) {
	segments := MeMailFolderChildFolderMessageRuleId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMailFolderChildFolderMessageRuleId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
