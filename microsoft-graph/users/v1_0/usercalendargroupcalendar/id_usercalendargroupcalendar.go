package usercalendargroupcalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarId{}

// UserCalendarGroupCalendarId is a struct representing the Resource ID for a User Calendar Group Calendar
type UserCalendarGroupCalendarId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
}

// NewUserCalendarGroupCalendarID returns a new UserCalendarGroupCalendarId struct
func NewUserCalendarGroupCalendarID(userId string, calendarGroupId string, calendarId string) UserCalendarGroupCalendarId {
	return UserCalendarGroupCalendarId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
	}
}

// ParseUserCalendarGroupCalendarID parses 'input' into a UserCalendarGroupCalendarId
func ParseUserCalendarGroupCalendarID(input string) (*UserCalendarGroupCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarGroupCalendarIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupCalendarId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupCalendarIDInsensitively(input string) (*UserCalendarGroupCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarGroupCalendarID checks that 'input' can be parsed as a User Calendar Group Calendar ID
func ValidateUserCalendarGroupCalendarID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupCalendarID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group Calendar ID
func (id UserCalendarGroupCalendarId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group Calendar ID
func (id UserCalendarGroupCalendarId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Group Calendar ID
func (id UserCalendarGroupCalendarId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
	}
	return fmt.Sprintf("User Calendar Group Calendar (%s)", strings.Join(components, "\n"))
}
