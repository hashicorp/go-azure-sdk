package calendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdEventIdAttachmentId{}

// MeCalendarIdEventIdAttachmentId is a struct representing the Resource ID for a Me Calendar Id Event Id Attachment
type MeCalendarIdEventIdAttachmentId struct {
	CalendarId   string
	EventId      string
	AttachmentId string
}

// NewMeCalendarIdEventIdAttachmentID returns a new MeCalendarIdEventIdAttachmentId struct
func NewMeCalendarIdEventIdAttachmentID(calendarId string, eventId string, attachmentId string) MeCalendarIdEventIdAttachmentId {
	return MeCalendarIdEventIdAttachmentId{
		CalendarId:   calendarId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarIdEventIdAttachmentID parses 'input' into a MeCalendarIdEventIdAttachmentId
func ParseMeCalendarIdEventIdAttachmentID(input string) (*MeCalendarIdEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdEventIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarIdEventIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdEventIdAttachmentIDInsensitively(input string) (*MeCalendarIdEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdEventIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarIdEventIdAttachmentID checks that 'input' can be parsed as a Me Calendar Id Event Id Attachment ID
func ValidateMeCalendarIdEventIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdEventIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Event Id Attachment ID
func (id MeCalendarIdEventIdAttachmentId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Event Id Attachment ID
func (id MeCalendarIdEventIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Event Id Attachment ID
func (id MeCalendarIdEventIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Id Event Id Attachment (%s)", strings.Join(components, "\n"))
}
