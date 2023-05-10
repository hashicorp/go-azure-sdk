package jitnetworkaccesspolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = JitNetworkAccessPolicyId{}

// JitNetworkAccessPolicyId is a struct representing the Resource ID for a Jit Network Access Policy
type JitNetworkAccessPolicyId struct {
	SubscriptionId             string
	ResourceGroupName          string
	LocationName               string
	JitNetworkAccessPolicyName string
}

// NewJitNetworkAccessPolicyID returns a new JitNetworkAccessPolicyId struct
func NewJitNetworkAccessPolicyID(subscriptionId string, resourceGroupName string, locationName string, jitNetworkAccessPolicyName string) JitNetworkAccessPolicyId {
	return JitNetworkAccessPolicyId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		LocationName:               locationName,
		JitNetworkAccessPolicyName: jitNetworkAccessPolicyName,
	}
}

// ParseJitNetworkAccessPolicyID parses 'input' into a JitNetworkAccessPolicyId
func ParseJitNetworkAccessPolicyID(input string) (*JitNetworkAccessPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(JitNetworkAccessPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := JitNetworkAccessPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.JitNetworkAccessPolicyName, ok = parsed.Parsed["jitNetworkAccessPolicyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jitNetworkAccessPolicyName", *parsed)
	}

	return &id, nil
}

// ParseJitNetworkAccessPolicyIDInsensitively parses 'input' case-insensitively into a JitNetworkAccessPolicyId
// note: this method should only be used for API response data and not user input
func ParseJitNetworkAccessPolicyIDInsensitively(input string) (*JitNetworkAccessPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(JitNetworkAccessPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := JitNetworkAccessPolicyId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.JitNetworkAccessPolicyName, ok = parsed.Parsed["jitNetworkAccessPolicyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "jitNetworkAccessPolicyName", *parsed)
	}

	return &id, nil
}

// ValidateJitNetworkAccessPolicyID checks that 'input' can be parsed as a Jit Network Access Policy ID
func ValidateJitNetworkAccessPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseJitNetworkAccessPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Jit Network Access Policy ID
func (id JitNetworkAccessPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/locations/%s/jitNetworkAccessPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.JitNetworkAccessPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Jit Network Access Policy ID
func (id JitNetworkAccessPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticJitNetworkAccessPolicies", "jitNetworkAccessPolicies", "jitNetworkAccessPolicies"),
		resourceids.UserSpecifiedSegment("jitNetworkAccessPolicyName", "jitNetworkAccessPolicyValue"),
	}
}

// String returns a human-readable description of this Jit Network Access Policy ID
func (id JitNetworkAccessPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Jit Network Access Policy Name: %q", id.JitNetworkAccessPolicyName),
	}
	return fmt.Sprintf("Jit Network Access Policy (%s)", strings.Join(components, "\n"))
}
