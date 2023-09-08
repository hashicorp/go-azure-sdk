package usermessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMessageId{}

// UserMessageId is a struct representing the Resource ID for a User Message
type UserMessageId struct {
	UserId    string
	MessageId string
}

// NewUserMessageID returns a new UserMessageId struct
func NewUserMessageID(userId string, messageId string) UserMessageId {
	return UserMessageId{
		UserId:    userId,
		MessageId: messageId,
	}
}

// ParseUserMessageID parses 'input' into a UserMessageId
func ParseUserMessageID(input string) (*UserMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMessageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ParseUserMessageIDInsensitively parses 'input' case-insensitively into a UserMessageId
// note: this method should only be used for API response data and not user input
func ParseUserMessageIDInsensitively(input string) (*UserMessageId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMessageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMessageId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	return &id, nil
}

// ValidateUserMessageID checks that 'input' can be parsed as a User Message ID
func ValidateUserMessageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMessageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Message ID
func (id UserMessageId) ID() string {
	fmtString := "/users/%s/messages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MessageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Message ID
func (id UserMessageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
	}
}

// String returns a human-readable description of this User Message ID
func (id UserMessageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Message: %q", id.MessageId),
	}
	return fmt.Sprintf("User Message (%s)", strings.Join(components, "\n"))
}
