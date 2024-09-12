package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarCalendarViewId{}

// UserIdCalendarCalendarViewId is a struct representing the Resource ID for a User Id Calendar Calendar View
type UserIdCalendarCalendarViewId struct {
	UserId  string
	EventId string
}

// NewUserIdCalendarCalendarViewID returns a new UserIdCalendarCalendarViewId struct
func NewUserIdCalendarCalendarViewID(userId string, eventId string) UserIdCalendarCalendarViewId {
	return UserIdCalendarCalendarViewId{
		UserId:  userId,
		EventId: eventId,
	}
}

// ParseUserIdCalendarCalendarViewID parses 'input' into a UserIdCalendarCalendarViewId
func ParseUserIdCalendarCalendarViewID(input string) (*UserIdCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarCalendarViewIDInsensitively parses 'input' case-insensitively into a UserIdCalendarCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarCalendarViewIDInsensitively(input string) (*UserIdCalendarCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarCalendarViewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateUserIdCalendarCalendarViewID checks that 'input' can be parsed as a User Id Calendar Calendar View ID
func ValidateUserIdCalendarCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Calendar View ID
func (id UserIdCalendarCalendarViewId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Calendar View ID
func (id UserIdCalendarCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Calendar View ID
func (id UserIdCalendarCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Id Calendar Calendar View (%s)", strings.Join(components, "\n"))
}
