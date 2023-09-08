package identityconditionalaccesauthenticationstrengthpolicy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesAuthenticationStrengthPolicyId{}

// IdentityConditionalAccesAuthenticationStrengthPolicyId is a struct representing the Resource ID for a Identity Conditional Acces Authentication Strength Policy
type IdentityConditionalAccesAuthenticationStrengthPolicyId struct {
	AuthenticationStrengthPolicyId string
}

// NewIdentityConditionalAccesAuthenticationStrengthPolicyID returns a new IdentityConditionalAccesAuthenticationStrengthPolicyId struct
func NewIdentityConditionalAccesAuthenticationStrengthPolicyID(authenticationStrengthPolicyId string) IdentityConditionalAccesAuthenticationStrengthPolicyId {
	return IdentityConditionalAccesAuthenticationStrengthPolicyId{
		AuthenticationStrengthPolicyId: authenticationStrengthPolicyId,
	}
}

// ParseIdentityConditionalAccesAuthenticationStrengthPolicyID parses 'input' into a IdentityConditionalAccesAuthenticationStrengthPolicyId
func ParseIdentityConditionalAccesAuthenticationStrengthPolicyID(input string) (*IdentityConditionalAccesAuthenticationStrengthPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityConditionalAccesAuthenticationStrengthPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityConditionalAccesAuthenticationStrengthPolicyId{}

	if id.AuthenticationStrengthPolicyId, ok = parsed.Parsed["authenticationStrengthPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", *parsed)
	}

	return &id, nil
}

// ParseIdentityConditionalAccesAuthenticationStrengthPolicyIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccesAuthenticationStrengthPolicyId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccesAuthenticationStrengthPolicyIDInsensitively(input string) (*IdentityConditionalAccesAuthenticationStrengthPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityConditionalAccesAuthenticationStrengthPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityConditionalAccesAuthenticationStrengthPolicyId{}

	if id.AuthenticationStrengthPolicyId, ok = parsed.Parsed["authenticationStrengthPolicyId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityConditionalAccesAuthenticationStrengthPolicyID checks that 'input' can be parsed as a Identity Conditional Acces Authentication Strength Policy ID
func ValidateIdentityConditionalAccesAuthenticationStrengthPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccesAuthenticationStrengthPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Acces Authentication Strength Policy ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationStrength/policies/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationStrengthPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Acces Authentication Strength Policy ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticConditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("staticAuthenticationStrength", "authenticationStrength", "authenticationStrength"),
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.UserSpecifiedSegment("authenticationStrengthPolicyId", "authenticationStrengthPolicyIdValue"),
	}
}

// String returns a human-readable description of this Identity Conditional Acces Authentication Strength Policy ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Strength Policy: %q", id.AuthenticationStrengthPolicyId),
	}
	return fmt.Sprintf("Identity Conditional Acces Authentication Strength Policy (%s)", strings.Join(components, "\n"))
}
