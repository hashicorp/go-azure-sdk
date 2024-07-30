package calendarcalendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdExtensionId{}

// MeCalendarCalendarViewIdExtensionId is a struct representing the Resource ID for a Me Calendar Calendar View Id Extension
type MeCalendarCalendarViewIdExtensionId struct {
	EventId     string
	ExtensionId string
}

// NewMeCalendarCalendarViewIdExtensionID returns a new MeCalendarCalendarViewIdExtensionId struct
func NewMeCalendarCalendarViewIdExtensionID(eventId string, extensionId string) MeCalendarCalendarViewIdExtensionId {
	return MeCalendarCalendarViewIdExtensionId{
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarCalendarViewIdExtensionID parses 'input' into a MeCalendarCalendarViewIdExtensionId
func ParseMeCalendarCalendarViewIdExtensionID(input string) (*MeCalendarCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdExtensionIDInsensitively(input string) (*MeCalendarCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarCalendarViewIdExtensionID checks that 'input' can be parsed as a Me Calendar Calendar View Id Extension ID
func ValidateMeCalendarCalendarViewIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Extension ID
func (id MeCalendarCalendarViewIdExtensionId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Extension ID
func (id MeCalendarCalendarViewIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Id Extension ID
func (id MeCalendarCalendarViewIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Extension (%s)", strings.Join(components, "\n"))
}
