package userapproval

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = UserApprovalId{}

// UserApprovalId is a struct representing the Resource ID for a User Approval
type UserApprovalId struct {
	UserId     string
	ApprovalId string
}

// NewUserApprovalID returns a new UserApprovalId struct
func NewUserApprovalID(userId string, approvalId string) UserApprovalId {
	return UserApprovalId{
		UserId:     userId,
		ApprovalId: approvalId,
	}
}

// ParseUserApprovalID parses 'input' into a UserApprovalId
func ParseUserApprovalID(input string) (*UserApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserApprovalId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserApprovalId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ParseUserApprovalIDInsensitively parses 'input' case-insensitively into a UserApprovalId
// note: this method should only be used for API response data and not user input
func ParseUserApprovalIDInsensitively(input string) (*UserApprovalId, error) {
	parser := resourceids.NewParserFromResourceIdType(UserApprovalId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := UserApprovalId{}

	if id.UserId, ok = parsed.Parsed["userId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "userId", *parsed)
	}

	if id.ApprovalId, ok = parsed.Parsed["approvalId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "approvalId", *parsed)
	}

	return &id, nil
}

// ValidateUserApprovalID checks that 'input' can be parsed as a User Approval ID
func ValidateUserApprovalID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserApprovalID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Approval ID
func (id UserApprovalId) ID() string {
	fmtString := "/users/%s/approvals/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ApprovalId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Approval ID
func (id UserApprovalId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticUsers", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userIdValue"),
		resourceids.StaticSegment("staticApprovals", "approvals", "approvals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalIdValue"),
	}
}

// String returns a human-readable description of this User Approval ID
func (id UserApprovalId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Approval: %q", id.ApprovalId),
	}
	return fmt.Sprintf("User Approval (%s)", strings.Join(components, "\n"))
}
