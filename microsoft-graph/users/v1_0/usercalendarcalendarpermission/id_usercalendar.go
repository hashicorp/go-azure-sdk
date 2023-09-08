package usercalendarcalendarpermission

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarId{}

// UserCalendarId is a struct representing the Resource ID for a User Calendar
type UserCalendarId struct {
	UserId     string
	CalendarId string
}

// NewUserCalendarID returns a new UserCalendarId struct
func NewUserCalendarID(userId string, calendarId string) UserCalendarId {
	return UserCalendarId{
		UserId:     userId,
		CalendarId: calendarId,
	}
}

// ParseUserCalendarID parses 'input' into a UserCalendarId
func ParseUserCalendarID(input string) (*UserCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarIDInsensitively parses 'input' case-insensitively into a UserCalendarId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarIDInsensitively(input string) (*UserCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarID checks that 'input' can be parsed as a User Calendar ID
func ValidateUserCalendarID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar ID
func (id UserCalendarId) ID() string {
	fmtString := "/users/%s/calendars/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar ID
func (id UserCalendarId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
	}
}

// String returns a human-readable description of this User Calendar ID
func (id UserCalendarId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
	}
	return fmt.Sprintf("User Calendar (%s)", strings.Join(components, "\n"))
}
