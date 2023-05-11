package computepolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ComputePolicyId{}

// ComputePolicyId is a struct representing the Resource ID for a Compute Policy
type ComputePolicyId struct {
	SubscriptionId    string
	ResourceGroupName string
	AccountName       string
	ComputePolicyName string
}

// NewComputePolicyID returns a new ComputePolicyId struct
func NewComputePolicyID(subscriptionId string, resourceGroupName string, accountName string, computePolicyName string) ComputePolicyId {
	return ComputePolicyId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		AccountName:       accountName,
		ComputePolicyName: computePolicyName,
	}
}

// ParseComputePolicyID parses 'input' into a ComputePolicyId
func ParseComputePolicyID(input string) (*ComputePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ComputePolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ComputePolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accountName", *parsed)
	}

	if id.ComputePolicyName, ok = parsed.Parsed["computePolicyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "computePolicyName", *parsed)
	}

	return &id, nil
}

// ParseComputePolicyIDInsensitively parses 'input' case-insensitively into a ComputePolicyId
// note: this method should only be used for API response data and not user input
func ParseComputePolicyIDInsensitively(input string) (*ComputePolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(ComputePolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ComputePolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "accountName", *parsed)
	}

	if id.ComputePolicyName, ok = parsed.Parsed["computePolicyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "computePolicyName", *parsed)
	}

	return &id, nil
}

// ValidateComputePolicyID checks that 'input' can be parsed as a Compute Policy ID
func ValidateComputePolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseComputePolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Compute Policy ID
func (id ComputePolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataLakeAnalytics/accounts/%s/computePolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ComputePolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Compute Policy ID
func (id ComputePolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataLakeAnalytics", "Microsoft.DataLakeAnalytics", "Microsoft.DataLakeAnalytics"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticComputePolicies", "computePolicies", "computePolicies"),
		resourceids.UserSpecifiedSegment("computePolicyName", "computePolicyValue"),
	}
}

// String returns a human-readable description of this Compute Policy ID
func (id ComputePolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Compute Policy Name: %q", id.ComputePolicyName),
	}
	return fmt.Sprintf("Compute Policy (%s)", strings.Join(components, "\n"))
}
