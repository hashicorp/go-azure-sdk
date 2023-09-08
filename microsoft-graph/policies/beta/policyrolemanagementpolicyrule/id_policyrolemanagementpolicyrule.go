package policyrolemanagementpolicyrule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyRoleManagementPolicyRuleId{}

// PolicyRoleManagementPolicyRuleId is a struct representing the Resource ID for a Policy Role Management Policy Rule
type PolicyRoleManagementPolicyRuleId struct {
	UnifiedRoleManagementPolicyId     string
	UnifiedRoleManagementPolicyRuleId string
}

// NewPolicyRoleManagementPolicyRuleID returns a new PolicyRoleManagementPolicyRuleId struct
func NewPolicyRoleManagementPolicyRuleID(unifiedRoleManagementPolicyId string, unifiedRoleManagementPolicyRuleId string) PolicyRoleManagementPolicyRuleId {
	return PolicyRoleManagementPolicyRuleId{
		UnifiedRoleManagementPolicyId:     unifiedRoleManagementPolicyId,
		UnifiedRoleManagementPolicyRuleId: unifiedRoleManagementPolicyRuleId,
	}
}

// ParsePolicyRoleManagementPolicyRuleID parses 'input' into a PolicyRoleManagementPolicyRuleId
func ParsePolicyRoleManagementPolicyRuleID(input string) (*PolicyRoleManagementPolicyRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyRoleManagementPolicyRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyRoleManagementPolicyRuleId{}

	if id.UnifiedRoleManagementPolicyId, ok = parsed.Parsed["unifiedRoleManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyId", *parsed)
	}

	if id.UnifiedRoleManagementPolicyRuleId, ok = parsed.Parsed["unifiedRoleManagementPolicyRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyRuleId", *parsed)
	}

	return &id, nil
}

// ParsePolicyRoleManagementPolicyRuleIDInsensitively parses 'input' case-insensitively into a PolicyRoleManagementPolicyRuleId
// note: this method should only be used for API response data and not user input
func ParsePolicyRoleManagementPolicyRuleIDInsensitively(input string) (*PolicyRoleManagementPolicyRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyRoleManagementPolicyRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyRoleManagementPolicyRuleId{}

	if id.UnifiedRoleManagementPolicyId, ok = parsed.Parsed["unifiedRoleManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyId", *parsed)
	}

	if id.UnifiedRoleManagementPolicyRuleId, ok = parsed.Parsed["unifiedRoleManagementPolicyRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyRuleId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyRoleManagementPolicyRuleID checks that 'input' can be parsed as a Policy Role Management Policy Rule ID
func ValidatePolicyRoleManagementPolicyRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyRoleManagementPolicyRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Role Management Policy Rule ID
func (id PolicyRoleManagementPolicyRuleId) ID() string {
	fmtString := "/policies/roleManagementPolicies/%s/rules/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementPolicyId, id.UnifiedRoleManagementPolicyRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Role Management Policy Rule ID
func (id PolicyRoleManagementPolicyRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticRoleManagementPolicies", "roleManagementPolicies", "roleManagementPolicies"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyId", "unifiedRoleManagementPolicyIdValue"),
		resourceids.StaticSegment("staticRules", "rules", "rules"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyRuleId", "unifiedRoleManagementPolicyRuleIdValue"),
	}
}

// String returns a human-readable description of this Policy Role Management Policy Rule ID
func (id PolicyRoleManagementPolicyRuleId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Policy: %q", id.UnifiedRoleManagementPolicyId),
		fmt.Sprintf("Unified Role Management Policy Rule: %q", id.UnifiedRoleManagementPolicyRuleId),
	}
	return fmt.Sprintf("Policy Role Management Policy Rule (%s)", strings.Join(components, "\n"))
}
