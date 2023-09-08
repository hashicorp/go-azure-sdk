package usercalendareventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventExceptionOccurrenceExtensionId{}

// UserCalendarEventExceptionOccurrenceExtensionId is a struct representing the Resource ID for a User Calendar Event Exception Occurrence Extension
type UserCalendarEventExceptionOccurrenceExtensionId struct {
	UserId      string
	CalendarId  string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewUserCalendarEventExceptionOccurrenceExtensionID returns a new UserCalendarEventExceptionOccurrenceExtensionId struct
func NewUserCalendarEventExceptionOccurrenceExtensionID(userId string, calendarId string, eventId string, eventId1 string, extensionId string) UserCalendarEventExceptionOccurrenceExtensionId {
	return UserCalendarEventExceptionOccurrenceExtensionId{
		UserId:      userId,
		CalendarId:  calendarId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseUserCalendarEventExceptionOccurrenceExtensionID parses 'input' into a UserCalendarEventExceptionOccurrenceExtensionId
func ParseUserCalendarEventExceptionOccurrenceExtensionID(input string) (*UserCalendarEventExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventExceptionOccurrenceExtensionId{}

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

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarEventExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a UserCalendarEventExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventExceptionOccurrenceExtensionIDInsensitively(input string) (*UserCalendarEventExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventExceptionOccurrenceExtensionId{}

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

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarEventExceptionOccurrenceExtensionID checks that 'input' can be parsed as a User Calendar Event Exception Occurrence Extension ID
func ValidateUserCalendarEventExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event Exception Occurrence Extension ID
func (id UserCalendarEventExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event Exception Occurrence Extension ID
func (id UserCalendarEventExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Event Exception Occurrence Extension ID
func (id UserCalendarEventExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Calendar Event Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
