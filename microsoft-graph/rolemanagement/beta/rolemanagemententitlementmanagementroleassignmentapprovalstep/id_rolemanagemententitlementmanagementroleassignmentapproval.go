package rolemanagemententitlementmanagementroleassignmentapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementEntitlementManagementRoleAssignmentApprovalId{}

// RoleManagementEntitlementManagementRoleAssignmentApprovalId is a struct representing the Resource ID for a Role Management Entitlement Management Role Assignment Approval
type RoleManagementEntitlementManagementRoleAssignmentApprovalId struct {
	ApprovalId string
}

// NewRoleManagementEntitlementManagementRoleAssignmentApprovalID returns a new RoleManagementEntitlementManagementRoleAssignmentApprovalId struct
func NewRoleManagementEntitlementManagementRoleAssignmentApprovalID(approvalId string) RoleManagementEntitlementManagementRoleAssignmentApprovalId {
	return RoleManagementEntitlementManagementRoleAssignmentApprovalId{
		ApprovalId: approvalId,
	}
}

// ParseRoleManagementEntitlementManagementRoleAssignmentApprovalID parses 'input' into a RoleManagementEntitlementManagementRoleAssignmentApprovalId
func ParseRoleManagementEntitlementManagementRoleAssignmentApprovalID(input string) (*RoleManagementEntitlementManagementRoleAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementRoleAssignmentApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementRoleAssignmentApprovalId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIDInsensitively parses 'input' case-insensitively into a RoleManagementEntitlementManagementRoleAssignmentApprovalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementEntitlementManagementRoleAssignmentApprovalIDInsensitively(input string) (*RoleManagementEntitlementManagementRoleAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementEntitlementManagementRoleAssignmentApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementEntitlementManagementRoleAssignmentApprovalId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementEntitlementManagementRoleAssignmentApprovalID checks that 'input' can be parsed as a Role Management Entitlement Management Role Assignment Approval ID
func ValidateRoleManagementEntitlementManagementRoleAssignmentApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementEntitlementManagementRoleAssignmentApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Entitlement Management Role Assignment Approval ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalId) ID() string {
	fmtString := "/roleManagement/entitlementManagement/roleAssignmentApprovals/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Entitlement Management Role Assignment Approval ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticEntitlementManagement", "entitlementManagement", "entitlementManagement"),
		resourceids.StaticSegment("staticRoleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
	}
}

// String returns a human-readable description of this Role Management Entitlement Management Role Assignment Approval ID
func (id RoleManagementEntitlementManagementRoleAssignmentApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Role Management Entitlement Management Role Assignment Approval (%s)", strings.Join(components, "\n"))
}
