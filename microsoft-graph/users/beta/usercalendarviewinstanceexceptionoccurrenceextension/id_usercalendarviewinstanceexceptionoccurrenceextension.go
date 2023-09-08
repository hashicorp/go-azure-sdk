package usercalendarviewinstanceexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarViewInstanceExceptionOccurrenceExtensionId{}

// UserCalendarViewInstanceExceptionOccurrenceExtensionId is a struct representing the Resource ID for a User Calendar View Instance Exception Occurrence Extension
type UserCalendarViewInstanceExceptionOccurrenceExtensionId struct {
	UserId      string
	EventId     string
	EventId1    string
	EventId2    string
	ExtensionId string
}

// NewUserCalendarViewInstanceExceptionOccurrenceExtensionID returns a new UserCalendarViewInstanceExceptionOccurrenceExtensionId struct
func NewUserCalendarViewInstanceExceptionOccurrenceExtensionID(userId string, eventId string, eventId1 string, eventId2 string, extensionId string) UserCalendarViewInstanceExceptionOccurrenceExtensionId {
	return UserCalendarViewInstanceExceptionOccurrenceExtensionId{
		UserId:      userId,
		EventId:     eventId,
		EventId1:    eventId1,
		EventId2:    eventId2,
		ExtensionId: extensionId,
	}
}

// ParseUserCalendarViewInstanceExceptionOccurrenceExtensionID parses 'input' into a UserCalendarViewInstanceExceptionOccurrenceExtensionId
func ParseUserCalendarViewInstanceExceptionOccurrenceExtensionID(input string) (*UserCalendarViewInstanceExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewInstanceExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewInstanceExceptionOccurrenceExtensionId{}

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

// ParseUserCalendarViewInstanceExceptionOccurrenceExtensionIDInsensitively parses 'input' case-insensitively into a UserCalendarViewInstanceExceptionOccurrenceExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarViewInstanceExceptionOccurrenceExtensionIDInsensitively(input string) (*UserCalendarViewInstanceExceptionOccurrenceExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewInstanceExceptionOccurrenceExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewInstanceExceptionOccurrenceExtensionId{}

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

// ValidateUserCalendarViewInstanceExceptionOccurrenceExtensionID checks that 'input' can be parsed as a User Calendar View Instance Exception Occurrence Extension ID
func ValidateUserCalendarViewInstanceExceptionOccurrenceExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarViewInstanceExceptionOccurrenceExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar View Instance Exception Occurrence Extension ID
func (id UserCalendarViewInstanceExceptionOccurrenceExtensionId) ID() string {
	fmtString := "/users/%s/calendarView/%s/instances/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.EventId2, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar View Instance Exception Occurrence Extension ID
func (id UserCalendarViewInstanceExceptionOccurrenceExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("staticExceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId2", "eventId2Value"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar View Instance Exception Occurrence Extension ID
func (id UserCalendarViewInstanceExceptionOccurrenceExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Event Id 2: %q", id.EventId2),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Calendar View Instance Exception Occurrence Extension (%s)", strings.Join(components, "\n"))
}
