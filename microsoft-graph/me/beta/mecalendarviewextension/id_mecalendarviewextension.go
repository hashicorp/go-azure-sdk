package mecalendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarViewExtensionId{}

// MeCalendarViewExtensionId is a struct representing the Resource ID for a Me Calendar View Extension
type MeCalendarViewExtensionId struct {
	EventId     string
	ExtensionId string
}

// NewMeCalendarViewExtensionID returns a new MeCalendarViewExtensionId struct
func NewMeCalendarViewExtensionID(eventId string, extensionId string) MeCalendarViewExtensionId {
	return MeCalendarViewExtensionId{
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarViewExtensionID parses 'input' into a MeCalendarViewExtensionId
func ParseMeCalendarViewExtensionID(input string) (*MeCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExtensionId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarViewExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarViewExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewExtensionIDInsensitively(input string) (*MeCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarViewExtensionId{}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarViewExtensionID checks that 'input' can be parsed as a Me Calendar View Extension ID
func ValidateMeCalendarViewExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Extension ID
func (id MeCalendarViewExtensionId) ID() string {
	fmtString := "/me/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Extension ID
func (id MeCalendarViewExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Extension ID
func (id MeCalendarViewExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar View Extension (%s)", strings.Join(components, "\n"))
}
