package usercalendarcalendarpermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarPermissionId{}

// UserCalendarCalendarPermissionId is a struct representing the Resource ID for a User Calendar Calendar Permission
type UserCalendarCalendarPermissionId struct {
	UserId               string
	CalendarId           string
	CalendarPermissionId string
}

// NewUserCalendarCalendarPermissionID returns a new UserCalendarCalendarPermissionId struct
func NewUserCalendarCalendarPermissionID(userId string, calendarId string, calendarPermissionId string) UserCalendarCalendarPermissionId {
	return UserCalendarCalendarPermissionId{
		UserId:               userId,
		CalendarId:           calendarId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseUserCalendarCalendarPermissionID parses 'input' into a UserCalendarCalendarPermissionId
func ParseUserCalendarCalendarPermissionID(input string) (*UserCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarPermissionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.CalendarPermissionId, ok = parsed.Parsed["calendarPermissionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarCalendarPermissionIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarPermissionIDInsensitively(input string) (*UserCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarPermissionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.CalendarPermissionId, ok = parsed.Parsed["calendarPermissionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarCalendarPermissionID checks that 'input' can be parsed as a User Calendar Calendar Permission ID
func ValidateUserCalendarCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar Permission ID
func (id UserCalendarCalendarPermissionId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar Permission ID
func (id UserCalendarCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Calendar Permission ID
func (id UserCalendarCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("User Calendar Calendar Permission (%s)", strings.Join(components, "\n"))
}
