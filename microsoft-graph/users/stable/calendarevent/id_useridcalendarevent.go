package calendarevent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarEventId{}

// UserIdCalendarEventId is a struct representing the Resource ID for a User Id Calendar Event
type UserIdCalendarEventId struct {
	UserId  string
	EventId string
}

// NewUserIdCalendarEventID returns a new UserIdCalendarEventId struct
func NewUserIdCalendarEventID(userId string, eventId string) UserIdCalendarEventId {
	return UserIdCalendarEventId{
		UserId:  userId,
		EventId: eventId,
	}
}

// ParseUserIdCalendarEventID parses 'input' into a UserIdCalendarEventId
func ParseUserIdCalendarEventID(input string) (*UserIdCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarEventIDInsensitively parses 'input' case-insensitively into a UserIdCalendarEventId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarEventIDInsensitively(input string) (*UserIdCalendarEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarEventId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	return nil
}

// ValidateUserIdCalendarEventID checks that 'input' can be parsed as a User Id Calendar Event ID
func ValidateUserIdCalendarEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Event ID
func (id UserIdCalendarEventId) ID() string {
	fmtString := "/users/%s/calendar/events/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Event ID
func (id UserIdCalendarEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Event ID
func (id UserIdCalendarEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Id Calendar Event (%s)", strings.Join(components, "\n"))
}
