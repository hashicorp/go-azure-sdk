package meeventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventExceptionOccurrenceExtensionId{}

// MeEventExceptionOccurrenceExtensionId is a struct representing the Resource ID for a Me Event Exception Occurrence Extension
type MeEventExceptionOccurrenceExtensionId struct {
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeEventExceptionOccurrenceExtensionID returns a new MeEventExceptionOccurrenceExtensionId struct
func NewMeEventExceptionOccurrenceExtensionID(eventId string, eventId1 string, extensionId string) MeEventExceptionOccurrenceExtensionId {
	return MeEventExceptionOccurrenceExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeEventExceptionOccurrenceExtensionID parses 'input' into a MeEventExceptionOccurrenceExtensionId
func ParseMeEventExceptionOccurrenceExtensionID(input string) (*MeEventExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventExceptionOccurrenceExtensionId{}

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

// ParseMeEventExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a MeEventExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeEventExceptionOccurrenceExtensionIDInsensitively(input string) (*MeEventExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventExceptionOccurrenceExtensionId{}

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

// ValidateMeEventExceptionOccurrenceExtensionID checks that 'input' can be parsed as a Me Event Exception Occurrence Extension ID
func ValidateMeEventExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Exception Occurrence Extension ID
func (id MeEventExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/me/events/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Exception Occurrence Extension ID
func (id MeEventExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Event Exception Occurrence Extension ID
func (id MeEventExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Event Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
