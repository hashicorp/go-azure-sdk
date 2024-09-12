package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarGroupIdCalendarIdCalendarPermissionId{}

// UserIdCalendarGroupIdCalendarIdCalendarPermissionId is a struct representing the Resource ID for a User Id Calendar Group Id Calendar Id Calendar Permission
type UserIdCalendarGroupIdCalendarIdCalendarPermissionId struct {
	UserId               string
	CalendarGroupId      string
	CalendarId           string
	CalendarPermissionId string
}

// NewUserIdCalendarGroupIdCalendarIdCalendarPermissionID returns a new UserIdCalendarGroupIdCalendarIdCalendarPermissionId struct
func NewUserIdCalendarGroupIdCalendarIdCalendarPermissionID(userId string, calendarGroupId string, calendarId string, calendarPermissionId string) UserIdCalendarGroupIdCalendarIdCalendarPermissionId {
	return UserIdCalendarGroupIdCalendarIdCalendarPermissionId{
		UserId:               userId,
		CalendarGroupId:      calendarGroupId,
		CalendarId:           calendarId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseUserIdCalendarGroupIdCalendarIdCalendarPermissionID parses 'input' into a UserIdCalendarGroupIdCalendarIdCalendarPermissionId
func ParseUserIdCalendarGroupIdCalendarIdCalendarPermissionID(input string) (*UserIdCalendarGroupIdCalendarIdCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupIdCalendarIdCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupIdCalendarIdCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarGroupIdCalendarIdCalendarPermissionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarGroupIdCalendarIdCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarGroupIdCalendarIdCalendarPermissionIDInsensitively(input string) (*UserIdCalendarGroupIdCalendarIdCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupIdCalendarIdCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupIdCalendarIdCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarGroupIdCalendarIdCalendarPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

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

// ValidateUserIdCalendarGroupIdCalendarIdCalendarPermissionID checks that 'input' can be parsed as a User Id Calendar Group Id Calendar Id Calendar Permission ID
func ValidateUserIdCalendarGroupIdCalendarIdCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarGroupIdCalendarIdCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Group Id Calendar Id Calendar Permission ID
func (id UserIdCalendarGroupIdCalendarIdCalendarPermissionId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Group Id Calendar Id Calendar Permission ID
func (id UserIdCalendarGroupIdCalendarIdCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Group Id Calendar Id Calendar Permission ID
func (id UserIdCalendarGroupIdCalendarIdCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("User Id Calendar Group Id Calendar Id Calendar Permission (%s)", strings.Join(components, "\n"))
}
