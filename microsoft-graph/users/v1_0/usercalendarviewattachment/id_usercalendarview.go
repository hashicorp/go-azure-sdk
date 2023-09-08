package usercalendarviewattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarViewId{}

// UserCalendarViewId is a struct representing the Resource ID for a User Calendar View
type UserCalendarViewId struct {
	UserId  string
	EventId string
}

// NewUserCalendarViewID returns a new UserCalendarViewId struct
func NewUserCalendarViewID(userId string, eventId string) UserCalendarViewId {
	return UserCalendarViewId{
		UserId:  userId,
		EventId: eventId,
	}
}

// ParseUserCalendarViewID parses 'input' into a UserCalendarViewId
func ParseUserCalendarViewID(input string) (*UserCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarViewIDInsensitively parses 'input' case-insensitively into a UserCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarViewIDInsensitively(input string) (*UserCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarViewID checks that 'input' can be parsed as a User Calendar View ID
func ValidateUserCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar View ID
func (id UserCalendarViewId) ID() string {
	fmtString := "/users/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar View ID
func (id UserCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Calendar View ID
func (id UserCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Calendar View (%s)", strings.Join(components, "\n"))
}
