package mecalendargroupcalendarevent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarId{}

// MeCalendarGroupCalendarId is a struct representing the Resource ID for a Me Calendar Group Calendar
type MeCalendarGroupCalendarId struct {
	CalendarGroupId string
	CalendarId      string
}

// NewMeCalendarGroupCalendarID returns a new MeCalendarGroupCalendarId struct
func NewMeCalendarGroupCalendarID(calendarGroupId string, calendarId string) MeCalendarGroupCalendarId {
	return MeCalendarGroupCalendarId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
	}
}

// ParseMeCalendarGroupCalendarID parses 'input' into a MeCalendarGroupCalendarId
func ParseMeCalendarGroupCalendarID(input string) (*MeCalendarGroupCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupCalendarIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarIDInsensitively(input string) (*MeCalendarGroupCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupCalendarID checks that 'input' can be parsed as a Me Calendar Group Calendar ID
func ValidateMeCalendarGroupCalendarID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar ID
func (id MeCalendarGroupCalendarId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar ID
func (id MeCalendarGroupCalendarId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar ID
func (id MeCalendarGroupCalendarId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
	}
	return fmt.Sprintf("Me Calendar Group Calendar (%s)", strings.Join(components, "\n"))
}
