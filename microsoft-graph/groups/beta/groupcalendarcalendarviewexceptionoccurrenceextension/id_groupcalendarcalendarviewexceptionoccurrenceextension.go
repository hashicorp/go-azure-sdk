package groupcalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarViewExceptionOccurrenceExtensionId{}

// GroupCalendarCalendarViewExceptionOccurrenceExtensionId is a struct representing the Resource ID for a Group Calendar Calendar View Exception Occurrence Extension
type GroupCalendarCalendarViewExceptionOccurrenceExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewGroupCalendarCalendarViewExceptionOccurrenceExtensionID returns a new GroupCalendarCalendarViewExceptionOccurrenceExtensionId struct
func NewGroupCalendarCalendarViewExceptionOccurrenceExtensionID(groupId string, eventId string, eventId1 string, extensionId string) GroupCalendarCalendarViewExceptionOccurrenceExtensionId {
	return GroupCalendarCalendarViewExceptionOccurrenceExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseGroupCalendarCalendarViewExceptionOccurrenceExtensionID parses 'input' into a GroupCalendarCalendarViewExceptionOccurrenceExtensionId
func ParseGroupCalendarCalendarViewExceptionOccurrenceExtensionID(input string) (*GroupCalendarCalendarViewExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewExceptionOccurrenceExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarCalendarViewExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarViewExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarViewExceptionOccurrenceExtensionIDInsensitively(input string) (*GroupCalendarCalendarViewExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarViewExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarViewExceptionOccurrenceExtensionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarCalendarViewExceptionOccurrenceExtensionID checks that 'input' can be parsed as a Group Calendar Calendar View Exception Occurrence Extension ID
func ValidateGroupCalendarCalendarViewExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarViewExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar View Exception Occurrence Extension ID
func (id GroupCalendarCalendarViewExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/groups/%s/calendar/calendarView/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar View Exception Occurrence Extension ID
func (id GroupCalendarCalendarViewExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar View Exception Occurrence Extension ID
func (id GroupCalendarCalendarViewExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Calendar Calendar View Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
