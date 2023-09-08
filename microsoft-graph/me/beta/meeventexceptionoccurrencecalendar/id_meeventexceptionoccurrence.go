package meeventexceptionoccurrencecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventExceptionOccurrenceId{}

// MeEventExceptionOccurrenceId is a struct representing the Resource ID for a Me Event Exception Occurrence
type MeEventExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
}

// NewMeEventExceptionOccurrenceID returns a new MeEventExceptionOccurrenceId struct
func NewMeEventExceptionOccurrenceID(eventId string, eventId1 string) MeEventExceptionOccurrenceId {
	return MeEventExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeEventExceptionOccurrenceID parses 'input' into a MeEventExceptionOccurrenceId
func ParseMeEventExceptionOccurrenceID(input string) (*MeEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventExceptionOccurrenceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseMeEventExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeEventExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeEventExceptionOccurrenceIDInsensitively(input string) (*MeEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventExceptionOccurrenceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateMeEventExceptionOccurrenceID checks that 'input' can be parsed as a Me Event Exception Occurrence ID
func ValidateMeEventExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Exception Occurrence ID
func (id MeEventExceptionOccurrenceId) ID() string {
	fmtString := "/me/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Exception Occurrence ID
func (id MeEventExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Event Exception Occurrence ID
func (id MeEventExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Event Exception Occurrence (%s)", strings.Join(components, "\n"))
}
