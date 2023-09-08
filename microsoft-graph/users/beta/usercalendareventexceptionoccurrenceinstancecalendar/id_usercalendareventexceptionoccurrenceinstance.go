package usercalendareventexceptionoccurrenceinstancecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventExceptionOccurrenceInstanceId{}

// UserCalendarEventExceptionOccurrenceInstanceId is a struct representing the Resource ID for a User Calendar Event Exception Occurrence Instance
type UserCalendarEventExceptionOccurrenceInstanceId struct {
	UserId   string
	EventId  string
	EventId1 string
	EventId2 string
}

// NewUserCalendarEventExceptionOccurrenceInstanceID returns a new UserCalendarEventExceptionOccurrenceInstanceId struct
func NewUserCalendarEventExceptionOccurrenceInstanceID(userId string, eventId string, eventId1 string, eventId2 string) UserCalendarEventExceptionOccurrenceInstanceId {
	return UserCalendarEventExceptionOccurrenceInstanceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
		EventId2: eventId2,
	}
}

// ParseUserCalendarEventExceptionOccurrenceInstanceID parses 'input' into a UserCalendarEventExceptionOccurrenceInstanceId
func ParseUserCalendarEventExceptionOccurrenceInstanceID(input string) (*UserCalendarEventExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventExceptionOccurrenceInstanceId{}

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

// ParseUserCalendarEventExceptionOccurrenceInstanceIDInsensitively parses 'input' case-insensitively into a UserCalendarEventExceptionOccurrenceInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventExceptionOccurrenceInstanceIDInsensitively(input string) (*UserCalendarEventExceptionOccurrenceInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventExceptionOccurrenceInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventExceptionOccurrenceInstanceId{}

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

// ValidateUserCalendarEventExceptionOccurrenceInstanceID checks that 'input' can be parsed as a User Calendar Event Exception Occurrence Instance ID
func ValidateUserCalendarEventExceptionOccurrenceInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventExceptionOccurrenceInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event Exception Occurrence Instance ID
func (id UserCalendarEventExceptionOccurrenceInstanceId) ID() string {
	fmtString := "/users/%s/calendar/events/%s/exceptionOccurrences/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event Exception Occurrence Instance ID
func (id UserCalendarEventExceptionOccurrenceInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Calendar Event Exception Occurrence Instance ID
func (id UserCalendarEventExceptionOccurrenceInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Calendar Event Exception Occurrence Instance (%s)", strings.Join(components, "\n"))
}
