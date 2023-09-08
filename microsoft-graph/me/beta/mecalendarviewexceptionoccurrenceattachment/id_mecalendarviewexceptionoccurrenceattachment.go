package mecalendarviewexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewExceptionOccurrenceAttachmentId{}

// MeCalendarViewExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a Me Calendar View Exception Occurrence Attachment
type MeCalendarViewExceptionOccurrenceAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarViewExceptionOccurrenceAttachmentID returns a new MeCalendarViewExceptionOccurrenceAttachmentId struct
func NewMeCalendarViewExceptionOccurrenceAttachmentID(eventId string, eventId1 string, attachmentId string) MeCalendarViewExceptionOccurrenceAttachmentId {
	return MeCalendarViewExceptionOccurrenceAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarViewExceptionOccurrenceAttachmentID parses 'input' into a MeCalendarViewExceptionOccurrenceAttachmentId
func ParseMeCalendarViewExceptionOccurrenceAttachmentID(input string) (*MeCalendarViewExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExceptionOccurrenceAttachmentId{}

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

// ParseMeCalendarViewExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarViewExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewExceptionOccurrenceAttachmentIDInsensitively(input string) (*MeCalendarViewExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExceptionOccurrenceAttachmentId{}

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

// ValidateMeCalendarViewExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a Me Calendar View Exception Occurrence Attachment ID
func ValidateMeCalendarViewExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Exception Occurrence Attachment ID
func (id MeCalendarViewExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/me/calendarView/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Exception Occurrence Attachment ID
func (id MeCalendarViewExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Exception Occurrence Attachment ID
func (id MeCalendarViewExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar View Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
