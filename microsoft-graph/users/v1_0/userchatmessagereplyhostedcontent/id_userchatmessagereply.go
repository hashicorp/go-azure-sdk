package userchatmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatMessageReplyId{}

// UserChatMessageReplyId is a struct representing the Resource ID for a User Chat Message Reply
type UserChatMessageReplyId struct {
	UserId         string
	ChatId         string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewUserChatMessageReplyID returns a new UserChatMessageReplyId struct
func NewUserChatMessageReplyID(userId string, chatId string, chatMessageId string, chatMessageId1 string) UserChatMessageReplyId {
	return UserChatMessageReplyId{
		UserId:         userId,
		ChatId:         chatId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseUserChatMessageReplyID parses 'input' into a UserChatMessageReplyId
func ParseUserChatMessageReplyID(input string) (*UserChatMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMessageReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMessageReplyId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ParseUserChatMessageReplyIDInsensitively parses 'input' case-insensitively into a UserChatMessageReplyId
// note: this method should only be used for API response data and not user input
func ParseUserChatMessageReplyIDInsensitively(input string) (*UserChatMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMessageReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMessageReplyId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	return &id, nil
}

// ValidateUserChatMessageReplyID checks that 'input' can be parsed as a User Chat Message Reply ID
func ValidateUserChatMessageReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatMessageReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Message Reply ID
func (id UserChatMessageReplyId) ID() string {
	fmtString := "/users/%s/chats/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Message Reply ID
func (id UserChatMessageReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
	}
}

// String returns a human-readable description of this User Chat Message Reply ID
func (id UserChatMessageReplyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("User Chat Message Reply (%s)", strings.Join(components, "\n"))
}
