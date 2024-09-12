package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarIdCalendarPermissionId{}

// MeCalendarIdCalendarPermissionId is a struct representing the Resource ID for a Me Calendar Id Calendar Permission
type MeCalendarIdCalendarPermissionId struct {
	CalendarId           string
	CalendarPermissionId string
}

// NewMeCalendarIdCalendarPermissionID returns a new MeCalendarIdCalendarPermissionId struct
func NewMeCalendarIdCalendarPermissionID(calendarId string, calendarPermissionId string) MeCalendarIdCalendarPermissionId {
	return MeCalendarIdCalendarPermissionId{
		CalendarId:           calendarId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseMeCalendarIdCalendarPermissionID parses 'input' into a MeCalendarIdCalendarPermissionId
func ParseMeCalendarIdCalendarPermissionID(input string) (*MeCalendarIdCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarIdCalendarPermissionIDInsensitively parses 'input' case-insensitively into a MeCalendarIdCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarIdCalendarPermissionIDInsensitively(input string) (*MeCalendarIdCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarIdCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarIdCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarIdCalendarPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.CalendarPermissionId, ok = input.Parsed["calendarPermissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", input)
	}

	return nil
}

// ValidateMeCalendarIdCalendarPermissionID checks that 'input' can be parsed as a Me Calendar Id Calendar Permission ID
func ValidateMeCalendarIdCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarIdCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Id Calendar Permission ID
func (id MeCalendarIdCalendarPermissionId) ID() string {
	fmtString := "/me/calendars/%s/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.CalendarId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Id Calendar Permission ID
func (id MeCalendarIdCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Id Calendar Permission ID
func (id MeCalendarIdCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("Me Calendar Id Calendar Permission (%s)", strings.Join(components, "\n"))
}
