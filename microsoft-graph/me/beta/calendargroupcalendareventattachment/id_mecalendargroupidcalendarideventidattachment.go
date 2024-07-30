package calendargroupcalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdEventIdAttachmentId{}

// MeCalendarGroupIdCalendarIdEventIdAttachmentId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Event Id Attachment
type MeCalendarGroupIdCalendarIdEventIdAttachmentId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	AttachmentId    string
}

// NewMeCalendarGroupIdCalendarIdEventIdAttachmentID returns a new MeCalendarGroupIdCalendarIdEventIdAttachmentId struct
func NewMeCalendarGroupIdCalendarIdEventIdAttachmentID(calendarGroupId string, calendarId string, eventId string, attachmentId string) MeCalendarGroupIdCalendarIdEventIdAttachmentId {
	return MeCalendarGroupIdCalendarIdEventIdAttachmentId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		AttachmentId:    attachmentId,
	}
}

// ParseMeCalendarGroupIdCalendarIdEventIdAttachmentID parses 'input' into a MeCalendarGroupIdCalendarIdEventIdAttachmentId
func ParseMeCalendarGroupIdCalendarIdEventIdAttachmentID(input string) (*MeCalendarGroupIdCalendarIdEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdEventIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdEventIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdEventIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdEventIdAttachmentIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdEventIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdEventIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarGroupId, ok = input.Parsed["calendarGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", input)
	}

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

// ValidateMeCalendarGroupIdCalendarIdEventIdAttachmentID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Event Id Attachment ID
func ValidateMeCalendarGroupIdCalendarIdEventIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdEventIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Event Id Attachment ID
func (id MeCalendarGroupIdCalendarIdEventIdAttachmentId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Event Id Attachment ID
func (id MeCalendarGroupIdCalendarIdEventIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Event Id Attachment ID
func (id MeCalendarGroupIdCalendarIdEventIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Event Id Attachment (%s)", strings.Join(components, "\n"))
}
