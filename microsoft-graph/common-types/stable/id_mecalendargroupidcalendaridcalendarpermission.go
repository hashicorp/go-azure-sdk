package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCalendarGroupIdCalendarIdCalendarPermissionId{}

// MeCalendarGroupIdCalendarIdCalendarPermissionId is a struct representing the Resource ID for a Me Calendar Group Id Calendar Id Calendar Permission
type MeCalendarGroupIdCalendarIdCalendarPermissionId struct {
	CalendarGroupId      string
	CalendarId           string
	CalendarPermissionId string
}

// NewMeCalendarGroupIdCalendarIdCalendarPermissionID returns a new MeCalendarGroupIdCalendarIdCalendarPermissionId struct
func NewMeCalendarGroupIdCalendarIdCalendarPermissionID(calendarGroupId string, calendarId string, calendarPermissionId string) MeCalendarGroupIdCalendarIdCalendarPermissionId {
	return MeCalendarGroupIdCalendarIdCalendarPermissionId{
		CalendarGroupId:      calendarGroupId,
		CalendarId:           calendarId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseMeCalendarGroupIdCalendarIdCalendarPermissionID parses 'input' into a MeCalendarGroupIdCalendarIdCalendarPermissionId
func ParseMeCalendarGroupIdCalendarIdCalendarPermissionID(input string) (*MeCalendarGroupIdCalendarIdCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCalendarGroupIdCalendarIdCalendarPermissionIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupIdCalendarIdCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupIdCalendarIdCalendarPermissionIDInsensitively(input string) (*MeCalendarGroupIdCalendarIdCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCalendarGroupIdCalendarIdCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCalendarGroupIdCalendarIdCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCalendarGroupIdCalendarIdCalendarPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CalendarGroupId, ok = input.Parsed["calendarGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.CalendarPermissionId, ok = input.Parsed["calendarPermissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", input)
	}

	return nil
}

// ValidateMeCalendarGroupIdCalendarIdCalendarPermissionID checks that 'input' can be parsed as a Me Calendar Group Id Calendar Id Calendar Permission ID
func ValidateMeCalendarGroupIdCalendarIdCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupIdCalendarIdCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Id Calendar Id Calendar Permission ID
func (id MeCalendarGroupIdCalendarIdCalendarPermissionId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Id Calendar Id Calendar Permission ID
func (id MeCalendarGroupIdCalendarIdCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Id Calendar Id Calendar Permission ID
func (id MeCalendarGroupIdCalendarIdCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("Me Calendar Group Id Calendar Id Calendar Permission (%s)", strings.Join(components, "\n"))
}
