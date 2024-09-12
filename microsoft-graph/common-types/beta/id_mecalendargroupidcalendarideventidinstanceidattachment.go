package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId{}

// MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Event Id Instance Id Attachment
type MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	AttachmentId    string
}

// NewMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID returns a new MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId struct
func NewMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID(calendarGroupId string, calendarId string, eventId string, eventId1 string, attachmentId string) MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId {
	return MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		AttachmentId:    attachmentId,
	}
}

// ParseMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID parses 'input' into a MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId
func ParseMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID(input string) (*MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Event Id Instance Id Attachment ID
func ValidateMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Event Id Instance Id Attachment ID
func (id MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/events/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Event Id Instance Id Attachment ID
func (id MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Event Id Instance Id Attachment ID
func (id MeCalendarGroupIdCalendarIdEventIdInstanceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Event Id Instance Id Attachment (%s)", strings.Join(components, "\n"))
}
