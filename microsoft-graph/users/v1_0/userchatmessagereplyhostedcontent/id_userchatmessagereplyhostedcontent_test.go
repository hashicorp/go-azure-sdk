package userchatmessagereplyhostedcontent

import (
	"testing"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatMessageReplyHostedContentId{}

func TestNewUserChatMessageReplyHostedContentID(t *testing.T) {
	id := NewUserChatMessageReplyHostedContentID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value", "chatMessageHostedContentIdValue")

	if id.UserId != "userIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'UserId'", id.UserId, "userIdValue")
	}

	if id.ChatId != "chatIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatId'", id.ChatId, "chatIdValue")
	}

	if id.ChatMessageId != "chatMessageIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId'", id.ChatMessageId, "chatMessageIdValue")
	}

	if id.ChatMessageId1 != "chatMessageId1Value" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageId1'", id.ChatMessageId1, "chatMessageId1Value")
	}

	if id.ChatMessageHostedContentId != "chatMessageHostedContentIdValue" {
		t.Fatalf("Expected %q but got %q for Segment 'ChatMessageHostedContentId'", id.ChatMessageHostedContentId, "chatMessageHostedContentIdValue")
	}
}

func TestFormatUserChatMessageReplyHostedContentID(t *testing.T) {
	actual := NewUserChatMessageReplyHostedContentID("userIdValue", "chatIdValue", "chatMessageIdValue", "chatMessageId1Value", "chatMessageHostedContentIdValue").ID()
	expected := "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue"
	if actual != expected {
		t.Fatalf("Expected the Formatted ID to be %q but got %q", expected, actual)
	}
}

func TestParseUserChatMessageReplyHostedContentID(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserChatMessageReplyHostedContentId
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
			Input: "/users/userIdValue/chats",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue",
			Expected: &UserChatMessageReplyHostedContentId{
				UserId:                     "userIdValue",
				ChatId:                     "chatIdValue",
				ChatMessageId:              "chatMessageIdValue",
				ChatMessageId1:             "chatMessageId1Value",
				ChatMessageHostedContentId: "chatMessageHostedContentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserChatMessageReplyHostedContentID(v.Input)
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

		if actual.ChatId != v.Expected.ChatId {
			t.Fatalf("Expected %q but got %q for ChatId", v.Expected.ChatId, actual.ChatId)
		}

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestParseUserChatMessageReplyHostedContentIDInsensitively(t *testing.T) {
	testData := []struct {
		Input    string
		Error    bool
		Expected *UserChatMessageReplyHostedContentId
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
			Input: "/users/userIdValue/chats",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/mEsSaGeS",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe",
			Error: true,
		},
		{
			// Incomplete URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents",
			Error: true,
		},
		{
			// Incomplete URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS",
			Error: true,
		},
		{
			// Valid URI
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue",
			Expected: &UserChatMessageReplyHostedContentId{
				UserId:                     "userIdValue",
				ChatId:                     "chatIdValue",
				ChatMessageId:              "chatMessageIdValue",
				ChatMessageId1:             "chatMessageId1Value",
				ChatMessageHostedContentId: "chatMessageHostedContentIdValue",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment)
			Input: "/users/userIdValue/chats/chatIdValue/messages/chatMessageIdValue/replies/chatMessageId1Value/hostedContents/chatMessageHostedContentIdValue/extra",
			Error: true,
		},
		{
			// Valid URI (mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe",
			Expected: &UserChatMessageReplyHostedContentId{
				UserId:                     "uSeRiDvAlUe",
				ChatId:                     "cHaTiDvAlUe",
				ChatMessageId:              "cHaTmEsSaGeIdVaLuE",
				ChatMessageId1:             "cHaTmEsSaGeId1vAlUe",
				ChatMessageHostedContentId: "cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe",
			},
		},
		{
			// Invalid (Valid Uri with Extra segment - mIxEd CaSe since this is insensitive)
			Input: "/uSeRs/uSeRiDvAlUe/cHaTs/cHaTiDvAlUe/mEsSaGeS/cHaTmEsSaGeIdVaLuE/rEpLiEs/cHaTmEsSaGeId1vAlUe/hOsTeDcOnTeNtS/cHaTmEsSaGeHoStEdCoNtEnTiDvAlUe/extra",
			Error: true,
		},
	}
	for _, v := range testData {
		t.Logf("[DEBUG] Testing %q", v.Input)

		actual, err := ParseUserChatMessageReplyHostedContentIDInsensitively(v.Input)
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

		if actual.ChatId != v.Expected.ChatId {
			t.Fatalf("Expected %q but got %q for ChatId", v.Expected.ChatId, actual.ChatId)
		}

		if actual.ChatMessageId != v.Expected.ChatMessageId {
			t.Fatalf("Expected %q but got %q for ChatMessageId", v.Expected.ChatMessageId, actual.ChatMessageId)
		}

		if actual.ChatMessageId1 != v.Expected.ChatMessageId1 {
			t.Fatalf("Expected %q but got %q for ChatMessageId1", v.Expected.ChatMessageId1, actual.ChatMessageId1)
		}

		if actual.ChatMessageHostedContentId != v.Expected.ChatMessageHostedContentId {
			t.Fatalf("Expected %q but got %q for ChatMessageHostedContentId", v.Expected.ChatMessageHostedContentId, actual.ChatMessageHostedContentId)
		}

	}
}

func TestSegmentsForUserChatMessageReplyHostedContentId(t *testing.T) {
	segments := UserChatMessageReplyHostedContentId{}.Segments()
	if len(segments) == 0 {
		t.Fatalf("UserChatMessageReplyHostedContentId has no segments")
	}

	uniqueNames := make(map[string]struct{}, 0)
	for _, segment := range segments {
		uniqueNames[segment.Name] = struct{}{}
	}
	if len(uniqueNames) != len(segments) {
		t.Fatalf("Expected the Segments to be unique but got %q unique segments and %d total segments", len(uniqueNames), len(segments))
	}
}
