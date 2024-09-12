package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{}

// UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId is a struct representing the Resource ID for a User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Extension
type UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId struct {
	UserId      string
	CalendarId  string
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID returns a new UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId struct
func NewUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(userId string, calendarId string, eventId string, eventId1 string, eventId2 string, extensionId string) UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId {
	return UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{
		UserId:      userId,
		CalendarId:  calendarId,
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID parses 'input' into a UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId
func ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(input string) (*UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionIDInsensitively(input string) (*UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.EventId2, ok = input.Parsed["eventId2"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId2", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID checks that 'input' can be parsed as a User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Extension ID
func ValidateUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Extension ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId) ID() string {
	fmtString := "/users/%s/calendars/%s/calendarView/%s/instances/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Extension ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Extension ID
func (id UserIdCalendarIdCalendarViewIdInstanceIdExceptionOccurrenceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Calendar Id Calendar View Id Instance Id Exception Occurrence Id Extension (%s)", strings.Join(components, "\n"))
}
