package attestations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProviderAttestationId{}

// ProviderAttestationId is a struct representing the Resource ID for a Provider Attestation
type ProviderAttestationId struct {
	SubscriptionId    string
	ResourceGroupName string
	AttestationName   string
}

// NewProviderAttestationID returns a new ProviderAttestationId struct
func NewProviderAttestationID(subscriptionId string, resourceGroupName string, attestationName string) ProviderAttestationId {
	return ProviderAttestationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AttestationName:   attestationName,
	}
}

// ParseProviderAttestationID parses 'input' into a ProviderAttestationId
func ParseProviderAttestationID(input string) (*ProviderAttestationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderAttestationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderAttestationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AttestationName, ok = parsed.Parsed["attestationName"]; !ok {
		return nil, fmt.Errorf("the segment 'attestationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviderAttestationIDInsensitively parses 'input' case-insensitively into a ProviderAttestationId
// note: this method should only be used for API response data and not user input
func ParseProviderAttestationIDInsensitively(input string) (*ProviderAttestationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderAttestationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderAttestationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AttestationName, ok = parsed.Parsed["attestationName"]; !ok {
		return nil, fmt.Errorf("the segment 'attestationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateProviderAttestationID checks that 'input' can be parsed as a Provider Attestation ID
func ValidateProviderAttestationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderAttestationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Attestation ID
func (id ProviderAttestationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.PolicyInsights/attestations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AttestationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Attestation ID
func (id ProviderAttestationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticAttestations", "attestations", "attestations"),
		resourceids.UserSpecifiedSegment("attestationName", "attestationValue"),
	}
}

// String returns a human-readable description of this Provider Attestation ID
func (id ProviderAttestationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Attestation Name: %q", id.AttestationName),
	}
	return fmt.Sprintf("Provider Attestation (%s)", strings.Join(components, "\n"))
}
