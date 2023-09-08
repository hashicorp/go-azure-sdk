package usercalendarcalendarviewexceptionoccurrenceinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarViewExceptionOccurrenceInstanceId{}

// UserCalendarCalendarViewExceptionOccurrenceInstanceId is a struct representing the Resource ID for a User Calendar Calendar View Exception Occurrence Instance
type UserCalendarCalendarViewExceptionOccurrenceInstanceId struct {
	UserId     string
	CalendarId string
	EventId    string
	EventId1   string
	EventId2   string
}

// NewUserCalendarCalendarViewExceptionOccurrenceInstanceID returns a new UserCalendarCalendarViewExceptionOccurrenceInstanceId struct
func NewUserCalendarCalendarViewExceptionOccurrenceInstanceID(userId string, calendarId string, eventId string, eventId1 string, eventId2 string) UserCalendarCalendarViewExceptionOccurrenceInstanceId {
	return UserCalendarCalendarViewExceptionOccurrenceInstanceId{
		UserId:     userId,
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
		EventId2:   eventId2,
	}
}

// ParseUserCalendarCalendarViewExceptionOccurrenceInstanceID parses 'input' into a UserCalendarCalendarViewExceptionOccurrenceInstanceId
func ParseUserCalendarCalendarViewExceptionOccurrenceInstanceID(input string) (*UserCalendarCalendarViewExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewExceptionOccurrenceInstanceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
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

// ParseUserCalendarCalendarViewExceptionOccurrenceInstanceIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarViewExceptionOccurrenceInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarViewExceptionOccurrenceInstanceIDInsensitively(input string) (*UserCalendarCalendarViewExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewExceptionOccurrenceInstanceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
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

// ValidateUserCalendarCalendarViewExceptionOccurrenceInstanceID checks that 'input' can be parsed as a User Calendar Calendar View Exception Occurrence Instance ID
func ValidateUserCalendarCalendarViewExceptionOccurrenceInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarViewExceptionOccurrenceInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar View Exception Occurrence Instance ID
func (id UserCalendarCalendarViewExceptionOccurrenceInstanceId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarView/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar View Exception Occurrence Instance ID
func (id UserCalendarCalendarViewExceptionOccurrenceInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Calendar Calendar View Exception Occurrence Instance ID
func (id UserCalendarCalendarViewExceptionOccurrenceInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Calendar Calendar View Exception Occurrence Instance (%s)", strings.Join(components, "\n"))
}
