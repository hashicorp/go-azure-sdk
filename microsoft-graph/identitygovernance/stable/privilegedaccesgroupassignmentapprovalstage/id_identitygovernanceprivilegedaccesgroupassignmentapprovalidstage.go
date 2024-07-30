package privilegedaccesgroupassignmentapprovalstage

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId{}

// IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId is a struct representing the Resource ID for a Identity Governance Privileged Acces Group Assignment Approval Id Stage
type IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId struct {
	ApprovalId      string
	ApprovalStageId string
}

// NewIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageID returns a new IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId struct
func NewIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageID(approvalId string, approvalStageId string) IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId {
	return IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId{
		ApprovalId:      approvalId,
		ApprovalStageId: approvalStageId,
	}
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageID parses 'input' into a IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageID(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageIDInsensitively parses 'input' case-insensitively into a IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageIDInsensitively(input string) (*IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStageId, ok = input.Parsed["approvalStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStageId", input)
	}

	return nil
}

// ValidateIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageID checks that 'input' can be parsed as a Identity Governance Privileged Acces Group Assignment Approval Id Stage ID
func ValidateIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Privileged Acces Group Assignment Approval Id Stage ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId) ID() string {
	fmtString := "/identityGovernance/privilegedAccess/group/assignmentApprovals/%s/stages/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Privileged Acces Group Assignment Approval Id Stage ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("privilegedAccess", "privilegedAccess", "privilegedAccess"),
		resourceids.StaticSegment("group", "group", "group"),
		resourceids.StaticSegment("assignmentApprovals", "assignmentApprovals", "assignmentApprovals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("approvalStageId", "approvalStageIdValue"),
	}
}

// String returns a human-readable description of this Identity Governance Privileged Acces Group Assignment Approval Id Stage ID
func (id IdentityGovernancePrivilegedAccesGroupAssignmentApprovalIdStageId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Stage: %q", id.ApprovalStageId),
	}
	return fmt.Sprintf("Identity Governance Privileged Acces Group Assignment Approval Id Stage (%s)", strings.Join(components, "\n"))
}
