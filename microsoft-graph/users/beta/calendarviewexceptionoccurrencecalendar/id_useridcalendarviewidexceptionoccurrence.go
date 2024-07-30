package calendarviewexceptionoccurrencecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarViewIdExceptionOccurrenceId{}

// UserIdCalendarViewIdExceptionOccurrenceId is a struct representing the Resource ID for a User Id Calendar View Id Exception Occurrence
type UserIdCalendarViewIdExceptionOccurrenceId struct {
	UserId   string
	EventId  string
	EventId1 string
}

// NewUserIdCalendarViewIdExceptionOccurrenceID returns a new UserIdCalendarViewIdExceptionOccurrenceId struct
func NewUserIdCalendarViewIdExceptionOccurrenceID(userId string, eventId string, eventId1 string) UserIdCalendarViewIdExceptionOccurrenceId {
	return UserIdCalendarViewIdExceptionOccurrenceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseUserIdCalendarViewIdExceptionOccurrenceID parses 'input' into a UserIdCalendarViewIdExceptionOccurrenceId
func ParseUserIdCalendarViewIdExceptionOccurrenceID(input string) (*UserIdCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarViewIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserIdCalendarViewIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarViewIdExceptionOccurrenceIDInsensitively(input string) (*UserIdCalendarViewIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarViewIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateUserIdCalendarViewIdExceptionOccurrenceID checks that 'input' can be parsed as a User Id Calendar View Id Exception Occurrence ID
func ValidateUserIdCalendarViewIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarViewIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar View Id Exception Occurrence ID
func (id UserIdCalendarViewIdExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar View Id Exception Occurrence ID
func (id UserIdCalendarViewIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Id Calendar View Id Exception Occurrence ID
func (id UserIdCalendarViewIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Id Calendar View Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
