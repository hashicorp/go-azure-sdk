package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdInstanceId{}

// MeCalendarCalendarViewIdInstanceId is a struct representing the Resource ID for a Me Calendar Calendar View Id Instance
type MeCalendarCalendarViewIdInstanceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarCalendarViewIdInstanceID returns a new MeCalendarCalendarViewIdInstanceId struct
func NewMeCalendarCalendarViewIdInstanceID(eventId string, eventId1 string) MeCalendarCalendarViewIdInstanceId {
	return MeCalendarCalendarViewIdInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarCalendarViewIdInstanceID parses 'input' into a MeCalendarCalendarViewIdInstanceId
func ParseMeCalendarCalendarViewIdInstanceID(input string) (*MeCalendarCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdInstanceIDInsensitively(input string) (*MeCalendarCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeCalendarCalendarViewIdInstanceID checks that 'input' can be parsed as a Me Calendar Calendar View Id Instance ID
func ValidateMeCalendarCalendarViewIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Instance ID
func (id MeCalendarCalendarViewIdInstanceId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Instance ID
func (id MeCalendarCalendarViewIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Id Instance ID
func (id MeCalendarCalendarViewIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Instance (%s)", strings.Join(components, "\n"))
}
