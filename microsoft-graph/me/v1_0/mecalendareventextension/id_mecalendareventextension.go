package mecalendareventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventExtensionId{}

// MeCalendarEventExtensionId is a struct representing the Resource ID for a Me Calendar Event Extension
type MeCalendarEventExtensionId struct {
	CalendarId  string
	EventId     string
	ExtensionId string
}

// NewMeCalendarEventExtensionID returns a new MeCalendarEventExtensionId struct
func NewMeCalendarEventExtensionID(calendarId string, eventId string, extensionId string) MeCalendarEventExtensionId {
	return MeCalendarEventExtensionId{
		CalendarId:  calendarId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarEventExtensionID parses 'input' into a MeCalendarEventExtensionId
func ParseMeCalendarEventExtensionID(input string) (*MeCalendarEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExtensionId{}

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

// ParseMeCalendarEventExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarEventExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventExtensionIDInsensitively(input string) (*MeCalendarEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExtensionId{}

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

// ValidateMeCalendarEventExtensionID checks that 'input' can be parsed as a Me Calendar Event Extension ID
func ValidateMeCalendarEventExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Extension ID
func (id MeCalendarEventExtensionId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Extension ID
func (id MeCalendarEventExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Extension ID
func (id MeCalendarEventExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Event Extension (%s)", strings.Join(components, "\n"))
}
