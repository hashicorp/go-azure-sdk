package groupcalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarEventId{}

// GroupCalendarEventId is a struct representing the Resource ID for a Group Calendar Event
type GroupCalendarEventId struct {
	GroupId string
	EventId string
}

// NewGroupCalendarEventID returns a new GroupCalendarEventId struct
func NewGroupCalendarEventID(groupId string, eventId string) GroupCalendarEventId {
	return GroupCalendarEventId{
		GroupId: groupId,
		EventId: eventId,
	}
}

// ParseGroupCalendarEventID parses 'input' into a GroupCalendarEventId
func ParseGroupCalendarEventID(input string) (*GroupCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarEventIDInsensitively parses 'input' case-insensitively into a GroupCalendarEventId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarEventIDInsensitively(input string) (*GroupCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarEventId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarEventID checks that 'input' can be parsed as a Group Calendar Event ID
func ValidateGroupCalendarEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Event ID
func (id GroupCalendarEventId) ID() string {
	fmtString := "/groups/%s/calendar/events/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Event ID
func (id GroupCalendarEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Event ID
func (id GroupCalendarEventId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Group Calendar Event (%s)", strings.Join(components, "\n"))
}
