package pricings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProviderPricingId{}

// ProviderPricingId is a struct representing the Resource ID for a Provider Pricing
type ProviderPricingId struct {
	SubscriptionId    string
	ResourceGroupName string
	PricingName       string
}

// NewProviderPricingID returns a new ProviderPricingId struct
func NewProviderPricingID(subscriptionId string, resourceGroupName string, pricingName string) ProviderPricingId {
	return ProviderPricingId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PricingName:       pricingName,
	}
}

// ParseProviderPricingID parses 'input' into a ProviderPricingId
func ParseProviderPricingID(input string) (*ProviderPricingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderPricingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderPricingId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviderPricingIDInsensitively parses 'input' case-insensitively into a ProviderPricingId
// note: this method should only be used for API response data and not user input
func ParseProviderPricingIDInsensitively(input string) (*ProviderPricingId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderPricingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProviderPricingId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProviderPricingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PricingName, ok = input.Parsed["pricingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "pricingName", input)
	}

	return nil
}

// ValidateProviderPricingID checks that 'input' can be parsed as a Provider Pricing ID
func ValidateProviderPricingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderPricingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Pricing ID
func (id ProviderPricingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/pricings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PricingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Pricing ID
func (id ProviderPricingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticPricings", "pricings", "pricings"),
		resourceids.UserSpecifiedSegment("pricingName", "pricingValue"),
	}
}

// String returns a human-readable description of this Provider Pricing ID
func (id ProviderPricingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Pricing Name: %q", id.PricingName),
	}
	return fmt.Sprintf("Provider Pricing (%s)", strings.Join(components, "\n"))
}
