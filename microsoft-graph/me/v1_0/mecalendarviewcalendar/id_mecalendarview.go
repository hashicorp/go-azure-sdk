package mecalendarviewcalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewId{}

// MeCalendarViewId is a struct representing the Resource ID for a Me Calendar View
type MeCalendarViewId struct {
	EventId string
}

// NewMeCalendarViewID returns a new MeCalendarViewId struct
func NewMeCalendarViewID(eventId string) MeCalendarViewId {
	return MeCalendarViewId{
		EventId: eventId,
	}
}

// ParseMeCalendarViewID parses 'input' into a MeCalendarViewId
func ParseMeCalendarViewID(input string) (*MeCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarViewIDInsensitively parses 'input' case-insensitively into a MeCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIDInsensitively(input string) (*MeCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarViewID checks that 'input' can be parsed as a Me Calendar View ID
func ValidateMeCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View ID
func (id MeCalendarViewId) ID() string {
	fmtString := "/me/calendarView/%s"
	return fmt.Sprintf(fmtString, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View ID
func (id MeCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View ID
func (id MeCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Calendar View (%s)", strings.Join(components, "\n"))
}
