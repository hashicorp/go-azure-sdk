package privilegedaccesgroupassignmentapprovalstage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId{}

// IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId is a struct representing the Resource ID for a Identity Governance Privileged Acces Group Assignment Approval
type IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId struct {
	ApprovalId string
}

// NewIdentityGovernancePrivilegedAccesGroupAssignmentApprovalID returns a new IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId struct
func NewIdentityGovernancePrivilegedAccesGroupAssignmentApprovalID(approvalId string) IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId {
	return IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId{
		ApprovalId: approvalId,
	}
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalID parses 'input' into a IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalID(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIDInsensitively(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccesGroupAssignmentApprovalID checks that 'input' can be parsed as a Identity Governance Privileged Acces Group Assignment Approval ID
func ValidateIdentityGovernancePrivilegedAccesGroupAssignmentApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Acces Group Assignment Approval ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentApprovals/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Acces Group Assignment Approval ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentApprovals", "assignmentApprovals", "assignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Acces Group Assignment Approval ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Identity Governance Privileged Acces Group Assignment Approval (%s)", strings.Join(components, "\n"))
}
