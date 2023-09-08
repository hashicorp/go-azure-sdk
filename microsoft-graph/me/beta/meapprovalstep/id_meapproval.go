package meapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MeApprovalId{}

// MeApprovalId is a struct representing the Resource ID for a Me Approval
type MeApprovalId struct {
	ApprovalId string
}

// NewMeApprovalID returns a new MeApprovalId struct
func NewMeApprovalID(approvalId string) MeApprovalId {
	return MeApprovalId{
		ApprovalId: approvalId,
	}
}

// ParseMeApprovalID parses 'input' into a MeApprovalId
func ParseMeApprovalID(input string) (*MeApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeApprovalId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ParseMeApprovalIDInsensitively parses 'input' case-insensitively into a MeApprovalId
// note: this method should only be used for API response data and not user input
func ParseMeApprovalIDInsensitively(input string) (*MeApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(MeApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MeApprovalId{}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ValidateMeApprovalID checks that 'input' can be parsed as a Me Approval ID
func ValidateMeApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Approval ID
func (id MeApprovalId) ID() string {
	fmtString := "/me/approvals/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Approval ID
func (id MeApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticMe", "me", "me"),
		resourceids.StaticSegment("staticApprovals", "approvals", "approvals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
	}
}

// String returns a human-readable description of this Me Approval ID
func (id MeApprovalId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("Me Approval (%s)", strings.Join(components, "\n"))
}
