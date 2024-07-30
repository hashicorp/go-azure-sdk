package calendarviewinstanceattachment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarViewIdInstanceId{}

// UserIdCalendarViewIdInstanceId is a struct representing the Resource ID for a User Id Calendar View Id Instance
type UserIdCalendarViewIdInstanceId struct {
	UserId   string
	EventId  string
	EventId1 string
}

// NewUserIdCalendarViewIdInstanceID returns a new UserIdCalendarViewIdInstanceId struct
func NewUserIdCalendarViewIdInstanceID(userId string, eventId string, eventId1 string) UserIdCalendarViewIdInstanceId {
	return UserIdCalendarViewIdInstanceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseUserIdCalendarViewIdInstanceID parses 'input' into a UserIdCalendarViewIdInstanceId
func ParseUserIdCalendarViewIdInstanceID(input string) (*UserIdCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarViewIdInstanceIDInsensitively parses 'input' case-insensitively into a UserIdCalendarViewIdInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarViewIdInstanceIDInsensitively(input string) (*UserIdCalendarViewIdInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarViewIdInstanceId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdCalendarViewIdInstanceID checks that 'input' can be parsed as a User Id Calendar View Id Instance ID
func ValidateUserIdCalendarViewIdInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarViewIdInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar View Id Instance ID
func (id UserIdCalendarViewIdInstanceId) ID() string {
	fmtString := "/users/%s/calendarView/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar View Id Instance ID
func (id UserIdCalendarViewIdInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Id Calendar View Id Instance ID
func (id UserIdCalendarViewIdInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Id Calendar View Id Instance (%s)", strings.Join(components, "\n"))
}
