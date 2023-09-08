package meeventinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventInstanceExtensionId{}

// MeEventInstanceExtensionId is a struct representing the Resource ID for a Me Event Instance Extension
type MeEventInstanceExtensionId struct {
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeEventInstanceExtensionID returns a new MeEventInstanceExtensionId struct
func NewMeEventInstanceExtensionID(eventId string, eventId1 string, extensionId string) MeEventInstanceExtensionId {
	return MeEventInstanceExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeEventInstanceExtensionID parses 'input' into a MeEventInstanceExtensionId
func ParseMeEventInstanceExtensionID(input string) (*MeEventInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceExtensionId{}

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

// ParseMeEventInstanceExtensionIDInsensitively parses 'input' case-insensitively into a MeEventInstanceExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeEventInstanceExtensionIDInsensitively(input string) (*MeEventInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventInstanceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventInstanceExtensionId{}

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

// ValidateMeEventInstanceExtensionID checks that 'input' can be parsed as a Me Event Instance Extension ID
func ValidateMeEventInstanceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventInstanceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Instance Extension ID
func (id MeEventInstanceExtensionId) ID() string {
	fmtString := "/me/events/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Instance Extension ID
func (id MeEventInstanceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Event Instance Extension ID
func (id MeEventInstanceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Event Instance Extension (%s)", strings.Join(components, "\n"))
}
