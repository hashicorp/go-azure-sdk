package usercalendarcalendarviewinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarViewInstanceExtensionId{}

// UserCalendarCalendarViewInstanceExtensionId is a struct representing the Resource ID for a User Calendar Calendar View Instance Extension
type UserCalendarCalendarViewInstanceExtensionId struct {
	UserId      string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewUserCalendarCalendarViewInstanceExtensionID returns a new UserCalendarCalendarViewInstanceExtensionId struct
func NewUserCalendarCalendarViewInstanceExtensionID(userId string, eventId string, eventId1 string, extensionId string) UserCalendarCalendarViewInstanceExtensionId {
	return UserCalendarCalendarViewInstanceExtensionId{
		UserId:      userId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseUserCalendarCalendarViewInstanceExtensionID parses 'input' into a UserCalendarCalendarViewInstanceExtensionId
func ParseUserCalendarCalendarViewInstanceExtensionID(input string) (*UserCalendarCalendarViewInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewInstanceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewInstanceExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarCalendarViewInstanceExtensionIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarViewInstanceExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarViewInstanceExtensionIDInsensitively(input string) (*UserCalendarCalendarViewInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewInstanceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewInstanceExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarCalendarViewInstanceExtensionID checks that 'input' can be parsed as a User Calendar Calendar View Instance Extension ID
func ValidateUserCalendarCalendarViewInstanceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarViewInstanceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar View Instance Extension ID
func (id UserCalendarCalendarViewInstanceExtensionId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar View Instance Extension ID
func (id UserCalendarCalendarViewInstanceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Calendar View Instance Extension ID
func (id UserCalendarCalendarViewInstanceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Calendar Calendar View Instance Extension (%s)", strings.Join(components, "\n"))
}
