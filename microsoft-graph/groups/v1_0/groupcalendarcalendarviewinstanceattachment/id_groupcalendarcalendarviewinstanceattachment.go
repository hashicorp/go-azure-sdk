package groupcalendarcalendarviewinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarViewInstanceAttachmentId{}

// GroupCalendarCalendarViewInstanceAttachmentId is a struct representing the Resource ID for a Group Calendar Calendar View Instance Attachment
type GroupCalendarCalendarViewInstanceAttachmentId struct {
	GroupId      string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewGroupCalendarCalendarViewInstanceAttachmentID returns a new GroupCalendarCalendarViewInstanceAttachmentId struct
func NewGroupCalendarCalendarViewInstanceAttachmentID(groupId string, eventId string, eventId1 string, attachmentId string) GroupCalendarCalendarViewInstanceAttachmentId {
	return GroupCalendarCalendarViewInstanceAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseGroupCalendarCalendarViewInstanceAttachmentID parses 'input' into a GroupCalendarCalendarViewInstanceAttachmentId
func ParseGroupCalendarCalendarViewInstanceAttachmentID(input string) (*GroupCalendarCalendarViewInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewInstanceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewInstanceAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ParseGroupCalendarCalendarViewInstanceAttachmentIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarViewInstanceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarViewInstanceAttachmentIDInsensitively(input string) (*GroupCalendarCalendarViewInstanceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewInstanceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewInstanceAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

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

// ValidateGroupCalendarCalendarViewInstanceAttachmentID checks that 'input' can be parsed as a Group Calendar Calendar View Instance Attachment ID
func ValidateGroupCalendarCalendarViewInstanceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarViewInstanceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar View Instance Attachment ID
func (id GroupCalendarCalendarViewInstanceAttachmentId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/instances/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar View Instance Attachment ID
func (id GroupCalendarCalendarViewInstanceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar View Instance Attachment ID
func (id GroupCalendarCalendarViewInstanceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Calendar Calendar View Instance Attachment (%s)", strings.Join(components, "\n"))
}
