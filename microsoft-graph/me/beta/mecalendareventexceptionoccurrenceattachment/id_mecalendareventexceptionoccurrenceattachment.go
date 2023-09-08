package mecalendareventexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventExceptionOccurrenceAttachmentId{}

// MeCalendarEventExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a Me Calendar Event Exception Occurrence Attachment
type MeCalendarEventExceptionOccurrenceAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarEventExceptionOccurrenceAttachmentID returns a new MeCalendarEventExceptionOccurrenceAttachmentId struct
func NewMeCalendarEventExceptionOccurrenceAttachmentID(eventId string, eventId1 string, attachmentId string) MeCalendarEventExceptionOccurrenceAttachmentId {
	return MeCalendarEventExceptionOccurrenceAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarEventExceptionOccurrenceAttachmentID parses 'input' into a MeCalendarEventExceptionOccurrenceAttachmentId
func ParseMeCalendarEventExceptionOccurrenceAttachmentID(input string) (*MeCalendarEventExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExceptionOccurrenceAttachmentId{}

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

// ParseMeCalendarEventExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarEventExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventExceptionOccurrenceAttachmentIDInsensitively(input string) (*MeCalendarEventExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExceptionOccurrenceAttachmentId{}

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

// ValidateMeCalendarEventExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a Me Calendar Event Exception Occurrence Attachment ID
func ValidateMeCalendarEventExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Exception Occurrence Attachment ID
func (id MeCalendarEventExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/me/calendar/events/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Exception Occurrence Attachment ID
func (id MeCalendarEventExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Exception Occurrence Attachment ID
func (id MeCalendarEventExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Event Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
