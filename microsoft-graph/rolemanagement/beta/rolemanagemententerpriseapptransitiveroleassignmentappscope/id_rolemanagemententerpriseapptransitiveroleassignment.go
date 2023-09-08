package rolemanagemententerpriseapptransitiveroleassignmentappscope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppTransitiveRoleAssignmentId{}

// RoleManagementEnterpriseAppTransitiveRoleAssignmentId is a struct representing the Resource ID for a Role Management Enterprise App Transitive Role Assignment
type RoleManagementEnterpriseAppTransitiveRoleAssignmentId struct {
	RbacApplicationId       string
	UnifiedRoleAssignmentId string
}

// NewRoleManagementEnterpriseAppTransitiveRoleAssignmentID returns a new RoleManagementEnterpriseAppTransitiveRoleAssignmentId struct
func NewRoleManagementEnterpriseAppTransitiveRoleAssignmentID(rbacApplicationId string, unifiedRoleAssignmentId string) RoleManagementEnterpriseAppTransitiveRoleAssignmentId {
	return RoleManagementEnterpriseAppTransitiveRoleAssignmentId{
		RbacApplicationId:       rbacApplicationId,
		UnifiedRoleAssignmentId: unifiedRoleAssignmentId,
	}
}

// ParseRoleManagementEnterpriseAppTransitiveRoleAssignmentID parses 'input' into a RoleManagementEnterpriseAppTransitiveRoleAssignmentId
func ParseRoleManagementEnterpriseAppTransitiveRoleAssignmentID(input string) (*RoleManagementEnterpriseAppTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppTransitiveRoleAssignmentId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentId, ok = parsed.Parsed["unifiedRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppTransitiveRoleAssignmentIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppTransitiveRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppTransitiveRoleAssignmentIDInsensitively(input string) (*RoleManagementEnterpriseAppTransitiveRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppTransitiveRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppTransitiveRoleAssignmentId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentId, ok = parsed.Parsed["unifiedRoleAssignmentId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppTransitiveRoleAssignmentID checks that 'input' can be parsed as a Role Management Enterprise App Transitive Role Assignment ID
func ValidateRoleManagementEnterpriseAppTransitiveRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppTransitiveRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Transitive Role Assignment ID
func (id RoleManagementEnterpriseAppTransitiveRoleAssignmentId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/transitiveRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Transitive Role Assignment ID
func (id RoleManagementEnterpriseAppTransitiveRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticTransitiveRoleAssignments", "transitiveRoleAssignments", "transitiveRoleAssignments"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentId", "unifiedRoleAssignmentIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Transitive Role Assignment ID
func (id RoleManagementEnterpriseAppTransitiveRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment: %q", id.UnifiedRoleAssignmentId),
	}
	return fmt.Sprintf("Role Management Enterprise App Transitive Role Assignment (%s)", strings.Join(components, "\n"))
}
