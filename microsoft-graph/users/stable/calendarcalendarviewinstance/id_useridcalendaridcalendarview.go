package calendarcalendarviewinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdCalendarViewId{}

// UserIdCalendarIdCalendarViewId is a struct representing the Resource ID for a User Id Calendar Id Calendar View
type UserIdCalendarIdCalendarViewId struct {
	UserId     string
	CalendarId string
	EventId    string
}

// NewUserIdCalendarIdCalendarViewID returns a new UserIdCalendarIdCalendarViewId struct
func NewUserIdCalendarIdCalendarViewID(userId string, calendarId string, eventId string) UserIdCalendarIdCalendarViewId {
	return UserIdCalendarIdCalendarViewId{
		UserId:     userId,
		CalendarId: calendarId,
		EventId:    eventId,
	}
}

// ParseUserIdCalendarIdCalendarViewID parses 'input' into a UserIdCalendarIdCalendarViewId
func ParseUserIdCalendarIdCalendarViewID(input string) (*UserIdCalendarIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdCalendarViewIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdCalendarViewIDInsensitively(input string) (*UserIdCalendarIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdCalendarViewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateUserIdCalendarIdCalendarViewID checks that 'input' can be parsed as a User Id Calendar Id Calendar View ID
func ValidateUserIdCalendarIdCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Calendar View ID
func (id UserIdCalendarIdCalendarViewId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Calendar View ID
func (id UserIdCalendarIdCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Calendar View ID
func (id UserIdCalendarIdCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Id Calendar Id Calendar View (%s)", strings.Join(components, "\n"))
}
