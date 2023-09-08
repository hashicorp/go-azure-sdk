package rolemanagemententerpriseapproleassignmentapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{}

// RoleManagementEnterpriseAppRoleAssignmentApprovalStepId is a struct representing the Resource ID for a Role Management Enterprise App Role Assignment Approval Step
type RoleManagementEnterpriseAppRoleAssignmentApprovalStepId struct {
	RbacApplicationId string
	ApprovalId        string
	ApprovalStepId    string
}

// NewRoleManagementEnterpriseAppRoleAssignmentApprovalStepID returns a new RoleManagementEnterpriseAppRoleAssignmentApprovalStepId struct
func NewRoleManagementEnterpriseAppRoleAssignmentApprovalStepID(rbacApplicationId string, approvalId string, approvalStepId string) RoleManagementEnterpriseAppRoleAssignmentApprovalStepId {
	return RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{
		RbacApplicationId: rbacApplicationId,
		ApprovalId:        approvalId,
		ApprovalStepId:    approvalStepId,
	}
}

// ParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepID parses 'input' into a RoleManagementEnterpriseAppRoleAssignmentApprovalStepId
func ParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepID(input string) (*RoleManagementEnterpriseAppRoleAssignmentApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepIDInsensitively parses 'input' case-insensitively into a RoleManagementEnterpriseAppRoleAssignmentApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepIDInsensitively(input string) (*RoleManagementEnterpriseAppRoleAssignmentApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEnterpriseAppRoleAssignmentApprovalStepId{}

	if id.RbacApplicationId, ok = parsed.Parsed["rbacApplicationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "rbacApplicationId", *parsed)
	}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEnterpriseAppRoleAssignmentApprovalStepID checks that 'input' can be parsed as a Role Management Enterprise App Role Assignment Approval Step ID
func ValidateRoleManagementEnterpriseAppRoleAssignmentApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEnterpriseAppRoleAssignmentApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Enterprise App Role Assignment Approval Step ID
func (id RoleManagementEnterpriseAppRoleAssignmentApprovalStepId) ID() string {
	fmtString := "/roleManagement/enterpriseApps/%s/roleAssignmentApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.RbacApplicationId, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Enterprise App Role Assignment Approval Step ID
func (id RoleManagementEnterpriseAppRoleAssignmentApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEnterpriseApps", "enterpriseApps", "enterpriseApps"),
		resourceids.UserSpecifiedSegment("rbacApplicationId", "rbacApplicationIdValue"),
		resourceids.StaticSegment("staticRoleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepIdValue"),
	}
}

// String returns a human-readable description of this Role Management Enterprise App Role Assignment Approval Step ID
func (id RoleManagementEnterpriseAppRoleAssignmentApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("Rbac Application: %q", id.RbacApplicationId),
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Role Management Enterprise App Role Assignment Approval Step (%s)", strings.Join(components, "\n"))
}
