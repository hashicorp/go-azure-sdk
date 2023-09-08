package groupcalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarEventAttachmentId{}

// GroupCalendarEventAttachmentId is a struct representing the Resource ID for a Group Calendar Event Attachment
type GroupCalendarEventAttachmentId struct {
	GroupId      string
	EventId      string
	AttachmentId string
}

// NewGroupCalendarEventAttachmentID returns a new GroupCalendarEventAttachmentId struct
func NewGroupCalendarEventAttachmentID(groupId string, eventId string, attachmentId string) GroupCalendarEventAttachmentId {
	return GroupCalendarEventAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseGroupCalendarEventAttachmentID parses 'input' into a GroupCalendarEventAttachmentId
func ParseGroupCalendarEventAttachmentID(input string) (*GroupCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarEventAttachmentIDInsensitively parses 'input' case-insensitively into a GroupCalendarEventAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarEventAttachmentIDInsensitively(input string) (*GroupCalendarEventAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarEventAttachmentID checks that 'input' can be parsed as a Group Calendar Event Attachment ID
func ValidateGroupCalendarEventAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarEventAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Event Attachment ID
func (id GroupCalendarEventAttachmentId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Event Attachment ID
func (id GroupCalendarEventAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Event Attachment ID
func (id GroupCalendarEventAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Calendar Event Attachment (%s)", strings.Join(components, "\n"))
}
