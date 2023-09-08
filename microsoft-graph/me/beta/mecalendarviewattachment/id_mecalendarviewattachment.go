package mecalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewAttachmentId{}

// MeCalendarViewAttachmentId is a struct representing the Resource ID for a Me Calendar View Attachment
type MeCalendarViewAttachmentId struct {
	EventId      string
	AttachmentId string
}

// NewMeCalendarViewAttachmentID returns a new MeCalendarViewAttachmentId struct
func NewMeCalendarViewAttachmentID(eventId string, attachmentId string) MeCalendarViewAttachmentId {
	return MeCalendarViewAttachmentId{
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarViewAttachmentID parses 'input' into a MeCalendarViewAttachmentId
func ParseMeCalendarViewAttachmentID(input string) (*MeCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarViewAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarViewAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewAttachmentIDInsensitively(input string) (*MeCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewAttachmentId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarViewAttachmentID checks that 'input' can be parsed as a Me Calendar View Attachment ID
func ValidateMeCalendarViewAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Attachment ID
func (id MeCalendarViewAttachmentId) ID() string {
	fmtString := "/me/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Attachment ID
func (id MeCalendarViewAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Attachment ID
func (id MeCalendarViewAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar View Attachment (%s)", strings.Join(components, "\n"))
}
