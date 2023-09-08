package mecalendarviewinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewInstanceId{}

// MeCalendarViewInstanceId is a struct representing the Resource ID for a Me Calendar View Instance
type MeCalendarViewInstanceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarViewInstanceID returns a new MeCalendarViewInstanceId struct
func NewMeCalendarViewInstanceID(eventId string, eventId1 string) MeCalendarViewInstanceId {
	return MeCalendarViewInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarViewInstanceID parses 'input' into a MeCalendarViewInstanceId
func ParseMeCalendarViewInstanceID(input string) (*MeCalendarViewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewInstanceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarViewInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarViewInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewInstanceIDInsensitively(input string) (*MeCalendarViewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewInstanceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarViewInstanceID checks that 'input' can be parsed as a Me Calendar View Instance ID
func ValidateMeCalendarViewInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Instance ID
func (id MeCalendarViewInstanceId) ID() string {
	fmtString := "/me/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Instance ID
func (id MeCalendarViewInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar View Instance ID
func (id MeCalendarViewInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar View Instance (%s)", strings.Join(components, "\n"))
}
