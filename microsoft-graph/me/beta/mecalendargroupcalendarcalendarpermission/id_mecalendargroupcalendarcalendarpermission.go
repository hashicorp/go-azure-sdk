package mecalendargroupcalendarcalendarpermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeCalendarGroupCalendarCalendarPermissionId{}

// MeCalendarGroupCalendarCalendarPermissionId is a struct representing the Resource ID for a Me Calendar Group Calendar Calendar Permission
type MeCalendarGroupCalendarCalendarPermissionId struct {
	CalendarGroupId      string
	CalendarId           string
	CalendarPermissionId string
}

// NewMeCalendarGroupCalendarCalendarPermissionID returns a new MeCalendarGroupCalendarCalendarPermissionId struct
func NewMeCalendarGroupCalendarCalendarPermissionID(calendarGroupId string, calendarId string, calendarPermissionId string) MeCalendarGroupCalendarCalendarPermissionId {
	return MeCalendarGroupCalendarCalendarPermissionId{
		CalendarGroupId:      calendarGroupId,
		CalendarId:           calendarId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseMeCalendarGroupCalendarCalendarPermissionID parses 'input' into a MeCalendarGroupCalendarCalendarPermissionId
func ParseMeCalendarGroupCalendarCalendarPermissionID(input string) (*MeCalendarGroupCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarPermissionId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.CalendarPermissionId, ok = parsed.Parsed["calendarPermissionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", *parsed)
	}

	return &id, nil
}

// ParseMeCalendarGroupCalendarCalendarPermissionIDInsensitively parses 'input' case-insensitively into a MeCalendarGroupCalendarCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseMeCalendarGroupCalendarCalendarPermissionIDInsensitively(input string) (*MeCalendarGroupCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeCalendarGroupCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeCalendarGroupCalendarCalendarPermissionId{}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.CalendarPermissionId, ok = parsed.Parsed["calendarPermissionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", *parsed)
	}

	return &id, nil
}

// ValidateMeCalendarGroupCalendarCalendarPermissionID checks that 'input' can be parsed as a Me Calendar Group Calendar Calendar Permission ID
func ValidateMeCalendarGroupCalendarCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCalendarGroupCalendarCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Calendar Group Calendar Calendar Permission ID
func (id MeCalendarGroupCalendarCalendarPermissionId) ID() string {
	fmtString := "/me/calendarGroups/%s/calendars/%s/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.CalendarGroupId, id.CalendarId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Calendar Group Calendar Calendar Permission ID
func (id MeCalendarGroupCalendarCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this Me Calendar Group Calendar Calendar Permission ID
func (id MeCalendarGroupCalendarCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("Me Calendar Group Calendar Calendar Permission (%s)", strings.Join(components, "\n"))
}
