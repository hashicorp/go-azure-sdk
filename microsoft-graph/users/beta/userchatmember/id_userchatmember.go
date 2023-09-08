package userchatmember

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatMemberId{}

// UserChatMemberId is a struct representing the Resource ID for a User Chat Member
type UserChatMemberId struct {
	UserId               string
	ChatId               string
	ConversationMemberId string
}

// NewUserChatMemberID returns a new UserChatMemberId struct
func NewUserChatMemberID(userId string, chatId string, conversationMemberId string) UserChatMemberId {
	return UserChatMemberId{
		UserId:               userId,
		ChatId:               chatId,
		ConversationMemberId: conversationMemberId,
	}
}

// ParseUserChatMemberID parses 'input' into a UserChatMemberId
func ParseUserChatMemberID(input string) (*UserChatMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMemberId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ParseUserChatMemberIDInsensitively parses 'input' case-insensitively into a UserChatMemberId
// note: this method should only be used for API response data and not user input
func ParseUserChatMemberIDInsensitively(input string) (*UserChatMemberId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMemberId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMemberId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ChatId, ok = parsed.Parsed["chatId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "chatId", *parsed)
	}

	if id.ConversationMemberId, ok = parsed.Parsed["conversationMemberId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "conversationMemberId", *parsed)
	}

	return &id, nil
}

// ValidateUserChatMemberID checks that 'input' can be parsed as a User Chat Member ID
func ValidateUserChatMemberID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatMemberID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Member ID
func (id UserChatMemberId) ID() string {
	fmtString := "/users/%s/chats/%s/members/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ConversationMemberId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Member ID
func (id UserChatMemberId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMembers", "members", "members"),
		resourceids.UserSpecifiedSegment("conversationMemberId", "conversationMemberIdValue"),
	}
}

// String returns a human-readable description of this User Chat Member ID
func (id UserChatMemberId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Conversation Member: %q", id.ConversationMemberId),
	}
	return fmt.Sprintf("User Chat Member (%s)", strings.Join(components, "\n"))
}
