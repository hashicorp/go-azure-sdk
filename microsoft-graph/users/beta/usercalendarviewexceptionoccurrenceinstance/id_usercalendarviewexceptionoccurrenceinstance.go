package usercalendarviewexceptionoccurrenceinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarViewExceptionOccurrenceInstanceId{}

// UserCalendarViewExceptionOccurrenceInstanceId is a struct representing the Resource ID for a User Calendar View Exception Occurrence Instance
type UserCalendarViewExceptionOccurrenceInstanceId struct {
	UserId   string
	EventId  string
	EventId1 string
	EventId2 string
}

// NewUserCalendarViewExceptionOccurrenceInstanceID returns a new UserCalendarViewExceptionOccurrenceInstanceId struct
func NewUserCalendarViewExceptionOccurrenceInstanceID(userId string, eventId string, eventId1 string, eventId2 string) UserCalendarViewExceptionOccurrenceInstanceId {
	return UserCalendarViewExceptionOccurrenceInstanceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseUserCalendarViewExceptionOccurrenceInstanceID parses 'input' into a UserCalendarViewExceptionOccurrenceInstanceId
func ParseUserCalendarViewExceptionOccurrenceInstanceID(input string) (*UserCalendarViewExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewExceptionOccurrenceInstanceId{}

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

// ParseUserCalendarViewExceptionOccurrenceInstanceIDInsensitively parses 'input' case-insensitively into a UserCalendarViewExceptionOccurrenceInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarViewExceptionOccurrenceInstanceIDInsensitively(input string) (*UserCalendarViewExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewExceptionOccurrenceInstanceId{}

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

// ValidateUserCalendarViewExceptionOccurrenceInstanceID checks that 'input' can be parsed as a User Calendar View Exception Occurrence Instance ID
func ValidateUserCalendarViewExceptionOccurrenceInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarViewExceptionOccurrenceInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar View Exception Occurrence Instance ID
func (id UserCalendarViewExceptionOccurrenceInstanceId) ID() string {
	fmtString := "/users/%s/calendarView/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar View Exception Occurrence Instance ID
func (id UserCalendarViewExceptionOccurrenceInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Calendar View Exception Occurrence Instance ID
func (id UserCalendarViewExceptionOccurrenceInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Calendar View Exception Occurrence Instance (%s)", strings.Join(components, "\n"))
}
