package memessage

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeMessageId{}

func TestNewMeMessageID(t *testing.T) {
	id := NewMeMessageID("messageIdValue")

	if id.MessageId != "messageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'MessageId'", id.MessageId, "messageIdValue")
	}
}

func TestFormatMeMessageID(t *testing.T) {
	actual := NewMeMessageID("messageIdValue").ID()
	expected := "/me/messages/messageIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseMeMessageID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMessageId
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
			// Valid URI
			Input: "/me/messages/messageIdValue",
			Expected: &MeMessageId{
				MessageId: "messageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/messages/messageIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMessageID(v.Input)
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

	}
}

func TestParseMeMessageIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *MeMessageId
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
			// Valid URI
			Input: "/me/messages/messageIdValue",
			Expected: &MeMessageId{
				MessageId: "messageIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/me/messages/messageIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeIdVaLuE",
			Expected: &MeMessageId{
				MessageId: "mEsSaGeIdVaLuE",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/mE/mEsSaGeS/mEsSaGeIdVaLuE/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseMeMessageIDInsensitively(v.Input)
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

	}
}

func TestSegmentsForMeMessageId(t *testing.T) {
	segments := MeMessageId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("MeMessageId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
