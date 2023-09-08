package mecalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventAttachmentId{}

// MeCalendarEventAttachmentId is a struct representing the Resource ID for a Me Calendar Event Attachment
type MeCalendarEventAttachmentId struct {
	EventId      string
	AttachmentId string
}

// NewMeCalendarEventAttachmentID returns a new MeCalendarEventAttachmentId struct
func NewMeCalendarEventAttachmentID(eventId string, attachmentId string) MeCalendarEventAttachmentId {
	return MeCalendarEventAttachmentId{
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarEventAttachmentID parses 'input' into a MeCalendarEventAttachmentId
func ParseMeCalendarEventAttachmentID(input string) (*MeCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarEventAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarEventAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventAttachmentIDInsensitively(input string) (*MeCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarEventAttachmentID checks that 'input' can be parsed as a Me Calendar Event Attachment ID
func ValidateMeCalendarEventAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Attachment ID
func (id MeCalendarEventAttachmentId) ID() string {
	fmtString := "/me/calendar/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Attachment ID
func (id MeCalendarEventAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Attachment ID
func (id MeCalendarEventAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Event Attachment (%s)", strings.Join(components, "\n"))
}
