package userapproleassignedresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserAppRoleAssignedResourceId{}

// UserAppRoleAssignedResourceId is a struct representing the Resource ID for a User App Role Assigned Resource
type UserAppRoleAssignedResourceId struct {
	UserId             string
	ServicePrincipalId string
}

// NewUserAppRoleAssignedResourceID returns a new UserAppRoleAssignedResourceId struct
func NewUserAppRoleAssignedResourceID(userId string, servicePrincipalId string) UserAppRoleAssignedResourceId {
	return UserAppRoleAssignedResourceId{
		UserId:             userId,
		ServicePrincipalId: servicePrincipalId,
	}
}

// ParseUserAppRoleAssignedResourceID parses 'input' into a UserAppRoleAssignedResourceId
func ParseUserAppRoleAssignedResourceID(input string) (*UserAppRoleAssignedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppRoleAssignedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppRoleAssignedResourceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	return &id, nil
}

// ParseUserAppRoleAssignedResourceIDInsensitively parses 'input' case-insensitively into a UserAppRoleAssignedResourceId
// note: this method should only be used for API response data and not user input
func ParseUserAppRoleAssignedResourceIDInsensitively(input string) (*UserAppRoleAssignedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserAppRoleAssignedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserAppRoleAssignedResourceId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ServicePrincipalId, ok = parsed.Parsed["servicePrincipalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "servicePrincipalId", *parsed)
	}

	return &id, nil
}

// ValidateUserAppRoleAssignedResourceID checks that 'input' can be parsed as a User App Role Assigned Resource ID
func ValidateUserAppRoleAssignedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserAppRoleAssignedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User App Role Assigned Resource ID
func (id UserAppRoleAssignedResourceId) ID() string {
	fmtString := "/users/%s/appRoleAssignedResources/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ServicePrincipalId)
}

// Segments returns a slice of Resource ID Segments which comprise this User App Role Assigned Resource ID
func (id UserAppRoleAssignedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticAppRoleAssignedResources", "appRoleAssignedResources", "appRoleAssignedResources"),
		resourceids.UserSpecifiedSegment("servicePrincipalId", "servicePrincipalIdValue"),
	}
}

// String returns a human-readable description of this User App Role Assigned Resource ID
func (id UserAppRoleAssignedResourceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Service Principal: %q", id.ServicePrincipalId),
	}
	return fmt.Sprintf("User App Role Assigned Resource (%s)", strings.Join(components, "\n"))
}
