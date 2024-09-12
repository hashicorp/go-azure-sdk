package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarCalendarViewIdExtensionId{}

// UserIdCalendarCalendarViewIdExtensionId is a struct representing the Resource ID for a User Id Calendar Calendar View Id Extension
type UserIdCalendarCalendarViewIdExtensionId struct {
	UserId      string
	EventId     string
	ExtensionId string
}

// NewUserIdCalendarCalendarViewIdExtensionID returns a new UserIdCalendarCalendarViewIdExtensionId struct
func NewUserIdCalendarCalendarViewIdExtensionID(userId string, eventId string, extensionId string) UserIdCalendarCalendarViewIdExtensionId {
	return UserIdCalendarCalendarViewIdExtensionId{
		UserId:      userId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdCalendarCalendarViewIdExtensionID parses 'input' into a UserIdCalendarCalendarViewIdExtensionId
func ParseUserIdCalendarCalendarViewIdExtensionID(input string) (*UserIdCalendarCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarCalendarViewIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarCalendarViewIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarCalendarViewIdExtensionIDInsensitively(input string) (*UserIdCalendarCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarCalendarViewIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.EventId, ok = input.Parsed["eventId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "eventId", input)
	}

	if id.ExtensionId, ok = input.Parsed["extensionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "extensionId", input)
	}

	return nil
}

// ValidateUserIdCalendarCalendarViewIdExtensionID checks that 'input' can be parsed as a User Id Calendar Calendar View Id Extension ID
func ValidateUserIdCalendarCalendarViewIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarCalendarViewIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar Calendar View Id Extension ID
func (id UserIdCalendarCalendarViewIdExtensionId) ID() string {
	fmtString := "/users/%s/calendar/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar Calendar View Id Extension ID
func (id UserIdCalendarCalendarViewIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendar", "calendar", "calendar"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar Calendar View Id Extension ID
func (id UserIdCalendarCalendarViewIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Calendar Calendar View Id Extension (%s)", strings.Join(components, "\n"))
}
