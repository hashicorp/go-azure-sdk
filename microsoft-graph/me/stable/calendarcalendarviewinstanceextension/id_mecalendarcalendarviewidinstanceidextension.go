package calendarcalendarviewinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarCalendarViewIdInstanceIdExtensionId{}

// MeCalendarCalendarViewIdInstanceIdExtensionId is a struct representing the Resource ID for a Me Calendar Calendar View Id Instance Id Extension
type MeCalendarCalendarViewIdInstanceIdExtensionId struct {
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewMeCalendarCalendarViewIdInstanceIdExtensionID returns a new MeCalendarCalendarViewIdInstanceIdExtensionId struct
func NewMeCalendarCalendarViewIdInstanceIdExtensionID(eventId string, eventId1 string, extensionId string) MeCalendarCalendarViewIdInstanceIdExtensionId {
	return MeCalendarCalendarViewIdInstanceIdExtensionId{
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseMeCalendarCalendarViewIdInstanceIdExtensionID parses 'input' into a MeCalendarCalendarViewIdInstanceIdExtensionId
func ParseMeCalendarCalendarViewIdInstanceIdExtensionID(input string) (*MeCalendarCalendarViewIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarCalendarViewIdInstanceIdExtensionIDInsensitively parses 'input' case-insensitively into a MeCalendarCalendarViewIdInstanceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarCalendarViewIdInstanceIdExtensionIDInsensitively(input string) (*MeCalendarCalendarViewIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarCalendarViewIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarCalendarViewIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarCalendarViewIdInstanceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateMeCalendarCalendarViewIdInstanceIdExtensionID checks that 'input' can be parsed as a Me Calendar Calendar View Id Instance Id Extension ID
func ValidateMeCalendarCalendarViewIdInstanceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarCalendarViewIdInstanceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Calendar View Id Instance Id Extension ID
func (id MeCalendarCalendarViewIdInstanceIdExtensionId) ID() string {
	fmtString := "/me/calendar/calendarView/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Calendar View Id Instance Id Extension ID
func (id MeCalendarCalendarViewIdInstanceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Calendar View Id Instance Id Extension ID
func (id MeCalendarCalendarViewIdInstanceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("Me Calendar Calendar View Id Instance Id Extension (%s)", strings.Join(components, "\n"))
}
