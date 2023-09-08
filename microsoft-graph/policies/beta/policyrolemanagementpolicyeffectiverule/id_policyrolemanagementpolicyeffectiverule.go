package policyrolemanagementpolicyeffectiverule

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyRoleManagementPolicyEffectiveRuleId{}

// PolicyRoleManagementPolicyEffectiveRuleId is a struct representing the Resource ID for a Policy Role Management Policy Effective Rule
type PolicyRoleManagementPolicyEffectiveRuleId struct {
	UnifiedRoleManagementPolicyId     string
	UnifiedRoleManagementPolicyRuleId string
}

// NewPolicyRoleManagementPolicyEffectiveRuleID returns a new PolicyRoleManagementPolicyEffectiveRuleId struct
func NewPolicyRoleManagementPolicyEffectiveRuleID(unifiedRoleManagementPolicyId string, unifiedRoleManagementPolicyRuleId string) PolicyRoleManagementPolicyEffectiveRuleId {
	return PolicyRoleManagementPolicyEffectiveRuleId{
		UnifiedRoleManagementPolicyId:     unifiedRoleManagementPolicyId,
		UnifiedRoleManagementPolicyRuleId: unifiedRoleManagementPolicyRuleId,
	}
}

// ParsePolicyRoleManagementPolicyEffectiveRuleID parses 'input' into a PolicyRoleManagementPolicyEffectiveRuleId
func ParsePolicyRoleManagementPolicyEffectiveRuleID(input string) (*PolicyRoleManagementPolicyEffectiveRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyRoleManagementPolicyEffectiveRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyRoleManagementPolicyEffectiveRuleId{}

	if id.UnifiedRoleManagementPolicyId, ok = parsed.Parsed["unifiedRoleManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyId", *parsed)
	}

	if id.UnifiedRoleManagementPolicyRuleId, ok = parsed.Parsed["unifiedRoleManagementPolicyRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyRuleId", *parsed)
	}

	return &id, nil
}

// ParsePolicyRoleManagementPolicyEffectiveRuleIDInsensitively parses 'input' case-insensitively into a PolicyRoleManagementPolicyEffectiveRuleId
// note: this method should only be used for API response data and not user input
func ParsePolicyRoleManagementPolicyEffectiveRuleIDInsensitively(input string) (*PolicyRoleManagementPolicyEffectiveRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyRoleManagementPolicyEffectiveRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyRoleManagementPolicyEffectiveRuleId{}

	if id.UnifiedRoleManagementPolicyId, ok = parsed.Parsed["unifiedRoleManagementPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyId", *parsed)
	}

	if id.UnifiedRoleManagementPolicyRuleId, ok = parsed.Parsed["unifiedRoleManagementPolicyRuleId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementPolicyRuleId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyRoleManagementPolicyEffectiveRuleID checks that 'input' can be parsed as a Policy Role Management Policy Effective Rule ID
func ValidatePolicyRoleManagementPolicyEffectiveRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyRoleManagementPolicyEffectiveRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Role Management Policy Effective Rule ID
func (id PolicyRoleManagementPolicyEffectiveRuleId) ID() string {
	fmtString := "/policies/roleManagementPolicies/%s/effectiveRules/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementPolicyId, id.UnifiedRoleManagementPolicyRuleId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Role Management Policy Effective Rule ID
func (id PolicyRoleManagementPolicyEffectiveRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticRoleManagementPolicies", "roleManagementPolicies", "roleManagementPolicies"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyId", "unifiedRoleManagementPolicyIdValue"),
		resourceids.StaticSegment("staticEffectiveRules", "effectiveRules", "effectiveRules"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementPolicyRuleId", "unifiedRoleManagementPolicyRuleIdValue"),
	}
}

// String returns a human-readable description of this Policy Role Management Policy Effective Rule ID
func (id PolicyRoleManagementPolicyEffectiveRuleId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Policy: %q", id.UnifiedRoleManagementPolicyId),
		fmt.Sprintf("Unified Role Management Policy Rule: %q", id.UnifiedRoleManagementPolicyRuleId),
	}
	return fmt.Sprintf("Policy Role Management Policy Effective Rule (%s)", strings.Join(components, "\n"))
}
