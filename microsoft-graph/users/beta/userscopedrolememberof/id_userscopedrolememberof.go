package userscopedrolememberof

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserScopedRoleMemberOfId{}

// UserScopedRoleMemberOfId is a struct representing the Resource ID for a User Scoped Role Member Of
type UserScopedRoleMemberOfId struct {
	UserId                 string
	ScopedRoleMembershipId string
}

// NewUserScopedRoleMemberOfID returns a new UserScopedRoleMemberOfId struct
func NewUserScopedRoleMemberOfID(userId string, scopedRoleMembershipId string) UserScopedRoleMemberOfId {
	return UserScopedRoleMemberOfId{
		UserId:                 userId,
		ScopedRoleMembershipId: scopedRoleMembershipId,
	}
}

// ParseUserScopedRoleMemberOfID parses 'input' into a UserScopedRoleMemberOfId
func ParseUserScopedRoleMemberOfID(input string) (*UserScopedRoleMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserScopedRoleMemberOfId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserScopedRoleMemberOfId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ScopedRoleMembershipId, ok = parsed.Parsed["scopedRoleMembershipId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", *parsed)
	}

	return &id, nil
}

// ParseUserScopedRoleMemberOfIDInsensitively parses 'input' case-insensitively into a UserScopedRoleMemberOfId
// note: this method should only be used for API response data and not user input
func ParseUserScopedRoleMemberOfIDInsensitively(input string) (*UserScopedRoleMemberOfId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserScopedRoleMemberOfId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserScopedRoleMemberOfId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ScopedRoleMembershipId, ok = parsed.Parsed["scopedRoleMembershipId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scopedRoleMembershipId", *parsed)
	}

	return &id, nil
}

// ValidateUserScopedRoleMemberOfID checks that 'input' can be parsed as a User Scoped Role Member Of ID
func ValidateUserScopedRoleMemberOfID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserScopedRoleMemberOfID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Scoped Role Member Of ID
func (id UserScopedRoleMemberOfId) ID() string {
	fmtString := "/users/%s/scopedRoleMemberOf/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ScopedRoleMembershipId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Scoped Role Member Of ID
func (id UserScopedRoleMemberOfId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticScopedRoleMemberOf", "scopedRoleMemberOf", "scopedRoleMemberOf"),
		resourceids.UserSpecifiedSegment("scopedRoleMembershipId", "scopedRoleMembershipIdValue"),
	}
}

// String returns a human-readable description of this User Scoped Role Member Of ID
func (id UserScopedRoleMemberOfId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Scoped Role Membership: %q", id.ScopedRoleMembershipId),
	}
	return fmt.Sprintf("User Scoped Role Member Of (%s)", strings.Join(components, "\n"))
}
