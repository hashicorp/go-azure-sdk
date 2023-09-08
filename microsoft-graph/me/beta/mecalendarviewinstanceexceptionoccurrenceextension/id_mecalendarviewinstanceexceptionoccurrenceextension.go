package mecalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewInstanceExceptionOccurrenceExtensionId{}

// MeCalendarViewInstanceExceptionOccurrenceExtensionId is a struct representing the Resource ID for a Me Calendar View Instance Exception Occurrence Extension
type MeCalendarViewInstanceExceptionOccurrenceExtensionId struct {
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewMeCalendarViewInstanceExceptionOccurrenceExtensionID returns a new MeCalendarViewInstanceExceptionOccurrenceExtensionId struct
func NewMeCalendarViewInstanceExceptionOccurrenceExtensionID(eventId string, eventId1 string, eventId2 string, extensionId string) MeCalendarViewInstanceExceptionOccurrenceExtensionId {
	return MeCalendarViewInstanceExceptionOccurrenceExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarViewInstanceExceptionOccurrenceExtensionID parses 'input' into a MeCalendarViewInstanceExceptionOccurrenceExtensionId
func ParseMeCalendarViewInstanceExceptionOccurrenceExtensionID(input string) (*MeCalendarViewInstanceExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewInstanceExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewInstanceExceptionOccurrenceExtensionId{}

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

// ParseMeCalendarViewInstanceExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarViewInstanceExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewInstanceExceptionOccurrenceExtensionIDInsensitively(input string) (*MeCalendarViewInstanceExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewInstanceExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewInstanceExceptionOccurrenceExtensionId{}

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

// ValidateMeCalendarViewInstanceExceptionOccurrenceExtensionID checks that 'input' can be parsed as a Me Calendar View Instance Exception Occurrence Extension ID
func ValidateMeCalendarViewInstanceExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewInstanceExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Instance Exception Occurrence Extension ID
func (id MeCalendarViewInstanceExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/me/calendarView/%s/instances/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Instance Exception Occurrence Extension ID
func (id MeCalendarViewInstanceExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Instance Exception Occurrence Extension ID
func (id MeCalendarViewInstanceExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar View Instance Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
