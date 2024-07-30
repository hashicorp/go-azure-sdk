package calendareventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdEventId{}

// UserIdCalendarIdEventId is a struct representing the Resource ID for a User Id Calendar Id Event
type UserIdCalendarIdEventId struct {
	UserId     string
	CalendarId string
	EventId    string
}

// NewUserIdCalendarIdEventID returns a new UserIdCalendarIdEventId struct
func NewUserIdCalendarIdEventID(userId string, calendarId string, eventId string) UserIdCalendarIdEventId {
	return UserIdCalendarIdEventId{
		UserId:     userId,
		CalendarId: calendarId,
		EventId:    eventId,
	}
}

// ParseUserIdCalendarIdEventID parses 'input' into a UserIdCalendarIdEventId
func ParseUserIdCalendarIdEventID(input string) (*UserIdCalendarIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdEventIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdEventId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdEventIDInsensitively(input string) (*UserIdCalendarIdEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateUserIdCalendarIdEventID checks that 'input' can be parsed as a User Id Calendar Id Event ID
func ValidateUserIdCalendarIdEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Event ID
func (id UserIdCalendarIdEventId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Event ID
func (id UserIdCalendarIdEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Event ID
func (id UserIdCalendarIdEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Id Calendar Id Event (%s)", strings.Join(components, "\n"))
}
