package calendarviewexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdExceptionOccurrenceIdExtensionId{}

// MeCalendarViewIdExceptionOccurrenceIdExtensionId is a struct representing the Resource ID for a Me Calendar View Id Exception Occurrence Id Extension
type MeCalendarViewIdExceptionOccurrenceIdExtensionId struct {
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeCalendarViewIdExceptionOccurrenceIdExtensionID returns a new MeCalendarViewIdExceptionOccurrenceIdExtensionId struct
func NewMeCalendarViewIdExceptionOccurrenceIdExtensionID(eventId string, eventId1 string, extensionId string) MeCalendarViewIdExceptionOccurrenceIdExtensionId {
	return MeCalendarViewIdExceptionOccurrenceIdExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarViewIdExceptionOccurrenceIdExtensionID parses 'input' into a MeCalendarViewIdExceptionOccurrenceIdExtensionId
func ParseMeCalendarViewIdExceptionOccurrenceIdExtensionID(input string) (*MeCalendarViewIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdExceptionOccurrenceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdExceptionOccurrenceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdExceptionOccurrenceIdExtensionIDInsensitively(input string) (*MeCalendarViewIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdExceptionOccurrenceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMeCalendarViewIdExceptionOccurrenceIdExtensionID checks that 'input' can be parsed as a Me Calendar View Id Exception Occurrence Id Extension ID
func ValidateMeCalendarViewIdExceptionOccurrenceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdExceptionOccurrenceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Exception Occurrence Id Extension ID
func (id MeCalendarViewIdExceptionOccurrenceIdExtensionId) ID() string {
	fmtString := "/me/calendarView/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Exception Occurrence Id Extension ID
func (id MeCalendarViewIdExceptionOccurrenceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Exception Occurrence Id Extension ID
func (id MeCalendarViewIdExceptionOccurrenceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar View Id Exception Occurrence Id Extension (%s)", strings.Join(components, "\n"))
}
