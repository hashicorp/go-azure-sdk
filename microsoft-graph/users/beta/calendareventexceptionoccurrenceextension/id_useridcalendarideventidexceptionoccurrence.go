package calendareventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdEventIdExceptionOccurrenceId{}

// UserIdCalendarIdEventIdExceptionOccurrenceId is a struct representing the Resource ID for a User Id Calendar Id Event Id Exception Occurrence
type UserIdCalendarIdEventIdExceptionOccurrenceId struct {
	UserId     string
	CalendarId string
	EventId    string
	EventId1   string
}

// NewUserIdCalendarIdEventIdExceptionOccurrenceID returns a new UserIdCalendarIdEventIdExceptionOccurrenceId struct
func NewUserIdCalendarIdEventIdExceptionOccurrenceID(userId string, calendarId string, eventId string, eventId1 string) UserIdCalendarIdEventIdExceptionOccurrenceId {
	return UserIdCalendarIdEventIdExceptionOccurrenceId{
		UserId:     userId,
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
	}
}

// ParseUserIdCalendarIdEventIdExceptionOccurrenceID parses 'input' into a UserIdCalendarIdEventIdExceptionOccurrenceId
func ParseUserIdCalendarIdEventIdExceptionOccurrenceID(input string) (*UserIdCalendarIdEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdEventIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdEventIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdEventIdExceptionOccurrenceIDInsensitively(input string) (*UserIdCalendarIdEventIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdEventIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdEventIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdEventIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdCalendarIdEventIdExceptionOccurrenceID checks that 'input' can be parsed as a User Id Calendar Id Event Id Exception Occurrence ID
func ValidateUserIdCalendarIdEventIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdEventIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Event Id Exception Occurrence ID
func (id UserIdCalendarIdEventIdExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Event Id Exception Occurrence ID
func (id UserIdCalendarIdEventIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Event Id Exception Occurrence ID
func (id UserIdCalendarIdEventIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Id Calendar Id Event Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
