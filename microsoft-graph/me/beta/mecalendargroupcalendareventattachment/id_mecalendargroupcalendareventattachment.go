package mecalendargroupcalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarEventAttachmentId{}

// MeCalendarGroupCalendarEventAttachmentId is a struct representing the Resource ID for a Me Calendar Group Calendar Event Attachment
type MeCalendarGroupCalendarEventAttachmentId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	AttachmentId    string
}

// NewMeCalendarGroupCalendarEventAttachmentID returns a new MeCalendarGroupCalendarEventAttachmentId struct
func NewMeCalendarGroupCalendarEventAttachmentID(calendarGroupId string, calendarId string, eventId string, attachmentId string) MeCalendarGroupCalendarEventAttachmentId {
	return MeCalendarGroupCalendarEventAttachmentId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		AttachmentId:    attachmentId,
	}
}

// ParseMeCalendarGroupCalendarEventAttachmentID parses 'input' into a MeCalendarGroupCalendarEventAttachmentId
func ParseMeCalendarGroupCalendarEventAttachmentID(input string) (*MeCalendarGroupCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarEventAttachmentId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupCalendarEventAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarEventAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarEventAttachmentIDInsensitively(input string) (*MeCalendarGroupCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarEventAttachmentId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupCalendarEventAttachmentID checks that 'input' can be parsed as a Me Calendar Group Calendar Event Attachment ID
func ValidateMeCalendarGroupCalendarEventAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarEventAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Event Attachment ID
func (id MeCalendarGroupCalendarEventAttachmentId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Event Attachment ID
func (id MeCalendarGroupCalendarEventAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Event Attachment ID
func (id MeCalendarGroupCalendarEventAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Event Attachment (%s)", strings.Join(components, "\n"))
}
