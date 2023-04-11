package adaptivenetworkhardenings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AdaptiveNetworkHardeningId{}

// AdaptiveNetworkHardeningId is a struct representing the Resource ID for a Adaptive Network Hardening
type AdaptiveNetworkHardeningId struct {
	SubscriptionId               string
	ResourceGroupName            string
	ProviderName                 string
	ResourceType                 string
	ResourceName                 string
	AdaptiveNetworkHardeningName string
}

// NewAdaptiveNetworkHardeningID returns a new AdaptiveNetworkHardeningId struct
func NewAdaptiveNetworkHardeningID(subscriptionId string, resourceGroupName string, providerName string, resourceType string, resourceName string, adaptiveNetworkHardeningName string) AdaptiveNetworkHardeningId {
	return AdaptiveNetworkHardeningId{
		SubscriptionId:               subscriptionId,
		ResourceGroupName:            resourceGroupName,
		ProviderName:                 providerName,
		ResourceType:                 resourceType,
		ResourceName:                 resourceName,
		AdaptiveNetworkHardeningName: adaptiveNetworkHardeningName,
	}
}

// ParseAdaptiveNetworkHardeningID parses 'input' into a AdaptiveNetworkHardeningId
func ParseAdaptiveNetworkHardeningID(input string) (*AdaptiveNetworkHardeningId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdaptiveNetworkHardeningId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AdaptiveNetworkHardeningId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceType' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.AdaptiveNetworkHardeningName, ok = parsed.Parsed["adaptiveNetworkHardeningName"]; !ok {
		return nil, fmt.Errorf("the segment 'adaptiveNetworkHardeningName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseAdaptiveNetworkHardeningIDInsensitively parses 'input' case-insensitively into a AdaptiveNetworkHardeningId
// note: this method should only be used for API response data and not user input
func ParseAdaptiveNetworkHardeningIDInsensitively(input string) (*AdaptiveNetworkHardeningId, error) {
	parser := resourceids.NewParserFromResourceIdType(AdaptiveNetworkHardeningId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AdaptiveNetworkHardeningId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceType' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.AdaptiveNetworkHardeningName, ok = parsed.Parsed["adaptiveNetworkHardeningName"]; !ok {
		return nil, fmt.Errorf("the segment 'adaptiveNetworkHardeningName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateAdaptiveNetworkHardeningID checks that 'input' can be parsed as a Adaptive Network Hardening ID
func ValidateAdaptiveNetworkHardeningID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAdaptiveNetworkHardeningID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Adaptive Network Hardening ID
func (id AdaptiveNetworkHardeningId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.Security/adaptiveNetworkHardenings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ResourceType, id.ResourceName, id.AdaptiveNetworkHardeningName)
}

// Segments returns a slice of Resource ID Segments which comprise this Adaptive Network Hardening ID
func (id AdaptiveNetworkHardeningId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("resourceType", "resourceTypeValue"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticAdaptiveNetworkHardenings", "adaptiveNetworkHardenings", "adaptiveNetworkHardenings"),
		resourceids.UserSpecifiedSegment("adaptiveNetworkHardeningName", "adaptiveNetworkHardeningValue"),
	}
}

// String returns a human-readable description of this Adaptive Network Hardening ID
func (id AdaptiveNetworkHardeningId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Resource Type: %q", id.ResourceType),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Adaptive Network Hardening Name: %q", id.AdaptiveNetworkHardeningName),
	}
	return fmt.Sprintf("Adaptive Network Hardening (%s)", strings.Join(components, "\n"))
}
