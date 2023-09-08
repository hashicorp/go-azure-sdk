package usercalendarviewexceptionoccurrenceinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarViewExceptionOccurrenceId{}

// UserCalendarViewExceptionOccurrenceId is a struct representing the Resource ID for a User Calendar View Exception Occurrence
type UserCalendarViewExceptionOccurrenceId struct {
	UserId   string
	EventId  string
	EventId1 string
}

// NewUserCalendarViewExceptionOccurrenceID returns a new UserCalendarViewExceptionOccurrenceId struct
func NewUserCalendarViewExceptionOccurrenceID(userId string, eventId string, eventId1 string) UserCalendarViewExceptionOccurrenceId {
	return UserCalendarViewExceptionOccurrenceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseUserCalendarViewExceptionOccurrenceID parses 'input' into a UserCalendarViewExceptionOccurrenceId
func ParseUserCalendarViewExceptionOccurrenceID(input string) (*UserCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewExceptionOccurrenceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarViewExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserCalendarViewExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarViewExceptionOccurrenceIDInsensitively(input string) (*UserCalendarViewExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewExceptionOccurrenceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarViewExceptionOccurrenceID checks that 'input' can be parsed as a User Calendar View Exception Occurrence ID
func ValidateUserCalendarViewExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarViewExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar View Exception Occurrence ID
func (id UserCalendarViewExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendarView/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar View Exception Occurrence ID
func (id UserCalendarViewExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Calendar View Exception Occurrence ID
func (id UserCalendarViewExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Calendar View Exception Occurrence (%s)", strings.Join(components, "\n"))
}
