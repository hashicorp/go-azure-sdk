package calendarviewexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarViewId{}

// UserIdCalendarViewId is a struct representing the Resource ID for a User Id Calendar View
type UserIdCalendarViewId struct {
	UserId  string
	EventId string
}

// NewUserIdCalendarViewID returns a new UserIdCalendarViewId struct
func NewUserIdCalendarViewID(userId string, eventId string) UserIdCalendarViewId {
	return UserIdCalendarViewId{
		UserId:  userId,
		EventId: eventId,
	}
}

// ParseUserIdCalendarViewID parses 'input' into a UserIdCalendarViewId
func ParseUserIdCalendarViewID(input string) (*UserIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarViewIDInsensitively parses 'input' case-insensitively into a UserIdCalendarViewId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarViewIDInsensitively(input string) (*UserIdCalendarViewId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarViewId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateUserIdCalendarViewID checks that 'input' can be parsed as a User Id Calendar View ID
func ValidateUserIdCalendarViewID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarViewID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar View ID
func (id UserIdCalendarViewId) ID() string {
	fmtString := "/users/%s/calendarView/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar View ID
func (id UserIdCalendarViewId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar View ID
func (id UserIdCalendarViewId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Id Calendar View (%s)", strings.Join(components, "\n"))
}
