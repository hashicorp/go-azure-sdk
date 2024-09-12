package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdAttachmentId{}

// MeCalendarViewIdAttachmentId is a struct representing the Resource ID for a Me Calendar View Id Attachment
type MeCalendarViewIdAttachmentId struct {
	EventId      string
	AttachmentId string
}

// NewMeCalendarViewIdAttachmentID returns a new MeCalendarViewIdAttachmentId struct
func NewMeCalendarViewIdAttachmentID(eventId string, attachmentId string) MeCalendarViewIdAttachmentId {
	return MeCalendarViewIdAttachmentId{
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarViewIdAttachmentID parses 'input' into a MeCalendarViewIdAttachmentId
func ParseMeCalendarViewIdAttachmentID(input string) (*MeCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdAttachmentIDInsensitively(input string) (*MeCalendarViewIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.AttachmentId, ok = input.Parsed["attachmentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", input)
	}

	return nil
}

// ValidateMeCalendarViewIdAttachmentID checks that 'input' can be parsed as a Me Calendar View Id Attachment ID
func ValidateMeCalendarViewIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Attachment ID
func (id MeCalendarViewIdAttachmentId) ID() string {
	fmtString := "/me/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Attachment ID
func (id MeCalendarViewIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Attachment ID
func (id MeCalendarViewIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar View Id Attachment (%s)", strings.Join(components, "\n"))
}
