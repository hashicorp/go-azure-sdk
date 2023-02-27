// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package marketplaceregistrationdefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ScopedMarketplaceRegistrationDefinitionId{}

// ScopedMarketplaceRegistrationDefinitionId is a struct representing the Resource ID for a Scoped Marketplace Registration Definition
type ScopedMarketplaceRegistrationDefinitionId struct {
	Scope                 string
	MarketplaceIdentifier string
}

// NewScopedMarketplaceRegistrationDefinitionID returns a new ScopedMarketplaceRegistrationDefinitionId struct
func NewScopedMarketplaceRegistrationDefinitionID(scope string, marketplaceIdentifier string) ScopedMarketplaceRegistrationDefinitionId {
	return ScopedMarketplaceRegistrationDefinitionId{
		Scope:                 scope,
		MarketplaceIdentifier: marketplaceIdentifier,
	}
}

// ParseScopedMarketplaceRegistrationDefinitionID parses 'input' into a ScopedMarketplaceRegistrationDefinitionId
func ParseScopedMarketplaceRegistrationDefinitionID(input string) (*ScopedMarketplaceRegistrationDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedMarketplaceRegistrationDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedMarketplaceRegistrationDefinitionId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, fmt.Errorf("the segment 'scope' was not found in the resource id %q", input)
	}

	if id.MarketplaceIdentifier, ok = parsed.Parsed["marketplaceIdentifier"]; !ok {
		return nil, fmt.Errorf("the segment 'marketplaceIdentifier' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseScopedMarketplaceRegistrationDefinitionIDInsensitively parses 'input' case-insensitively into a ScopedMarketplaceRegistrationDefinitionId
// note: this method should only be used for API response data and not user input
func ParseScopedMarketplaceRegistrationDefinitionIDInsensitively(input string) (*ScopedMarketplaceRegistrationDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedMarketplaceRegistrationDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedMarketplaceRegistrationDefinitionId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, fmt.Errorf("the segment 'scope' was not found in the resource id %q", input)
	}

	if id.MarketplaceIdentifier, ok = parsed.Parsed["marketplaceIdentifier"]; !ok {
		return nil, fmt.Errorf("the segment 'marketplaceIdentifier' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateScopedMarketplaceRegistrationDefinitionID checks that 'input' can be parsed as a Scoped Marketplace Registration Definition ID
func ValidateScopedMarketplaceRegistrationDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedMarketplaceRegistrationDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Marketplace Registration Definition ID
func (id ScopedMarketplaceRegistrationDefinitionId) ID() string {
	fmtString := "/%s/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.MarketplaceIdentifier)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Marketplace Registration Definition ID
func (id ScopedMarketplaceRegistrationDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagedServices", "Microsoft.ManagedServices", "Microsoft.ManagedServices"),
		resourceids.StaticSegment("staticMarketplaceRegistrationDefinitions", "marketplaceRegistrationDefinitions", "marketplaceRegistrationDefinitions"),
		resourceids.UserSpecifiedSegment("marketplaceIdentifier", "marketplaceIdentifierValue"),
	}
}

// String returns a human-readable description of this Scoped Marketplace Registration Definition ID
func (id ScopedMarketplaceRegistrationDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Marketplace Identifier: %q", id.MarketplaceIdentifier),
	}
	return fmt.Sprintf("Scoped Marketplace Registration Definition (%s)", strings.Join(components, "\n"))
}
