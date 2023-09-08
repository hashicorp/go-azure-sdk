package mecalendareventinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventInstanceId{}

// MeCalendarEventInstanceId is a struct representing the Resource ID for a Me Calendar Event Instance
type MeCalendarEventInstanceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarEventInstanceID returns a new MeCalendarEventInstanceId struct
func NewMeCalendarEventInstanceID(eventId string, eventId1 string) MeCalendarEventInstanceId {
	return MeCalendarEventInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarEventInstanceID parses 'input' into a MeCalendarEventInstanceId
func ParseMeCalendarEventInstanceID(input string) (*MeCalendarEventInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventInstanceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarEventInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarEventInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventInstanceIDInsensitively(input string) (*MeCalendarEventInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventInstanceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarEventInstanceID checks that 'input' can be parsed as a Me Calendar Event Instance ID
func ValidateMeCalendarEventInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Instance ID
func (id MeCalendarEventInstanceId) ID() string {
	fmtString := "/me/calendar/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Instance ID
func (id MeCalendarEventInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Event Instance ID
func (id MeCalendarEventInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Event Instance (%s)", strings.Join(components, "\n"))
}
