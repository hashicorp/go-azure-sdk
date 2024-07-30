package calendareventexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarEventIdExceptionOccurrenceIdAttachmentId{}

// MeCalendarEventIdExceptionOccurrenceIdAttachmentId is a struct representing the Resource ID for a Me Calendar Event Id Exception Occurrence Id Attachment
type MeCalendarEventIdExceptionOccurrenceIdAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarEventIdExceptionOccurrenceIdAttachmentID returns a new MeCalendarEventIdExceptionOccurrenceIdAttachmentId struct
func NewMeCalendarEventIdExceptionOccurrenceIdAttachmentID(eventId string, eventId1 string, attachmentId string) MeCalendarEventIdExceptionOccurrenceIdAttachmentId {
	return MeCalendarEventIdExceptionOccurrenceIdAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarEventIdExceptionOccurrenceIdAttachmentID parses 'input' into a MeCalendarEventIdExceptionOccurrenceIdAttachmentId
func ParseMeCalendarEventIdExceptionOccurrenceIdAttachmentID(input string) (*MeCalendarEventIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarEventIdExceptionOccurrenceIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarEventIdExceptionOccurrenceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventIdExceptionOccurrenceIdAttachmentIDInsensitively(input string) (*MeCalendarEventIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarEventIdExceptionOccurrenceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeCalendarEventIdExceptionOccurrenceIdAttachmentID checks that 'input' can be parsed as a Me Calendar Event Id Exception Occurrence Id Attachment ID
func ValidateMeCalendarEventIdExceptionOccurrenceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventIdExceptionOccurrenceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Id Exception Occurrence Id Attachment ID
func (id MeCalendarEventIdExceptionOccurrenceIdAttachmentId) ID() string {
	fmtString := "/me/calendar/events/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Id Exception Occurrence Id Attachment ID
func (id MeCalendarEventIdExceptionOccurrenceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Id Exception Occurrence Id Attachment ID
func (id MeCalendarEventIdExceptionOccurrenceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Event Id Exception Occurrence Id Attachment (%s)", strings.Join(components, "\n"))
}
