package usercalendareventinstanceexceptionoccurrence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventInstanceExceptionOccurrenceId{}

// UserCalendarEventInstanceExceptionOccurrenceId is a struct representing the Resource ID for a User Calendar Event Instance Exception Occurrence
type UserCalendarEventInstanceExceptionOccurrenceId struct {
	UserId   string
	EventId  string
	EventId1 string
	EventId2 string
}

// NewUserCalendarEventInstanceExceptionOccurrenceID returns a new UserCalendarEventInstanceExceptionOccurrenceId struct
func NewUserCalendarEventInstanceExceptionOccurrenceID(userId string, eventId string, eventId1 string, eventId2 string) UserCalendarEventInstanceExceptionOccurrenceId {
	return UserCalendarEventInstanceExceptionOccurrenceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseUserCalendarEventInstanceExceptionOccurrenceID parses 'input' into a UserCalendarEventInstanceExceptionOccurrenceId
func ParseUserCalendarEventInstanceExceptionOccurrenceID(input string) (*UserCalendarEventInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventInstanceExceptionOccurrenceId{}

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

// ParseUserCalendarEventInstanceExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserCalendarEventInstanceExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventInstanceExceptionOccurrenceIDInsensitively(input string) (*UserCalendarEventInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventInstanceExceptionOccurrenceId{}

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

// ValidateUserCalendarEventInstanceExceptionOccurrenceID checks that 'input' can be parsed as a User Calendar Event Instance Exception Occurrence ID
func ValidateUserCalendarEventInstanceExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventInstanceExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event Instance Exception Occurrence ID
func (id UserCalendarEventInstanceExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendar/events/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event Instance Exception Occurrence ID
func (id UserCalendarEventInstanceExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Calendar Event Instance Exception Occurrence ID
func (id UserCalendarEventInstanceExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Calendar Event Instance Exception Occurrence (%s)", strings.Join(components, "\n"))
}
