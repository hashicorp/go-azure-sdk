package userapprovalstep

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserApprovalStepId{}

// UserApprovalStepId is a struct representing the Resource ID for a User Approval Step
type UserApprovalStepId struct {
	UserId         string
	ApprovalId     string
	ApprovalStepId string
}

// NewUserApprovalStepID returns a new UserApprovalStepId struct
func NewUserApprovalStepID(userId string, approvalId string, approvalStepId string) UserApprovalStepId {
	return UserApprovalStepId{
		UserId:         userId,
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseUserApprovalStepID parses 'input' into a UserApprovalStepId
func ParseUserApprovalStepID(input string) (*UserApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserApprovalStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserApprovalStepId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ParseUserApprovalStepIDInsensitively parses 'input' case-insensitively into a UserApprovalStepId
// note: this method should only be used for API response data and not user input
func ParseUserApprovalStepIDInsensitively(input string) (*UserApprovalStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserApprovalStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserApprovalStepId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	if id.ApprovalStepId, ok = parsed.Parsed["approvalStepId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", *parsed)
	}

	return &id, nil
}

// ValidateUserApprovalStepID checks that 'input' can be parsed as a User Approval Step ID
func ValidateUserApprovalStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserApprovalStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Approval Step ID
func (id UserApprovalStepId) ID() string {
	fmtString := "/users/%s/approvals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Approval Step ID
func (id UserApprovalStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticApprovals", "approvals", "approvals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
		resourceids.StaticSegment("staticSteps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepIdValue"),
	}
}

// String returns a human-readable description of this User Approval Step ID
func (id UserApprovalStepId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("User Approval Step (%s)", strings.Join(components, "\n"))
}
