package rolemanagemententerpriseapproleassignmentscheduleactivatedusing

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentScheduleId{}

// RoleManagementEnterpriseAppRoleAssignmentScheduleId is a struct representing the Resource ID for a Role Management Enterprise App Role Assignment Schedule
type RoleManagementEnterpriseAppRoleAssignmentScheduleId struct {
	RbacApplicationId               string
	UnifiedRoleAssignmentScheduleId string
}

// NewRoleManagementEnterpriseAppRoleAssignmentScheduleID returns a new RoleManagementEnterpriseAppRoleAssignmentScheduleId struct
func NewRoleManagementEnterpriseAppRoleAssignmentScheduleID(rbacApplicationId string, unifiedRoleAssignmentScheduleId string) RoleManagementEnterpriseAppRoleAssignmentScheduleId {
	return RoleManagementEnterpriseAppRoleAssignmentScheduleId{
		RbacApplicationId:               rbacApplicationId,
		UnifiedRoleAssignmentScheduleId: unifiedRoleAssignmentScheduleId,
	}
}

// ParseRoleManagementEnterpriseAppRoleAssignmentScheduleID parses 'input' into a RoleManagementEnterpriseAppRoleAssignmentScheduleId
func ParseRoleManagementEnterpriseAppRoleAssignmentScheduleID(input string) (*RoleManagementEnterpriseAppRoleAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentScheduleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentScheduleId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentScheduleId, ok = parsed.Parsed["unifiedRoleAssignmentScheduleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleAssignmentScheduleIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleAssignmentScheduleId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleAssignmentScheduleIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleAssignmentScheduleId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentScheduleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentScheduleId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentScheduleId, ok = parsed.Parsed["unifiedRoleAssignmentScheduleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleAssignmentScheduleID checks that 'input' can be parsed as a Role Management Enterprise App Role Assignment Schedule ID
func ValidateRoleManagementEnterpriseAppRoleAssignmentScheduleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleAssignmentScheduleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Assignment Schedule ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentSchedules/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentScheduleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Assignment Schedule ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleAssignmentSchedules", "roleAssignmentSchedules", "roleAssignmentSchedules"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleId", "unifiedRoleAssignmentScheduleIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Assignment Schedule ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment Schedule: %q", id.UnifiedRoleAssignmentScheduleId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Assignment Schedule (%s)", strings.Join(components, "\n"))
}
