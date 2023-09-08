package mecalendarcalendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewExtensionId{}

// MeCalendarCalendarViewExtensionId is a struct representing the Resource ID for a Me Calendar Calendar View Extension
type MeCalendarCalendarViewExtensionId struct {
	CalendarId  string
	EventId     string
	ExtensionId string
}

// NewMeCalendarCalendarViewExtensionID returns a new MeCalendarCalendarViewExtensionId struct
func NewMeCalendarCalendarViewExtensionID(calendarId string, eventId string, extensionId string) MeCalendarCalendarViewExtensionId {
	return MeCalendarCalendarViewExtensionId{
		CalendarId:  calendarId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarCalendarViewExtensionID parses 'input' into a MeCalendarCalendarViewExtensionId
func ParseMeCalendarCalendarViewExtensionID(input string) (*MeCalendarCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewExtensionId{}

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

// ParseMeCalendarCalendarViewExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewExtensionIDInsensitively(input string) (*MeCalendarCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewExtensionId{}

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

// ValidateMeCalendarCalendarViewExtensionID checks that 'input' can be parsed as a Me Calendar Calendar View Extension ID
func ValidateMeCalendarCalendarViewExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Extension ID
func (id MeCalendarCalendarViewExtensionId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Extension ID
func (id MeCalendarCalendarViewExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Extension ID
func (id MeCalendarCalendarViewExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Extension (%s)", strings.Join(components, "\n"))
}
