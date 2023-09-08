package policytokenlifetimepolicyappliesto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyTokenLifetimePolicyAppliesToId{}

// PolicyTokenLifetimePolicyAppliesToId is a struct representing the Resource ID for a Policy Token Lifetime Policy Applies To
type PolicyTokenLifetimePolicyAppliesToId struct {
	TokenLifetimePolicyId string
	DirectoryObjectId     string
}

// NewPolicyTokenLifetimePolicyAppliesToID returns a new PolicyTokenLifetimePolicyAppliesToId struct
func NewPolicyTokenLifetimePolicyAppliesToID(tokenLifetimePolicyId string, directoryObjectId string) PolicyTokenLifetimePolicyAppliesToId {
	return PolicyTokenLifetimePolicyAppliesToId{
		TokenLifetimePolicyId: tokenLifetimePolicyId,
		DirectoryObjectId:     directoryObjectId,
	}
}

// ParsePolicyTokenLifetimePolicyAppliesToID parses 'input' into a PolicyTokenLifetimePolicyAppliesToId
func ParsePolicyTokenLifetimePolicyAppliesToID(input string) (*PolicyTokenLifetimePolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyTokenLifetimePolicyAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyTokenLifetimePolicyAppliesToId{}

	if id.TokenLifetimePolicyId, ok = parsed.Parsed["tokenLifetimePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParsePolicyTokenLifetimePolicyAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyTokenLifetimePolicyAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyTokenLifetimePolicyAppliesToIDInsensitively(input string) (*PolicyTokenLifetimePolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyTokenLifetimePolicyAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyTokenLifetimePolicyAppliesToId{}

	if id.TokenLifetimePolicyId, ok = parsed.Parsed["tokenLifetimePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenLifetimePolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyTokenLifetimePolicyAppliesToID checks that 'input' can be parsed as a Policy Token Lifetime Policy Applies To ID
func ValidatePolicyTokenLifetimePolicyAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyTokenLifetimePolicyAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Token Lifetime Policy Applies To ID
func (id PolicyTokenLifetimePolicyAppliesToId) ID() string {
	fmtString := "/policies/tokenLifetimePolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.TokenLifetimePolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Token Lifetime Policy Applies To ID
func (id PolicyTokenLifetimePolicyAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticTokenLifetimePolicies", "tokenLifetimePolicies", "tokenLifetimePolicies"),
		resourceids.UserSpecifiedSegment("tokenLifetimePolicyId", "tokenLifetimePolicyIdValue"),
		resourceids.StaticSegment("staticAppliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Policy Token Lifetime Policy Applies To ID
func (id PolicyTokenLifetimePolicyAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Token Lifetime Policy: %q", id.TokenLifetimePolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Token Lifetime Policy Applies To (%s)", strings.Join(components, "\n"))
}
