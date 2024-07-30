package calendareventinstanceexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId{}

// MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId is a struct representing the Resource ID for a Me Calendar Id Event Id Instance Id Exception Occurrence Id Extension
type MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId struct {
	CalendarId  string
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionID returns a new MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId struct
func NewMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionID(calendarId string, eventId string, eventId1 string, eventId2 string, extensionId string) MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId {
	return MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId{
		CalendarId:  calendarId,
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionID parses 'input' into a MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId
func ParseMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionID(input string) (*MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively(input string) (*MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionID checks that 'input' can be parsed as a Me Calendar Id Event Id Instance Id Exception Occurrence Id Extension ID
func ValidateMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Event Id Instance Id Exception Occurrence Id Extension ID
func (id MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/instances/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Event Id Instance Id Exception Occurrence Id Extension ID
func (id MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Event Id Instance Id Exception Occurrence Id Extension ID
func (id MeCalendarIdEventIdInstanceIdExceptionOccurrenceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Id Event Id Instance Id Exception Occurrence Id Extension (%s)", strings.Join(components, "\n"))
}
