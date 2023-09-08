package identityconditionalaccesauthenticationstrengthpolicycombinationconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{}

// IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId is a struct representing the Resource ID for a Identity Conditional Acces Authentication Strength Policy Combination Configuration
type IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId struct {
	AuthenticationStrengthPolicyId           string
	AuthenticationCombinationConfigurationId string
}

// NewIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID returns a new IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId struct
func NewIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID(authenticationStrengthPolicyId string, authenticationCombinationConfigurationId string) IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId {
	return IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{
		AuthenticationStrengthPolicyId:           authenticationStrengthPolicyId,
		AuthenticationCombinationConfigurationId: authenticationCombinationConfigurationId,
	}
}

// ParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID parses 'input' into a IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId
func ParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID(input string) (*IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{}

	if id.AuthenticationStrengthPolicyId, ok = parsed.Parsed["authenticationStrengthPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", *parsed)
	}

	if id.AuthenticationCombinationConfigurationId, ok = parsed.Parsed["authenticationCombinationConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationCombinationConfigurationId", *parsed)
	}

	return &id, nil
}

// ParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationIDInsensitively(input string) (*IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId{}

	if id.AuthenticationStrengthPolicyId, ok = parsed.Parsed["authenticationStrengthPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", *parsed)
	}

	if id.AuthenticationCombinationConfigurationId, ok = parsed.Parsed["authenticationCombinationConfigurationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationCombinationConfigurationId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID checks that 'input' can be parsed as a Identity Conditional Acces Authentication Strength Policy Combination Configuration ID
func ValidateIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Acces Authentication Strength Policy Combination Configuration ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationStrength/policies/%s/combinationConfigurations/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationStrengthPolicyId, id.AuthenticationCombinationConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Acces Authentication Strength Policy Combination Configuration ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticConditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("staticAuthenticationStrength", "authenticationStrength", "authenticationStrength"),
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.UserSpecifiedSegment("authenticationStrengthPolicyId", "authenticationStrengthPolicyIdValue"),
		resourceids.StaticSegment("staticCombinationConfigurations", "combinationConfigurations", "combinationConfigurations"),
		resourceids.UserSpecifiedSegment("authenticationCombinationConfigurationId", "authenticationCombinationConfigurationIdValue"),
	}
}

// String returns a human-readable description of this Identity Conditional Acces Authentication Strength Policy Combination Configuration ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyCombinationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Strength Policy: %q", id.AuthenticationStrengthPolicyId),
		fmt.Sprintf("Authentication Combination Configuration: %q", id.AuthenticationCombinationConfigurationId),
	}
	return fmt.Sprintf("Identity Conditional Acces Authentication Strength Policy Combination Configuration (%s)", strings.Join(components, "\n"))
}
