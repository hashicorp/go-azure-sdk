package usereventexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserEventExceptionOccurrenceAttachmentId{}

// UserEventExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a User Event Exception Occurrence Attachment
type UserEventExceptionOccurrenceAttachmentId struct {
	UserId       string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewUserEventExceptionOccurrenceAttachmentID returns a new UserEventExceptionOccurrenceAttachmentId struct
func NewUserEventExceptionOccurrenceAttachmentID(userId string, eventId string, eventId1 string, attachmentId string) UserEventExceptionOccurrenceAttachmentId {
	return UserEventExceptionOccurrenceAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseUserEventExceptionOccurrenceAttachmentID parses 'input' into a UserEventExceptionOccurrenceAttachmentId
func ParseUserEventExceptionOccurrenceAttachmentID(input string) (*UserEventExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventExceptionOccurrenceAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserEventExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a UserEventExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserEventExceptionOccurrenceAttachmentIDInsensitively(input string) (*UserEventExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventExceptionOccurrenceAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserEventExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a User Event Exception Occurrence Attachment ID
func ValidateUserEventExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserEventExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Event Exception Occurrence Attachment ID
func (id UserEventExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/users/%s/events/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Event Exception Occurrence Attachment ID
func (id UserEventExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Event Exception Occurrence Attachment ID
func (id UserEventExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Event Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
