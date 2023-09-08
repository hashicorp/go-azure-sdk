package usercalendareventexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventExceptionOccurrenceAttachmentId{}

// UserCalendarEventExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a User Calendar Event Exception Occurrence Attachment
type UserCalendarEventExceptionOccurrenceAttachmentId struct {
	UserId       string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewUserCalendarEventExceptionOccurrenceAttachmentID returns a new UserCalendarEventExceptionOccurrenceAttachmentId struct
func NewUserCalendarEventExceptionOccurrenceAttachmentID(userId string, eventId string, eventId1 string, attachmentId string) UserCalendarEventExceptionOccurrenceAttachmentId {
	return UserCalendarEventExceptionOccurrenceAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseUserCalendarEventExceptionOccurrenceAttachmentID parses 'input' into a UserCalendarEventExceptionOccurrenceAttachmentId
func ParseUserCalendarEventExceptionOccurrenceAttachmentID(input string) (*UserCalendarEventExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventExceptionOccurrenceAttachmentId{}

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

// ParseUserCalendarEventExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarEventExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventExceptionOccurrenceAttachmentIDInsensitively(input string) (*UserCalendarEventExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventExceptionOccurrenceAttachmentId{}

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

// ValidateUserCalendarEventExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a User Calendar Event Exception Occurrence Attachment ID
func ValidateUserCalendarEventExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event Exception Occurrence Attachment ID
func (id UserCalendarEventExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/users/%s/calendar/events/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event Exception Occurrence Attachment ID
func (id UserCalendarEventExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Event Exception Occurrence Attachment ID
func (id UserCalendarEventExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar Event Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
