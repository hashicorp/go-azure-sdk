package usercalendargroupcalendareventinstanceexceptionoccurrencecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId{}

// UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId is a struct representing the Resource ID for a User Calendar Group Calendar Event Instance Exception Occurrence
type UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
	EventId2        string
}

// NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID returns a new UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId struct
func NewUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID(userId string, calendarGroupId string, calendarId string, eventId string, eventId1 string, eventId2 string) UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId {
	return UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
		EventId2:        eventId2,
	}
}

// ParseUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID parses 'input' into a UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId
func ParseUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID(input string) (*UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
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

// ParseUserCalendarGroupCalendarEventInstanceExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupCalendarEventInstanceExceptionOccurrenceIDInsensitively(input string) (*UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarGroupId, ok = parsed.Parsed["calendarGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", *parsed)
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

// ValidateUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID checks that 'input' can be parsed as a User Calendar Group Calendar Event Instance Exception Occurrence ID
func ValidateUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupCalendarEventInstanceExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group Calendar Event Instance Exception Occurrence ID
func (id UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/events/%s/instances/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1, id.EventId2)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group Calendar Event Instance Exception Occurrence ID
func (id UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
	}
}

// String returns a human-readable description of this User Calendar Group Calendar Event Instance Exception Occurrence ID
func (id UserCalendarGroupCalendarEventInstanceExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
	}
	return fmt.Sprintf("User Calendar Group Calendar Event Instance Exception Occurrence (%s)", strings.Join(components, "\n"))
}
