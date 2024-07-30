package calendarcalendarpermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarCalendarPermissionId{}

// UserIdCalendarCalendarPermissionId is a struct representing the Resource ID for a User Id Calendar Calendar Permission
type UserIdCalendarCalendarPermissionId struct {
	UserId               string
	CalendarPermissionId string
}

// NewUserIdCalendarCalendarPermissionID returns a new UserIdCalendarCalendarPermissionId struct
func NewUserIdCalendarCalendarPermissionID(userId string, calendarPermissionId string) UserIdCalendarCalendarPermissionId {
	return UserIdCalendarCalendarPermissionId{
		UserId:               userId,
		CalendarPermissionId: calendarPermissionId,
	}
}

// ParseUserIdCalendarCalendarPermissionID parses 'input' into a UserIdCalendarCalendarPermissionId
func ParseUserIdCalendarCalendarPermissionID(input string) (*UserIdCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarCalendarPermissionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarCalendarPermissionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarCalendarPermissionIDInsensitively(input string) (*UserIdCalendarCalendarPermissionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarCalendarPermissionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarCalendarPermissionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarCalendarPermissionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarPermissionId, ok = input.Parsed["calendarPermissionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarPermissionId", input)
	}

	return nil
}

// ValidateUserIdCalendarCalendarPermissionID checks that 'input' can be parsed as a User Id Calendar Calendar Permission ID
func ValidateUserIdCalendarCalendarPermissionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarCalendarPermissionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Calendar Permission ID
func (id UserIdCalendarCalendarPermissionId) ID() string {
	fmtString := "/users/%s/calendar/calendarPermissions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarPermissionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Calendar Permission ID
func (id UserIdCalendarCalendarPermissionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarPermissions", "calendarPermissions", "calendarPermissions"),
		resourceids.UserSpecifiedSegment("calendarPermissionId", "calendarPermissionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Calendar Permission ID
func (id UserIdCalendarCalendarPermissionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Permission: %q", id.CalendarPermissionId),
	}
	return fmt.Sprintf("User Id Calendar Calendar Permission (%s)", strings.Join(components, "\n"))
}
