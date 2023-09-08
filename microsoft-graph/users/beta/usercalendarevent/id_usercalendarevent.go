package usercalendarevent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventId{}

// UserCalendarEventId is a struct representing the Resource ID for a User Calendar Event
type UserCalendarEventId struct {
	UserId  string
	EventId string
}

// NewUserCalendarEventID returns a new UserCalendarEventId struct
func NewUserCalendarEventID(userId string, eventId string) UserCalendarEventId {
	return UserCalendarEventId{
		UserId:  userId,
		EventId: eventId,
	}
}

// ParseUserCalendarEventID parses 'input' into a UserCalendarEventId
func ParseUserCalendarEventID(input string) (*UserCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarEventIDInsensitively parses 'input' case-insensitively into a UserCalendarEventId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventIDInsensitively(input string) (*UserCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarEventID checks that 'input' can be parsed as a User Calendar Event ID
func ValidateUserCalendarEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event ID
func (id UserCalendarEventId) ID() string {
	fmtString := "/users/%s/calendar/events/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event ID
func (id UserCalendarEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Event ID
func (id UserCalendarEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Calendar Event (%s)", strings.Join(components, "\n"))
}
