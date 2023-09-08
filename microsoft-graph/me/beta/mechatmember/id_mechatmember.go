package mechatmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeChatMemberId{}

// MeChatMemberId is a struct representing the Resource ID for a Me Chat Member
type MeChatMemberId struct {
	ChatId               string
	ConversationMemberId string
}

// NewMeChatMemberID returns a new MeChatMemberId struct
func NewMeChatMemberID(chatId string, conversationMemberId string) MeChatMemberId {
	return MeChatMemberId{
		ChatId:               chatId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseMeChatMemberID parses 'input' into a MeChatMemberId
func ParseMeChatMemberID(input string) (*MeChatMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMemberId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseMeChatMemberIDInsensitively parses 'input' case-insensitively into a MeChatMemberId
// note: this method should only be used for API response data and not user input
func ParseMeChatMemberIDInsensitively(input string) (*MeChatMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeChatMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeChatMemberId{}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateMeChatMemberID checks that 'input' can be parsed as a Me Chat Member ID
func ValidateMeChatMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeChatMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Chat Member ID
func (id MeChatMemberId) ID() string {
	fmtString := "/me/chats/%s/members/%s"
	return fmt.Sprintf(fmtString, id.ChatId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Chat Member ID
func (id MeChatMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this Me Chat Member ID
func (id MeChatMemberId) String() string {
	components := []string{
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("Me Chat Member (%s)", strings.Join(components, "\n"))
}
