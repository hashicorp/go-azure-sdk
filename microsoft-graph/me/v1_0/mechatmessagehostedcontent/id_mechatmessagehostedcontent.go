package mechatmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatMessageHostedContentId{}

// MeChatMessageHostedContentId is a struct representing the Resource ID for a Me Chat Message Hosted Content
type MeChatMessageHostedContentId struct {
	ChatId                     string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewMeChatMessageHostedContentID returns a new MeChatMessageHostedContentId struct
func NewMeChatMessageHostedContentID(chatId string, chatMessageId string, chatMessageHostedContentId string) MeChatMessageHostedContentId {
	return MeChatMessageHostedContentId{
		ChatId:                     chatId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseMeChatMessageHostedContentID parses 'input' into a MeChatMessageHostedContentId
func ParseMeChatMessageHostedContentID(input string) (*MeChatMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMessageHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMessageHostedContentId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ParseMeChatMessageHostedContentIDInsensitively parses 'input' case-insensitively into a MeChatMessageHostedContentId
// note: this method should only be used for API response data and not user input
func ParseMeChatMessageHostedContentIDInsensitively(input string) (*MeChatMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMessageHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMessageHostedContentId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ChatMessageId, ok = parsed.Parsed["chatMessageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageId", *parsed)
	}

	if id.ChatMessageHostedContentId, ok = parsed.Parsed["chatMessageHostedContentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatMessageHostedContentId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatMessageHostedContentID checks that 'input' can be parsed as a Me Chat Message Hosted Content ID
func ValidateMeChatMessageHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatMessageHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Message Hosted Content ID
func (id MeChatMessageHostedContentId) ID() string {
	fmtString := "/me/chats/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Message Hosted Content ID
func (id MeChatMessageHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Message Hosted Content ID
func (id MeChatMessageHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("Me Chat Message Hosted Content (%s)", strings.Join(components, "\n"))
}
