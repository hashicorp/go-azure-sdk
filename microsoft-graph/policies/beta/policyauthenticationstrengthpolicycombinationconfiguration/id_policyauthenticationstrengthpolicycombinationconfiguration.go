package policyauthenticationstrengthpolicycombinationconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicyAuthenticationStrengthPolicyCombinationConfigurationId{}

// PolicyAuthenticationStrengthPolicyCombinationConfigurationId is a struct representing the Resource ID for a Policy Authentication Strength Policy Combination Configuration
type PolicyAuthenticationStrengthPolicyCombinationConfigurationId struct {
	AuthenticationStrengthPolicyId           string
	AuthenticationCombinationConfigurationId string
}

// NewPolicyAuthenticationStrengthPolicyCombinationConfigurationID returns a new PolicyAuthenticationStrengthPolicyCombinationConfigurationId struct
func NewPolicyAuthenticationStrengthPolicyCombinationConfigurationID(authenticationStrengthPolicyId string, authenticationCombinationConfigurationId string) PolicyAuthenticationStrengthPolicyCombinationConfigurationId {
	return PolicyAuthenticationStrengthPolicyCombinationConfigurationId{
		AuthenticationStrengthPolicyId:           authenticationStrengthPolicyId,
		AuthenticationCombinationConfigurationId: authenticationCombinationConfigurationId,
	}
}

// ParsePolicyAuthenticationStrengthPolicyCombinationConfigurationID parses 'input' into a PolicyAuthenticationStrengthPolicyCombinationConfigurationId
func ParsePolicyAuthenticationStrengthPolicyCombinationConfigurationID(input string) (*PolicyAuthenticationStrengthPolicyCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyAuthenticationStrengthPolicyCombinationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyAuthenticationStrengthPolicyCombinationConfigurationId{}

	if id.AuthenticationStrengthPolicyId, ok = parsed.Parsed["authenticationStrengthPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", *parsed)
	}

	if id.AuthenticationCombinationConfigurationId, ok = parsed.Parsed["authenticationCombinationConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationCombinationConfigurationId", *parsed)
	}

	return &id, nil
}

// ParsePolicyAuthenticationStrengthPolicyCombinationConfigurationIDInsensitively parses 'input' case-insensitively into a PolicyAuthenticationStrengthPolicyCombinationConfigurationId
// note: this method should only be used for API response data and not user input
func ParsePolicyAuthenticationStrengthPolicyCombinationConfigurationIDInsensitively(input string) (*PolicyAuthenticationStrengthPolicyCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicyAuthenticationStrengthPolicyCombinationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicyAuthenticationStrengthPolicyCombinationConfigurationId{}

	if id.AuthenticationStrengthPolicyId, ok = parsed.Parsed["authenticationStrengthPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", *parsed)
	}

	if id.AuthenticationCombinationConfigurationId, ok = parsed.Parsed["authenticationCombinationConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationCombinationConfigurationId", *parsed)
	}

	return &id, nil
}

// ValidatePolicyAuthenticationStrengthPolicyCombinationConfigurationID checks that 'input' can be parsed as a Policy Authentication Strength Policy Combination Configuration ID
func ValidatePolicyAuthenticationStrengthPolicyCombinationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAuthenticationStrengthPolicyCombinationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Authentication Strength Policy Combination Configuration ID
func (id PolicyAuthenticationStrengthPolicyCombinationConfigurationId) ID() string {
	fmtString := "/policies/authenticationStrengthPolicies/%s/combinationConfigurations/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationStrengthPolicyId, id.AuthenticationCombinationConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Authentication Strength Policy Combination Configuration ID
func (id PolicyAuthenticationStrengthPolicyCombinationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.StaticSegment("staticAuthenticationStrengthPolicies", "authenticationStrengthPolicies", "authenticationStrengthPolicies"),
		resourceids.UserSpecifiedSegment("authenticationStrengthPolicyId", "authenticationStrengthPolicyIdValue"),
		resourceids.StaticSegment("staticCombinationConfigurations", "combinationConfigurations", "combinationConfigurations"),
		resourceids.UserSpecifiedSegment("authenticationCombinationConfigurationId", "authenticationCombinationConfigurationIdValue"),
	}
}

// String returns a human-readable description of this Policy Authentication Strength Policy Combination Configuration ID
func (id PolicyAuthenticationStrengthPolicyCombinationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Strength Policy: %q", id.AuthenticationStrengthPolicyId),
		fmt.Sprintf("Authentication Combination Configuration: %q", id.AuthenticationCombinationConfigurationId),
	}
	return fmt.Sprintf("Policy Authentication Strength Policy Combination Configuration (%s)", strings.Join(components, "\n"))
}
