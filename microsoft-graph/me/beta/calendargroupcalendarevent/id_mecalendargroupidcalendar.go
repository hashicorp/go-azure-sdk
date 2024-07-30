package calendargroupcalendarevent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarId{}

// MeCalendarGroupIdCalendarId is a struct representing the Resource ID for a Me Calendar Group Id Calendar
type MeCalendarGroupIdCalendarId struct {
	CalendarGroupId string
	CalendarId      string
}

// NewMeCalendarGroupIdCalendarID returns a new MeCalendarGroupIdCalendarId struct
func NewMeCalendarGroupIdCalendarID(calendarGroupId string, calendarId string) MeCalendarGroupIdCalendarId {
	return MeCalendarGroupIdCalendarId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
	}
}

// ParseMeCalendarGroupIdCalendarID parses 'input' into a MeCalendarGroupIdCalendarId
func ParseMeCalendarGroupIdCalendarID(input string) (*MeCalendarGroupIdCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIDInsensitively(input string) (*MeCalendarGroupIdCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarGroupId, ok = input.Parsed["calendarGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	return nil
}

// ValidateMeCalendarGroupIdCalendarID checks that 'input' can be parsed as a Me Calendar Group Id Calendar ID
func ValidateMeCalendarGroupIdCalendarID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar ID
func (id MeCalendarGroupIdCalendarId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar ID
func (id MeCalendarGroupIdCalendarId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar ID
func (id MeCalendarGroupIdCalendarId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar (%s)", strings.Join(components, "\n"))
}
