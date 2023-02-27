// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package globalreachconnections

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = GlobalReachConnectionId{}

// GlobalReachConnectionId is a struct representing the Resource ID for a Global Reach Connection
type GlobalReachConnectionId struct {
	SubscriptionId            string
	ResourceGroupName         string
	PrivateCloudName          string
	GlobalReachConnectionName string
}

// NewGlobalReachConnectionID returns a new GlobalReachConnectionId struct
func NewGlobalReachConnectionID(subscriptionId string, resourceGroupName string, privateCloudName string, globalReachConnectionName string) GlobalReachConnectionId {
	return GlobalReachConnectionId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		PrivateCloudName:          privateCloudName,
		GlobalReachConnectionName: globalReachConnectionName,
	}
}

// ParseGlobalReachConnectionID parses 'input' into a GlobalReachConnectionId
func ParseGlobalReachConnectionID(input string) (*GlobalReachConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GlobalReachConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GlobalReachConnectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.GlobalReachConnectionName, ok = parsed.Parsed["globalReachConnectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'globalReachConnectionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseGlobalReachConnectionIDInsensitively parses 'input' case-insensitively into a GlobalReachConnectionId
// note: this method should only be used for API response data and not user input
func ParseGlobalReachConnectionIDInsensitively(input string) (*GlobalReachConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(GlobalReachConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := GlobalReachConnectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.GlobalReachConnectionName, ok = parsed.Parsed["globalReachConnectionName"]; !ok {
		return nil, fmt.Errorf("the segment 'globalReachConnectionName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateGlobalReachConnectionID checks that 'input' can be parsed as a Global Reach Connection ID
func ValidateGlobalReachConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGlobalReachConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Global Reach Connection ID
func (id GlobalReachConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/globalReachConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.GlobalReachConnectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Global Reach Connection ID
func (id GlobalReachConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticGlobalReachConnections", "globalReachConnections", "globalReachConnections"),
		resourceids.UserSpecifiedSegment("globalReachConnectionName", "globalReachConnectionValue"),
	}
}

// String returns a human-readable description of this Global Reach Connection ID
func (id GlobalReachConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Global Reach Connection Name: %q", id.GlobalReachConnectionName),
	}
	return fmt.Sprintf("Global Reach Connection (%s)", strings.Join(components, "\n"))
}
