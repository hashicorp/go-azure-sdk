package usercalendargroupcalendarcalendarpermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarCalendarPermissionId{}

// UserCalendarGroupCalendarCalendarPermissionId is a struct representing the Resource ID for a User Calendar Group Calendar Calendar Permission
type UserCalendarGroupCalendarCalendarPermissionId struct {
	UserId               string
	CalendarGroupId      string
	CalendarId           string
	CalendarPermissionId string
}

// NewUserCalendarGroupCalendarCalendarPermissionID returns a new UserCalendarGroupCalendarCalendarPermissionId struct
func NewUserCalendarGroupCalendarCalendarPermissionID(userId string, calendarGroupId string, calendarId string, calendarPermissionId string) UserCalendarGroupCalendarCalendarPermissionId {
	return UserCalendarGroupCalendarCalendarPermissionId{
		UserId:               userId,
		CalendarGroupId:      calendarGroupId,
		CalendarId:           calendarId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseUserCalendarGroupCalendarCalendarPermissionID parses 'input' into a UserCalendarGroupCalendarCalendarPermissionId
func ParseUserCalendarGroupCalendarCalendarPermissionID(input string) (*UserCalendarGroupCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarCalendarPermissionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ParseUserCalendarGroupCalendarCalendarPermissionIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupCalendarCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupCalendarCalendarPermissionIDInsensitively(input string) (*UserCalendarGroupCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarCalendarPermissionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

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

// ValidateUserCalendarGroupCalendarCalendarPermissionID checks that 'input' can be parsed as a User Calendar Group Calendar Calendar Permission ID
func ValidateUserCalendarGroupCalendarCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupCalendarCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group Calendar Calendar Permission ID
func (id UserCalendarGroupCalendarCalendarPermissionId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group Calendar Calendar Permission ID
func (id UserCalendarGroupCalendarCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Group Calendar Calendar Permission ID
func (id UserCalendarGroupCalendarCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("User Calendar Group Calendar Calendar Permission (%s)", strings.Join(components, "\n"))
}
