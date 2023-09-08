package userchattabteamsapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatTabId{}

// UserChatTabId is a struct representing the Resource ID for a User Chat Tab
type UserChatTabId struct {
	UserId     string
	ChatId     string
	TeamsTabId string
}

// NewUserChatTabID returns a new UserChatTabId struct
func NewUserChatTabID(userId string, chatId string, teamsTabId string) UserChatTabId {
	return UserChatTabId{
		UserId:     userId,
		ChatId:     chatId,
		TeamsTabId: teamsTabId,
	}
}

// ParseUserChatTabID parses 'input' into a UserChatTabId
func ParseUserChatTabID(input string) (*UserChatTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatTabId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ParseUserChatTabIDInsensitively parses 'input' case-insensitively into a UserChatTabId
// note: this method should only be used for API response data and not user input
func ParseUserChatTabIDInsensitively(input string) (*UserChatTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatTabId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ValidateUserChatTabID checks that 'input' can be parsed as a User Chat Tab ID
func ValidateUserChatTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Tab ID
func (id UserChatTabId) ID() string {
	fmtString := "/users/%s/chats/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Tab ID
func (id UserChatTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticTabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabIdValue"),
	}
}

// String returns a human-readable description of this User Chat Tab ID
func (id UserChatTabId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("User Chat Tab (%s)", strings.Join(components, "\n"))
}
