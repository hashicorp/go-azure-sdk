package rolemanagemententitlementmanagementroleassignmentapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEntitlementManagementRoleAssignmentApprovalStepId{}

// RoleManagementEntitlementManagementRoleAssignmentApprovalStepId is a struct representing the Resource ID for a Role Management Entitlement Management Role Assignment Approval Step
type RoleManagementEntitlementManagementRoleAssignmentApprovalStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewRoleManagementEntitlementManagementRoleAssignmentApprovalStepID returns a new RoleManagementEntitlementManagementRoleAssignmentApprovalStepId struct
func NewRoleManagementEntitlementManagementRoleAssignmentApprovalStepID(approvalId string, approvalStepId string) RoleManagementEntitlementManagementRoleAssignmentApprovalStepId {
	return RoleManagementEntitlementManagementRoleAssignmentApprovalStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseRoleManagementEntitlementManagementRoleAssignmentApprovalStepID parses 'input' into a RoleManagementEntitlementManagementRoleAssignmentApprovalStepId
func ParseRoleManagementEntitlementManagementRoleAssignmentApprovalStepID(input string) (*RoleManagementEntitlementManagementRoleAssignmentApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementRoleAssignmentApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementRoleAssignmentApprovalStepId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleAssignmentApprovalStepIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleAssignmentApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleAssignmentApprovalStepIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleAssignmentApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementRoleAssignmentApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementRoleAssignmentApprovalStepId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEntitlementManagementRoleAssignmentApprovalStepID checks that 'input' can be parsed as a Role Management Entitlement Management Role Assignment Approval Step ID
func ValidateRoleManagementEntitlementManagementRoleAssignmentApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleAssignmentApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Assignment Approval Step ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalStepId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleAssignmentApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Assignment Approval Step ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEntitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("staticRoleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepIdValue"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Assignment Approval Step ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Assignment Approval Step (%s)", strings.Join(components, "\n"))
}
