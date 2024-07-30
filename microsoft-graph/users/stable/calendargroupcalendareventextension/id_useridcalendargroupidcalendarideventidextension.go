package calendargroupcalendareventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarGroupIdCalendarIdEventIdExtensionId{}

// UserIdCalendarGroupIdCalendarIdEventIdExtensionId is a struct representing the Resource ID for a User Id Calendar Group Id Calendar Id Event Id Extension
type UserIdCalendarGroupIdCalendarIdEventIdExtensionId struct {
	UserId          string
	CalendarGroupId string
	CalendarId      string
	EventId         string
	ExtensionId     string
}

// NewUserIdCalendarGroupIdCalendarIdEventIdExtensionID returns a new UserIdCalendarGroupIdCalendarIdEventIdExtensionId struct
func NewUserIdCalendarGroupIdCalendarIdEventIdExtensionID(userId string, calendarGroupId string, calendarId string, eventId string, extensionId string) UserIdCalendarGroupIdCalendarIdEventIdExtensionId {
	return UserIdCalendarGroupIdCalendarIdEventIdExtensionId{
		UserId:          userId,
		CalendarGroupId: calendarGroupId,
		CalendarId:      calendarId,
		EventId:         eventId,
		ExtensionId:     extensionId,
	}
}

// ParseUserIdCalendarGroupIdCalendarIdEventIdExtensionID parses 'input' into a UserIdCalendarGroupIdCalendarIdEventIdExtensionId
func ParseUserIdCalendarGroupIdCalendarIdEventIdExtensionID(input string) (*UserIdCalendarGroupIdCalendarIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupIdCalendarIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupIdCalendarIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarGroupIdCalendarIdEventIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarGroupIdCalendarIdEventIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarGroupIdCalendarIdEventIdExtensionIDInsensitively(input string) (*UserIdCalendarGroupIdCalendarIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarGroupIdCalendarIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarGroupIdCalendarIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarGroupIdCalendarIdEventIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CalendarGroupId, ok = input.Parsed["calendarGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarGroupId", input)
	}

	if id.CalendarId, ok = input.Parsed["calendarId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "calendarId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdCalendarGroupIdCalendarIdEventIdExtensionID checks that 'input' can be parsed as a User Id Calendar Group Id Calendar Id Event Id Extension ID
func ValidateUserIdCalendarGroupIdCalendarIdEventIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarGroupIdCalendarIdEventIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Group Id Calendar Id Event Id Extension ID
func (id UserIdCalendarGroupIdCalendarIdEventIdExtensionId) ID() string {
	fmtString := "/users/%s/calendarGroups/%s/calendars/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarGroupId, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Group Id Calendar Id Event Id Extension ID
func (id UserIdCalendarGroupIdCalendarIdEventIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarGroups", "calendarGroups", "calendarGroups"),
		resourceids.UserSpecifiedSegment("calendarGroupId", "calendarGroupIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Group Id Calendar Id Event Id Extension ID
func (id UserIdCalendarGroupIdCalendarIdEventIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar Group: %q", id.CalendarGroupId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Calendar Group Id Calendar Id Event Id Extension (%s)", strings.Join(components, "\n"))
}
