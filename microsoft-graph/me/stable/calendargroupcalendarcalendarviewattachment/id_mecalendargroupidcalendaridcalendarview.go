package calendargroupcalendarcalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdCalendarViewId{}

// MeCalendarGroupIdCalendarIdCalendarViewId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Calendar View
type MeCalendarGroupIdCalendarIdCalendarViewId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
}

// NewMeCalendarGroupIdCalendarIdCalendarViewID returns a new MeCalendarGroupIdCalendarIdCalendarViewId struct
func NewMeCalendarGroupIdCalendarIdCalendarViewID(calendarGroupId string, calendarId string, eventId string) MeCalendarGroupIdCalendarIdCalendarViewId {
	return MeCalendarGroupIdCalendarIdCalendarViewId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
	}
}

// ParseMeCalendarGroupIdCalendarIdCalendarViewID parses 'input' into a MeCalendarGroupIdCalendarIdCalendarViewId
func ParseMeCalendarGroupIdCalendarIdCalendarViewID(input string) (*MeCalendarGroupIdCalendarIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdCalendarViewIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdCalendarViewIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdCalendarViewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarGroupId, ok = input.Parsed["calendarGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateMeCalendarGroupIdCalendarIdCalendarViewID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Calendar View ID
func ValidateMeCalendarGroupIdCalendarIdCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Calendar View ID
func (id MeCalendarGroupIdCalendarIdCalendarViewId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Calendar View ID
func (id MeCalendarGroupIdCalendarIdCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Calendar View ID
func (id MeCalendarGroupIdCalendarIdCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Calendar View (%s)", strings.Join(components, "\n"))
}
