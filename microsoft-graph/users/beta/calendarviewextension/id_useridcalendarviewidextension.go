package calendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCalendarViewIdExtensionId{}

// UserIdCalendarViewIdExtensionId is a struct representing the Resource ID for a User Id Calendar View Id Extension
type UserIdCalendarViewIdExtensionId struct {
	UserId      string
	EventId     string
	ExtensionId string
}

// NewUserIdCalendarViewIdExtensionID returns a new UserIdCalendarViewIdExtensionId struct
func NewUserIdCalendarViewIdExtensionID(userId string, eventId string, extensionId string) UserIdCalendarViewIdExtensionId {
	return UserIdCalendarViewIdExtensionId{
		UserId:      userId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseUserIdCalendarViewIdExtensionID parses 'input' into a UserIdCalendarViewIdExtensionId
func ParseUserIdCalendarViewIdExtensionID(input string) (*UserIdCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCalendarViewIdExtensionIDInsensitively parses 'input' case-insensitively into a UserIdCalendarViewIdExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserIdCalendarViewIdExtensionIDInsensitively(input string) (*UserIdCalendarViewIdExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCalendarViewIdExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCalendarViewIdExtensionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCalendarViewIdExtensionId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateUserIdCalendarViewIdExtensionID checks that 'input' can be parsed as a User Id Calendar View Id Extension ID
func ValidateUserIdCalendarViewIdExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCalendarViewIdExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Calendar View Id Extension ID
func (id UserIdCalendarViewIdExtensionId) ID() string {
	fmtString := "/users/%s/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Calendar View Id Extension ID
func (id UserIdCalendarViewIdExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("calendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("extensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Id Calendar View Id Extension ID
func (id UserIdCalendarViewIdExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Id Calendar View Id Extension (%s)", strings.Join(components, "\n"))
}
