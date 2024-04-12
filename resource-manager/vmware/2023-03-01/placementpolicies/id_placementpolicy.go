package placementpolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PlacementPolicyId{})
}

var _ resourceids.ResourceId = &PlacementPolicyId{}

// PlacementPolicyId is a struct representing the Resource ID for a Placement Policy
type PlacementPolicyId struct {
	SubscriptionId      string
	ResourceGroupName   string
	PrivateCloudName    string
	ClusterName         string
	PlacementPolicyName string
}

// NewPlacementPolicyID returns a new PlacementPolicyId struct
func NewPlacementPolicyID(subscriptionId string, resourceGroupName string, privateCloudName string, clusterName string, placementPolicyName string) PlacementPolicyId {
	return PlacementPolicyId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		PrivateCloudName:    privateCloudName,
		ClusterName:         clusterName,
		PlacementPolicyName: placementPolicyName,
	}
}

// ParsePlacementPolicyID parses 'input' into a PlacementPolicyId
func ParsePlacementPolicyID(input string) (*PlacementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PlacementPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PlacementPolicyId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePlacementPolicyIDInsensitively parses 'input' case-insensitively into a PlacementPolicyId
// note: this method should only be used for API response data and not user input
func ParsePlacementPolicyIDInsensitively(input string) (*PlacementPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PlacementPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PlacementPolicyId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PlacementPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.PrivateCloudName, ok = input.Parsed["privateCloudName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateCloudName", input)
	}

	if id.ClusterName, ok = input.Parsed["clusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterName", input)
	}

	if id.PlacementPolicyName, ok = input.Parsed["placementPolicyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "placementPolicyName", input)
	}

	return nil
}

// ValidatePlacementPolicyID checks that 'input' can be parsed as a Placement Policy ID
func ValidatePlacementPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePlacementPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Placement Policy ID
func (id PlacementPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/clusters/%s/placementPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.ClusterName, id.PlacementPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Placement Policy ID
func (id PlacementPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticPlacementPolicies", "placementPolicies", "placementPolicies"),
		resourceids.UserSpecifiedSegment("placementPolicyName", "placementPolicyValue"),
	}
}

// String returns a human-readable description of this Placement Policy ID
func (id PlacementPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Placement Policy Name: %q", id.PlacementPolicyName),
	}
	return fmt.Sprintf("Placement Policy (%s)", strings.Join(components, "\n"))
}
