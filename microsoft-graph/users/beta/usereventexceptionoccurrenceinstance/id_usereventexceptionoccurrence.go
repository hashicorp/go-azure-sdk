package usereventexceptionoccurrenceinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserEventExceptionOccurrenceId{}

// UserEventExceptionOccurrenceId is a struct representing the Resource ID for a User Event Exception Occurrence
type UserEventExceptionOccurrenceId struct {
	UserId   string
	EventId  string
	EventId1 string
}

// NewUserEventExceptionOccurrenceID returns a new UserEventExceptionOccurrenceId struct
func NewUserEventExceptionOccurrenceID(userId string, eventId string, eventId1 string) UserEventExceptionOccurrenceId {
	return UserEventExceptionOccurrenceId{
		UserId:   userId,
		EventId:  eventId,
		EventId1: eventId1,
	}
}

// ParseUserEventExceptionOccurrenceID parses 'input' into a UserEventExceptionOccurrenceId
func ParseUserEventExceptionOccurrenceID(input string) (*UserEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventExceptionOccurrenceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ParseUserEventExceptionOccurrenceIDInsensitively parses 'input' case-insensitively into a UserEventExceptionOccurrenceId
// note: this method should only be used for API response data and not user input
func ParseUserEventExceptionOccurrenceIDInsensitively(input string) (*UserEventExceptionOccurrenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventExceptionOccurrenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventExceptionOccurrenceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	return &id, nil
}

// ValidateUserEventExceptionOccurrenceID checks that 'input' can be parsed as a User Event Exception Occurrence ID
func ValidateUserEventExceptionOccurrenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserEventExceptionOccurrenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Event Exception Occurrence ID
func (id UserEventExceptionOccurrenceId) ID() string {
	fmtString := "/users/%s/events/%s/exceptionOccurrences/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1)
}

// Segments returns a slice of Resource ID Segments which comprise this User Event Exception Occurrence ID
func (id UserEventExceptionOccurrenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
	}
}

// String returns a human-readable description of this User Event Exception Occurrence ID
func (id UserEventExceptionOccurrenceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
	}
	return fmt.Sprintf("User Event Exception Occurrence (%s)", strings.Join(components, "\n"))
}
