package calendarviewinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdInstanceIdAttachmentId{}

// MeCalendarViewIdInstanceIdAttachmentId is a struct representing the Resource ID for a Me Calendar View Id Instance Id Attachment
type MeCalendarViewIdInstanceIdAttachmentId struct {
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewMeCalendarViewIdInstanceIdAttachmentID returns a new MeCalendarViewIdInstanceIdAttachmentId struct
func NewMeCalendarViewIdInstanceIdAttachmentID(eventId string, eventId1 string, attachmentId string) MeCalendarViewIdInstanceIdAttachmentId {
	return MeCalendarViewIdInstanceIdAttachmentId{
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseMeCalendarViewIdInstanceIdAttachmentID parses 'input' into a MeCalendarViewIdInstanceIdAttachmentId
func ParseMeCalendarViewIdInstanceIdAttachmentID(input string) (*MeCalendarViewIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdInstanceIdAttachmentIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdInstanceIdAttachmentId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdInstanceIdAttachmentIDInsensitively(input string) (*MeCalendarViewIdInstanceIdAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdInstanceIdAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdInstanceIdAttachmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdInstanceIdAttachmentId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarViewIdInstanceIdAttachmentID checks that 'input' can be parsed as a Me Calendar View Id Instance Id Attachment ID
func ValidateMeCalendarViewIdInstanceIdAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdInstanceIdAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Instance Id Attachment ID
func (id MeCalendarViewIdInstanceIdAttachmentId) ID() string {
	fmtString := "/me/calendarView/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Instance Id Attachment ID
func (id MeCalendarViewIdInstanceIdAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("attachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Instance Id Attachment ID
func (id MeCalendarViewIdInstanceIdAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Me Calendar View Id Instance Id Attachment (%s)", strings.Join(components, "\n"))
}
