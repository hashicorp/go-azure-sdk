package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeMessageIdMentionId{}

func TestNewMeMessageIdMentionID(t *testing.T) {
	id := NewMeMessageIdMentionID("messageId", "mentionId")

	if id.MessageId != "messageId" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageId")
	}

	if id.MentionId != "mentionId" {
		t.Fatalf("Expected %q but got %q for Segment 'MentionId'", id.MentionId, "mentionId")
	}
}

func TestFormatMeMessageIdMentionID(t *testing.T) {
	actual := NewMeMessageIdMentionID("messageId", "mentionId").ID()
	expected := "/me/messages/messageId/mentions/mentionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMessageIdMentionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMessageIdMentionId
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
			Input: "/me/messages/messageId/mentions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/messages/messageId/mentions/mentionId",
			Expected: &MeMessageIdMentionId{
				MessageId: "messageId",
				MentionId: "mentionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/messages/messageId/mentions/mentionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMessageIdMentionID(v.Input)
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

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestParseMeMessageIdMentionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMessageIdMentionId
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
			Input: "/me/messages/messageId/mentions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeId/mEnTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/messages/messageId/mentions/mentionId",
			Expected: &MeMessageIdMentionId{
				MessageId: "messageId",
				MentionId: "mentionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/messages/messageId/mentions/mentionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeId/mEnTiOnS/mEnTiOnId",
			Expected: &MeMessageIdMentionId{
				MessageId: "mEsSaGeId",
				MentionId: "mEnTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeId/mEnTiOnS/mEnTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMessageIdMentionIDInsensitively(v.Input)
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

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestSegmentsForMeMessageIdMentionId(t *testing.T) {
	segments := MeMessageIdMentionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMessageIdMentionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
