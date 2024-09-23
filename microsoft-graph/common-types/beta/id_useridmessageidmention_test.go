package beta

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdMessageIdMentionId{}

func TestNewUserIdMessageIdMentionID(t *testing.T) {
	id := NewUserIdMessageIdMentionID("userId", "messageId", "mentionId")

	if id.UserId != "userId" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userId")
	}

	if id.MessageId != "messageId" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageId")
	}

	if id.MentionId != "mentionId" {
		t.Fatalf("Expected %q but got %q for Segment 'MentionId'", id.MentionId, "mentionId")
	}
}

func TestFormatUserIdMessageIdMentionID(t *testing.T) {
	actual := NewUserIdMessageIdMentionID("userId", "messageId", "mentionId").ID()
	expected := "/users/userId/messages/messageId/mentions/mentionId"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserIdMessageIdMentionID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMessageIdMentionId
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
			Input: "/users/userId/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/messages/messageId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/messages/messageId/mentions",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/messages/messageId/mentions/mentionId",
			Expected: &UserIdMessageIdMentionId{
				UserId:    "userId",
				MessageId: "messageId",
				MentionId: "mentionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/messages/messageId/mentions/mentionId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMessageIdMentionID(v.Input)
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

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestParseUserIdMessageIdMentionIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserIdMessageIdMentionId
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
			Input: "/users/userId/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/messages/messageId",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mEsSaGeS/mEsSaGeId",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userId/messages/messageId/mentions",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mEsSaGeS/mEsSaGeId/mEnTiOnS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userId/messages/messageId/mentions/mentionId",
			Expected: &UserIdMessageIdMentionId{
				UserId:    "userId",
				MessageId: "messageId",
				MentionId: "mentionId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userId/messages/messageId/mentions/mentionId/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mEsSaGeS/mEsSaGeId/mEnTiOnS/mEnTiOnId",
			Expected: &UserIdMessageIdMentionId{
				UserId:    "uSeRiD",
				MessageId: "mEsSaGeId",
				MentionId: "mEnTiOnId",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiD/mEsSaGeS/mEsSaGeId/mEnTiOnS/mEnTiOnId/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserIdMessageIdMentionIDInsensitively(v.Input)
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

		if actual.MessageId != v.Expected.MessageId {
			t.Fatalf("Expected %q but got %q for MessageId", v.Expected.MessageId, actual.MessageId)
		}

		if actual.MentionId != v.Expected.MentionId {
			t.Fatalf("Expected %q but got %q for MentionId", v.Expected.MentionId, actual.MentionId)
		}

	}
}

func TestSegmentsForUserIdMessageIdMentionId(t *testing.T) {
	segments := UserIdMessageIdMentionId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserIdMessageIdMentionId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
