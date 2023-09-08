package identityb2xuserflowidentityprovider

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IdentityB2xUserFlowIdentityProviderId{}

// IdentityB2xUserFlowIdentityProviderId is a struct representing the Resource ID for a Identity B 2x User Flow Identity Provider
type IdentityB2xUserFlowIdentityProviderId struct {
	B2xIdentityUserFlowId string
	IdentityProviderId    string
}

// NewIdentityB2xUserFlowIdentityProviderID returns a new IdentityB2xUserFlowIdentityProviderId struct
func NewIdentityB2xUserFlowIdentityProviderID(b2xIdentityUserFlowId string, identityProviderId string) IdentityB2xUserFlowIdentityProviderId {
	return IdentityB2xUserFlowIdentityProviderId{
		B2xIdentityUserFlowId: b2xIdentityUserFlowId,
		IdentityProviderId:    identityProviderId,
	}
}

// ParseIdentityB2xUserFlowIdentityProviderID parses 'input' into a IdentityB2xUserFlowIdentityProviderId
func ParseIdentityB2xUserFlowIdentityProviderID(input string) (*IdentityB2xUserFlowIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowIdentityProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowIdentityProviderId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.IdentityProviderId, ok = parsed.Parsed["identityProviderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityProviderId", *parsed)
	}

	return &id, nil
}

// ParseIdentityB2xUserFlowIdentityProviderIDInsensitively parses 'input' case-insensitively into a IdentityB2xUserFlowIdentityProviderId
// note: this method should only be used for API response data and not user input
func ParseIdentityB2xUserFlowIdentityProviderIDInsensitively(input string) (*IdentityB2xUserFlowIdentityProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(IdentityB2xUserFlowIdentityProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IdentityB2xUserFlowIdentityProviderId{}

	if id.B2xIdentityUserFlowId, ok = parsed.Parsed["b2xIdentityUserFlowId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "b2xIdentityUserFlowId", *parsed)
	}

	if id.IdentityProviderId, ok = parsed.Parsed["identityProviderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "identityProviderId", *parsed)
	}

	return &id, nil
}

// ValidateIdentityB2xUserFlowIdentityProviderID checks that 'input' can be parsed as a Identity B 2x User Flow Identity Provider ID
func ValidateIdentityB2xUserFlowIdentityProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityB2xUserFlowIdentityProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity B 2x User Flow Identity Provider ID
func (id IdentityB2xUserFlowIdentityProviderId) ID() string {
	fmtString := "/identity/b2xUserFlows/%s/identityProviders/%s"
	return fmt.Sprintf(fmtString, id.B2xIdentityUserFlowId, id.IdentityProviderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity B 2x User Flow Identity Provider ID
func (id IdentityB2xUserFlowIdentityProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticIdentity", "identity", "identity"),
		resourceids.StaticSegment("staticB2xUserFlows", "b2xUserFlows", "b2xUserFlows"),
		resourceids.UserSpecifiedSegment("b2xIdentityUserFlowId", "b2xIdentityUserFlowIdValue"),
		resourceids.StaticSegment("staticIdentityProviders", "identityProviders", "identityProviders"),
		resourceids.UserSpecifiedSegment("identityProviderId", "identityProviderIdValue"),
	}
}

// String returns a human-readable description of this Identity B 2x User Flow Identity Provider ID
func (id IdentityB2xUserFlowIdentityProviderId) String() string {
	components := []string{
		fmt.Sprintf("B 2x Identity User Flow: %q", id.B2xIdentityUserFlowId),
		fmt.Sprintf("Identity Provider: %q", id.IdentityProviderId),
	}
	return fmt.Sprintf("Identity B 2x User Flow Identity Provider (%s)", strings.Join(components, "\n"))
}
