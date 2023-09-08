package usertransitivememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserTransitiveMemberOfId{}

// UserTransitiveMemberOfId is a struct representing the Resource ID for a User Transitive Member Of
type UserTransitiveMemberOfId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserTransitiveMemberOfID returns a new UserTransitiveMemberOfId struct
func NewUserTransitiveMemberOfID(userId string, directoryObjectId string) UserTransitiveMemberOfId {
	return UserTransitiveMemberOfId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserTransitiveMemberOfID parses 'input' into a UserTransitiveMemberOfId
func ParseUserTransitiveMemberOfID(input string) (*UserTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTransitiveMemberOfId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParseUserTransitiveMemberOfIDInsensitively parses 'input' case-insensitively into a UserTransitiveMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserTransitiveMemberOfIDInsensitively(input string) (*UserTransitiveMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserTransitiveMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserTransitiveMemberOfId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidateUserTransitiveMemberOfID checks that 'input' can be parsed as a User Transitive Member Of ID
func ValidateUserTransitiveMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserTransitiveMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Transitive Member Of ID
func (id UserTransitiveMemberOfId) ID() string {
	fmtString := "/users/%s/transitiveMemberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Transitive Member Of ID
func (id UserTransitiveMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticTransitiveMemberOf", "transitiveMemberOf", "transitiveMemberOf"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this User Transitive Member Of ID
func (id UserTransitiveMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Transitive Member Of (%s)", strings.Join(components, "\n"))
}
