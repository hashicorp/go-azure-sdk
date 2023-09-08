package usercalendarcalendarviewexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarCalendarViewExceptionOccurrenceExtensionId{}

// UserCalendarCalendarViewExceptionOccurrenceExtensionId is a struct representing the Resource ID for a User Calendar Calendar View Exception Occurrence Extension
type UserCalendarCalendarViewExceptionOccurrenceExtensionId struct {
	UserId      string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewUserCalendarCalendarViewExceptionOccurrenceExtensionID returns a new UserCalendarCalendarViewExceptionOccurrenceExtensionId struct
func NewUserCalendarCalendarViewExceptionOccurrenceExtensionID(userId string, eventId string, eventId1 string, extensionId string) UserCalendarCalendarViewExceptionOccurrenceExtensionId {
	return UserCalendarCalendarViewExceptionOccurrenceExtensionId{
		UserId:      userId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseUserCalendarCalendarViewExceptionOccurrenceExtensionID parses 'input' into a UserCalendarCalendarViewExceptionOccurrenceExtensionId
func ParseUserCalendarCalendarViewExceptionOccurrenceExtensionID(input string) (*UserCalendarCalendarViewExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewExceptionOccurrenceExtensionId{}

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

// ParseUserCalendarCalendarViewExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a UserCalendarCalendarViewExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarCalendarViewExceptionOccurrenceExtensionIDInsensitively(input string) (*UserCalendarCalendarViewExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarCalendarViewExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarCalendarViewExceptionOccurrenceExtensionId{}

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

// ValidateUserCalendarCalendarViewExceptionOccurrenceExtensionID checks that 'input' can be parsed as a User Calendar Calendar View Exception Occurrence Extension ID
func ValidateUserCalendarCalendarViewExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarCalendarViewExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Calendar View Exception Occurrence Extension ID
func (id UserCalendarCalendarViewExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Calendar View Exception Occurrence Extension ID
func (id UserCalendarCalendarViewExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendar", "calendar", "calendar"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Calendar View Exception Occurrence Extension ID
func (id UserCalendarCalendarViewExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Calendar Calendar View Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
