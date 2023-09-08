package groupcalendareventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarEventExceptionOccurrenceExtensionId{}

// GroupCalendarEventExceptionOccurrenceExtensionId is a struct representing the Resource ID for a Group Calendar Event Exception Occurrence Extension
type GroupCalendarEventExceptionOccurrenceExtensionId struct {
	GroupId     string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewGroupCalendarEventExceptionOccurrenceExtensionID returns a new GroupCalendarEventExceptionOccurrenceExtensionId struct
func NewGroupCalendarEventExceptionOccurrenceExtensionID(groupId string, eventId string, eventId1 string, extensionId string) GroupCalendarEventExceptionOccurrenceExtensionId {
	return GroupCalendarEventExceptionOccurrenceExtensionId{
		GroupId:     groupId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseGroupCalendarEventExceptionOccurrenceExtensionID parses 'input' into a GroupCalendarEventExceptionOccurrenceExtensionId
func ParseGroupCalendarEventExceptionOccurrenceExtensionID(input string) (*GroupCalendarEventExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventExceptionOccurrenceExtensionId{}

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

// ParseGroupCalendarEventExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a GroupCalendarEventExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarEventExceptionOccurrenceExtensionIDInsensitively(input string) (*GroupCalendarEventExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventExceptionOccurrenceExtensionId{}

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

// ValidateGroupCalendarEventExceptionOccurrenceExtensionID checks that 'input' can be parsed as a Group Calendar Event Exception Occurrence Extension ID
func ValidateGroupCalendarEventExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarEventExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Event Exception Occurrence Extension ID
func (id GroupCalendarEventExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Event Exception Occurrence Extension ID
func (id GroupCalendarEventExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Event Exception Occurrence Extension ID
func (id GroupCalendarEventExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Group Calendar Event Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
