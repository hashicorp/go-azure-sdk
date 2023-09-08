package mecalendarcalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewAttachmentId{}

// MeCalendarCalendarViewAttachmentId is a struct representing the Resource ID for a Me Calendar Calendar View Attachment
type MeCalendarCalendarViewAttachmentId struct {
	EventId      string
	AttachmentId string
}

// NewMeCalendarCalendarViewAttachmentID returns a new MeCalendarCalendarViewAttachmentId struct
func NewMeCalendarCalendarViewAttachmentID(eventId string, attachmentId string) MeCalendarCalendarViewAttachmentId {
	return MeCalendarCalendarViewAttachmentId{
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarCalendarViewAttachmentID parses 'input' into a MeCalendarCalendarViewAttachmentId
func ParseMeCalendarCalendarViewAttachmentID(input string) (*MeCalendarCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewAttachmentIDInsensitively(input string) (*MeCalendarCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarCalendarViewAttachmentID checks that 'input' can be parsed as a Me Calendar Calendar View Attachment ID
func ValidateMeCalendarCalendarViewAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Attachment ID
func (id MeCalendarCalendarViewAttachmentId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Attachment ID
func (id MeCalendarCalendarViewAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Attachment ID
func (id MeCalendarCalendarViewAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Attachment (%s)", strings.Join(components, "\n"))
}
