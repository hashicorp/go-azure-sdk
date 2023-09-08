package meeventexceptionoccurrenceinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventExceptionOccurrenceInstanceId{}

// MeEventExceptionOccurrenceInstanceId is a struct representing the Resource ID for a Me Event Exception Occurrence Instance
type MeEventExceptionOccurrenceInstanceId struct {
	EventId  string
	EventId1 string
	EventId2 string
}

// NewMeEventExceptionOccurrenceInstanceID returns a new MeEventExceptionOccurrenceInstanceId struct
func NewMeEventExceptionOccurrenceInstanceID(eventId string, eventId1 string, eventId2 string) MeEventExceptionOccurrenceInstanceId {
	return MeEventExceptionOccurrenceInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseMeEventExceptionOccurrenceInstanceID parses 'input' into a MeEventExceptionOccurrenceInstanceId
func ParseMeEventExceptionOccurrenceInstanceID(input string) (*MeEventExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventExceptionOccurrenceInstanceId{}

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

// ParseMeEventExceptionOccurrenceInstanceIDInsensitively parses 'input' case-insensitively into a MeEventExceptionOccurrenceInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeEventExceptionOccurrenceInstanceIDInsensitively(input string) (*MeEventExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventExceptionOccurrenceInstanceId{}

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

// ValidateMeEventExceptionOccurrenceInstanceID checks that 'input' can be parsed as a Me Event Exception Occurrence Instance ID
func ValidateMeEventExceptionOccurrenceInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventExceptionOccurrenceInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Exception Occurrence Instance ID
func (id MeEventExceptionOccurrenceInstanceId) ID() string {
	fmtString := "/me/events/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Exception Occurrence Instance ID
func (id MeEventExceptionOccurrenceInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this Me Event Exception Occurrence Instance ID
func (id MeEventExceptionOccurrenceInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("Me Event Exception Occurrence Instance (%s)", strings.Join(components, "\n"))
}
