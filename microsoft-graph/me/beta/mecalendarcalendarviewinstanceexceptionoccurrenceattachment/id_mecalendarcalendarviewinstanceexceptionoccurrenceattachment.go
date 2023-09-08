package mecalendarcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{}

// MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a Me Calendar Calendar View Instance Exception Occurrence Attachment
type MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId struct {
	CalendarId   string
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID returns a new MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId struct
func NewMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID(calendarId string, eventId string, eventId1 string, eventId2 string, attachmentId string) MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId {
	return MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{
		CalendarId:   calendarId,
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID parses 'input' into a MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId
func ParseMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID(input string) (*MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
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

// ParseMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentIDInsensitively(input string) (*MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
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

// ValidateMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a Me Calendar Calendar View Instance Exception Occurrence Attachment ID
func ValidateMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Instance Exception Occurrence Attachment ID
func (id MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/instances/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Instance Exception Occurrence Attachment ID
func (id MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
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

// String returns a human-readable description of this Me Calendar Calendar View Instance Exception Occurrence Attachment ID
func (id MeCalendarCalendarViewInstanceExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Instance Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
