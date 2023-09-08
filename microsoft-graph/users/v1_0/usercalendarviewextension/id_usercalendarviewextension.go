package usercalendarviewextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCalendarViewExtensionId{}

// UserCalendarViewExtensionId is a struct representing the Resource ID for a User Calendar View Extension
type UserCalendarViewExtensionId struct {
	UserId      string
	EventId     string
	ExtensionId string
}

// NewUserCalendarViewExtensionID returns a new UserCalendarViewExtensionId struct
func NewUserCalendarViewExtensionID(userId string, eventId string, extensionId string) UserCalendarViewExtensionId {
	return UserCalendarViewExtensionId{
		UserId:      userId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseUserCalendarViewExtensionID parses 'input' into a UserCalendarViewExtensionId
func ParseUserCalendarViewExtensionID(input string) (*UserCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserCalendarViewExtensionIDInsensitively parses 'input' case-insensitively into a UserCalendarViewExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserCalendarViewExtensionIDInsensitively(input string) (*UserCalendarViewExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCalendarViewExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCalendarViewExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.EventId, ok = parsed.Parsed["eventId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "eventId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserCalendarViewExtensionID checks that 'input' can be parsed as a User Calendar View Extension ID
func ValidateUserCalendarViewExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCalendarViewExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Calendar View Extension ID
func (id UserCalendarViewExtensionId) ID() string {
	fmtString := "/users/%s/calendarView/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Calendar View Extension ID
func (id UserCalendarViewExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCalendarView", "calendarView", "calendarView"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Calendar View Extension ID
func (id UserCalendarViewExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Calendar View Extension (%s)", strings.Join(components, "\n"))
}
