package userchatpinnedmessagemessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatPinnedMessageId{}

// UserChatPinnedMessageId is a struct representing the Resource ID for a User Chat Pinned Message
type UserChatPinnedMessageId struct {
	UserId                  string
	ChatId                  string
	PinnedChatMessageInfoId string
}

// NewUserChatPinnedMessageID returns a new UserChatPinnedMessageId struct
func NewUserChatPinnedMessageID(userId string, chatId string, pinnedChatMessageInfoId string) UserChatPinnedMessageId {
	return UserChatPinnedMessageId{
		UserId:                  userId,
		ChatId:                  chatId,
		PinnedChatMessageInfoId: pinnedChatMessageInfoId,
	}
}

// ParseUserChatPinnedMessageID parses 'input' into a UserChatPinnedMessageId
func ParseUserChatPinnedMessageID(input string) (*UserChatPinnedMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatPinnedMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatPinnedMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.PinnedChatMessageInfoId, ok = parsed.Parsed["pinnedChatMessageInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "pinnedChatMessageInfoId", *parsed)
	}

	return &id, nil
}

// ParseUserChatPinnedMessageIDInsensitively parses 'input' case-insensitively into a UserChatPinnedMessageId
// note: this method should only be used for API response data and not user input
func ParseUserChatPinnedMessageIDInsensitively(input string) (*UserChatPinnedMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatPinnedMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatPinnedMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.PinnedChatMessageInfoId, ok = parsed.Parsed["pinnedChatMessageInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "pinnedChatMessageInfoId", *parsed)
	}

	return &id, nil
}

// ValidateUserChatPinnedMessageID checks that 'input' can be parsed as a User Chat Pinned Message ID
func ValidateUserChatPinnedMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatPinnedMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Pinned Message ID
func (id UserChatPinnedMessageId) ID() string {
	fmtString := "/users/%s/chats/%s/pinnedMessages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.PinnedChatMessageInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Pinned Message ID
func (id UserChatPinnedMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticPinnedMessages", "pinnedMessages", "pinnedMessages"),
		resourceids.UserSpecifiedSegment("pinnedChatMessageInfoId", "pinnedChatMessageInfoIdValue"),
	}
}

// String returns a human-readable description of this User Chat Pinned Message ID
func (id UserChatPinnedMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Pinned Chat Message Info: %q", id.PinnedChatMessageInfoId),
	}
	return fmt.Sprintf("User Chat Pinned Message (%s)", strings.Join(components, "\n"))
}
