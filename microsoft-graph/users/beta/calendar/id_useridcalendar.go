package calendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarId{}

// UserIdCalendarId is a struct representing the Resource ID for a User Id Calendar
type UserIdCalendarId struct {
	UserId     string
	CalendarId string
}

// NewUserIdCalendarID returns a new UserIdCalendarId struct
func NewUserIdCalendarID(userId string, calendarId string) UserIdCalendarId {
	return UserIdCalendarId{
		UserId:     userId,
		CalendarId: calendarId,
	}
}

// ParseUserIdCalendarID parses 'input' into a UserIdCalendarId
func ParseUserIdCalendarID(input string) (*UserIdCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIDInsensitively parses 'input' case-insensitively into a UserIdCalendarId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIDInsensitively(input string) (*UserIdCalendarId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	return nil
}

// ValidateUserIdCalendarID checks that 'input' can be parsed as a User Id Calendar ID
func ValidateUserIdCalendarID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar ID
func (id UserIdCalendarId) ID() string {
	fmtString := "/users/%s/calendars/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar ID
func (id UserIdCalendarId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar ID
func (id UserIdCalendarId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
	}
	return fmt.Sprintf("User Id Calendar (%s)", strings.Join(components, "\n"))
}
