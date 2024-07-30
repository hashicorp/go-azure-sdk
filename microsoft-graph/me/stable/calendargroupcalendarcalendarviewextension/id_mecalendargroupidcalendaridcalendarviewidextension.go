package calendargroupcalendarcalendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId{}

// MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Calendar View Id Extension
type MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	ExtensionId     string
}

// NewMeCalendarGroupIdCalendarIdCalendarViewIdExtensionID returns a new MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId struct
func NewMeCalendarGroupIdCalendarIdCalendarViewIdExtensionID(calendarGroupId string, calendarId string, eventId string, extensionId string) MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId {
	return MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		ExtensionId:     extensionId,
	}
}

// ParseMeCalendarGroupIdCalendarIdCalendarViewIdExtensionID parses 'input' into a MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId
func ParseMeCalendarGroupIdCalendarIdCalendarViewIdExtensionID(input string) (*MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdCalendarViewIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdCalendarViewIdExtensionIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarGroupIdCalendarIdCalendarViewIdExtensionID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Calendar View Id Extension ID
func ValidateMeCalendarGroupIdCalendarIdCalendarViewIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdCalendarViewIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Calendar View Id Extension ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Calendar View Id Extension ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Calendar View Id Extension ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Calendar View Id Extension (%s)", strings.Join(components, "\n"))
}
