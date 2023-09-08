package usercreatedobject

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserCreatedObjectId{}

// UserCreatedObjectId is a struct representing the Resource ID for a User Created Object
type UserCreatedObjectId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserCreatedObjectID returns a new UserCreatedObjectId struct
func NewUserCreatedObjectID(userId string, directoryObjectId string) UserCreatedObjectId {
	return UserCreatedObjectId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserCreatedObjectID parses 'input' into a UserCreatedObjectId
func ParseUserCreatedObjectID(input string) (*UserCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCreatedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCreatedObjectId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserCreatedObjectIDInsensitively parses 'input' case-insensitively into a UserCreatedObjectId
// note: this method should only be used for API response data and not user input
func ParseUserCreatedObjectIDInsensitively(input string) (*UserCreatedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserCreatedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserCreatedObjectId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserCreatedObjectID checks that 'input' can be parsed as a User Created Object ID
func ValidateUserCreatedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserCreatedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Created Object ID
func (id UserCreatedObjectId) ID() string {
	fmtString := "/users/%s/createdObjects/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Created Object ID
func (id UserCreatedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticCreatedObjects", "createdObjects", "createdObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Created Object ID
func (id UserCreatedObjectId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Created Object (%s)", strings.Join(components, "\n"))
}
