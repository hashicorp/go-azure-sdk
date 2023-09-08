package mecalendarcalendarviewinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewInstanceAttachmentId{}

// MeCalendarCalendarViewInstanceAttachmentId is a struct representing the Resource ID for a Me Calendar Calendar View Instance Attachment
type MeCalendarCalendarViewInstanceAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarCalendarViewInstanceAttachmentID returns a new MeCalendarCalendarViewInstanceAttachmentId struct
func NewMeCalendarCalendarViewInstanceAttachmentID(eventId string, eventId1 string, attachmentId string) MeCalendarCalendarViewInstanceAttachmentId {
	return MeCalendarCalendarViewInstanceAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarCalendarViewInstanceAttachmentID parses 'input' into a MeCalendarCalendarViewInstanceAttachmentId
func ParseMeCalendarCalendarViewInstanceAttachmentID(input string) (*MeCalendarCalendarViewInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewInstanceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewInstanceAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewInstanceAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewInstanceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewInstanceAttachmentIDInsensitively(input string) (*MeCalendarCalendarViewInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewInstanceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewInstanceAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarCalendarViewInstanceAttachmentID checks that 'input' can be parsed as a Me Calendar Calendar View Instance Attachment ID
func ValidateMeCalendarCalendarViewInstanceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewInstanceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Instance Attachment ID
func (id MeCalendarCalendarViewInstanceAttachmentId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Instance Attachment ID
func (id MeCalendarCalendarViewInstanceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Instance Attachment ID
func (id MeCalendarCalendarViewInstanceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Instance Attachment (%s)", strings.Join(components, "\n"))
}
