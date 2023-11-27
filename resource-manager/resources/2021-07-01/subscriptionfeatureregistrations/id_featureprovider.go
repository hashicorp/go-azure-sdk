package subscriptionfeatureregistrations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FeatureProviderId{}

// FeatureProviderId is a struct representing the Resource ID for a Feature Provider
type FeatureProviderId struct {
	SubscriptionId      string
	FeatureProviderName string
}

// NewFeatureProviderID returns a new FeatureProviderId struct
func NewFeatureProviderID(subscriptionId string, featureProviderName string) FeatureProviderId {
	return FeatureProviderId{
		SubscriptionId:      subscriptionId,
		FeatureProviderName: featureProviderName,
	}
}

// ParseFeatureProviderID parses 'input' into a FeatureProviderId
func ParseFeatureProviderID(input string) (*FeatureProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(FeatureProviderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureProviderId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFeatureProviderIDInsensitively parses 'input' case-insensitively into a FeatureProviderId
// note: this method should only be used for API response data and not user input
func ParseFeatureProviderIDInsensitively(input string) (*FeatureProviderId, error) {
	parser := resourceids.NewParserFromResourceIdType(FeatureProviderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureProviderId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FeatureProviderId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.FeatureProviderName, ok = input.Parsed["featureProviderName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "featureProviderName", input)
	}

	return nil
}

// ValidateFeatureProviderID checks that 'input' can be parsed as a Feature Provider ID
func ValidateFeatureProviderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFeatureProviderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Feature Provider ID
func (id FeatureProviderId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Features/featureProviders/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.FeatureProviderName)
}

// Segments returns a slice of Resource ID Segments which comprise this Feature Provider ID
func (id FeatureProviderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftFeatures", "Microsoft.Features", "Microsoft.Features"),
		resourceids.StaticSegment("staticFeatureProviders", "featureProviders", "featureProviders"),
		resourceids.UserSpecifiedSegment("featureProviderName", "featureProviderValue"),
	}
}

// String returns a human-readable description of this Feature Provider ID
func (id FeatureProviderId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Feature Provider Name: %q", id.FeatureProviderName),
	}
	return fmt.Sprintf("Feature Provider (%s)", strings.Join(components, "\n"))
}
