package rolemanagemententerpriseapproleassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentId{}

// RoleManagementEnterpriseAppRoleAssignmentId is a struct representing the Resource ID for a Role Management Enterprise App Role Assignment
type RoleManagementEnterpriseAppRoleAssignmentId struct {
	RbacApplicationId       string
	UnifiedRoleAssignmentId string
}

// NewRoleManagementEnterpriseAppRoleAssignmentID returns a new RoleManagementEnterpriseAppRoleAssignmentId struct
func NewRoleManagementEnterpriseAppRoleAssignmentID(rbacApplicationId string, unifiedRoleAssignmentId string) RoleManagementEnterpriseAppRoleAssignmentId {
	return RoleManagementEnterpriseAppRoleAssignmentId{
		RbacApplicationId:       rbacApplicationId,
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementEnterpriseAppRoleAssignmentID parses 'input' into a RoleManagementEnterpriseAppRoleAssignmentId
func ParseRoleManagementEnterpriseAppRoleAssignmentID(input string) (*RoleManagementEnterpriseAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentId, ok = parsed.Parsed["unifiedRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleAssignmentIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentId, ok = parsed.Parsed["unifiedRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleAssignmentID checks that 'input' can be parsed as a Role Management Enterprise App Role Assignment ID
func ValidateRoleManagementEnterpriseAppRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Assignment ID
func (id RoleManagementEnterpriseAppRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignments/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Assignment ID
func (id RoleManagementEnterpriseAppRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleAssignments", "roleAssignments", "roleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Assignment ID
func (id RoleManagementEnterpriseAppRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Assignment (%s)", strings.Join(components, "\n"))
}
