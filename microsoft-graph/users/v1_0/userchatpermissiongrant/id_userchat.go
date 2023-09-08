package userchatpermissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatId{}

// UserChatId is a struct representing the Resource ID for a User Chat
type UserChatId struct {
	UserId string
	ChatId string
}

// NewUserChatID returns a new UserChatId struct
func NewUserChatID(userId string, chatId string) UserChatId {
	return UserChatId{
		UserId: userId,
		ChatId: chatId,
	}
}

// ParseUserChatID parses 'input' into a UserChatId
func ParseUserChatID(input string) (*UserChatId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	return &id, nil
}

// ParseUserChatIDInsensitively parses 'input' case-insensitively into a UserChatId
// note: this method should only be used for API response data and not user input
func ParseUserChatIDInsensitively(input string) (*UserChatId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	return &id, nil
}

// ValidateUserChatID checks that 'input' can be parsed as a User Chat ID
func ValidateUserChatID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat ID
func (id UserChatId) ID() string {
	fmtString := "/users/%s/chats/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat ID
func (id UserChatId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
	}
}

// String returns a human-readable description of this User Chat ID
func (id UserChatId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
	}
	return fmt.Sprintf("User Chat (%s)", strings.Join(components, "\n"))
}
