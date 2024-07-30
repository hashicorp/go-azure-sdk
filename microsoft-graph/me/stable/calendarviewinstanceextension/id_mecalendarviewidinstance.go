package calendarviewinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdInstanceId{}

// MeCalendarViewIdInstanceId is a struct representing the Resource ID for a Me Calendar View Id Instance
type MeCalendarViewIdInstanceId struct {
	EventId  string
	EventId1 string
}

// NewMeCalendarViewIdInstanceID returns a new MeCalendarViewIdInstanceId struct
func NewMeCalendarViewIdInstanceID(eventId string, eventId1 string) MeCalendarViewIdInstanceId {
	return MeCalendarViewIdInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeCalendarViewIdInstanceID parses 'input' into a MeCalendarViewIdInstanceId
func ParseMeCalendarViewIdInstanceID(input string) (*MeCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdInstanceIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdInstanceIDInsensitively(input string) (*MeCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateMeCalendarViewIdInstanceID checks that 'input' can be parsed as a Me Calendar View Id Instance ID
func ValidateMeCalendarViewIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Instance ID
func (id MeCalendarViewIdInstanceId) ID() string {
	fmtString := "/me/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Instance ID
func (id MeCalendarViewIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Instance ID
func (id MeCalendarViewIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Calendar View Id Instance (%s)", strings.Join(components, "\n"))
}
