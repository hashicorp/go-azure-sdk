package userpermissiongrant

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserPermissionGrantId{}

// UserPermissionGrantId is a struct representing the Resource ID for a User Permission Grant
type UserPermissionGrantId struct {
	UserId                            string
	ResourceSpecificPermissionGrantId string
}

// NewUserPermissionGrantID returns a new UserPermissionGrantId struct
func NewUserPermissionGrantID(userId string, resourceSpecificPermissionGrantId string) UserPermissionGrantId {
	return UserPermissionGrantId{
		UserId:                            userId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseUserPermissionGrantID parses 'input' into a UserPermissionGrantId
func ParseUserPermissionGrantID(input string) (*UserPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPermissionGrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ParseUserPermissionGrantIDInsensitively parses 'input' case-insensitively into a UserPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserPermissionGrantIDInsensitively(input string) (*UserPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserPermissionGrantId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ResourceSpecificPermissionGrantId, ok = parsed.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", *parsed)
	}

	return &id, nil
}

// ValidateUserPermissionGrantID checks that 'input' can be parsed as a User Permission Grant ID
func ValidateUserPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Permission Grant ID
func (id UserPermissionGrantId) ID() string {
	fmtString := "/users/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Permission Grant ID
func (id UserPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticPermissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantIdValue"),
	}
}

// String returns a human-readable description of this User Permission Grant ID
func (id UserPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("User Permission Grant (%s)", strings.Join(components, "\n"))
}
