package usercalendareventinstancecalendar

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventInstanceId{}

// UserCalendarEventInstanceId is a struct representing the Resource ID for a User Calendar Event Instance
type UserCalendarEventInstanceId struct {
	UserId     string
	CalendarId string
	EventId    string
	EventId1   string
}

// NewUserCalendarEventInstanceID returns a new UserCalendarEventInstanceId struct
func NewUserCalendarEventInstanceID(userId string, calendarId string, eventId string, eventId1 string) UserCalendarEventInstanceId {
	return UserCalendarEventInstanceId{
		UserId:     userId,
		CalendarId: calendarId,
		EventId:    eventId,
		EventId1:   eventId1,
	}
}

// ParseUserCalendarEventInstanceID parses 'input' into a UserCalendarEventInstanceId
func ParseUserCalendarEventInstanceID(input string) (*UserCalendarEventInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventInstanceId{}

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

	return &id, nil
}

// ParseUserCalendarEventInstanceIDInsensitively parses 'input' case-insensitively into a UserCalendarEventInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventInstanceIDInsensitively(input string) (*UserCalendarEventInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventInstanceId{}

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

	return &id, nil
}

// ValidateUserCalendarEventInstanceID checks that 'input' can be parsed as a User Calendar Event Instance ID
func ValidateUserCalendarEventInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event Instance ID
func (id UserCalendarEventInstanceId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event Instance ID
func (id UserCalendarEventInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Calendar Event Instance ID
func (id UserCalendarEventInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Calendar Event Instance (%s)", strings.Join(components, "\n"))
}
