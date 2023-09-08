package mecalendareventexceptionoccurrenceinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventExceptionOccurrenceInstanceAttachmentId{}

// MeCalendarEventExceptionOccurrenceInstanceAttachmentId is a struct representing the Resource ID for a Me Calendar Event Exception Occurrence Instance Attachment
type MeCalendarEventExceptionOccurrenceInstanceAttachmentId struct {
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewMeCalendarEventExceptionOccurrenceInstanceAttachmentID returns a new MeCalendarEventExceptionOccurrenceInstanceAttachmentId struct
func NewMeCalendarEventExceptionOccurrenceInstanceAttachmentID(eventId string, eventId1 string, eventId2 string, attachmentId string) MeCalendarEventExceptionOccurrenceInstanceAttachmentId {
	return MeCalendarEventExceptionOccurrenceInstanceAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarEventExceptionOccurrenceInstanceAttachmentID parses 'input' into a MeCalendarEventExceptionOccurrenceInstanceAttachmentId
func ParseMeCalendarEventExceptionOccurrenceInstanceAttachmentID(input string) (*MeCalendarEventExceptionOccurrenceInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExceptionOccurrenceInstanceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExceptionOccurrenceInstanceAttachmentId{}

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

// ParseMeCalendarEventExceptionOccurrenceInstanceAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarEventExceptionOccurrenceInstanceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventExceptionOccurrenceInstanceAttachmentIDInsensitively(input string) (*MeCalendarEventExceptionOccurrenceInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExceptionOccurrenceInstanceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExceptionOccurrenceInstanceAttachmentId{}

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

// ValidateMeCalendarEventExceptionOccurrenceInstanceAttachmentID checks that 'input' can be parsed as a Me Calendar Event Exception Occurrence Instance Attachment ID
func ValidateMeCalendarEventExceptionOccurrenceInstanceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventExceptionOccurrenceInstanceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Exception Occurrence Instance Attachment ID
func (id MeCalendarEventExceptionOccurrenceInstanceAttachmentId) ID() string {
	fmtString := "/me/calendar/events/%s/exceptionOccurrences/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Exception Occurrence Instance Attachment ID
func (id MeCalendarEventExceptionOccurrenceInstanceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Exception Occurrence Instance Attachment ID
func (id MeCalendarEventExceptionOccurrenceInstanceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Event Exception Occurrence Instance Attachment (%s)", strings.Join(components, "\n"))
}
