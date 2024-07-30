package calendareventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdEventIdExceptionOccurrenceIdExtensionId{}

// MeCalendarIdEventIdExceptionOccurrenceIdExtensionId is a struct representing the Resource ID for a Me Calendar Id Event Id Exception Occurrence Id Extension
type MeCalendarIdEventIdExceptionOccurrenceIdExtensionId struct {
	CalendarId  string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeCalendarIdEventIdExceptionOccurrenceIdExtensionID returns a new MeCalendarIdEventIdExceptionOccurrenceIdExtensionId struct
func NewMeCalendarIdEventIdExceptionOccurrenceIdExtensionID(calendarId string, eventId string, eventId1 string, extensionId string) MeCalendarIdEventIdExceptionOccurrenceIdExtensionId {
	return MeCalendarIdEventIdExceptionOccurrenceIdExtensionId{
		CalendarId:  calendarId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarIdEventIdExceptionOccurrenceIdExtensionID parses 'input' into a MeCalendarIdEventIdExceptionOccurrenceIdExtensionId
func ParseMeCalendarIdEventIdExceptionOccurrenceIdExtensionID(input string) (*MeCalendarIdEventIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdEventIdExceptionOccurrenceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarIdEventIdExceptionOccurrenceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdEventIdExceptionOccurrenceIdExtensionIDInsensitively(input string) (*MeCalendarIdEventIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdEventIdExceptionOccurrenceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarIdEventIdExceptionOccurrenceIdExtensionID checks that 'input' can be parsed as a Me Calendar Id Event Id Exception Occurrence Id Extension ID
func ValidateMeCalendarIdEventIdExceptionOccurrenceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdEventIdExceptionOccurrenceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Event Id Exception Occurrence Id Extension ID
func (id MeCalendarIdEventIdExceptionOccurrenceIdExtensionId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Event Id Exception Occurrence Id Extension ID
func (id MeCalendarIdEventIdExceptionOccurrenceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Event Id Exception Occurrence Id Extension ID
func (id MeCalendarIdEventIdExceptionOccurrenceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Id Event Id Exception Occurrence Id Extension (%s)", strings.Join(components, "\n"))
}
