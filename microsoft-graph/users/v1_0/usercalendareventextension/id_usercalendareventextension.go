package usercalendareventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarEventExtensionId{}

// UserCalendarEventExtensionId is a struct representing the Resource ID for a User Calendar Event Extension
type UserCalendarEventExtensionId struct {
	UserId      string
	CalendarId  string
	EventId     string
	ExtensionId string
}

// NewUserCalendarEventExtensionID returns a new UserCalendarEventExtensionId struct
func NewUserCalendarEventExtensionID(userId string, calendarId string, eventId string, extensionId string) UserCalendarEventExtensionId {
	return UserCalendarEventExtensionId{
		UserId:      userId,
		CalendarId:  calendarId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseUserCalendarEventExtensionID parses 'input' into a UserCalendarEventExtensionId
func ParseUserCalendarEventExtensionID(input string) (*UserCalendarEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarEventExtensionIDInsensitively parses 'input' case-insensitively into a UserCalendarEventExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarEventExtensionIDInsensitively(input string) (*UserCalendarEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarEventExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarEventExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.CalendarId, ok = parsed.Parsed["calendarId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "calendarId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarEventExtensionID checks that 'input' can be parsed as a User Calendar Event Extension ID
func ValidateUserCalendarEventExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarEventExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar Event Extension ID
func (id UserCalendarEventExtensionId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar Event Extension ID
func (id UserCalendarEventExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar Event Extension ID
func (id UserCalendarEventExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Calendar Event Extension (%s)", strings.Join(components, "\n"))
}
