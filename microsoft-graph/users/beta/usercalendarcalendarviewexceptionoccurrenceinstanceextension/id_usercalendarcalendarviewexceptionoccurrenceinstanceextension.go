package usercalendarcalendarviewexceptionoccurrenceinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId{}

// UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId is a struct representing the Resource ID for a User Calendar Calendar View Exception Occurrence Instance Extension
type UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId struct {
	UserId      string
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionID returns a new UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId struct
func NewUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionID(userId string, eventId string, eventId1 string, eventId2 string, extensionId string) UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId {
	return UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId{
		UserId:      userId,
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionID parses 'input' into a UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId
func ParseUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionID(input string) (*UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionIDInsensitively(input string) (*UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.EventId1, ok = parsed.Parsed["eventId1"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId1", *parsed)
	}

	if id.EventId2, ok = parsed.Parsed["eventId2"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId2", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionID checks that 'input' can be parsed as a User Calendar Calendar View Exception Occurrence Instance Extension ID
func ValidateUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarViewExceptionOccurrenceInstanceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar View Exception Occurrence Instance Extension ID
func (id UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s/exceptionOccurrences/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar View Exception Occurrence Instance Extension ID
func (id UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Calendar View Exception Occurrence Instance Extension ID
func (id UserCalendarCalendarViewExceptionOccurrenceInstanceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Calendar Calendar View Exception Occurrence Instance Extension (%s)", strings.Join(components, "\n"))
}
