package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdAttachmentId{}

// MeCalendarCalendarViewIdAttachmentId is a struct representing the Resource ID for a Me Calendar Calendar View Id Attachment
type MeCalendarCalendarViewIdAttachmentId struct {
	EventId      string
	AttachmentId string
}

// NewMeCalendarCalendarViewIdAttachmentID returns a new MeCalendarCalendarViewIdAttachmentId struct
func NewMeCalendarCalendarViewIdAttachmentID(eventId string, attachmentId string) MeCalendarCalendarViewIdAttachmentId {
	return MeCalendarCalendarViewIdAttachmentId{
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarCalendarViewIdAttachmentID parses 'input' into a MeCalendarCalendarViewIdAttachmentId
func ParseMeCalendarCalendarViewIdAttachmentID(input string) (*MeCalendarCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdAttachmentIDInsensitively(input string) (*MeCalendarCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeCalendarCalendarViewIdAttachmentID checks that 'input' can be parsed as a Me Calendar Calendar View Id Attachment ID
func ValidateMeCalendarCalendarViewIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Attachment ID
func (id MeCalendarCalendarViewIdAttachmentId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Attachment ID
func (id MeCalendarCalendarViewIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Id Attachment ID
func (id MeCalendarCalendarViewIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Attachment (%s)", strings.Join(components, "\n"))
}
