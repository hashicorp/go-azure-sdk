package userchatmessagehostedcontent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserChatMessageHostedContentId{}

// UserChatMessageHostedContentId is a struct representing the Resource ID for a User Chat Message Hosted Content
type UserChatMessageHostedContentId struct {
	UserId                     string
	ChatId                     string
	ChatMessageId              string
	ChatMessageHostedContentId string
}

// NewUserChatMessageHostedContentID returns a new UserChatMessageHostedContentId struct
func NewUserChatMessageHostedContentID(userId string, chatId string, chatMessageId string, chatMessageHostedContentId string) UserChatMessageHostedContentId {
	return UserChatMessageHostedContentId{
		UserId:                     userId,
		ChatId:                     chatId,
		ChatMessageId:              chatMessageId,
		ChatMessageHostedContentId: chatMessageHostedContentId,
	}
}

// ParseUserChatMessageHostedContentID parses 'input' into a UserChatMessageHostedContentId
func ParseUserChatMessageHostedContentID(input string) (*UserChatMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMessageHostedContentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMessageHostedContentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserChatMessageHostedContentIDInsensitively parses 'input' case-insensitively into a UserChatMessageHostedContentId
// note: this method should only be used for API response data and not user input
func ParseUserChatMessageHostedContentIDInsensitively(input string) (*UserChatMessageHostedContentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserChatMessageHostedContentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserChatMessageHostedContentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserChatMessageHostedContentID checks that 'input' can be parsed as a User Chat Message Hosted Content ID
func ValidateUserChatMessageHostedContentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserChatMessageHostedContentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Chat Message Hosted Content ID
func (id UserChatMessageHostedContentId) ID() string {
	fmtString := "/users/%s/chats/%s/messages/%s/hostedContents/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ChatId, id.ChatMessageId, id.ChatMessageHostedContentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Chat Message Hosted Content ID
func (id UserChatMessageHostedContentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticChats", "chats", "chats"),
		resourceids.UserSpecifiedSegment("chatId", "chatIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("chatMessageId", "chatMessageIdValue"),
		resourceids.StaticSegment("staticHostedContents", "hostedContents", "hostedContents"),
		resourceids.UserSpecifiedSegment("chatMessageHostedContentId", "chatMessageHostedContentIdValue"),
	}
}

// String returns a human-readable description of this User Chat Message Hosted Content ID
func (id UserChatMessageHostedContentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Chat: %q", id.ChatId),
		fmt.Sprintf("Chat Message: %q", id.ChatMessageId),
		fmt.Sprintf("Chat Message Hosted Content: %q", id.ChatMessageHostedContentId),
	}
	return fmt.Sprintf("User Chat Message Hosted Content (%s)", strings.Join(components, "\n"))
}
