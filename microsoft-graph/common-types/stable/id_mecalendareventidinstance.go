package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarEventIdInstanceId{}

// MeCalendarEventIdInstanceId is a struct representing the Resource ID for a Me Calendar Event Id Instance
type MeCalendarEventIdInstanceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarEventIdInstanceID returns a new MeCalendarEventIdInstanceId struct
func NewMeCalendarEventIdInstanceID(eventId string, eventId1 string) MeCalendarEventIdInstanceId {
	return MeCalendarEventIdInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarEventIdInstanceID parses 'input' into a MeCalendarEventIdInstanceId
func ParseMeCalendarEventIdInstanceID(input string) (*MeCalendarEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarEventIdInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarEventIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventIdInstanceIDInsensitively(input string) (*MeCalendarEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarEventIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeCalendarEventIdInstanceID checks that 'input' can be parsed as a Me Calendar Event Id Instance ID
func ValidateMeCalendarEventIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Id Instance ID
func (id MeCalendarEventIdInstanceId) ID() string {
	fmtString := "/me/calendar/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Id Instance ID
func (id MeCalendarEventIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Event Id Instance ID
func (id MeCalendarEventIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Event Id Instance (%s)", strings.Join(components, "\n"))
}
