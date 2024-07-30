package calendarcalendarviewinstanceexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId{}

// UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId is a struct representing the Resource ID for a User Id Calendar Id Calendar View Id Instance Id Exception Occurrence
type UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId struct {
	UserId     string
	CalendarId string
	EventId    string
	EventId1   string
	EventId2   string
}

// NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID returns a new UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId struct
func NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID(userId string, calendarId string, eventId string, eventId1 string, eventId2 string) UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId {
	return UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId{
		UserId:     userId,
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
		EventId2:   eventId2,
	}
}

// ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID parses 'input' into a UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId
func ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID(input string) (*UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively(input string) (*UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	return nil
}

// ValidateUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID checks that 'input' can be parsed as a User Id Calendar Id Calendar View Id Instance Id Exception Occurrence ID
func ValidateUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Calendar View Id Instance Id Exception Occurrence ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarView/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Calendar View Id Instance Id Exception Occurrence ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Calendar View Id Instance Id Exception Occurrence ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Id Calendar Id Calendar View Id Instance Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
