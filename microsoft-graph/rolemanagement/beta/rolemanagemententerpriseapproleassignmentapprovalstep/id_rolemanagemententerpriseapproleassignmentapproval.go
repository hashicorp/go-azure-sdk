package rolemanagemententerpriseapproleassignmentapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentApprovalId{}

// RoleManagementEnterpriseAppRoleAssignmentApprovalId is a struct representing the Resource ID for a Role Management Enterprise App Role Assignment Approval
type RoleManagementEnterpriseAppRoleAssignmentApprovalId struct {
	RbacApplicationId string
	ApprovalId        string
}

// NewRoleManagementEnterpriseAppRoleAssignmentApprovalID returns a new RoleManagementEnterpriseAppRoleAssignmentApprovalId struct
func NewRoleManagementEnterpriseAppRoleAssignmentApprovalID(rbacApplicationId string, approvalId string) RoleManagementEnterpriseAppRoleAssignmentApprovalId {
	return RoleManagementEnterpriseAppRoleAssignmentApprovalId{
		RbacApplicationId: rbacApplicationId,
		ApprovalId:        approvalId,
	}
}

// ParseRoleManagementEnterpriseAppRoleAssignmentApprovalID parses 'input' into a RoleManagementEnterpriseAppRoleAssignmentApprovalId
func ParseRoleManagementEnterpriseAppRoleAssignmentApprovalID(input string) (*RoleManagementEnterpriseAppRoleAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentApprovalId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleAssignmentApprovalIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleAssignmentApprovalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleAssignmentApprovalIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentApprovalId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleAssignmentApprovalID checks that 'input' can be parsed as a Role Management Enterprise App Role Assignment Approval ID
func ValidateRoleManagementEnterpriseAppRoleAssignmentApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleAssignmentApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Assignment Approval ID
func (id RoleManagementEnterpriseAppRoleAssignmentApprovalId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentApprovals/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Assignment Approval ID
func (id RoleManagementEnterpriseAppRoleAssignmentApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Assignment Approval ID
func (id RoleManagementEnterpriseAppRoleAssignmentApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Assignment Approval (%s)", strings.Join(components, "\n"))
}
