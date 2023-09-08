package userevent

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserEventId{}

// UserEventId is a struct representing the Resource ID for a User Event
type UserEventId struct {
	UserId  string
	EventId string
}

// NewUserEventID returns a new UserEventId struct
func NewUserEventID(userId string, eventId string) UserEventId {
	return UserEventId{
		UserId:  userId,
		EventId: eventId,
	}
}

// ParseUserEventID parses 'input' into a UserEventId
func ParseUserEventID(input string) (*UserEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ParseUserEventIDInsensitively parses 'input' case-insensitively into a UserEventId
// note: this method should only be used for API response data and not user input
func ParseUserEventIDInsensitively(input string) (*UserEventId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	return &id, nil
}

// ValidateUserEventID checks that 'input' can be parsed as a User Event ID
func ValidateUserEventID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserEventID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Event ID
func (id UserEventId) ID() string {
	fmtString := "/users/%s/events/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Event ID
func (id UserEventId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
	}
}

// String returns a human-readable description of this User Event ID
func (id UserEventId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
	}
	return fmt.Sprintf("User Event (%s)", strings.Join(components, "\n"))
}
