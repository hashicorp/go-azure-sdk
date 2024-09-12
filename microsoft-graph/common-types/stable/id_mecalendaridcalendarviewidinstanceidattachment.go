package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarViewIdInstanceIdAttachmentId{}

// MeCalendarIdCalendarViewIdInstanceIdAttachmentId is a struct representing the Resource ID for a Me Calendar Id Calendar View Id Instance Id Attachment
type MeCalendarIdCalendarViewIdInstanceIdAttachmentId struct {
	CalendarId   string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarIdCalendarViewIdInstanceIdAttachmentID returns a new MeCalendarIdCalendarViewIdInstanceIdAttachmentId struct
func NewMeCalendarIdCalendarViewIdInstanceIdAttachmentID(calendarId string, eventId string, eventId1 string, attachmentId string) MeCalendarIdCalendarViewIdInstanceIdAttachmentId {
	return MeCalendarIdCalendarViewIdInstanceIdAttachmentId{
		CalendarId:   calendarId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarIdCalendarViewIdInstanceIdAttachmentID parses 'input' into a MeCalendarIdCalendarViewIdInstanceIdAttachmentId
func ParseMeCalendarIdCalendarViewIdInstanceIdAttachmentID(input string) (*MeCalendarIdCalendarViewIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdCalendarViewIdInstanceIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarIdCalendarViewIdInstanceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdCalendarViewIdInstanceIdAttachmentIDInsensitively(input string) (*MeCalendarIdCalendarViewIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdCalendarViewIdInstanceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

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

// ValidateMeCalendarIdCalendarViewIdInstanceIdAttachmentID checks that 'input' can be parsed as a Me Calendar Id Calendar View Id Instance Id Attachment ID
func ValidateMeCalendarIdCalendarViewIdInstanceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdCalendarViewIdInstanceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Calendar View Id Instance Id Attachment ID
func (id MeCalendarIdCalendarViewIdInstanceIdAttachmentId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Calendar View Id Instance Id Attachment ID
func (id MeCalendarIdCalendarViewIdInstanceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Calendar View Id Instance Id Attachment ID
func (id MeCalendarIdCalendarViewIdInstanceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Id Calendar View Id Instance Id Attachment (%s)", strings.Join(components, "\n"))
}
