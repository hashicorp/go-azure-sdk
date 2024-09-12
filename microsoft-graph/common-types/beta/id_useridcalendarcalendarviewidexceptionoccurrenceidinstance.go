package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId{}

// UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId is a struct representing the Resource ID for a User Id Calendar Calendar View Id Exception Occurrence Id Instance
type UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId struct {
	UserId   string
	EventId  string
	EventId1 string
	EventId2 string
}

// NewUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID returns a new UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId struct
func NewUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID(userId string, eventId string, eventId1 string, eventId2 string) UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId {
	return UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID parses 'input' into a UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId
func ParseUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID(input string) (*UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceIDInsensitively parses 'input' case-insensitively into a UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceIDInsensitively(input string) (*UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID checks that 'input' can be parsed as a User Id Calendar Calendar View Id Exception Occurrence Id Instance ID
func ValidateUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Calendar View Id Exception Occurrence Id Instance ID
func (id UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Calendar View Id Exception Occurrence Id Instance ID
func (id UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Id Calendar Calendar View Id Exception Occurrence Id Instance ID
func (id UserIdCalendarCalendarViewIdExceptionOccurrenceIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Id Calendar Calendar View Id Exception Occurrence Id Instance (%s)", strings.Join(components, "\n"))
}
