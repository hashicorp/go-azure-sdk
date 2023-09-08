package groupcalendarcalendarpermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = GroupCalendarCalendarPermissionId{}

// GroupCalendarCalendarPermissionId is a struct representing the Resource ID for a Group Calendar Calendar Permission
type GroupCalendarCalendarPermissionId struct {
	GroupId              string
	CalendarPermissionId string
}

// NewGroupCalendarCalendarPermissionID returns a new GroupCalendarCalendarPermissionId struct
func NewGroupCalendarCalendarPermissionID(groupId string, calendarPermissionId string) GroupCalendarCalendarPermissionId {
	return GroupCalendarCalendarPermissionId{
		GroupId:              groupId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseGroupCalendarCalendarPermissionID parses 'input' into a GroupCalendarCalendarPermissionId
func ParseGroupCalendarCalendarPermissionID(input string) (*GroupCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarPermissionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.CalendarPermissionId, ok = parsed.Parsed["calendarPermissionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", *parsed)
	}

	return &id, nil
}

// ParseGroupCalendarCalendarPermissionIDInsensitively parses 'input' case-insensitively into a GroupCalendarCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseGroupCalendarCalendarPermissionIDInsensitively(input string) (*GroupCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GroupCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GroupCalendarCalendarPermissionId{}

	if id.GroupId, ok = parsed.Parsed["groupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupId", *parsed)
	}

	if id.CalendarPermissionId, ok = parsed.Parsed["calendarPermissionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", *parsed)
	}

	return &id, nil
}

// ValidateGroupCalendarCalendarPermissionID checks that 'input' can be parsed as a Group Calendar Calendar Permission ID
func ValidateGroupCalendarCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupCalendarCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Calendar Calendar Permission ID
func (id GroupCalendarCalendarPermissionId) ID() string {
	fmtString := "/groups/%s/calendar/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Calendar Calendar Permission ID
func (id GroupCalendarCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this Group Calendar Calendar Permission ID
func (id GroupCalendarCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("Group Calendar Calendar Permission (%s)", strings.Join(components, "\n"))
}
