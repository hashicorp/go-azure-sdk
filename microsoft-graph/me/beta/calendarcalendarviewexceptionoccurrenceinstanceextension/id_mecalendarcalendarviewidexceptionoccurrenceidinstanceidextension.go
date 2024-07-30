package calendarcalendarviewexceptionoccurrenceinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{}

// MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId is a struct representing the Resource ID for a Me Calendar Calendar View Id Exception Occurrence Id Instance Id Extension
type MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId struct {
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID returns a new MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId struct
func NewMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID(eventId string, eventId1 string, eventId2 string, extensionId string) MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId {
	return MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID parses 'input' into a MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId
func ParseMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID(input string) (*MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionIDInsensitively(input string) (*MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID checks that 'input' can be parsed as a Me Calendar Calendar View Id Exception Occurrence Id Instance Id Extension ID
func ValidateMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Exception Occurrence Id Instance Id Extension ID
func (id MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/exceptionOccurrences/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Exception Occurrence Id Instance Id Extension ID
func (id MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
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

// String returns a human-readable description of this Me Calendar Calendar View Id Exception Occurrence Id Instance Id Extension ID
func (id MeCalendarCalendarViewIdExceptionOccurrenceIdInstanceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Exception Occurrence Id Instance Id Extension (%s)", strings.Join(components, "\n"))
}
