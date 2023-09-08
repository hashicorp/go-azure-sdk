package mecalendargroupcalendarcalendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarCalendarViewExtensionId{}

// MeCalendarGroupCalendarCalendarViewExtensionId is a struct representing the Resource ID for a Me Calendar Group Calendar Calendar View Extension
type MeCalendarGroupCalendarCalendarViewExtensionId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	ExtensionId     string
}

// NewMeCalendarGroupCalendarCalendarViewExtensionID returns a new MeCalendarGroupCalendarCalendarViewExtensionId struct
func NewMeCalendarGroupCalendarCalendarViewExtensionID(calendarGroupId string, calendarId string, eventId string, extensionId string) MeCalendarGroupCalendarCalendarViewExtensionId {
	return MeCalendarGroupCalendarCalendarViewExtensionId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		ExtensionId:     extensionId,
	}
}

// ParseMeCalendarGroupCalendarCalendarViewExtensionID parses 'input' into a MeCalendarGroupCalendarCalendarViewExtensionId
func ParseMeCalendarGroupCalendarCalendarViewExtensionID(input string) (*MeCalendarGroupCalendarCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewExtensionId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupCalendarCalendarViewExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarCalendarViewExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarCalendarViewExtensionIDInsensitively(input string) (*MeCalendarGroupCalendarCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarViewExtensionId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupCalendarCalendarViewExtensionID checks that 'input' can be parsed as a Me Calendar Group Calendar Calendar View Extension ID
func ValidateMeCalendarGroupCalendarCalendarViewExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarCalendarViewExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Calendar View Extension ID
func (id MeCalendarGroupCalendarCalendarViewExtensionId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Calendar View Extension ID
func (id MeCalendarGroupCalendarCalendarViewExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Calendar View Extension ID
func (id MeCalendarGroupCalendarCalendarViewExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Calendar View Extension (%s)", strings.Join(components, "\n"))
}
