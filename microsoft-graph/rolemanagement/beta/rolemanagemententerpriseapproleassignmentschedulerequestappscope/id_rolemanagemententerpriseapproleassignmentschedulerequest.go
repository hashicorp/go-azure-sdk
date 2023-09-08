package rolemanagemententerpriseapproleassignmentschedulerequestappscope

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId{}

// RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId is a struct representing the Resource ID for a Role Management Enterprise App Role Assignment Schedule Request
type RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId struct {
	RbacApplicationId                      string
	UnifiedRoleAssignmentScheduleRequestId string
}

// NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestID returns a new RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId struct
func NewRoleManagementEnterpriseAppRoleAssignmentScheduleRequestID(rbacApplicationId string, unifiedRoleAssignmentScheduleRequestId string) RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId {
	return RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId{
		RbacApplicationId:                      rbacApplicationId,
		UnifiedRoleAssignmentScheduleRequestId: unifiedRoleAssignmentScheduleRequestId,
	}
}

// ParseRoleManagementEnterpriseAppRoleAssignmentScheduleRequestID parses 'input' into a RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId
func ParseRoleManagementEnterpriseAppRoleAssignmentScheduleRequestID(input string) (*RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentScheduleRequestId, ok = parsed.Parsed["unifiedRoleAssignmentScheduleRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleRequestId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleAssignmentScheduleRequestIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleAssignmentScheduleRequestIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.UnifiedRoleAssignmentScheduleRequestId, ok = parsed.Parsed["unifiedRoleAssignmentScheduleRequestId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleAssignmentScheduleRequestId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleAssignmentScheduleRequestID checks that 'input' can be parsed as a Role Management Enterprise App Role Assignment Schedule Request ID
func ValidateRoleManagementEnterpriseAppRoleAssignmentScheduleRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleAssignmentScheduleRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Assignment Schedule Request ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentScheduleRequests/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.UnifiedRoleAssignmentScheduleRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Assignment Schedule Request ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleAssignmentScheduleRequests", "roleAssignmentScheduleRequests", "roleAssignmentScheduleRequests"),
		resourceids.UserSpecifiedSegment("unifiedRoleAssignmentScheduleRequestId", "unifiedRoleAssignmentScheduleRequestIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Assignment Schedule Request ID
func (id RoleManagementEnterpriseAppRoleAssignmentScheduleRequestId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Unified Role Assignment Schedule Request: %q", id.UnifiedRoleAssignmentScheduleRequestId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Assignment Schedule Request (%s)", strings.Join(components, "\n"))
}
