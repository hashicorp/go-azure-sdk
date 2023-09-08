package groupcalendarviewinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarViewInstanceExceptionOccurrenceAttachmentId{}

// GroupCalendarViewInstanceExceptionOccurrenceAttachmentId is a struct representing the Resource ID for a Group Calendar View Instance Exception Occurrence Attachment
type GroupCalendarViewInstanceExceptionOccurrenceAttachmentId struct {
	GroupId      string
	EventId      string
	EventId1     string
	EventId2     string
	AttachmentId string
}

// NewGroupCalendarViewInstanceExceptionOccurrenceAttachmentID returns a new GroupCalendarViewInstanceExceptionOccurrenceAttachmentId struct
func NewGroupCalendarViewInstanceExceptionOccurrenceAttachmentID(groupId string, eventId string, eventId1 string, eventId2 string, attachmentId string) GroupCalendarViewInstanceExceptionOccurrenceAttachmentId {
	return GroupCalendarViewInstanceExceptionOccurrenceAttachmentId{
		GroupId:      groupId,
		EventId:      eventId,
		EventId1:     eventId1,
		EventId2:     eventId2,
		AttachmentId: attachmentId,
	}
}

// ParseGroupCalendarViewInstanceExceptionOccurrenceAttachmentID parses 'input' into a GroupCalendarViewInstanceExceptionOccurrenceAttachmentId
func ParseGroupCalendarViewInstanceExceptionOccurrenceAttachmentID(input string) (*GroupCalendarViewInstanceExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewInstanceExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewInstanceExceptionOccurrenceAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarViewInstanceExceptionOccurrenceAttachmentIDInsensitively parses 'input' case-insensitively into a GroupCalendarViewInstanceExceptionOccurrenceAttachmentId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarViewInstanceExceptionOccurrenceAttachmentIDInsensitively(input string) (*GroupCalendarViewInstanceExceptionOccurrenceAttachmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewInstanceExceptionOccurrenceAttachmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewInstanceExceptionOccurrenceAttachmentId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.AttachmentId, ok = parsed.Parsed["attachmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "attachmentId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarViewInstanceExceptionOccurrenceAttachmentID checks that 'input' can be parsed as a Group Calendar View Instance Exception Occurrence Attachment ID
func ValidateGroupCalendarViewInstanceExceptionOccurrenceAttachmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarViewInstanceExceptionOccurrenceAttachmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar View Instance Exception Occurrence Attachment ID
func (id GroupCalendarViewInstanceExceptionOccurrenceAttachmentId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/instances/%s/exceptionOccurrences/%s/attachments/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.EventId2, id.AttachmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar View Instance Exception Occurrence Attachment ID
func (id GroupCalendarViewInstanceExceptionOccurrenceAttachmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticAttachments", "attachments", "attachments"),
		resourceids.UserSpecifiedSegment("attachmentId", "attachmentIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar View Instance Exception Occurrence Attachment ID
func (id GroupCalendarViewInstanceExceptionOccurrenceAttachmentId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Attachment: %q", id.AttachmentId),
	}
	return fmt.Sprintf("Group Calendar View Instance Exception Occurrence Attachment (%s)", strings.Join(components, "\n"))
}
