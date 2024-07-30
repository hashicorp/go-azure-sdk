package calendarcalendarviewinstanceexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdCalendarViewIdInstanceId{}

// UserIdCalendarIdCalendarViewIdInstanceId is a struct representing the Resource ID for a User Id Calendar Id Calendar View Id Instance
type UserIdCalendarIdCalendarViewIdInstanceId struct {
	UserId     string
	CalendarId string
	EventId    string
	EventId1   string
}

// NewUserIdCalendarIdCalendarViewIdInstanceID returns a new UserIdCalendarIdCalendarViewIdInstanceId struct
func NewUserIdCalendarIdCalendarViewIdInstanceID(userId string, calendarId string, eventId string, eventId1 string) UserIdCalendarIdCalendarViewIdInstanceId {
	return UserIdCalendarIdCalendarViewIdInstanceId{
		UserId:     userId,
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
	}
}

// ParseUserIdCalendarIdCalendarViewIdInstanceID parses 'input' into a UserIdCalendarIdCalendarViewIdInstanceId
func ParseUserIdCalendarIdCalendarViewIdInstanceID(input string) (*UserIdCalendarIdCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdCalendarViewIdInstanceIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdCalendarViewIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdCalendarViewIdInstanceIDInsensitively(input string) (*UserIdCalendarIdCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdCalendarViewIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateUserIdCalendarIdCalendarViewIdInstanceID checks that 'input' can be parsed as a User Id Calendar Id Calendar View Id Instance ID
func ValidateUserIdCalendarIdCalendarViewIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdCalendarViewIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Calendar View Id Instance ID
func (id UserIdCalendarIdCalendarViewIdInstanceId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Calendar View Id Instance ID
func (id UserIdCalendarIdCalendarViewIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Calendar View Id Instance ID
func (id UserIdCalendarIdCalendarViewIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Id Calendar Id Calendar View Id Instance (%s)", strings.Join(components, "\n"))
}
