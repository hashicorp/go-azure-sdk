package groupcalendarviewexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarViewExceptionOccurrenceExtensionId{}

// GroupCalendarViewExceptionOccurrenceExtensionId is a struct representing the Resource ID for a Group Calendar View Exception Occurrence Extension
type GroupCalendarViewExceptionOccurrenceExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewGroupCalendarViewExceptionOccurrenceExtensionID returns a new GroupCalendarViewExceptionOccurrenceExtensionId struct
func NewGroupCalendarViewExceptionOccurrenceExtensionID(groupId string, eventId string, eventId1 string, extensionId string) GroupCalendarViewExceptionOccurrenceExtensionId {
	return GroupCalendarViewExceptionOccurrenceExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseGroupCalendarViewExceptionOccurrenceExtensionID parses 'input' into a GroupCalendarViewExceptionOccurrenceExtensionId
func ParseGroupCalendarViewExceptionOccurrenceExtensionID(input string) (*GroupCalendarViewExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewExceptionOccurrenceExtensionId{}

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

// ParseGroupCalendarViewExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a GroupCalendarViewExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarViewExceptionOccurrenceExtensionIDInsensitively(input string) (*GroupCalendarViewExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarViewExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarViewExceptionOccurrenceExtensionId{}

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

// ValidateGroupCalendarViewExceptionOccurrenceExtensionID checks that 'input' can be parsed as a Group Calendar View Exception Occurrence Extension ID
func ValidateGroupCalendarViewExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarViewExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar View Exception Occurrence Extension ID
func (id GroupCalendarViewExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/groups/%s/calendarView/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar View Exception Occurrence Extension ID
func (id GroupCalendarViewExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar View Exception Occurrence Extension ID
func (id GroupCalendarViewExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Calendar View Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
