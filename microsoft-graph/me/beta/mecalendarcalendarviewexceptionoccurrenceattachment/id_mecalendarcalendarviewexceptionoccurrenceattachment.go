package mecalendarcalendarviewexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewExceptionOccurrenceAttachmentId{}

// MeCalendarCalendarViewExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a Me Calendar Calendar View Exception Occurrence Attachment
type MeCalendarCalendarViewExceptionOccurrenceAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarCalendarViewExceptionOccurrenceAttachmentID returns a new MeCalendarCalendarViewExceptionOccurrenceAttachmentId struct
func NewMeCalendarCalendarViewExceptionOccurrenceAttachmentID(eventId string, eventId1 string, attachmentId string) MeCalendarCalendarViewExceptionOccurrenceAttachmentId {
	return MeCalendarCalendarViewExceptionOccurrenceAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarCalendarViewExceptionOccurrenceAttachmentID parses 'input' into a MeCalendarCalendarViewExceptionOccurrenceAttachmentId
func ParseMeCalendarCalendarViewExceptionOccurrenceAttachmentID(input string) (*MeCalendarCalendarViewExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewExceptionOccurrenceAttachmentId{}

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

// ParseMeCalendarCalendarViewExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewExceptionOccurrenceAttachmentIDInsensitively(input string) (*MeCalendarCalendarViewExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewExceptionOccurrenceAttachmentId{}

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

// ValidateMeCalendarCalendarViewExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a Me Calendar Calendar View Exception Occurrence Attachment ID
func ValidateMeCalendarCalendarViewExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Exception Occurrence Attachment ID
func (id MeCalendarCalendarViewExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Exception Occurrence Attachment ID
func (id MeCalendarCalendarViewExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Exception Occurrence Attachment ID
func (id MeCalendarCalendarViewExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
