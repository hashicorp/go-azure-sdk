// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package marketplaceregistrationdefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = MarketplaceRegistrationDefinitionId{}

// MarketplaceRegistrationDefinitionId is a struct representing the Resource ID for a Marketplace Registration Definition
type MarketplaceRegistrationDefinitionId struct {
	MarketplaceIdentifier string
}

// NewMarketplaceRegistrationDefinitionID returns a new MarketplaceRegistrationDefinitionId struct
func NewMarketplaceRegistrationDefinitionID(marketplaceIdentifier string) MarketplaceRegistrationDefinitionId {
	return MarketplaceRegistrationDefinitionId{
		MarketplaceIdentifier: marketplaceIdentifier,
	}
}

// ParseMarketplaceRegistrationDefinitionID parses 'input' into a MarketplaceRegistrationDefinitionId
func ParseMarketplaceRegistrationDefinitionID(input string) (*MarketplaceRegistrationDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MarketplaceRegistrationDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MarketplaceRegistrationDefinitionId{}

	if id.MarketplaceIdentifier, ok = parsed.Parsed["marketplaceIdentifier"]; !ok {
		return nil, fmt.Errorf("the segment 'marketplaceIdentifier' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseMarketplaceRegistrationDefinitionIDInsensitively parses 'input' case-insensitively into a MarketplaceRegistrationDefinitionId
// note: this method should only be used for API response data and not user input
func ParseMarketplaceRegistrationDefinitionIDInsensitively(input string) (*MarketplaceRegistrationDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MarketplaceRegistrationDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MarketplaceRegistrationDefinitionId{}

	if id.MarketplaceIdentifier, ok = parsed.Parsed["marketplaceIdentifier"]; !ok {
		return nil, fmt.Errorf("the segment 'marketplaceIdentifier' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateMarketplaceRegistrationDefinitionID checks that 'input' can be parsed as a Marketplace Registration Definition ID
func ValidateMarketplaceRegistrationDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMarketplaceRegistrationDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Marketplace Registration Definition ID
func (id MarketplaceRegistrationDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.ManagedServices/marketplaceRegistrationDefinitions/%s"
	return fmt.Sprintf(fmtString, id.MarketplaceIdentifier)
}

// Segments returns a slice of Resource ID Segments which comprise this Marketplace Registration Definition ID
func (id MarketplaceRegistrationDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagedServices", "Microsoft.ManagedServices", "Microsoft.ManagedServices"),
		resourceids.StaticSegment("staticMarketplaceRegistrationDefinitions", "marketplaceRegistrationDefinitions", "marketplaceRegistrationDefinitions"),
		resourceids.UserSpecifiedSegment("marketplaceIdentifier", "marketplaceIdentifierValue"),
	}
}

// String returns a human-readable description of this Marketplace Registration Definition ID
func (id MarketplaceRegistrationDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Marketplace Identifier: %q", id.MarketplaceIdentifier),
	}
	return fmt.Sprintf("Marketplace Registration Definition (%s)", strings.Join(components, "\n"))
}
