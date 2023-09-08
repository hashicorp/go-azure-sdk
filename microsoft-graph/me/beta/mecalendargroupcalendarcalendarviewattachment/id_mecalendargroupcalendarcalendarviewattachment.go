package mecalendargroupcalendarcalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarCalendarViewAttachmentId{}

// MeCalendarGroupCalendarCalendarViewAttachmentId is a struct representing the Resource ID for a Me Calendar Group Calendar Calendar View Attachment
type MeCalendarGroupCalendarCalendarViewAttachmentId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	AttachmentId    string
}

// NewMeCalendarGroupCalendarCalendarViewAttachmentID returns a new MeCalendarGroupCalendarCalendarViewAttachmentId struct
func NewMeCalendarGroupCalendarCalendarViewAttachmentID(calendarGroupId string, calendarId string, eventId string, attachmentId string) MeCalendarGroupCalendarCalendarViewAttachmentId {
	return MeCalendarGroupCalendarCalendarViewAttachmentId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		AttachmentId:    attachmentId,
	}
}

// ParseMeCalendarGroupCalendarCalendarViewAttachmentID parses 'input' into a MeCalendarGroupCalendarCalendarViewAttachmentId
func ParseMeCalendarGroupCalendarCalendarViewAttachmentID(input string) (*MeCalendarGroupCalendarCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewAttachmentId{}

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

// ParseMeCalendarGroupCalendarCalendarViewAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarCalendarViewAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarCalendarViewAttachmentIDInsensitively(input string) (*MeCalendarGroupCalendarCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewAttachmentId{}

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

// ValidateMeCalendarGroupCalendarCalendarViewAttachmentID checks that 'input' can be parsed as a Me Calendar Group Calendar Calendar View Attachment ID
func ValidateMeCalendarGroupCalendarCalendarViewAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarCalendarViewAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Calendar View Attachment ID
func (id MeCalendarGroupCalendarCalendarViewAttachmentId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Calendar View Attachment ID
func (id MeCalendarGroupCalendarCalendarViewAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Calendar View Attachment ID
func (id MeCalendarGroupCalendarCalendarViewAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Calendar View Attachment (%s)", strings.Join(components, "\n"))
}
