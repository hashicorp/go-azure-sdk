package mecalendarcalendarview

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarId{}

// MeCalendarId is a struct representing the Resource ID for a Me Calendar
type MeCalendarId struct {
	CalendarId string
}

// NewMeCalendarID returns a new MeCalendarId struct
func NewMeCalendarID(calendarId string) MeCalendarId {
	return MeCalendarId{
		CalendarId: calendarId,
	}
}

// ParseMeCalendarID parses 'input' into a MeCalendarId
func ParseMeCalendarID(input string) (*MeCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarIDInsensitively parses 'input' case-insensitively into a MeCalendarId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIDInsensitively(input string) (*MeCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarId{}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarID checks that 'input' can be parsed as a Me Calendar ID
func ValidateMeCalendarID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar ID
func (id MeCalendarId) ID() string {
	fmtString := "/me/calendars/%s"
	return fmt.Sprintf(fmtString, id.CalendarId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar ID
func (id MeCalendarId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar ID
func (id MeCalendarId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
	}
	return fmt.Sprintf("Me Calendar (%s)", strings.Join(components, "\n"))
}
