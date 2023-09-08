package mechatmessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatMessageId{}

// MeChatMessageId is a struct representing the Resource ID for a Me Chat Message
type MeChatMessageId struct {
	ChatId        string
	ChatMessageId string
}

// NewMeChatMessageID returns a new MeChatMessageId struct
func NewMeChatMessageID(chatId string, chatMessageId string) MeChatMessageId {
	return MeChatMessageId{
		ChatId:        chatId,
		ChatMessageId: chatMessageId,
	}
}

// ParseMeChatMessageID parses 'input' into a MeChatMessageId
func ParseMeChatMessageID(input string) (*MeChatMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMessageId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ParseMeChatMessageIDInsensitively parses 'input' case-insensitively into a MeChatMessageId
// note: this method should only be used for API response data and not user input
func ParseMeChatMessageIDInsensitively(input string) (*MeChatMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMessageId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatMessageID checks that 'input' can be parsed as a Me Chat Message ID
func ValidateMeChatMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Message ID
func (id MeChatMessageId) ID() string {
	fmtString := "/me/chats/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ChatMessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Message ID
func (id MeChatMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Message ID
func (id MeChatMessageId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
	}
	return fmt.Sprintf("Me Chat Message (%s)", strings.Join(components, "\n"))
}
