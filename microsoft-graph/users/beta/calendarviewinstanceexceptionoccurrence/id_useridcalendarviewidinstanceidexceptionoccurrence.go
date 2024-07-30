package calendarviewinstanceexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarViewIdInstanceIdExceptionOccurrenceId{}

// UserIdCalendarViewIdInstanceIdExceptionOccurrenceId is a struct representing the Resource ID for a User Id Calendar View Id Instance Id Exception Occurrence
type UserIdCalendarViewIdInstanceIdExceptionOccurrenceId struct {
	UserId   string
	EventId  string
	EventId1 string
	EventId2 string
}

// NewUserIdCalendarViewIdInstanceIdExceptionOccurrenceID returns a new UserIdCalendarViewIdInstanceIdExceptionOccurrenceId struct
func NewUserIdCalendarViewIdInstanceIdExceptionOccurrenceID(userId string, eventId string, eventId1 string, eventId2 string) UserIdCalendarViewIdInstanceIdExceptionOccurrenceId {
	return UserIdCalendarViewIdInstanceIdExceptionOccurrenceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseUserIdCalendarViewIdInstanceIdExceptionOccurrenceID parses 'input' into a UserIdCalendarViewIdInstanceIdExceptionOccurrenceId
func ParseUserIdCalendarViewIdInstanceIdExceptionOccurrenceID(input string) (*UserIdCalendarViewIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserIdCalendarViewIdInstanceIdExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarViewIdInstanceIdExceptionOccurrenceIDInsensitively(input string) (*UserIdCalendarViewIdInstanceIdExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdInstanceIdExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdInstanceIdExceptionOccurrenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarViewIdInstanceIdExceptionOccurrenceId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	return nil
}

// ValidateUserIdCalendarViewIdInstanceIdExceptionOccurrenceID checks that 'input' can be parsed as a User Id Calendar View Id Instance Id Exception Occurrence ID
func ValidateUserIdCalendarViewIdInstanceIdExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarViewIdInstanceIdExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar View Id Instance Id Exception Occurrence ID
func (id UserIdCalendarViewIdInstanceIdExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendarView/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar View Id Instance Id Exception Occurrence ID
func (id UserIdCalendarViewIdInstanceIdExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Id Calendar View Id Instance Id Exception Occurrence ID
func (id UserIdCalendarViewIdInstanceIdExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Id Calendar View Id Instance Id Exception Occurrence (%s)", strings.Join(components, "\n"))
}
