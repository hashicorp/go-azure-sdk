package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarEventIdExtensionId{}

// MeCalendarEventIdExtensionId is a struct representing the Resource ID for a Me Calendar Event Id Extension
type MeCalendarEventIdExtensionId struct {
	EventId     string
	ExtensionId string
}

// NewMeCalendarEventIdExtensionID returns a new MeCalendarEventIdExtensionId struct
func NewMeCalendarEventIdExtensionID(eventId string, extensionId string) MeCalendarEventIdExtensionId {
	return MeCalendarEventIdExtensionId{
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarEventIdExtensionID parses 'input' into a MeCalendarEventIdExtensionId
func ParseMeCalendarEventIdExtensionID(input string) (*MeCalendarEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarEventIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarEventIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarEventIdExtensionIDInsensitively(input string) (*MeCalendarEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarEventIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarEventIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarEventIdExtensionID checks that 'input' can be parsed as a Me Calendar Event Id Extension ID
func ValidateMeCalendarEventIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarEventIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Event Id Extension ID
func (id MeCalendarEventIdExtensionId) ID() string {
	fmtString := "/me/calendar/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Event Id Extension ID
func (id MeCalendarEventIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Event Id Extension ID
func (id MeCalendarEventIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Event Id Extension (%s)", strings.Join(components, "\n"))
}
