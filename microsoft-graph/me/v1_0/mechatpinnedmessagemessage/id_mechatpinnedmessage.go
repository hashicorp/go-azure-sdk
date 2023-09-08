package mechatpinnedmessagemessage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatPinnedMessageId{}

// MeChatPinnedMessageId is a struct representing the Resource ID for a Me Chat Pinned Message
type MeChatPinnedMessageId struct {
	ChatId                  string
	PinnedChatMessageInfoId string
}

// NewMeChatPinnedMessageID returns a new MeChatPinnedMessageId struct
func NewMeChatPinnedMessageID(chatId string, pinnedChatMessageInfoId string) MeChatPinnedMessageId {
	return MeChatPinnedMessageId{
		ChatId:                  chatId,
		PinnedChatMessageInfoId: pinnedChatMessageInfoId,
	}
}

// ParseMeChatPinnedMessageID parses 'input' into a MeChatPinnedMessageId
func ParseMeChatPinnedMessageID(input string) (*MeChatPinnedMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatPinnedMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatPinnedMessageId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.PinnedChatMessageInfoId, ok = parsed.Parsed["pinnedChatMessageInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "pinnedChatMessageInfoId", *parsed)
	}

	return &id, nil
}

// ParseMeChatPinnedMessageIDInsensitively parses 'input' case-insensitively into a MeChatPinnedMessageId
// note: this method should only be used for API response data and not user input
func ParseMeChatPinnedMessageIDInsensitively(input string) (*MeChatPinnedMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatPinnedMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatPinnedMessageId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.PinnedChatMessageInfoId, ok = parsed.Parsed["pinnedChatMessageInfoId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "pinnedChatMessageInfoId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatPinnedMessageID checks that 'input' can be parsed as a Me Chat Pinned Message ID
func ValidateMeChatPinnedMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatPinnedMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Pinned Message ID
func (id MeChatPinnedMessageId) ID() string {
	fmtString := "/me/chats/%s/pinnedMessages/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.PinnedChatMessageInfoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Pinned Message ID
func (id MeChatPinnedMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticPinnedMessages", "pinnedMessages", "pinnedMessages"),
		resourceids.UserSpecifiedSegment("pinnedChatMessageInfoId", "pinnedChatMessageInfoIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Pinned Message ID
func (id MeChatPinnedMessageId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Pinned Chat Message Info: %q", id.PinnedChatMessageInfoId),
	}
	return fmt.Sprintf("Me Chat Pinned Message (%s)", strings.Join(components, "\n"))
}
