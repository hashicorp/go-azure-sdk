package policytokenissuancepolicyappliesto

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyTokenIssuancePolicyAppliesToId{}

// PolicyTokenIssuancePolicyAppliesToId is a struct representing the Resource ID for a Policy Token Issuance Policy Applies To
type PolicyTokenIssuancePolicyAppliesToId struct {
	TokenIssuancePolicyId string
	DirectoryObjectId     string
}

// NewPolicyTokenIssuancePolicyAppliesToID returns a new PolicyTokenIssuancePolicyAppliesToId struct
func NewPolicyTokenIssuancePolicyAppliesToID(tokenIssuancePolicyId string, directoryObjectId string) PolicyTokenIssuancePolicyAppliesToId {
	return PolicyTokenIssuancePolicyAppliesToId{
		TokenIssuancePolicyId: tokenIssuancePolicyId,
		DirectoryObjectId:     directoryObjectId,
	}
}

// ParsePolicyTokenIssuancePolicyAppliesToID parses 'input' into a PolicyTokenIssuancePolicyAppliesToId
func ParsePolicyTokenIssuancePolicyAppliesToID(input string) (*PolicyTokenIssuancePolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyTokenIssuancePolicyAppliesToId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyTokenIssuancePolicyAppliesToId{}

	if id.TokenIssuancePolicyId, ok = parsed.Parsed["tokenIssuancePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ParsePolicyTokenIssuancePolicyAppliesToIDInsensitively parses 'input' case-insensitively into a PolicyTokenIssuancePolicyAppliesToId
// note: this method should only be used for API response data and not user input
func ParsePolicyTokenIssuancePolicyAppliesToIDInsensitively(input string) (*PolicyTokenIssuancePolicyAppliesToId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyTokenIssuancePolicyAppliesToId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyTokenIssuancePolicyAppliesToId{}

	if id.TokenIssuancePolicyId, ok = parsed.Parsed["tokenIssuancePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", *parsed)
	}

	if id.DirectoryObjectId, ok = parsed.Parsed["directoryObjectId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyTokenIssuancePolicyAppliesToID checks that 'input' can be parsed as a Policy Token Issuance Policy Applies To ID
func ValidatePolicyTokenIssuancePolicyAppliesToID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyTokenIssuancePolicyAppliesToID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Token Issuance Policy Applies To ID
func (id PolicyTokenIssuancePolicyAppliesToId) ID() string {
	fmtString := "/policies/tokenIssuancePolicies/%s/appliesTo/%s"
	return fmt.Sprintf(fmtString, id.TokenIssuancePolicyId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Token Issuance Policy Applies To ID
func (id PolicyTokenIssuancePolicyAppliesToId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticTokenIssuancePolicies", "tokenIssuancePolicies", "tokenIssuancePolicies"),
		resourceids.UserSpecifiedSegment("tokenIssuancePolicyId", "tokenIssuancePolicyIdValue"),
		resourceids.StaticSegment("staticAppliesTo", "appliesTo", "appliesTo"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectIdValue"),
	}
}

// String returns a human-readable description of this Policy Token Issuance Policy Applies To ID
func (id PolicyTokenIssuancePolicyAppliesToId) String() string {
	components := []string{
		fmt.Sprintf("Token Issuance Policy: %q", id.TokenIssuancePolicyId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Policy Token Issuance Policy Applies To (%s)", strings.Join(components, "\n"))
}
