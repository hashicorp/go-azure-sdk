package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarEventIdInstanceId{}

// UserIdCalendarEventIdInstanceId is a struct representing the Resource ID for a User Id Calendar Event Id Instance
type UserIdCalendarEventIdInstanceId struct {
	UserId   string
	EventId  string
	EventId1 string
}

// NewUserIdCalendarEventIdInstanceID returns a new UserIdCalendarEventIdInstanceId struct
func NewUserIdCalendarEventIdInstanceID(userId string, eventId string, eventId1 string) UserIdCalendarEventIdInstanceId {
	return UserIdCalendarEventIdInstanceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseUserIdCalendarEventIdInstanceID parses 'input' into a UserIdCalendarEventIdInstanceId
func ParseUserIdCalendarEventIdInstanceID(input string) (*UserIdCalendarEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarEventIdInstanceIDInsensitively parses 'input' case-insensitively into a UserIdCalendarEventIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarEventIdInstanceIDInsensitively(input string) (*UserIdCalendarEventIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarEventIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	return nil
}

// ValidateUserIdCalendarEventIdInstanceID checks that 'input' can be parsed as a User Id Calendar Event Id Instance ID
func ValidateUserIdCalendarEventIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarEventIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Event Id Instance ID
func (id UserIdCalendarEventIdInstanceId) ID() string {
	fmtString := "/users/%s/calendar/events/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Event Id Instance ID
func (id UserIdCalendarEventIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Id Calendar Event Id Instance ID
func (id UserIdCalendarEventIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Id Calendar Event Id Instance (%s)", strings.Join(components, "\n"))
}
