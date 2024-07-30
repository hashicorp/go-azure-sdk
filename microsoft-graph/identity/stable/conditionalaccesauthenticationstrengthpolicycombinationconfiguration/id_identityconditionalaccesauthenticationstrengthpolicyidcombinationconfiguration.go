package conditionalaccesauthenticationstrengthpolicycombinationconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId{}

// IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId is a struct representing the Resource ID for a Identity Conditional Acces Authentication Strength Policy Id Combination Configuration
type IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId struct {
	AuthenticationStrengthPolicyId           string
	AuthenticationCombinationConfigurationId string
}

// NewIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationID returns a new IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId struct
func NewIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationID(authenticationStrengthPolicyId string, authenticationCombinationConfigurationId string) IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId {
	return IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId{
		AuthenticationStrengthPolicyId:           authenticationStrengthPolicyId,
		AuthenticationCombinationConfigurationId: authenticationCombinationConfigurationId,
	}
}

// ParseIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationID parses 'input' into a IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId
func ParseIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationID(input string) (*IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationIDInsensitively(input string) (*IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AuthenticationStrengthPolicyId, ok = input.Parsed["authenticationStrengthPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationStrengthPolicyId", input)
	}

	if id.AuthenticationCombinationConfigurationId, ok = input.Parsed["authenticationCombinationConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "authenticationCombinationConfigurationId", input)
	}

	return nil
}

// ValidateIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationID checks that 'input' can be parsed as a Identity Conditional Acces Authentication Strength Policy Id Combination Configuration ID
func ValidateIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Acces Authentication Strength Policy Id Combination Configuration ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationStrength/policies/%s/combinationConfigurations/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationStrengthPolicyId, id.AuthenticationCombinationConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Acces Authentication Strength Policy Id Combination Configuration ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identity", "identity", "identity"),
		resourceids.StaticSegment("conditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("authenticationStrength", "authenticationStrength", "authenticationStrength"),
		resourceids.StaticSegment("policies", "policies", "policies"),
		resourceids.UserSpecifiedSegment("authenticationStrengthPolicyId", "authenticationStrengthPolicyIdValue"),
		resourceids.StaticSegment("combinationConfigurations", "combinationConfigurations", "combinationConfigurations"),
		resourceids.UserSpecifiedSegment("authenticationCombinationConfigurationId", "authenticationCombinationConfigurationIdValue"),
	}
}

// String returns a human-readable description of this Identity Conditional Acces Authentication Strength Policy Id Combination Configuration ID
func (id IdentityConditionalAccesAuthenticationStrengthPolicyIdCombinationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Strength Policy: %q", id.AuthenticationStrengthPolicyId),
		fmt.Sprintf("Authentication Combination Configuration: %q", id.AuthenticationCombinationConfigurationId),
	}
	return fmt.Sprintf("Identity Conditional Acces Authentication Strength Policy Id Combination Configuration (%s)", strings.Join(components, "\n"))
}
