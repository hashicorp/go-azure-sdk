package calendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{}

// MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId is a struct representing the Resource ID for a Me Calendar View Id Exception Occurrence Id Instance Id Attachment
type MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId struct {
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID returns a new MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId struct
func NewMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID(eventId string, eventId1 string, eventId2 string, attachmentId string) MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId {
	return MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID parses 'input' into a MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId
func ParseMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID(input string) (*MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentIDInsensitively(input string) (*MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID checks that 'input' can be parsed as a Me Calendar View Id Exception Occurrence Id Instance Id Attachment ID
func ValidateMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Exception Occurrence Id Instance Id Attachment ID
func (id MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId) ID() string {
	fmtString := "/me/calendarView/%s/exceptionOccurrences/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Exception Occurrence Id Instance Id Attachment ID
func (id MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Exception Occurrence Id Instance Id Attachment ID
func (id MeCalendarViewIdExceptionOccurrenceIdInstanceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar View Id Exception Occurrence Id Instance Id Attachment (%s)", strings.Join(components, "\n"))
}
