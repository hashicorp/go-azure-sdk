package mechattabteamsapp

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatTabId{}

// MeChatTabId is a struct representing the Resource ID for a Me Chat Tab
type MeChatTabId struct {
	ChatId     string
	TeamsTabId string
}

// NewMeChatTabID returns a new MeChatTabId struct
func NewMeChatTabID(chatId string, teamsTabId string) MeChatTabId {
	return MeChatTabId{
		ChatId:     chatId,
		TeamsTabId: teamsTabId,
	}
}

// ParseMeChatTabID parses 'input' into a MeChatTabId
func ParseMeChatTabID(input string) (*MeChatTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatTabId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatTabId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ParseMeChatTabIDInsensitively parses 'input' case-insensitively into a MeChatTabId
// note: this method should only be used for API response data and not user input
func ParseMeChatTabIDInsensitively(input string) (*MeChatTabId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatTabId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatTabId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.TeamsTabId, ok = parsed.Parsed["teamsTabId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "teamsTabId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatTabID checks that 'input' can be parsed as a Me Chat Tab ID
func ValidateMeChatTabID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatTabID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Tab ID
func (id MeChatTabId) ID() string {
	fmtString := "/me/chats/%s/tabs/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.TeamsTabId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Tab ID
func (id MeChatTabId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticTabs", "tabs", "tabs"),
		resourceids.UserSpecifiedSegment("teamsTabId", "teamsTabIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Tab ID
func (id MeChatTabId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Teams Tab: %q", id.TeamsTabId),
	}
	return fmt.Sprintf("Me Chat Tab (%s)", strings.Join(components, "\n"))
}
