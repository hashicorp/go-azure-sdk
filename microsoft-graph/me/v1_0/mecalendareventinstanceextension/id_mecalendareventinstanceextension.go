package mecalendareventinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventInstanceExtensionId{}

// MeCalendarEventInstanceExtensionId is a struct representing the Resource ID for a Me Calendar Event Instance Extension
type MeCalendarEventInstanceExtensionId struct {
	CalendarId  string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeCalendarEventInstanceExtensionID returns a new MeCalendarEventInstanceExtensionId struct
func NewMeCalendarEventInstanceExtensionID(calendarId string, eventId string, eventId1 string, extensionId string) MeCalendarEventInstanceExtensionId {
	return MeCalendarEventInstanceExtensionId{
		CalendarId:  calendarId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarEventInstanceExtensionID parses 'input' into a MeCalendarEventInstanceExtensionId
func ParseMeCalendarEventInstanceExtensionID(input string) (*MeCalendarEventInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventInstanceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventInstanceExtensionId{}

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

// ParseMeCalendarEventInstanceExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarEventInstanceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventInstanceExtensionIDInsensitively(input string) (*MeCalendarEventInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventInstanceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventInstanceExtensionId{}

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

// ValidateMeCalendarEventInstanceExtensionID checks that 'input' can be parsed as a Me Calendar Event Instance Extension ID
func ValidateMeCalendarEventInstanceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventInstanceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Instance Extension ID
func (id MeCalendarEventInstanceExtensionId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Instance Extension ID
func (id MeCalendarEventInstanceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Instance Extension ID
func (id MeCalendarEventInstanceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Event Instance Extension (%s)", strings.Join(components, "\n"))
}
