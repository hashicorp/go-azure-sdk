package mecalendargroupcalendarcalendarviewinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarCalendarViewInstanceExtensionId{}

// MeCalendarGroupCalendarCalendarViewInstanceExtensionId is a struct representing the Resource ID for a Me Calendar Group Calendar Calendar View Instance Extension
type MeCalendarGroupCalendarCalendarViewInstanceExtensionId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	ExtensionId     string
}

// NewMeCalendarGroupCalendarCalendarViewInstanceExtensionID returns a new MeCalendarGroupCalendarCalendarViewInstanceExtensionId struct
func NewMeCalendarGroupCalendarCalendarViewInstanceExtensionID(calendarGroupId string, calendarId string, eventId string, eventId1 string, extensionId string) MeCalendarGroupCalendarCalendarViewInstanceExtensionId {
	return MeCalendarGroupCalendarCalendarViewInstanceExtensionId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		ExtensionId:     extensionId,
	}
}

// ParseMeCalendarGroupCalendarCalendarViewInstanceExtensionID parses 'input' into a MeCalendarGroupCalendarCalendarViewInstanceExtensionId
func ParseMeCalendarGroupCalendarCalendarViewInstanceExtensionID(input string) (*MeCalendarGroupCalendarCalendarViewInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewInstanceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewInstanceExtensionId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupCalendarCalendarViewInstanceExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarCalendarViewInstanceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarCalendarViewInstanceExtensionIDInsensitively(input string) (*MeCalendarGroupCalendarCalendarViewInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewInstanceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewInstanceExtensionId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupCalendarCalendarViewInstanceExtensionID checks that 'input' can be parsed as a Me Calendar Group Calendar Calendar View Instance Extension ID
func ValidateMeCalendarGroupCalendarCalendarViewInstanceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarCalendarViewInstanceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Calendar View Instance Extension ID
func (id MeCalendarGroupCalendarCalendarViewInstanceExtensionId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Calendar View Instance Extension ID
func (id MeCalendarGroupCalendarCalendarViewInstanceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Calendar View Instance Extension ID
func (id MeCalendarGroupCalendarCalendarViewInstanceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Calendar View Instance Extension (%s)", strings.Join(components, "\n"))
}
