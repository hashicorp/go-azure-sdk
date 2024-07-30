package calendareventinstanceextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdEventIdInstanceIdExtensionId{}

// UserIdCalendarIdEventIdInstanceIdExtensionId is a struct representing the Resource ID for a User Id Calendar Id Event Id Instance Id Extension
type UserIdCalendarIdEventIdInstanceIdExtensionId struct {
	UserId      string
	CalendarId  string
	EventId     string
	EventId1    string
	ExtensionId string
}

// NewUserIdCalendarIdEventIdInstanceIdExtensionID returns a new UserIdCalendarIdEventIdInstanceIdExtensionId struct
func NewUserIdCalendarIdEventIdInstanceIdExtensionID(userId string, calendarId string, eventId string, eventId1 string, extensionId string) UserIdCalendarIdEventIdInstanceIdExtensionId {
	return UserIdCalendarIdEventIdInstanceIdExtensionId{
		UserId:      userId,
		CalendarId:  calendarId,
		EventId:     eventId,
		EventId1:    eventId1,
		ExtensionId: extensionId,
	}
}

// ParseUserIdCalendarIdEventIdInstanceIdExtensionID parses 'input' into a UserIdCalendarIdEventIdInstanceIdExtensionId
func ParseUserIdCalendarIdEventIdInstanceIdExtensionID(input string) (*UserIdCalendarIdEventIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdEventIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdEventIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdEventIdInstanceIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdEventIdInstanceIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdEventIdInstanceIdExtensionIDInsensitively(input string) (*UserIdCalendarIdEventIdInstanceIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdEventIdInstanceIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdEventIdInstanceIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdEventIdInstanceIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdCalendarIdEventIdInstanceIdExtensionID checks that 'input' can be parsed as a User Id Calendar Id Event Id Instance Id Extension ID
func ValidateUserIdCalendarIdEventIdInstanceIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdEventIdInstanceIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Event Id Instance Id Extension ID
func (id UserIdCalendarIdEventIdInstanceIdExtensionId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s/instances/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.EventId1, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Event Id Instance Id Extension ID
func (id UserIdCalendarIdEventIdInstanceIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("eventId1", "eventId1Value"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Event Id Instance Id Extension ID
func (id UserIdCalendarIdEventIdInstanceIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Event Id 1: %q", id.EventId1),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Calendar Id Event Id Instance Id Extension (%s)", strings.Join(components, "\n"))
}
