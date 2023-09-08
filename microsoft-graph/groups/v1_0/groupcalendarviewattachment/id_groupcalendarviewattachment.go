package groupcalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarViewAttachmentId{}

// GroupCalendarViewAttachmentId is a struct representing the Resource ID for a Group Calendar View Attachment
type GroupCalendarViewAttachmentId struct {
	GroupId      string
	EventId      string
	AttachmentId string
}

// NewGroupCalendarViewAttachmentID returns a new GroupCalendarViewAttachmentId struct
func NewGroupCalendarViewAttachmentID(groupId string, eventId string, attachmentId string) GroupCalendarViewAttachmentId {
	return GroupCalendarViewAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		AttachmentId: attachmentId,
	}
}

// ParseGroupCalendarViewAttachmentID parses 'input' into a GroupCalendarViewAttachmentId
func ParseGroupCalendarViewAttachmentID(input string) (*GroupCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewAttachmentId{}

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

// ParseGroupCalendarViewAttachmentIDInsensitively parses 'input' case-insensitively into a GroupCalendarViewAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarViewAttachmentIDInsensitively(input string) (*GroupCalendarViewAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewAttachmentId{}

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

// ValidateGroupCalendarViewAttachmentID checks that 'input' can be parsed as a Group Calendar View Attachment ID
func ValidateGroupCalendarViewAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarViewAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar View Attachment ID
func (id GroupCalendarViewAttachmentId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar View Attachment ID
func (id GroupCalendarViewAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar View Attachment ID
func (id GroupCalendarViewAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Calendar View Attachment (%s)", strings.Join(components, "\n"))
}
