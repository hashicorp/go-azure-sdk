package usermessageattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserMessageAttachmentId{}

// UserMessageAttachmentId is a struct representing the Resource ID for a User Message Attachment
type UserMessageAttachmentId struct {
	UserId       string
	MessageId    string
	AttachmentId string
}

// NewUserMessageAttachmentID returns a new UserMessageAttachmentId struct
func NewUserMessageAttachmentID(userId string, messageId string, attachmentId string) UserMessageAttachmentId {
	return UserMessageAttachmentId{
		UserId:       userId,
		MessageId:    messageId,
		AttachmentId: attachmentId,
	}
}

// ParseUserMessageAttachmentID parses 'input' into a UserMessageAttachmentId
func ParseUserMessageAttachmentID(input string) (*UserMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMessageAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMessageAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserMessageAttachmentIDInsensitively parses 'input' case-insensitively into a UserMessageAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserMessageAttachmentIDInsensitively(input string) (*UserMessageAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserMessageAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserMessageAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.MessageId, ok = parsed.Parsed["messageId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "messageId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserMessageAttachmentID checks that 'input' can be parsed as a User Message Attachment ID
func ValidateUserMessageAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserMessageAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Message Attachment ID
func (id UserMessageAttachmentId) ID() string {
	fmtString := "/users/%s/messages/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.MessageId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Message Attachment ID
func (id UserMessageAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticMessages", "messages", "messages"),
		resourceids.UserSpecifiedSegment("messageId", "messageIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Message Attachment ID
func (id UserMessageAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Message: %q", id.MessageId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Message Attachment (%s)", strings.Join(components, "\n"))
}
