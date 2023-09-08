package usermessagemention

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMessageMentionId{}

// UserMessageMentionId is a struct representing the Resource ID for a User Message Mention
type UserMessageMentionId struct {
	UserId    string
	MessageId string
	MentionId string
}

// NewUserMessageMentionID returns a new UserMessageMentionId struct
func NewUserMessageMentionID(userId string, messageId string, mentionId string) UserMessageMentionId {
	return UserMessageMentionId{
		UserId:    userId,
		MessageId: messageId,
		MentionId: mentionId,
	}
}

// ParseUserMessageMentionID parses 'input' into a UserMessageMentionId
func ParseUserMessageMentionID(input string) (*UserMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMessageMentionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMessageMentionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ParseUserMessageMentionIDInsensitively parses 'input' case-insensitively into a UserMessageMentionId
// note: this method should only be used for API response data and not user input
func ParseUserMessageMentionIDInsensitively(input string) (*UserMessageMentionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMessageMentionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMessageMentionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.MentionId, ok = parsed.Parsed["mentionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "mentionId", *parsed)
	}

	return &id, nil
}

// ValidateUserMessageMentionID checks that 'input' can be parsed as a User Message Mention ID
func ValidateUserMessageMentionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMessageMentionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Message Mention ID
func (id UserMessageMentionId) ID() string {
	fmtString := "/users/%s/messages/%s/mentions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MessageId, id.MentionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Message Mention ID
func (id UserMessageMentionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticMentions", "mentions", "mentions"),
		resourceids.UserSpecifiedSegment("mentionId", "mentionIdValue"),
	}
}

// String returns a human-readable description of this User Message Mention ID
func (id UserMessageMentionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Mention: %q", id.MentionId),
	}
	return fmt.Sprintf("User Message Mention (%s)", strings.Join(components, "\n"))
}
