package mechatmessagereply

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatMessageReplyId{}

// MeChatMessageReplyId is a struct representing the Resource ID for a Me Chat Message Reply
type MeChatMessageReplyId struct {
	ChatId         string
	ChatMessageId  string
	ChatMessageId1 string
}

// NewMeChatMessageReplyID returns a new MeChatMessageReplyId struct
func NewMeChatMessageReplyID(chatId string, chatMessageId string, chatMessageId1 string) MeChatMessageReplyId {
	return MeChatMessageReplyId{
		ChatId:         chatId,
		ChatMessageId:  chatMessageId,
		ChatMessageId1: chatMessageId1,
	}
}

// ParseMeChatMessageReplyID parses 'input' into a MeChatMessageReplyId
func ParseMeChatMessageReplyID(input string) (*MeChatMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMessageReplyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMessageReplyId{}

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

// ParseMeChatMessageReplyIDInsensitively parses 'input' case-insensitively into a MeChatMessageReplyId
// note: this method should only be used for API response data and not user input
func ParseMeChatMessageReplyIDInsensitively(input string) (*MeChatMessageReplyId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMessageReplyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMessageReplyId{}

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

// ValidateMeChatMessageReplyID checks that 'input' can be parsed as a Me Chat Message Reply ID
func ValidateMeChatMessageReplyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatMessageReplyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Message Reply ID
func (id MeChatMessageReplyId) ID() string {
	fmtString := "/me/chats/%s/messages/%s/replies/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ChatMessageId, id.ChatMessageId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Message Reply ID
func (id MeChatMessageReplyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
	}
}

// String returns a human-readable description of this Me Chat Message Reply ID
func (id MeChatMessageReplyId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
	}
	return fmt.Sprintf("Me Chat Message Reply (%s)", strings.Join(components, "\n"))
}
