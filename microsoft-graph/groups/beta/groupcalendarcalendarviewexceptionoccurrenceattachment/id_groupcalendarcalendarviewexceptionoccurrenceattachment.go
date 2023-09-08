package groupcalendarcalendarviewexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarViewExceptionOccurrenceAttachmentId{}

// GroupCalendarCalendarViewExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a Group Calendar Calendar View Exception Occurrence Attachment
type GroupCalendarCalendarViewExceptionOccurrenceAttachmentId struct {
	GroupId      string
	EventId      string
	EventId1     string
	AttachmentId string
}

// NewGroupCalendarCalendarViewExceptionOccurrenceAttachmentID returns a new GroupCalendarCalendarViewExceptionOccurrenceAttachmentId struct
func NewGroupCalendarCalendarViewExceptionOccurrenceAttachmentID(groupId string, eventId string, eventId1 string, attachmentId string) GroupCalendarCalendarViewExceptionOccurrenceAttachmentId {
	return GroupCalendarCalendarViewExceptionOccurrenceAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		EventId1:     eventId1,
		AttachmentId: attachmentId,
	}
}

// ParseGroupCalendarCalendarViewExceptionOccurrenceAttachmentID parses 'input' into a GroupCalendarCalendarViewExceptionOccurrenceAttachmentId
func ParseGroupCalendarCalendarViewExceptionOccurrenceAttachmentID(input string) (*GroupCalendarCalendarViewExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewExceptionOccurrenceAttachmentId{}

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

// ParseGroupCalendarCalendarViewExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarViewExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarViewExceptionOccurrenceAttachmentIDInsensitively(input string) (*GroupCalendarCalendarViewExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewExceptionOccurrenceAttachmentId{}

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

// ValidateGroupCalendarCalendarViewExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a Group Calendar Calendar View Exception Occurrence Attachment ID
func ValidateGroupCalendarCalendarViewExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarViewExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar View Exception Occurrence Attachment ID
func (id GroupCalendarCalendarViewExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar View Exception Occurrence Attachment ID
func (id GroupCalendarCalendarViewExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar View Exception Occurrence Attachment ID
func (id GroupCalendarCalendarViewExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Calendar Calendar View Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
