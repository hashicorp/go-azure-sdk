package usercalendargroupcalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupId{}

// UserCalendarGroupId is a struct representing the Resource ID for a User Calendar Group
type UserCalendarGroupId struct {
	UserId          string
	CalendarGroupId string
}

// NewUserCalendarGroupID returns a new UserCalendarGroupId struct
func NewUserCalendarGroupID(userId string, calendarGroupId string) UserCalendarGroupId {
	return UserCalendarGroupId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
	}
}

// ParseUserCalendarGroupID parses 'input' into a UserCalendarGroupId
func ParseUserCalendarGroupID(input string) (*UserCalendarGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarGroupIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupIDInsensitively(input string) (*UserCalendarGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarGroupID checks that 'input' can be parsed as a User Calendar Group ID
func ValidateUserCalendarGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group ID
func (id UserCalendarGroupId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group ID
func (id UserCalendarGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Group ID
func (id UserCalendarGroupId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
	}
	return fmt.Sprintf("User Calendar Group (%s)", strings.Join(components, "\n"))
}
