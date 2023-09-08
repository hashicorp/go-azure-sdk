package userchatoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatOperationId{}

// UserChatOperationId is a struct representing the Resource ID for a User Chat Operation
type UserChatOperationId struct {
	UserId                string
	ChatId                string
	TeamsAsyncOperationId string
}

// NewUserChatOperationID returns a new UserChatOperationId struct
func NewUserChatOperationID(userId string, chatId string, teamsAsyncOperationId string) UserChatOperationId {
	return UserChatOperationId{
		UserId:                userId,
		ChatId:                chatId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseUserChatOperationID parses 'input' into a UserChatOperationId
func ParseUserChatOperationID(input string) (*UserChatOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatOperationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ParseUserChatOperationIDInsensitively parses 'input' case-insensitively into a UserChatOperationId
// note: this method should only be used for API response data and not user input
func ParseUserChatOperationIDInsensitively(input string) (*UserChatOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatOperationId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ValidateUserChatOperationID checks that 'input' can be parsed as a User Chat Operation ID
func ValidateUserChatOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Operation ID
func (id UserChatOperationId) ID() string {
	fmtString := "/users/%s/chats/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Operation ID
func (id UserChatOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationIdValue"),
	}
}

// String returns a human-readable description of this User Chat Operation ID
func (id UserChatOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("User Chat Operation (%s)", strings.Join(components, "\n"))
}
