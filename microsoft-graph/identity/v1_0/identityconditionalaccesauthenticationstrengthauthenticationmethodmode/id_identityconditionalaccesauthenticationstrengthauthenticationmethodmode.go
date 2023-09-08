package identityconditionalaccesauthenticationstrengthauthenticationmethodmode

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId{}

// IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId is a struct representing the Resource ID for a Identity Conditional Acces Authentication Strength Authentication Method Mode
type IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId struct {
	AuthenticationMethodModeDetailId string
}

// NewIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeID returns a new IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId struct
func NewIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeID(authenticationMethodModeDetailId string) IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId {
	return IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId{
		AuthenticationMethodModeDetailId: authenticationMethodModeDetailId,
	}
}

// ParseIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeID parses 'input' into a IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId
func ParseIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeID(input string) (*IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId{}

	if id.AuthenticationMethodModeDetailId, ok = parsed.Parsed["authenticationMethodModeDetailId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationMethodModeDetailId", *parsed)
	}

	return &id, nil
}

// ParseIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeIDInsensitively parses 'input' case-insensitively into a IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId
// note: this method should only be used for API response data and not user input
func ParseIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeIDInsensitively(input string) (*IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId{}

	if id.AuthenticationMethodModeDetailId, ok = parsed.Parsed["authenticationMethodModeDetailId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authenticationMethodModeDetailId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeID checks that 'input' can be parsed as a Identity Conditional Acces Authentication Strength Authentication Method Mode ID
func ValidateIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Conditional Acces Authentication Strength Authentication Method Mode ID
func (id IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId) ID() string {
	fmtString := "/identity/conditionalAccess/authenticationStrength/authenticationMethodModes/%s"
	return fmt.Sprintf(fmtString, id.AuthenticationMethodModeDetailId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Conditional Acces Authentication Strength Authentication Method Mode ID
func (id IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticConditionalAccess", "conditionalAccess", "conditionalAccess"),
		resourceids.StaticSegment("staticAuthenticationStrength", "authenticationStrength", "authenticationStrength"),
		resourceids.StaticSegment("staticAuthenticationMethodModes", "authenticationMethodModes", "authenticationMethodModes"),
		resourceids.UserSpecifiedSegment("authenticationMethodModeDetailId", "authenticationMethodModeDetailIdValue"),
	}
}

// String returns a human-readable description of this Identity Conditional Acces Authentication Strength Authentication Method Mode ID
func (id IdentityConditionalAccesAuthenticationStrengthAuthenticationMethodModeId) String() string {
	components := []string{
		fmt.Sprintf("Authentication Method Mode Detail: %q", id.AuthenticationMethodModeDetailId),
	}
	return fmt.Sprintf("Identity Conditional Acces Authentication Strength Authentication Method Mode (%s)", strings.Join(components, "\n"))
}
