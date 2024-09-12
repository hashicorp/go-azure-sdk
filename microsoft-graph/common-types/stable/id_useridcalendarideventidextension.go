package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarIdEventIdExtensionId{}

// UserIdCalendarIdEventIdExtensionId is a struct representing the Resource ID for a User Id Calendar Id Event Id Extension
type UserIdCalendarIdEventIdExtensionId struct {
	UserId      string
	CalendarId  string
	EventId     string
	ExtensionId string
}

// NewUserIdCalendarIdEventIdExtensionID returns a new UserIdCalendarIdEventIdExtensionId struct
func NewUserIdCalendarIdEventIdExtensionID(userId string, calendarId string, eventId string, extensionId string) UserIdCalendarIdEventIdExtensionId {
	return UserIdCalendarIdEventIdExtensionId{
		UserId:      userId,
		CalendarId:  calendarId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdCalendarIdEventIdExtensionID parses 'input' into a UserIdCalendarIdEventIdExtensionId
func ParseUserIdCalendarIdEventIdExtensionID(input string) (*UserIdCalendarIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarIdEventIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarIdEventIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarIdEventIdExtensionIDInsensitively(input string) (*UserIdCalendarIdEventIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarIdEventIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarIdEventIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarIdEventIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdCalendarIdEventIdExtensionID checks that 'input' can be parsed as a User Id Calendar Id Event Id Extension ID
func ValidateUserIdCalendarIdEventIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarIdEventIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Id Event Id Extension ID
func (id UserIdCalendarIdEventIdExtensionId) ID() string {
	fmtString := "/users/%s/calendars/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CalendarId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Id Event Id Extension ID
func (id UserIdCalendarIdEventIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendars", "calendars", "calendars"),
		resourceids.UserSpecifiedSegment("calendarId", "calendarIdValue"),
		resourceids.StaticSegment("events", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Id Event Id Extension ID
func (id UserIdCalendarIdEventIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Calendar: %q", id.CalendarId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Calendar Id Event Id Extension (%s)", strings.Join(components, "\n"))
}
