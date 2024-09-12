package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarEventIdAttachmentId{}

// MeCalendarEventIdAttachmentId is a struct representing the Resource ID for a Me Calendar Event Id Attachment
type MeCalendarEventIdAttachmentId struct {
	EventId      string
	AttachmentId string
}

// NewMeCalendarEventIdAttachmentID returns a new MeCalendarEventIdAttachmentId struct
func NewMeCalendarEventIdAttachmentID(eventId string, attachmentId string) MeCalendarEventIdAttachmentId {
	return MeCalendarEventIdAttachmentId{
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarEventIdAttachmentID parses 'input' into a MeCalendarEventIdAttachmentId
func ParseMeCalendarEventIdAttachmentID(input string) (*MeCalendarEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarEventIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarEventIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventIdAttachmentIDInsensitively(input string) (*MeCalendarEventIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarEventIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeCalendarEventIdAttachmentID checks that 'input' can be parsed as a Me Calendar Event Id Attachment ID
func ValidateMeCalendarEventIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Id Attachment ID
func (id MeCalendarEventIdAttachmentId) ID() string {
	fmtString := "/me/calendar/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Id Attachment ID
func (id MeCalendarEventIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Id Attachment ID
func (id MeCalendarEventIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Event Id Attachment (%s)", strings.Join(components, "\n"))
}
