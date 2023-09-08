package usermessageextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMessageExtensionId{}

// UserMessageExtensionId is a struct representing the Resource ID for a User Message Extension
type UserMessageExtensionId struct {
	UserId      string
	MessageId   string
	ExtensionId string
}

// NewUserMessageExtensionID returns a new UserMessageExtensionId struct
func NewUserMessageExtensionID(userId string, messageId string, extensionId string) UserMessageExtensionId {
	return UserMessageExtensionId{
		UserId:      userId,
		MessageId:   messageId,
		ExtensionId: extensionId,
	}
}

// ParseUserMessageExtensionID parses 'input' into a UserMessageExtensionId
func ParseUserMessageExtensionID(input string) (*UserMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMessageExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMessageExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserMessageExtensionIDInsensitively parses 'input' case-insensitively into a UserMessageExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserMessageExtensionIDInsensitively(input string) (*UserMessageExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMessageExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMessageExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserMessageExtensionID checks that 'input' can be parsed as a User Message Extension ID
func ValidateUserMessageExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMessageExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Message Extension ID
func (id UserMessageExtensionId) ID() string {
	fmtString := "/users/%s/messages/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MessageId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Message Extension ID
func (id UserMessageExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Message Extension ID
func (id UserMessageExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Message Extension (%s)", strings.Join(components, "\n"))
}
