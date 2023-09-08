package mecalendarviewexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewExceptionOccurrenceExtensionId{}

// MeCalendarViewExceptionOccurrenceExtensionId is a struct representing the Resource ID for a Me Calendar View Exception Occurrence Extension
type MeCalendarViewExceptionOccurrenceExtensionId struct {
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeCalendarViewExceptionOccurrenceExtensionID returns a new MeCalendarViewExceptionOccurrenceExtensionId struct
func NewMeCalendarViewExceptionOccurrenceExtensionID(eventId string, eventId1 string, extensionId string) MeCalendarViewExceptionOccurrenceExtensionId {
	return MeCalendarViewExceptionOccurrenceExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarViewExceptionOccurrenceExtensionID parses 'input' into a MeCalendarViewExceptionOccurrenceExtensionId
func ParseMeCalendarViewExceptionOccurrenceExtensionID(input string) (*MeCalendarViewExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExceptionOccurrenceExtensionId{}

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

// ParseMeCalendarViewExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarViewExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewExceptionOccurrenceExtensionIDInsensitively(input string) (*MeCalendarViewExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExceptionOccurrenceExtensionId{}

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

// ValidateMeCalendarViewExceptionOccurrenceExtensionID checks that 'input' can be parsed as a Me Calendar View Exception Occurrence Extension ID
func ValidateMeCalendarViewExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Exception Occurrence Extension ID
func (id MeCalendarViewExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/me/calendarView/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Exception Occurrence Extension ID
func (id MeCalendarViewExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Exception Occurrence Extension ID
func (id MeCalendarViewExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar View Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
