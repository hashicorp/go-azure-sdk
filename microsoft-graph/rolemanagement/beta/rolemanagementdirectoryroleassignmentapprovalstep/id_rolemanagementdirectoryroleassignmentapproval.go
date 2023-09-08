package rolemanagementdirectoryroleassignmentapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDirectoryRoleAssignmentApprovalId{}

// RoleManagementDirectoryRoleAssignmentApprovalId is a struct representing the Resource ID for a Role Management Directory Role Assignment Approval
type RoleManagementDirectoryRoleAssignmentApprovalId struct {
	ApprovalId string
}

// NewRoleManagementDirectoryRoleAssignmentApprovalID returns a new RoleManagementDirectoryRoleAssignmentApprovalId struct
func NewRoleManagementDirectoryRoleAssignmentApprovalID(approvalId string) RoleManagementDirectoryRoleAssignmentApprovalId {
	return RoleManagementDirectoryRoleAssignmentApprovalId{
		ApprovalId: approvalId,
	}
}

// ParseRoleManagementDirectoryRoleAssignmentApprovalID parses 'input' into a RoleManagementDirectoryRoleAssignmentApprovalId
func ParseRoleManagementDirectoryRoleAssignmentApprovalID(input string) (*RoleManagementDirectoryRoleAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryRoleAssignmentApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryRoleAssignmentApprovalId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleAssignmentApprovalIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleAssignmentApprovalId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleAssignmentApprovalIDInsensitively(input string) (*RoleManagementDirectoryRoleAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryRoleAssignmentApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryRoleAssignmentApprovalId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDirectoryRoleAssignmentApprovalID checks that 'input' can be parsed as a Role Management Directory Role Assignment Approval ID
func ValidateRoleManagementDirectoryRoleAssignmentApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleAssignmentApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Assignment Approval ID
func (id RoleManagementDirectoryRoleAssignmentApprovalId) ID() string {
	fmtString := "/roleManagement/directory/roleAssignmentApprovals/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Assignment Approval ID
func (id RoleManagementDirectoryRoleAssignmentApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticRoleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Assignment Approval ID
func (id RoleManagementDirectoryRoleAssignmentApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Role Management Directory Role Assignment Approval (%s)", strings.Join(components, "\n"))
}
