package usercalendargroupcalendareventexceptionoccurrenceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarEventExceptionOccurrenceId{}

// UserCalendarGroupCalendarEventExceptionOccurrenceId is a struct representing the Resource ID for a User Calendar Group Calendar Event Exception Occurrence
type UserCalendarGroupCalendarEventExceptionOccurrenceId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
	EventId1        string
}

// NewUserCalendarGroupCalendarEventExceptionOccurrenceID returns a new UserCalendarGroupCalendarEventExceptionOccurrenceId struct
func NewUserCalendarGroupCalendarEventExceptionOccurrenceID(userId string, calendarGroupId string, calendarId string, eventId string, eventId1 string) UserCalendarGroupCalendarEventExceptionOccurrenceId {
	return UserCalendarGroupCalendarEventExceptionOccurrenceId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		EventId1:        eventId1,
	}
}

// ParseUserCalendarGroupCalendarEventExceptionOccurrenceID parses 'input' into a UserCalendarGroupCalendarEventExceptionOccurrenceId
func ParseUserCalendarGroupCalendarEventExceptionOccurrenceID(input string) (*UserCalendarGroupCalendarEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarEventExceptionOccurrenceId{}

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

	return &id, nil
}

// ParseUserCalendarGroupCalendarEventExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupCalendarEventExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupCalendarEventExceptionOccurrenceIDInsensitively(input string) (*UserCalendarGroupCalendarEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarEventExceptionOccurrenceId{}

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

	return &id, nil
}

// ValidateUserCalendarGroupCalendarEventExceptionOccurrenceID checks that 'input' can be parsed as a User Calendar Group Calendar Event Exception Occurrence ID
func ValidateUserCalendarGroupCalendarEventExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupCalendarEventExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group Calendar Event Exception Occurrence ID
func (id UserCalendarGroupCalendarEventExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group Calendar Event Exception Occurrence ID
func (id UserCalendarGroupCalendarEventExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Calendar Group Calendar Event Exception Occurrence ID
func (id UserCalendarGroupCalendarEventExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Calendar Group Calendar Event Exception Occurrence (%s)", strings.Join(components, "\n"))
}
