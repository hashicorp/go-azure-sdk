package policytokenissuancepolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyTokenIssuancePolicyId{}

// PolicyTokenIssuancePolicyId is a struct representing the Resource ID for a Policy Token Issuance Policy
type PolicyTokenIssuancePolicyId struct {
	TokenIssuancePolicyId string
}

// NewPolicyTokenIssuancePolicyID returns a new PolicyTokenIssuancePolicyId struct
func NewPolicyTokenIssuancePolicyID(tokenIssuancePolicyId string) PolicyTokenIssuancePolicyId {
	return PolicyTokenIssuancePolicyId{
		TokenIssuancePolicyId: tokenIssuancePolicyId,
	}
}

// ParsePolicyTokenIssuancePolicyID parses 'input' into a PolicyTokenIssuancePolicyId
func ParsePolicyTokenIssuancePolicyID(input string) (*PolicyTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyTokenIssuancePolicyId{}

	if id.TokenIssuancePolicyId, ok = parsed.Parsed["tokenIssuancePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", *parsed)
	}

	return &id, nil
}

// ParsePolicyTokenIssuancePolicyIDInsensitively parses 'input' case-insensitively into a PolicyTokenIssuancePolicyId
// note: this method should only be used for API response data and not user input
func ParsePolicyTokenIssuancePolicyIDInsensitively(input string) (*PolicyTokenIssuancePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyTokenIssuancePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyTokenIssuancePolicyId{}

	if id.TokenIssuancePolicyId, ok = parsed.Parsed["tokenIssuancePolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tokenIssuancePolicyId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyTokenIssuancePolicyID checks that 'input' can be parsed as a Policy Token Issuance Policy ID
func ValidatePolicyTokenIssuancePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyTokenIssuancePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Token Issuance Policy ID
func (id PolicyTokenIssuancePolicyId) ID() string {
	fmtString := "/policies/tokenIssuancePolicies/%s"
	return fmt.Sprintf(fmtString, id.TokenIssuancePolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Token Issuance Policy ID
func (id PolicyTokenIssuancePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticTokenIssuancePolicies", "tokenIssuancePolicies", "tokenIssuancePolicies"),
		resourceids.UserSpecifiedSegment("tokenIssuancePolicyId", "tokenIssuancePolicyIdValue"),
	}
}

// String returns a human-readable description of this Policy Token Issuance Policy ID
func (id PolicyTokenIssuancePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Token Issuance Policy: %q", id.TokenIssuancePolicyId),
	}
	return fmt.Sprintf("Policy Token Issuance Policy (%s)", strings.Join(components, "\n"))
}
