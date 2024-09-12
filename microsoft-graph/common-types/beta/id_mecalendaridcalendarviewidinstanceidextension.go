package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarViewIdInstanceIdExtensionId{}

// MeCalendarIdCalendarViewIdInstanceIdExtensionId is a struct representing the Resource ID for a Me Calendar Id Calendar View Id Instance Id Extension
type MeCalendarIdCalendarViewIdInstanceIdExtensionId struct {
	CalendarId  string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeCalendarIdCalendarViewIdInstanceIdExtensionID returns a new MeCalendarIdCalendarViewIdInstanceIdExtensionId struct
func NewMeCalendarIdCalendarViewIdInstanceIdExtensionID(calendarId string, eventId string, eventId1 string, extensionId string) MeCalendarIdCalendarViewIdInstanceIdExtensionId {
	return MeCalendarIdCalendarViewIdInstanceIdExtensionId{
		CalendarId:  calendarId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarIdCalendarViewIdInstanceIdExtensionID parses 'input' into a MeCalendarIdCalendarViewIdInstanceIdExtensionId
func ParseMeCalendarIdCalendarViewIdInstanceIdExtensionID(input string) (*MeCalendarIdCalendarViewIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdCalendarViewIdInstanceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarIdCalendarViewIdInstanceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdCalendarViewIdInstanceIdExtensionIDInsensitively(input string) (*MeCalendarIdCalendarViewIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdCalendarViewIdInstanceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarIdCalendarViewIdInstanceIdExtensionID checks that 'input' can be parsed as a Me Calendar Id Calendar View Id Instance Id Extension ID
func ValidateMeCalendarIdCalendarViewIdInstanceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdCalendarViewIdInstanceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Calendar View Id Instance Id Extension ID
func (id MeCalendarIdCalendarViewIdInstanceIdExtensionId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Calendar View Id Instance Id Extension ID
func (id MeCalendarIdCalendarViewIdInstanceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Calendar View Id Instance Id Extension ID
func (id MeCalendarIdCalendarViewIdInstanceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Id Calendar View Id Instance Id Extension (%s)", strings.Join(components, "\n"))
}
