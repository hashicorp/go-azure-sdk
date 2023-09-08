package usercalendarcalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarViewInstanceExceptionOccurrenceId{}

// UserCalendarCalendarViewInstanceExceptionOccurrenceId is a struct representing the Resource ID for a User Calendar Calendar View Instance Exception Occurrence
type UserCalendarCalendarViewInstanceExceptionOccurrenceId struct {
	UserId   string
	EventId  string
	EventId1 string
	EventId2 string
}

// NewUserCalendarCalendarViewInstanceExceptionOccurrenceID returns a new UserCalendarCalendarViewInstanceExceptionOccurrenceId struct
func NewUserCalendarCalendarViewInstanceExceptionOccurrenceID(userId string, eventId string, eventId1 string, eventId2 string) UserCalendarCalendarViewInstanceExceptionOccurrenceId {
	return UserCalendarCalendarViewInstanceExceptionOccurrenceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseUserCalendarCalendarViewInstanceExceptionOccurrenceID parses 'input' into a UserCalendarCalendarViewInstanceExceptionOccurrenceId
func ParseUserCalendarCalendarViewInstanceExceptionOccurrenceID(input string) (*UserCalendarCalendarViewInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewInstanceExceptionOccurrenceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarCalendarViewInstanceExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarViewInstanceExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarViewInstanceExceptionOccurrenceIDInsensitively(input string) (*UserCalendarCalendarViewInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewInstanceExceptionOccurrenceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarCalendarViewInstanceExceptionOccurrenceID checks that 'input' can be parsed as a User Calendar Calendar View Instance Exception Occurrence ID
func ValidateUserCalendarCalendarViewInstanceExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarViewInstanceExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar View Instance Exception Occurrence ID
func (id UserCalendarCalendarViewInstanceExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar View Instance Exception Occurrence ID
func (id UserCalendarCalendarViewInstanceExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Calendar Calendar View Instance Exception Occurrence ID
func (id UserCalendarCalendarViewInstanceExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Calendar Calendar View Instance Exception Occurrence (%s)", strings.Join(components, "\n"))
}
