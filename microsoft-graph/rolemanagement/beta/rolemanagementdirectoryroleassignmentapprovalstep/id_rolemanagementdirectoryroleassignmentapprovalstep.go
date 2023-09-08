package rolemanagementdirectoryroleassignmentapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RoleManagementDirectoryRoleAssignmentApprovalStepId{}

// RoleManagementDirectoryRoleAssignmentApprovalStepId is a struct representing the Resource ID for a Role Management Directory Role Assignment Approval Step
type RoleManagementDirectoryRoleAssignmentApprovalStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewRoleManagementDirectoryRoleAssignmentApprovalStepID returns a new RoleManagementDirectoryRoleAssignmentApprovalStepId struct
func NewRoleManagementDirectoryRoleAssignmentApprovalStepID(approvalId string, approvalStepId string) RoleManagementDirectoryRoleAssignmentApprovalStepId {
	return RoleManagementDirectoryRoleAssignmentApprovalStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseRoleManagementDirectoryRoleAssignmentApprovalStepID parses 'input' into a RoleManagementDirectoryRoleAssignmentApprovalStepId
func ParseRoleManagementDirectoryRoleAssignmentApprovalStepID(input string) (*RoleManagementDirectoryRoleAssignmentApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryRoleAssignmentApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryRoleAssignmentApprovalStepId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ParseRoleManagementDirectoryRoleAssignmentApprovalStepIDInsensitively parses 'input' case-insensitively into a RoleManagementDirectoryRoleAssignmentApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementDirectoryRoleAssignmentApprovalStepIDInsensitively(input string) (*RoleManagementDirectoryRoleAssignmentApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(RoleManagementDirectoryRoleAssignmentApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RoleManagementDirectoryRoleAssignmentApprovalStepId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ValidateRoleManagementDirectoryRoleAssignmentApprovalStepID checks that 'input' can be parsed as a Role Management Directory Role Assignment Approval Step ID
func ValidateRoleManagementDirectoryRoleAssignmentApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementDirectoryRoleAssignmentApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Directory Role Assignment Approval Step ID
func (id RoleManagementDirectoryRoleAssignmentApprovalStepId) ID() string {
	fmtString := "/roleManagement/directory/roleAssignmentApprovals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Directory Role Assignment Approval Step ID
func (id RoleManagementDirectoryRoleAssignmentApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticRoleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("staticDirectory", "directory", "directory"),
		resourceids.StaticSegment("staticRoleAssignmentApprovals", "roleAssignmentApprovals", "roleAssignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepIdValue"),
	}
}

// String returns a human-readable description of this Role Management Directory Role Assignment Approval Step ID
func (id RoleManagementDirectoryRoleAssignmentApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Role Management Directory Role Assignment Approval Step (%s)", strings.Join(components, "\n"))
}
