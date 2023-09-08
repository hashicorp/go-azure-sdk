package userownedobject

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserOwnedObjectId{}

// UserOwnedObjectId is a struct representing the Resource ID for a User Owned Object
type UserOwnedObjectId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserOwnedObjectID returns a new UserOwnedObjectId struct
func NewUserOwnedObjectID(userId string, directoryObjectId string) UserOwnedObjectId {
	return UserOwnedObjectId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserOwnedObjectID parses 'input' into a UserOwnedObjectId
func ParseUserOwnedObjectID(input string) (*UserOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOwnedObjectId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOwnedObjectId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserOwnedObjectIDInsensitively parses 'input' case-insensitively into a UserOwnedObjectId
// note: this method should only be used for API response data and not user input
func ParseUserOwnedObjectIDInsensitively(input string) (*UserOwnedObjectId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserOwnedObjectId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserOwnedObjectId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserOwnedObjectID checks that 'input' can be parsed as a User Owned Object ID
func ValidateUserOwnedObjectID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserOwnedObjectID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Owned Object ID
func (id UserOwnedObjectId) ID() string {
	fmtString := "/users/%s/ownedObjects/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Owned Object ID
func (id UserOwnedObjectId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticOwnedObjects", "ownedObjects", "ownedObjects"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Owned Object ID
func (id UserOwnedObjectId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Owned Object (%s)", strings.Join(components, "\n"))
}
