// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package remediations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ProviderRemediationId{}

// ProviderRemediationId is a struct representing the Resource ID for a Provider Remediation
type ProviderRemediationId struct {
	SubscriptionId    string
	ResourceGroupName string
	RemediationName   string
}

// NewProviderRemediationID returns a new ProviderRemediationId struct
func NewProviderRemediationID(subscriptionId string, resourceGroupName string, remediationName string) ProviderRemediationId {
	return ProviderRemediationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		RemediationName:   remediationName,
	}
}

// ParseProviderRemediationID parses 'input' into a ProviderRemediationId
func ParseProviderRemediationID(input string) (*ProviderRemediationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderRemediationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderRemediationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.RemediationName, ok = parsed.Parsed["remediationName"]; !ok {
		return nil, fmt.Errorf("the segment 'remediationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviderRemediationIDInsensitively parses 'input' case-insensitively into a ProviderRemediationId
// note: this method should only be used for API response data and not user input
func ParseProviderRemediationIDInsensitively(input string) (*ProviderRemediationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderRemediationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderRemediationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.RemediationName, ok = parsed.Parsed["remediationName"]; !ok {
		return nil, fmt.Errorf("the segment 'remediationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateProviderRemediationID checks that 'input' can be parsed as a Provider Remediation ID
func ValidateProviderRemediationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderRemediationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Remediation ID
func (id ProviderRemediationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.PolicyInsights/remediations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.RemediationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Remediation ID
func (id ProviderRemediationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticRemediations", "remediations", "remediations"),
		resourceids.UserSpecifiedSegment("remediationName", "remediationValue"),
	}
}

// String returns a human-readable description of this Provider Remediation ID
func (id ProviderRemediationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Remediation Name: %q", id.RemediationName),
	}
	return fmt.Sprintf("Provider Remediation (%s)", strings.Join(components, "\n"))
}
