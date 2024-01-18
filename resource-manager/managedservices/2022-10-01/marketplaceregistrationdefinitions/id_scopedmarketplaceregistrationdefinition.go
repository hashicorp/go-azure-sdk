package marketplaceregistrationdefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedMarketplaceRegistrationDefinitionId{}

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
	parser := resourceids.NewParserFromResourceIdType(&ScopedMarketplaceRegistrationDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedMarketplaceRegistrationDefinitionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedMarketplaceRegistrationDefinitionIDInsensitively parses 'input' case-insensitively into a ScopedMarketplaceRegistrationDefinitionId
// note: this method should only be used for API response data and not user input
func ParseScopedMarketplaceRegistrationDefinitionIDInsensitively(input string) (*ScopedMarketplaceRegistrationDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedMarketplaceRegistrationDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedMarketplaceRegistrationDefinitionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedMarketplaceRegistrationDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.MarketplaceIdentifier, ok = input.Parsed["marketplaceIdentifier"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "marketplaceIdentifier", input)
	}

	return nil
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
