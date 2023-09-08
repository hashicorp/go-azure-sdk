package usereventextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserEventExtensionId{}

// UserEventExtensionId is a struct representing the Resource ID for a User Event Extension
type UserEventExtensionId struct {
	UserId      string
	EventId     string
	ExtensionId string
}

// NewUserEventExtensionID returns a new UserEventExtensionId struct
func NewUserEventExtensionID(userId string, eventId string, extensionId string) UserEventExtensionId {
	return UserEventExtensionId{
		UserId:      userId,
		EventId:     eventId,
		ExtensionId: extensionId,
	}
}

// ParseUserEventExtensionID parses 'input' into a UserEventExtensionId
func ParseUserEventExtensionID(input string) (*UserEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventExtensionId{}

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

// ParseUserEventExtensionIDInsensitively parses 'input' case-insensitively into a UserEventExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserEventExtensionIDInsensitively(input string) (*UserEventExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserEventExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserEventExtensionId{}

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

// ValidateUserEventExtensionID checks that 'input' can be parsed as a User Event Extension ID
func ValidateUserEventExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserEventExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Event Extension ID
func (id UserEventExtensionId) ID() string {
	fmtString := "/users/%s/events/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.EventId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Event Extension ID
func (id UserEventExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticEvents", "events", "events"),
		resourceids.UserSpecifiedSegment("eventId", "eventIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Event Extension ID
func (id UserEventExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Event: %q", id.EventId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Event Extension (%s)", strings.Join(components, "\n"))
}
