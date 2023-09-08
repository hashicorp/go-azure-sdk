package mecalendareventexceptionoccurrenceinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarEventExceptionOccurrenceInstanceExtensionId{}

// MeCalendarEventExceptionOccurrenceInstanceExtensionId is a struct representing the Resource ID for a Me Calendar Event Exception Occurrence Instance Extension
type MeCalendarEventExceptionOccurrenceInstanceExtensionId struct {
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewMeCalendarEventExceptionOccurrenceInstanceExtensionID returns a new MeCalendarEventExceptionOccurrenceInstanceExtensionId struct
func NewMeCalendarEventExceptionOccurrenceInstanceExtensionID(eventId string, eventId1 string, eventId2 string, extensionId string) MeCalendarEventExceptionOccurrenceInstanceExtensionId {
	return MeCalendarEventExceptionOccurrenceInstanceExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarEventExceptionOccurrenceInstanceExtensionID parses 'input' into a MeCalendarEventExceptionOccurrenceInstanceExtensionId
func ParseMeCalendarEventExceptionOccurrenceInstanceExtensionID(input string) (*MeCalendarEventExceptionOccurrenceInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExceptionOccurrenceInstanceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExceptionOccurrenceInstanceExtensionId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarEventExceptionOccurrenceInstanceExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarEventExceptionOccurrenceInstanceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventExceptionOccurrenceInstanceExtensionIDInsensitively(input string) (*MeCalendarEventExceptionOccurrenceInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarEventExceptionOccurrenceInstanceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarEventExceptionOccurrenceInstanceExtensionId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarEventExceptionOccurrenceInstanceExtensionID checks that 'input' can be parsed as a Me Calendar Event Exception Occurrence Instance Extension ID
func ValidateMeCalendarEventExceptionOccurrenceInstanceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventExceptionOccurrenceInstanceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Exception Occurrence Instance Extension ID
func (id MeCalendarEventExceptionOccurrenceInstanceExtensionId) ID() string {
	fmtString := "/me/calendar/events/%s/exceptionOccurrences/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Exception Occurrence Instance Extension ID
func (id MeCalendarEventExceptionOccurrenceInstanceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Exception Occurrence Instance Extension ID
func (id MeCalendarEventExceptionOccurrenceInstanceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Event Exception Occurrence Instance Extension (%s)", strings.Join(components, "\n"))
}
