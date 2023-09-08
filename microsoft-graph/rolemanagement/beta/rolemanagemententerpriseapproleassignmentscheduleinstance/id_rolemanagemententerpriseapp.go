package rolemanagemententerpriseapproleassignmentscheduleinstance

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppId{}

// RoleManagementEnterpriseAppId is a struct representing the Resource ID for a Role Management Enterprise App
type RoleManagementEnterpriseAppId struct {
	RbacApplicationId string
}

// NewRoleManagementEnterpriseAppID returns a new RoleManagementEnterpriseAppId struct
func NewRoleManagementEnterpriseAppID(rbacApplicationId string) RoleManagementEnterpriseAppId {
	return RoleManagementEnterpriseAppId{
		RbacApplicationId: rbacApplicationId,
	}
}

// ParseRoleManagementEnterpriseAppID parses 'input' into a RoleManagementEnterpriseAppId
func ParseRoleManagementEnterpriseAppID(input string) (*RoleManagementEnterpriseAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppIDInsensitively(input string) (*RoleManagementEnterpriseAppId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppID checks that 'input' can be parsed as a Role Management Enterprise App ID
func ValidateRoleManagementEnterpriseAppID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App ID
func (id RoleManagementEnterpriseAppId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App ID
func (id RoleManagementEnterpriseAppId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App ID
func (id RoleManagementEnterpriseAppId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
	}
	return fmt.Sprintf("Role Management Enterprise App (%s)", strings.Join(components, "\n"))
}
