package usercalendareventinstanceexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventInstanceExceptionOccurrenceExtensionId{}

// UserCalendarEventInstanceExceptionOccurrenceExtensionId is a struct representing the Resource ID for a User Calendar Event Instance Exception Occurrence Extension
type UserCalendarEventInstanceExceptionOccurrenceExtensionId struct {
	UserId      string
	CalendarId  string
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewUserCalendarEventInstanceExceptionOccurrenceExtensionID returns a new UserCalendarEventInstanceExceptionOccurrenceExtensionId struct
func NewUserCalendarEventInstanceExceptionOccurrenceExtensionID(userId string, calendarId string, eventId string, eventId1 string, eventId2 string, extensionId string) UserCalendarEventInstanceExceptionOccurrenceExtensionId {
	return UserCalendarEventInstanceExceptionOccurrenceExtensionId{
		UserId:      userId,
		CalendarId:  calendarId,
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseUserCalendarEventInstanceExceptionOccurrenceExtensionID parses 'input' into a UserCalendarEventInstanceExceptionOccurrenceExtensionId
func ParseUserCalendarEventInstanceExceptionOccurrenceExtensionID(input string) (*UserCalendarEventInstanceExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventInstanceExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventInstanceExceptionOccurrenceExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
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

// ParseUserCalendarEventInstanceExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a UserCalendarEventInstanceExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventInstanceExceptionOccurrenceExtensionIDInsensitively(input string) (*UserCalendarEventInstanceExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventInstanceExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventInstanceExceptionOccurrenceExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
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

// ValidateUserCalendarEventInstanceExceptionOccurrenceExtensionID checks that 'input' can be parsed as a User Calendar Event Instance Exception Occurrence Extension ID
func ValidateUserCalendarEventInstanceExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventInstanceExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event Instance Exception Occurrence Extension ID
func (id UserCalendarEventInstanceExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s/instances/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event Instance Exception Occurrence Extension ID
func (id UserCalendarEventInstanceExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Event Instance Exception Occurrence Extension ID
func (id UserCalendarEventInstanceExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Calendar Event Instance Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
