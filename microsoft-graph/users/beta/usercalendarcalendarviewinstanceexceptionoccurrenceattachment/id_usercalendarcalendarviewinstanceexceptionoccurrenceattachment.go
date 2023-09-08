package usercalendarcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{}

// UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a User Calendar Calendar View Instance Exception Occurrence Attachment
type UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId struct {
	UserId       string
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID returns a new UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId struct
func NewUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID(userId string, eventId string, eventId1 string, eventId2 string, attachmentId string) UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId {
	return UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID parses 'input' into a UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId
func ParseUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID(input string) (*UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentIDInsensitively(input string) (*UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a User Calendar Calendar View Instance Exception Occurrence Attachment ID
func ValidateUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar View Instance Exception Occurrence Attachment ID
func (id UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s/instances/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar View Instance Exception Occurrence Attachment ID
func (id UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Calendar View Instance Exception Occurrence Attachment ID
func (id UserCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar Calendar View Instance Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
