package mechatmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatMessageReplyHostedContentId{}

// MeChatMessageReplyHostedContentId is a struct representing the Resource ID for a Me Chat Message Reply Hosted Content
type MeChatMessageReplyHostedContentId struct {
	ChatId                     string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewMeChatMessageReplyHostedContentID returns a new MeChatMessageReplyHostedContentId struct
func NewMeChatMessageReplyHostedContentID(chatId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) MeChatMessageReplyHostedContentId {
	return MeChatMessageReplyHostedContentId{
		ChatId:                     chatId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeChatMessageReplyHostedContentID parses 'input' into a MeChatMessageReplyHostedContentId
func ParseMeChatMessageReplyHostedContentID(input string) (*MeChatMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMessageReplyHostedContentId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ParseMeChatMessageReplyHostedContentIDInsensitively parses 'input' case-insensitively into a MeChatMessageReplyHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeChatMessageReplyHostedContentIDInsensitively(input string) (*MeChatMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMessageReplyHostedContentId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageId1, ok = parsed.Parsed["chatMessageId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId1", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatMessageReplyHostedContentID checks that 'input' can be parsed as a Me Chat Message Reply Hosted Content ID
func ValidateMeChatMessageReplyHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatMessageReplyHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Message Reply Hosted Content ID
func (id MeChatMessageReplyHostedContentId) ID() string {
	fmtString := "/me/chats/%s/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Message Reply Hosted Content ID
func (id MeChatMessageReplyHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticReplies", "replies", "replies"),
		resourceids.UserSpecifiedSegment("chatMessageId1", "chatMessageId1Value"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Message Reply Hosted Content ID
func (id MeChatMessageReplyHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Chat Message Reply Hosted Content (%s)", strings.Join(components, "\n"))
}
