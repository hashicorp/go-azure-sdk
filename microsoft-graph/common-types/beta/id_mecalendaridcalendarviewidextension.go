package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarViewIdExtensionId{}

// MeCalendarIdCalendarViewIdExtensionId is a struct representing the Resource ID for a Me Calendar Id Calendar View Id Extension
type MeCalendarIdCalendarViewIdExtensionId struct {
	CalendarId  string
	EventId     string
	ExtensionId string
}

// NewMeCalendarIdCalendarViewIdExtensionID returns a new MeCalendarIdCalendarViewIdExtensionId struct
func NewMeCalendarIdCalendarViewIdExtensionID(calendarId string, eventId string, extensionId string) MeCalendarIdCalendarViewIdExtensionId {
	return MeCalendarIdCalendarViewIdExtensionId{
		CalendarId:  calendarId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarIdCalendarViewIdExtensionID parses 'input' into a MeCalendarIdCalendarViewIdExtensionId
func ParseMeCalendarIdCalendarViewIdExtensionID(input string) (*MeCalendarIdCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdCalendarViewIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarIdCalendarViewIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdCalendarViewIdExtensionIDInsensitively(input string) (*MeCalendarIdCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdCalendarViewIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarIdCalendarViewIdExtensionID checks that 'input' can be parsed as a Me Calendar Id Calendar View Id Extension ID
func ValidateMeCalendarIdCalendarViewIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdCalendarViewIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Calendar View Id Extension ID
func (id MeCalendarIdCalendarViewIdExtensionId) ID() string {
	fmtString := "/me/calendars/%s/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Calendar View Id Extension ID
func (id MeCalendarIdCalendarViewIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Calendar View Id Extension ID
func (id MeCalendarIdCalendarViewIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Id Calendar View Id Extension (%s)", strings.Join(components, "\n"))
}
