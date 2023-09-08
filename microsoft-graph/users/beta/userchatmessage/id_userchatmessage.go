package userchatmessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatMessageId{}

// UserChatMessageId is a struct representing the Resource ID for a User Chat Message
type UserChatMessageId struct {
	UserId        string
	ChatId        string
	ChatMessageId string
}

// NewUserChatMessageID returns a new UserChatMessageId struct
func NewUserChatMessageID(userId string, chatId string, chatMessageId string) UserChatMessageId {
	return UserChatMessageId{
		UserId:        userId,
		ChatId:        chatId,
		ChatMessageId: chatMessageId,
	}
}

// ParseUserChatMessageID parses 'input' into a UserChatMessageId
func ParseUserChatMessageID(input string) (*UserChatMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ParseUserChatMessageIDInsensitively parses 'input' case-insensitively into a UserChatMessageId
// note: this method should only be used for API response data and not user input
func ParseUserChatMessageIDInsensitively(input string) (*UserChatMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ValidateUserChatMessageID checks that 'input' can be parsed as a User Chat Message ID
func ValidateUserChatMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Message ID
func (id UserChatMessageId) ID() string {
	fmtString := "/users/%s/chats/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Message ID
func (id UserChatMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
	}
}

// String returns a human-readable description of this User Chat Message ID
func (id UserChatMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("User Chat Message (%s)", strings.Join(components, "\n"))
}
