package meeventinstanceexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventInstanceId{}

// MeEventInstanceId is a struct representing the Resource ID for a Me Event Instance
type MeEventInstanceId struct {
	EventId  string
	EventId1 string
}

// NewMeEventInstanceID returns a new MeEventInstanceId struct
func NewMeEventInstanceID(eventId string, eventId1 string) MeEventInstanceId {
	return MeEventInstanceId{
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseMeEventInstanceID parses 'input' into a MeEventInstanceId
func ParseMeEventInstanceID(input string) (*MeEventInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseMeEventInstanceIDInsensitively parses 'input' case-insensitively into a MeEventInstanceId
// note: this method should only be used for API response data and not user input
func ParseMeEventInstanceIDInsensitively(input string) (*MeEventInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateMeEventInstanceID checks that 'input' can be parsed as a Me Event Instance ID
func ValidateMeEventInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Instance ID
func (id MeEventInstanceId) ID() string {
	fmtString := "/me/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Instance ID
func (id MeEventInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this Me Event Instance ID
func (id MeEventInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("Me Event Instance (%s)", strings.Join(components, "\n"))
}
