package usercalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarViewExceptionOccurrenceInstanceAttachmentId{}

// UserCalendarViewExceptionOccurrenceInstanceAttachmentId is a struct representing the Resource ID for a User Calendar View Exception Occurrence Instance Attachment
type UserCalendarViewExceptionOccurrenceInstanceAttachmentId struct {
	UserId       string
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewUserCalendarViewExceptionOccurrenceInstanceAttachmentID returns a new UserCalendarViewExceptionOccurrenceInstanceAttachmentId struct
func NewUserCalendarViewExceptionOccurrenceInstanceAttachmentID(userId string, eventId string, eventId1 string, eventId2 string, attachmentId string) UserCalendarViewExceptionOccurrenceInstanceAttachmentId {
	return UserCalendarViewExceptionOccurrenceInstanceAttachmentId{
		UserId:       userId,
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseUserCalendarViewExceptionOccurrenceInstanceAttachmentID parses 'input' into a UserCalendarViewExceptionOccurrenceInstanceAttachmentId
func ParseUserCalendarViewExceptionOccurrenceInstanceAttachmentID(input string) (*UserCalendarViewExceptionOccurrenceInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewExceptionOccurrenceInstanceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewExceptionOccurrenceInstanceAttachmentId{}

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

// ParseUserCalendarViewExceptionOccurrenceInstanceAttachmentIDInsensitively parses 'input' case-insensitively into a UserCalendarViewExceptionOccurrenceInstanceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarViewExceptionOccurrenceInstanceAttachmentIDInsensitively(input string) (*UserCalendarViewExceptionOccurrenceInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewExceptionOccurrenceInstanceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewExceptionOccurrenceInstanceAttachmentId{}

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

// ValidateUserCalendarViewExceptionOccurrenceInstanceAttachmentID checks that 'input' can be parsed as a User Calendar View Exception Occurrence Instance Attachment ID
func ValidateUserCalendarViewExceptionOccurrenceInstanceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarViewExceptionOccurrenceInstanceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar View Exception Occurrence Instance Attachment ID
func (id UserCalendarViewExceptionOccurrenceInstanceAttachmentId) ID() string {
	fmtString := "/users/%s/calendarView/%s/exceptionOccurrences/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar View Exception Occurrence Instance Attachment ID
func (id UserCalendarViewExceptionOccurrenceInstanceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this User Calendar View Exception Occurrence Instance Attachment ID
func (id UserCalendarViewExceptionOccurrenceInstanceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("User Calendar View Exception Occurrence Instance Attachment (%s)", strings.Join(components, "\n"))
}
