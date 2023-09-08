package userextension

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserExtensionId{}

// UserExtensionId is a struct representing the Resource ID for a User Extension
type UserExtensionId struct {
	UserId      string
	ExtensionId string
}

// NewUserExtensionID returns a new UserExtensionId struct
func NewUserExtensionID(userId string, extensionId string) UserExtensionId {
	return UserExtensionId{
		UserId:      userId,
		ExtensionId: extensionId,
	}
}

// ParseUserExtensionID parses 'input' into a UserExtensionId
func ParseUserExtensionID(input string) (*UserExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserExtensionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ParseUserExtensionIDInsensitively parses 'input' case-insensitively into a UserExtensionId
// note: this method should only be used for API response data and not user input
func ParseUserExtensionIDInsensitively(input string) (*UserExtensionId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserExtensionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserExtensionId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ExtensionId, ok = parsed.Parsed["extensionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "extensionId", *parsed)
	}

	return &id, nil
}

// ValidateUserExtensionID checks that 'input' can be parsed as a User Extension ID
func ValidateUserExtensionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserExtensionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Extension ID
func (id UserExtensionId) ID() string {
	fmtString := "/users/%s/extensions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ExtensionId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Extension ID
func (id UserExtensionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticExtensions", "extensions", "extensions"),
		resourceids.UserSpecifiedSegment("extensionId", "extensionIdValue"),
	}
}

// String returns a human-readable description of this User Extension ID
func (id UserExtensionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Extension: %q", id.ExtensionId),
	}
	return fmt.Sprintf("User Extension (%s)", strings.Join(components, "\n"))
}
