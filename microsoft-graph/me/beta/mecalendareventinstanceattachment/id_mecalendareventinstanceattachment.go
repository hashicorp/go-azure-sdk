package mecalendareventinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventInstanceAttachmentId{}

// MeCalendarEventInstanceAttachmentId is a struct representing the Resource ID for a Me Calendar Event Instance Attachment
type MeCalendarEventInstanceAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarEventInstanceAttachmentID returns a new MeCalendarEventInstanceAttachmentId struct
func NewMeCalendarEventInstanceAttachmentID(eventId string, eventId1 string, attachmentId string) MeCalendarEventInstanceAttachmentId {
	return MeCalendarEventInstanceAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarEventInstanceAttachmentID parses 'input' into a MeCalendarEventInstanceAttachmentId
func ParseMeCalendarEventInstanceAttachmentID(input string) (*MeCalendarEventInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventInstanceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventInstanceAttachmentId{}

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

// ParseMeCalendarEventInstanceAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarEventInstanceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventInstanceAttachmentIDInsensitively(input string) (*MeCalendarEventInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventInstanceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventInstanceAttachmentId{}

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

// ValidateMeCalendarEventInstanceAttachmentID checks that 'input' can be parsed as a Me Calendar Event Instance Attachment ID
func ValidateMeCalendarEventInstanceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventInstanceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Instance Attachment ID
func (id MeCalendarEventInstanceAttachmentId) ID() string {
	fmtString := "/me/calendar/events/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Instance Attachment ID
func (id MeCalendarEventInstanceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Instance Attachment ID
func (id MeCalendarEventInstanceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Event Instance Attachment (%s)", strings.Join(components, "\n"))
}
