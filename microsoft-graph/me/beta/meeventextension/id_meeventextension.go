package meeventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeEventExtensionId{}

// MeEventExtensionId is a struct representing the Resource ID for a Me Event Extension
type MeEventExtensionId struct {
	EventId     string
	ExtensionId string
}

// NewMeEventExtensionID returns a new MeEventExtensionId struct
func NewMeEventExtensionID(eventId string, extensionId string) MeEventExtensionId {
	return MeEventExtensionId{
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeEventExtensionID parses 'input' into a MeEventExtensionId
func ParseMeEventExtensionID(input string) (*MeEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventExtensionId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeEventExtensionIDInsensitively parses 'input' case-insensitively into a MeEventExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeEventExtensionIDInsensitively(input string) (*MeEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeEventExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeEventExtensionId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeEventExtensionID checks that 'input' can be parsed as a Me Event Extension ID
func ValidateMeEventExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeEventExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Event Extension ID
func (id MeEventExtensionId) ID() string {
	fmtString := "/me/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Event Extension ID
func (id MeEventExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Event Extension ID
func (id MeEventExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Event Extension (%s)", strings.Join(components, "\n"))
}
