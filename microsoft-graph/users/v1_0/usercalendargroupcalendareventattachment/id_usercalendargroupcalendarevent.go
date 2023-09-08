package usercalendargroupcalendareventattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarGroupCalendarEventId{}

// UserCalendarGroupCalendarEventId is a struct representing the Resource ID for a User Calendar Group Calendar Event
type UserCalendarGroupCalendarEventId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
}

// NewUserCalendarGroupCalendarEventID returns a new UserCalendarGroupCalendarEventId struct
func NewUserCalendarGroupCalendarEventID(userId string, calendarGroupId string, calendarId string, eventId string) UserCalendarGroupCalendarEventId {
	return UserCalendarGroupCalendarEventId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
	}
}

// ParseUserCalendarGroupCalendarEventID parses 'input' into a UserCalendarGroupCalendarEventId
func ParseUserCalendarGroupCalendarEventID(input string) (*UserCalendarGroupCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarEventId{}

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

	return &id, nil
}

// ParseUserCalendarGroupCalendarEventIDInsensitively parses 'input' case-insensitively into a UserCalendarGroupCalendarEventId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarGroupCalendarEventIDInsensitively(input string) (*UserCalendarGroupCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarGroupCalendarEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarGroupCalendarEventId{}

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

	return &id, nil
}

// ValidateUserCalendarGroupCalendarEventID checks that 'input' can be parsed as a User Calendar Group Calendar Event ID
func ValidateUserCalendarGroupCalendarEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarGroupCalendarEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Group Calendar Event ID
func (id UserCalendarGroupCalendarEventId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/events/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Group Calendar Event ID
func (id UserCalendarGroupCalendarEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Group Calendar Event ID
func (id UserCalendarGroupCalendarEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Calendar Group Calendar Event (%s)", strings.Join(components, "\n"))
}
