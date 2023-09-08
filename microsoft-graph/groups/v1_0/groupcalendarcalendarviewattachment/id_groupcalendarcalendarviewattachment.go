package groupcalendarcalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarViewAttachmentId{}

// GroupCalendarCalendarViewAttachmentId is a struct representing the Resource ID for a Group Calendar Calendar View Attachment
type GroupCalendarCalendarViewAttachmentId struct {
	GroupId      string
	EventId      string
	AttachmentId string
}

// NewGroupCalendarCalendarViewAttachmentID returns a new GroupCalendarCalendarViewAttachmentId struct
func NewGroupCalendarCalendarViewAttachmentID(groupId string, eventId string, attachmentId string) GroupCalendarCalendarViewAttachmentId {
	return GroupCalendarCalendarViewAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseGroupCalendarCalendarViewAttachmentID parses 'input' into a GroupCalendarCalendarViewAttachmentId
func ParseGroupCalendarCalendarViewAttachmentID(input string) (*GroupCalendarCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewAttachmentId{}

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

// ParseGroupCalendarCalendarViewAttachmentIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarViewAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarViewAttachmentIDInsensitively(input string) (*GroupCalendarCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewAttachmentId{}

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

// ValidateGroupCalendarCalendarViewAttachmentID checks that 'input' can be parsed as a Group Calendar Calendar View Attachment ID
func ValidateGroupCalendarCalendarViewAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarViewAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar View Attachment ID
func (id GroupCalendarCalendarViewAttachmentId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar View Attachment ID
func (id GroupCalendarCalendarViewAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar View Attachment ID
func (id GroupCalendarCalendarViewAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Calendar Calendar View Attachment (%s)", strings.Join(components, "\n"))
}
