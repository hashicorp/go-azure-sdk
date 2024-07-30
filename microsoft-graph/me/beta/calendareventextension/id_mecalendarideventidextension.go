package calendareventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdEventIdExtensionId{}

// MeCalendarIdEventIdExtensionId is a struct representing the Resource ID for a Me Calendar Id Event Id Extension
type MeCalendarIdEventIdExtensionId struct {
	CalendarId  string
	EventId     string
	ExtensionId string
}

// NewMeCalendarIdEventIdExtensionID returns a new MeCalendarIdEventIdExtensionId struct
func NewMeCalendarIdEventIdExtensionID(calendarId string, eventId string, extensionId string) MeCalendarIdEventIdExtensionId {
	return MeCalendarIdEventIdExtensionId{
		CalendarId:  calendarId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarIdEventIdExtensionID parses 'input' into a MeCalendarIdEventIdExtensionId
func ParseMeCalendarIdEventIdExtensionID(input string) (*MeCalendarIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdEventIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarIdEventIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdEventIdExtensionIDInsensitively(input string) (*MeCalendarIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdEventIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateMeCalendarIdEventIdExtensionID checks that 'input' can be parsed as a Me Calendar Id Event Id Extension ID
func ValidateMeCalendarIdEventIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdEventIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Event Id Extension ID
func (id MeCalendarIdEventIdExtensionId) ID() string {
	fmtString := "/me/calendars/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Event Id Extension ID
func (id MeCalendarIdEventIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Event Id Extension ID
func (id MeCalendarIdEventIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Id Event Id Extension (%s)", strings.Join(components, "\n"))
}
