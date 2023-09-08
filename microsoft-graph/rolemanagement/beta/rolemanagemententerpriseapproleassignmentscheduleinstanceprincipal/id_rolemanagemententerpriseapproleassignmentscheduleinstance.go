package rolemanagemententerpriseapproleassignmentscheduleinstanceprincipal

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{}

// RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId is a struct representing the Resource ID for a Role Management Enterprise App Role Assignment Schedule Instance
type RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId struct {
	RbacApplicationId                       string
	UnifiedRoleAssignmentScheduleInstanceId string
}

// NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID returns a new RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId struct
func NewRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID(rbacApplicationId string, unifiedRoleAssignmentScheduleInstanceId string) RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId {
	return RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{
		RbacApplicationId:                       rbacApplicationId,
		UnifiedRoleAssignmentScheduleInstanceId: unifiedRoleAssignmentScheduleInstanceId,
	}
}

// ParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID parses 'input' into a RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId
func ParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID(input string) (*RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentScheduleInstanceId, ok = parsed.Parsed["unifiedRoleAssignmentScheduleInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleInstanceId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentScheduleInstanceId, ok = parsed.Parsed["unifiedRoleAssignmentScheduleInstanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleInstanceId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID checks that 'input' can be parsed as a Role Management Enterprise App Role Assignment Schedule Instance ID
func ValidateRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleAssignmentScheduleInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Assignment Schedule Instance ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentScheduleInstances/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentScheduleInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Assignment Schedule Instance ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleAssignmentScheduleInstances", "roleAssignmentScheduleInstances", "roleAssignmentScheduleInstances"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleInstanceId", "unifiedRoleAssignmentScheduleInstanceIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Assignment Schedule Instance ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment Schedule Instance: %q", id.UnifiedRoleAssignmentScheduleInstanceId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Assignment Schedule Instance (%s)", strings.Join(components, "\n"))
}
