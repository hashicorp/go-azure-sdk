package mecalendarcalendarviewinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarCalendarViewInstanceExtensionId{}

// MeCalendarCalendarViewInstanceExtensionId is a struct representing the Resource ID for a Me Calendar Calendar View Instance Extension
type MeCalendarCalendarViewInstanceExtensionId struct {
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeCalendarCalendarViewInstanceExtensionID returns a new MeCalendarCalendarViewInstanceExtensionId struct
func NewMeCalendarCalendarViewInstanceExtensionID(eventId string, eventId1 string, extensionId string) MeCalendarCalendarViewInstanceExtensionId {
	return MeCalendarCalendarViewInstanceExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarCalendarViewInstanceExtensionID parses 'input' into a MeCalendarCalendarViewInstanceExtensionId
func ParseMeCalendarCalendarViewInstanceExtensionID(input string) (*MeCalendarCalendarViewInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewInstanceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewInstanceExtensionId{}

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

// ParseMeCalendarCalendarViewInstanceExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewInstanceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewInstanceExtensionIDInsensitively(input string) (*MeCalendarCalendarViewInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarCalendarViewInstanceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarCalendarViewInstanceExtensionId{}

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

// ValidateMeCalendarCalendarViewInstanceExtensionID checks that 'input' can be parsed as a Me Calendar Calendar View Instance Extension ID
func ValidateMeCalendarCalendarViewInstanceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewInstanceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Instance Extension ID
func (id MeCalendarCalendarViewInstanceExtensionId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Instance Extension ID
func (id MeCalendarCalendarViewInstanceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Instance Extension ID
func (id MeCalendarCalendarViewInstanceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Instance Extension (%s)", strings.Join(components, "\n"))
}
