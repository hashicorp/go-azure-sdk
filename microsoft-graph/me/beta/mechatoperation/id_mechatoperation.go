package mechatoperation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatOperationId{}

// MeChatOperationId is a struct representing the Resource ID for a Me Chat Operation
type MeChatOperationId struct {
	ChatId                string
	TeamsAsyncOperationId string
}

// NewMeChatOperationID returns a new MeChatOperationId struct
func NewMeChatOperationID(chatId string, teamsAsyncOperationId string) MeChatOperationId {
	return MeChatOperationId{
		ChatId:                chatId,
		TeamsAsyncOperationId: teamsAsyncOperationId,
	}
}

// ParseMeChatOperationID parses 'input' into a MeChatOperationId
func ParseMeChatOperationID(input string) (*MeChatOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatOperationId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ParseMeChatOperationIDInsensitively parses 'input' case-insensitively into a MeChatOperationId
// note: this method should only be used for API response data and not user input
func ParseMeChatOperationIDInsensitively(input string) (*MeChatOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatOperationId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsAsyncOperationId, ok = parsed.Parsed["teamsAsyncOperationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsAsyncOperationId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatOperationID checks that 'input' can be parsed as a Me Chat Operation ID
func ValidateMeChatOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Operation ID
func (id MeChatOperationId) ID() string {
	fmtString := "/me/chats/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.TeamsAsyncOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Operation ID
func (id MeChatOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("teamsAsyncOperationId", "teamsAsyncOperationIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Operation ID
func (id MeChatOperationId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams Async Operation: %q", id.TeamsAsyncOperationId),
	}
	return fmt.Sprintf("Me Chat Operation (%s)", strings.Join(components, "\n"))
}
