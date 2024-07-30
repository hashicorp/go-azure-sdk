package calendareventexceptionoccurrenceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarEventIdExceptionOccurrenceIdExtensionId{}

// UserIdCalendarEventIdExceptionOccurrenceIdExtensionId is a struct representing the Resource ID for a User Id Calendar Event Id Exception Occurrence Id Extension
type UserIdCalendarEventIdExceptionOccurrenceIdExtensionId struct {
	UserId      string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewUserIdCalendarEventIdExceptionOccurrenceIdExtensionID returns a new UserIdCalendarEventIdExceptionOccurrenceIdExtensionId struct
func NewUserIdCalendarEventIdExceptionOccurrenceIdExtensionID(userId string, eventId string, eventId1 string, extensionId string) UserIdCalendarEventIdExceptionOccurrenceIdExtensionId {
	return UserIdCalendarEventIdExceptionOccurrenceIdExtensionId{
		UserId:      userId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseUserIdCalendarEventIdExceptionOccurrenceIdExtensionID parses 'input' into a UserIdCalendarEventIdExceptionOccurrenceIdExtensionId
func ParseUserIdCalendarEventIdExceptionOccurrenceIdExtensionID(input string) (*UserIdCalendarEventIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarEventIdExceptionOccurrenceIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarEventIdExceptionOccurrenceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarEventIdExceptionOccurrenceIdExtensionIDInsensitively(input string) (*UserIdCalendarEventIdExceptionOccurrenceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarEventIdExceptionOccurrenceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarEventIdExceptionOccurrenceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarEventIdExceptionOccurrenceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.EventId1, ok = input.Parsed["eventId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId1", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdCalendarEventIdExceptionOccurrenceIdExtensionID checks that 'input' can be parsed as a User Id Calendar Event Id Exception Occurrence Id Extension ID
func ValidateUserIdCalendarEventIdExceptionOccurrenceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarEventIdExceptionOccurrenceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Event Id Exception Occurrence Id Extension ID
func (id UserIdCalendarEventIdExceptionOccurrenceIdExtensionId) ID() string {
	fmtString := "/users/%s/calendar/events/%s/exceptionOccurrences/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Event Id Exception Occurrence Id Extension ID
func (id UserIdCalendarEventIdExceptionOccurrenceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("exceptionOccurrences", "exceptionOccurrences", "exceptionOccurrences"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Event Id Exception Occurrence Id Extension ID
func (id UserIdCalendarEventIdExceptionOccurrenceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Calendar Event Id Exception Occurrence Id Extension (%s)", strings.Join(components, "\n"))
}
