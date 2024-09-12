package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdInstanceIdAttachmentId{}

// MeCalendarCalendarViewIdInstanceIdAttachmentId is a struct representing the Resource ID for a Me Calendar Calendar View Id Instance Id Attachment
type MeCalendarCalendarViewIdInstanceIdAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarCalendarViewIdInstanceIdAttachmentID returns a new MeCalendarCalendarViewIdInstanceIdAttachmentId struct
func NewMeCalendarCalendarViewIdInstanceIdAttachmentID(eventId string, eventId1 string, attachmentId string) MeCalendarCalendarViewIdInstanceIdAttachmentId {
	return MeCalendarCalendarViewIdInstanceIdAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarCalendarViewIdInstanceIdAttachmentID parses 'input' into a MeCalendarCalendarViewIdInstanceIdAttachmentId
func ParseMeCalendarCalendarViewIdInstanceIdAttachmentID(input string) (*MeCalendarCalendarViewIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdInstanceIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdInstanceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdInstanceIdAttachmentIDInsensitively(input string) (*MeCalendarCalendarViewIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdInstanceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeCalendarCalendarViewIdInstanceIdAttachmentID checks that 'input' can be parsed as a Me Calendar Calendar View Id Instance Id Attachment ID
func ValidateMeCalendarCalendarViewIdInstanceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdInstanceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Instance Id Attachment ID
func (id MeCalendarCalendarViewIdInstanceIdAttachmentId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Instance Id Attachment ID
func (id MeCalendarCalendarViewIdInstanceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Id Instance Id Attachment ID
func (id MeCalendarCalendarViewIdInstanceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Instance Id Attachment (%s)", strings.Join(components, "\n"))
}
