package userchatmessagereplyhostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatMessageReplyHostedContentId{}

// UserChatMessageReplyHostedContentId is a struct representing the Resource ID for a User Chat Message Reply Hosted Content
type UserChatMessageReplyHostedContentId struct {
	UserId                     string
	ChatId                     string
	ChatMessageId              string
	ChatMessageId1             string
	ChatMessageHostedContentId string
}

// NewUserChatMessageReplyHostedContentID returns a new UserChatMessageReplyHostedContentId struct
func NewUserChatMessageReplyHostedContentID(userId string, chatId string, chatMessageId string, chatMessageId1 string, chatMessageHostedContentId string) UserChatMessageReplyHostedContentId {
	return UserChatMessageReplyHostedContentId{
		UserId:                     userId,
		ChatId:                     chatId,
		ChatMessageId:              chatMessageId,
		ChatMessageId1:             chatMessageId1,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserChatMessageReplyHostedContentID parses 'input' into a UserChatMessageReplyHostedContentId
func ParseUserChatMessageReplyHostedContentID(input string) (*UserChatMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMessageReplyHostedContentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserChatMessageReplyHostedContentIDInsensitively parses 'input' case-insensitively into a UserChatMessageReplyHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserChatMessageReplyHostedContentIDInsensitively(input string) (*UserChatMessageReplyHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMessageReplyHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMessageReplyHostedContentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserChatMessageReplyHostedContentID checks that 'input' can be parsed as a User Chat Message Reply Hosted Content ID
func ValidateUserChatMessageReplyHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatMessageReplyHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Message Reply Hosted Content ID
func (id UserChatMessageReplyHostedContentId) ID() string {
	fmtString := "/users/%s/chats/%s/messages/%s/replies/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ChatMessageId, id.ChatMessageId1, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Message Reply Hosted Content ID
func (id UserChatMessageReplyHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
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

// String returns a human-readable description of this User Chat Message Reply Hosted Content ID
func (id UserChatMessageReplyHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Id 1: %q", id.ChatMessageId1),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Chat Message Reply Hosted Content (%s)", strings.Join(components, "\n"))
}
