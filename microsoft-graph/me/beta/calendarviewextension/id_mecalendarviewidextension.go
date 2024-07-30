package calendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarViewIdExtensionId{}

// MeCalendarViewIdExtensionId is a struct representing the Resource ID for a Me Calendar View Id Extension
type MeCalendarViewIdExtensionId struct {
	EventId     string
	ExtensionId string
}

// NewMeCalendarViewIdExtensionID returns a new MeCalendarViewIdExtensionId struct
func NewMeCalendarViewIdExtensionID(eventId string, extensionId string) MeCalendarViewIdExtensionId {
	return MeCalendarViewIdExtensionId{
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarViewIdExtensionID parses 'input' into a MeCalendarViewIdExtensionId
func ParseMeCalendarViewIdExtensionID(input string) (*MeCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarViewIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarViewIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarViewIdExtensionIDInsensitively(input string) (*MeCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarViewIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarViewIdExtensionID checks that 'input' can be parsed as a Me Calendar View Id Extension ID
func ValidateMeCalendarViewIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarViewIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar View Id Extension ID
func (id MeCalendarViewIdExtensionId) ID() string {
	fmtString := "/me/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar View Id Extension ID
func (id MeCalendarViewIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar View Id Extension ID
func (id MeCalendarViewIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar View Id Extension (%s)", strings.Join(components, "\n"))
}
