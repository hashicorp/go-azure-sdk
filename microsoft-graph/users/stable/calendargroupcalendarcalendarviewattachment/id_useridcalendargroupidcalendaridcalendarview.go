package calendargroupcalendarcalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarGroupIdCalendarIdCalendarViewId{}

// UserIdCalendarGroupIdCalendarIdCalendarViewId is a struct representing the Resource ID for a User Id Calendar Group Id Calendar Id Calendar View
type UserIdCalendarGroupIdCalendarIdCalendarViewId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
}

// NewUserIdCalendarGroupIdCalendarIdCalendarViewID returns a new UserIdCalendarGroupIdCalendarIdCalendarViewId struct
func NewUserIdCalendarGroupIdCalendarIdCalendarViewID(userId string, calendarGroupId string, calendarId string, eventId string) UserIdCalendarGroupIdCalendarIdCalendarViewId {
	return UserIdCalendarGroupIdCalendarIdCalendarViewId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
	}
}

// ParseUserIdCalendarGroupIdCalendarIdCalendarViewID parses 'input' into a UserIdCalendarGroupIdCalendarIdCalendarViewId
func ParseUserIdCalendarGroupIdCalendarIdCalendarViewID(input string) (*UserIdCalendarGroupIdCalendarIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupIdCalendarIdCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupIdCalendarIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarGroupIdCalendarIdCalendarViewIDInsensitively parses 'input' case-insensitively into a UserIdCalendarGroupIdCalendarIdCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarGroupIdCalendarIdCalendarViewIDInsensitively(input string) (*UserIdCalendarGroupIdCalendarIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupIdCalendarIdCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupIdCalendarIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarGroupIdCalendarIdCalendarViewId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateUserIdCalendarGroupIdCalendarIdCalendarViewID checks that 'input' can be parsed as a User Id Calendar Group Id Calendar Id Calendar View ID
func ValidateUserIdCalendarGroupIdCalendarIdCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarGroupIdCalendarIdCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Group Id Calendar Id Calendar View ID
func (id UserIdCalendarGroupIdCalendarIdCalendarViewId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Group Id Calendar Id Calendar View ID
func (id UserIdCalendarGroupIdCalendarIdCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Group Id Calendar Id Calendar View ID
func (id UserIdCalendarGroupIdCalendarIdCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Id Calendar Group Id Calendar Id Calendar View (%s)", strings.Join(components, "\n"))
}
