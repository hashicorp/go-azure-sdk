package memessagemention

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMessageMentionId{}

func TestNewMeMessageMentionID(t *testing.T) {
	id := NewMeMessageMentionID("messageIdValue", "mentionIdValue")

	if id.MessageId != "messageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageIdValue")
	}

	if id.MentionId != "mentionIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MentionId'", id.MentionId, "mentionIdValue")
	}
}

func TestFormatMeMessageMentionID(t *testing.T) {
	actual := NewMeMessageMentionID("messageIdValue", "mentionIdValue").ID()
	expected := "/me/messages/messageIdValue/mentions/mentionIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMessageMentionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMessageMentionId
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
			Input: "/me/messages/messageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/messages/messageIdValue/mentions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/messages/messageIdValue/mentions/mentionIdValue",
			Expected: &MeMessageMentionId{
				MessageId: "messageIdValue",
				MentionId: "mentionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/messages/messageIdValue/mentions/mentionIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMessageMentionID(v.Input)
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

func TestParseMeMessageMentionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMessageMentionId
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
			Input: "/me/messages/messageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/me/messages/messageIdValue/mentions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeIdVaLuE/mEnTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/me/messages/messageIdValue/mentions/mentionIdValue",
			Expected: &MeMessageMentionId{
				MessageId: "messageIdValue",
				MentionId: "mentionIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/messages/messageIdValue/mentions/mentionIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeIdVaLuE/mEnTiOnS/mEnTiOnIdVaLuE",
			Expected: &MeMessageMentionId{
				MessageId: "mEsSaGeIdVaLuE",
				MentionId: "mEnTiOnIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeIdVaLuE/mEnTiOnS/mEnTiOnIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMessageMentionIDInsensitively(v.Input)
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

func TestSegmentsForMeMessageMentionId(t *testing.T) {
	segments := MeMessageMentionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMessageMentionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
