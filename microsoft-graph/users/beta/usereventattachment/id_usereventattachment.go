package usereventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserEventAttachmentId{}

// UserEventAttachmentId is a struct representing the Resource ID for a User Event Attachment
type UserEventAttachmentId struct {
	UserId       string
	EventId      string
	AttachmentId string
}

// NewUserEventAttachmentID returns a new UserEventAttachmentId struct
func NewUserEventAttachmentID(userId string, eventId string, attachmentId string) UserEventAttachmentId {
	return UserEventAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseUserEventAttachmentID parses 'input' into a UserEventAttachmentId
func ParseUserEventAttachmentID(input string) (*UserEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserEventAttachmentIDInsensitively parses 'input' case-insensitively into a UserEventAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserEventAttachmentIDInsensitively(input string) (*UserEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserEventAttachmentID checks that 'input' can be parsed as a User Event Attachment ID
func ValidateUserEventAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserEventAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Event Attachment ID
func (id UserEventAttachmentId) ID() string {
	fmtString := "/users/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Event Attachment ID
func (id UserEventAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Event Attachment ID
func (id UserEventAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Event Attachment (%s)", strings.Join(components, "\n"))
}
