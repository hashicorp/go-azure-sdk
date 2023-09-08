package meeventinstanceexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventInstanceExceptionOccurrenceId{}

// MeEventInstanceExceptionOccurrenceId is a struct representing the Resource ID for a Me Event Instance Exception Occurrence
type MeEventInstanceExceptionOccurrenceId struct {
	EventId  string
	EventId1 string
	EventId2 string
}

// NewMeEventInstanceExceptionOccurrenceID returns a new MeEventInstanceExceptionOccurrenceId struct
func NewMeEventInstanceExceptionOccurrenceID(eventId string, eventId1 string, eventId2 string) MeEventInstanceExceptionOccurrenceId {
	return MeEventInstanceExceptionOccurrenceId{
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseMeEventInstanceExceptionOccurrenceID parses 'input' into a MeEventInstanceExceptionOccurrenceId
func ParseMeEventInstanceExceptionOccurrenceID(input string) (*MeEventInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceExceptionOccurrenceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	return &id, nil
}

// ParseMeEventInstanceExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a MeEventInstanceExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseMeEventInstanceExceptionOccurrenceIDInsensitively(input string) (*MeEventInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceExceptionOccurrenceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	return &id, nil
}

// ValidateMeEventInstanceExceptionOccurrenceID checks that 'input' can be parsed as a Me Event Instance Exception Occurrence ID
func ValidateMeEventInstanceExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventInstanceExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Instance Exception Occurrence ID
func (id MeEventInstanceExceptionOccurrenceId) ID() string {
	fmtString := "/me/events/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Instance Exception Occurrence ID
func (id MeEventInstanceExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Event Instance Exception Occurrence ID
func (id MeEventInstanceExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Event Instance Exception Occurrence (%s)", strings.Join(components, "\n"))
}
