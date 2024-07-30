package calendargroupcalendarcalendarviewexceptionoccurrenceinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{}

// MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence Id Instance Id Extension
type MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId struct {
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	EventId2        string
	ExtensionId     string
}

// NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID returns a new MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId struct
func NewMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID(calendarGroupId string, calendarId string, eventId string, eventId1 string, eventId2 string, extensionId string) MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId {
	return MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		EventId2:        eventId2,
		ExtensionId:     extensionId,
	}
}

// ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID parses 'input' into a MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId
func ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID(input string) (*MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence Id Instance Id Extension ID
func ValidateMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence Id Instance Id Extension ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarView/%s/exceptionOccurrences/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence Id Instance Id Extension ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence Id Instance Id Extension ID
func (id MeCalendarGroupIdCalendarIdCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Calendar View Id Exception Occurrence Id Instance Id Extension (%s)", strings.Join(components, "\n"))
}
