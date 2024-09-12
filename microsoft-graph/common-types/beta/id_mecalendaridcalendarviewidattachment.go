package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarViewIdAttachmentId{}

// MeCalendarIdCalendarViewIdAttachmentId is a struct representing the Resource ID for a Me Calendar Id Calendar View Id Attachment
type MeCalendarIdCalendarViewIdAttachmentId struct {
	CalendarId   string
	EventId      string
	AttachmentId string
}

// NewMeCalendarIdCalendarViewIdAttachmentID returns a new MeCalendarIdCalendarViewIdAttachmentId struct
func NewMeCalendarIdCalendarViewIdAttachmentID(calendarId string, eventId string, attachmentId string) MeCalendarIdCalendarViewIdAttachmentId {
	return MeCalendarIdCalendarViewIdAttachmentId{
		CalendarId:   calendarId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarIdCalendarViewIdAttachmentID parses 'input' into a MeCalendarIdCalendarViewIdAttachmentId
func ParseMeCalendarIdCalendarViewIdAttachmentID(input string) (*MeCalendarIdCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdCalendarViewIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarIdCalendarViewIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdCalendarViewIdAttachmentIDInsensitively(input string) (*MeCalendarIdCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdCalendarViewIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeCalendarIdCalendarViewIdAttachmentID checks that 'input' can be parsed as a Me Calendar Id Calendar View Id Attachment ID
func ValidateMeCalendarIdCalendarViewIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdCalendarViewIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Calendar View Id Attachment ID
func (id MeCalendarIdCalendarViewIdAttachmentId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Calendar View Id Attachment ID
func (id MeCalendarIdCalendarViewIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Calendar View Id Attachment ID
func (id MeCalendarIdCalendarViewIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Id Calendar View Id Attachment (%s)", strings.Join(components, "\n"))
}
