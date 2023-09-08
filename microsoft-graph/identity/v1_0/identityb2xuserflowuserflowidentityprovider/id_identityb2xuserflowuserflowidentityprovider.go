package identityb2xuserflowuserflowidentityprovider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityB2xUserFlowUserFlowIdentityProviderId{}

// IdentityB2xUserFlowUserFlowIdentityProviderId is a struct representing the Resource ID for a Identity B 2x User Flow User Flow Identity Provider
type IdentityB2xUserFlowUserFlowIdentityProviderId struct {
	B2xIdentityUserFlowId  string
	IdentityProviderBaseId string
}

// NewIdentityB2xUserFlowUserFlowIdentityProviderID returns a new IdentityB2xUserFlowUserFlowIdentityProviderId struct
func NewIdentityB2xUserFlowUserFlowIdentityProviderID(b2xIdentityUserFlowId string, identityProviderBaseId string) IdentityB2xUserFlowUserFlowIdentityProviderId {
	return IdentityB2xUserFlowUserFlowIdentityProviderId{
		B2xIdentityUserFlowId:  b2xIdentityUserFlowId,
		IdentityProviderBaseId: identityProviderBaseId,
	}
}

// ParseIdentityB2xUserFlowUserFlowIdentityProviderID parses 'input' into a IdentityB2xUserFlowUserFlowIdentityProviderId
func ParseIdentityB2xUserFlowUserFlowIdentityProviderID(input string) (*IdentityB2xUserFlowUserFlowIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowUserFlowIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowUserFlowIdentityProviderId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.IdentityProviderBaseId, ok = parsed.Parsed["identityProviderBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", *parsed)
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowUserFlowIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowUserFlowIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowUserFlowIdentityProviderIDInsensitively(input string) (*IdentityB2xUserFlowUserFlowIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowUserFlowIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowUserFlowIdentityProviderId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.IdentityProviderBaseId, ok = parsed.Parsed["identityProviderBaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityProviderBaseId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityB2xUserFlowUserFlowIdentityProviderID checks that 'input' can be parsed as a Identity B 2x User Flow User Flow Identity Provider ID
func ValidateIdentityB2xUserFlowUserFlowIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowUserFlowIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2x User Flow User Flow Identity Provider ID
func (id IdentityB2xUserFlowUserFlowIdentityProviderId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/userFlowIdentityProviders/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.IdentityProviderBaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2x User Flow User Flow Identity Provider ID
func (id IdentityB2xUserFlowUserFlowIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticB2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowIdValue"),
		resourceids.StaticSegment("staticUserFlowIdentityProviders", "userFlowIdentityProviders", "userFlowIdentityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderBaseId", "identityProviderBaseIdValue"),
	}
}

// String returns a human-readable description of this Identity B 2x User Flow User Flow Identity Provider ID
func (id IdentityB2xUserFlowUserFlowIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("B 2x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("Identity Provider Base: %q", id.IdentityProviderBaseId),
	}
	return fmt.Sprintf("Identity B 2x User Flow User Flow Identity Provider (%s)", strings.Join(components, "\n"))
}
