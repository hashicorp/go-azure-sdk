package calendarcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{}

// MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId is a struct representing the Resource ID for a Me Calendar Calendar View Id Instance Id Exception Occurrence Id Attachment
type MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId struct {
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID returns a new MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId struct
func NewMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(eventId string, eventId1 string, eventId2 string, attachmentId string) MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId {
	return MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID parses 'input' into a MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId
func ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(input string) (*MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentIDInsensitively(input string) (*MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID checks that 'input' can be parsed as a Me Calendar Calendar View Id Instance Id Exception Occurrence Id Attachment ID
func ValidateMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Instance Id Exception Occurrence Id Attachment ID
func (id MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/instances/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Instance Id Exception Occurrence Id Attachment ID
func (id MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Id Instance Id Exception Occurrence Id Attachment ID
func (id MeCalendarCalendarViewIdInstanceIdExceptionOccurrenceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Instance Id Exception Occurrence Id Attachment (%s)", strings.Join(components, "\n"))
}
