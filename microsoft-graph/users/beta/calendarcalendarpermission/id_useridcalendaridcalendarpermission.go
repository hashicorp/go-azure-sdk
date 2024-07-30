package calendarcalendarpermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdCalendarPermissionId{}

// UserIdCalendarIdCalendarPermissionId is a struct representing the Resource ID for a User Id Calendar Id Calendar Permission
type UserIdCalendarIdCalendarPermissionId struct {
	UserId               string
	CalendarId           string
	CalendarPermissionId string
}

// NewUserIdCalendarIdCalendarPermissionID returns a new UserIdCalendarIdCalendarPermissionId struct
func NewUserIdCalendarIdCalendarPermissionID(userId string, calendarId string, calendarPermissionId string) UserIdCalendarIdCalendarPermissionId {
	return UserIdCalendarIdCalendarPermissionId{
		UserId:               userId,
		CalendarId:           calendarId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseUserIdCalendarIdCalendarPermissionID parses 'input' into a UserIdCalendarIdCalendarPermissionId
func ParseUserIdCalendarIdCalendarPermissionID(input string) (*UserIdCalendarIdCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdCalendarPermissionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdCalendarPermissionIDInsensitively(input string) (*UserIdCalendarIdCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdCalendarPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.CalendarPermissionId, ok = input.Parsed["calendarPermissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", input)
	}

	return nil
}

// ValidateUserIdCalendarIdCalendarPermissionID checks that 'input' can be parsed as a User Id Calendar Id Calendar Permission ID
func ValidateUserIdCalendarIdCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Calendar Permission ID
func (id UserIdCalendarIdCalendarPermissionId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Calendar Permission ID
func (id UserIdCalendarIdCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Calendar Permission ID
func (id UserIdCalendarIdCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("User Id Calendar Id Calendar Permission (%s)", strings.Join(components, "\n"))
}
