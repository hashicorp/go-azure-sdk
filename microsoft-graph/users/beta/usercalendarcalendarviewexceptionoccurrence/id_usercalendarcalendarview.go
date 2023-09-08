package usercalendarcalendarviewexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarViewId{}

// UserCalendarCalendarViewId is a struct representing the Resource ID for a User Calendar Calendar View
type UserCalendarCalendarViewId struct {
	UserId  string
	EventId string
}

// NewUserCalendarCalendarViewID returns a new UserCalendarCalendarViewId struct
func NewUserCalendarCalendarViewID(userId string, eventId string) UserCalendarCalendarViewId {
	return UserCalendarCalendarViewId{
		UserId:  userId,
		EventId: eventId,
	}
}

// ParseUserCalendarCalendarViewID parses 'input' into a UserCalendarCalendarViewId
func ParseUserCalendarCalendarViewID(input string) (*UserCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarCalendarViewIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarViewIDInsensitively(input string) (*UserCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarCalendarViewID checks that 'input' can be parsed as a User Calendar Calendar View ID
func ValidateUserCalendarCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar View ID
func (id UserCalendarCalendarViewId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar View ID
func (id UserCalendarCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Calendar View ID
func (id UserCalendarCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Calendar Calendar View (%s)", strings.Join(components, "\n"))
}
