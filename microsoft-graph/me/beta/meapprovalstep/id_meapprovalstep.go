package meapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeApprovalStepId{}

// MeApprovalStepId is a struct representing the Resource ID for a Me Approval Step
type MeApprovalStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewMeApprovalStepID returns a new MeApprovalStepId struct
func NewMeApprovalStepID(approvalId string, approvalStepId string) MeApprovalStepId {
	return MeApprovalStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseMeApprovalStepID parses 'input' into a MeApprovalStepId
func ParseMeApprovalStepID(input string) (*MeApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeApprovalStepId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ParseMeApprovalStepIDInsensitively parses 'input' case-insensitively into a MeApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseMeApprovalStepIDInsensitively(input string) (*MeApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeApprovalStepId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ValidateMeApprovalStepID checks that 'input' can be parsed as a Me Approval Step ID
func ValidateMeApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Approval Step ID
func (id MeApprovalStepId) ID() string {
	fmtString := "/me/approvals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Approval Step ID
func (id MeApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticApprovals", "approvals", "approvals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepIdValue"),
	}
}

// String returns a human-readable description of this Me Approval Step ID
func (id MeApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Me Approval Step (%s)", strings.Join(components, "\n"))
}
