package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BmcKeySetId{})
}

var _ resourceids.ResourceId = &BmcKeySetId{}

// BmcKeySetId is a struct representing the Resource ID for a Bmc Key Set
type BmcKeySetId struct {
	SubscriptionId    string
	ResourceGroupName string
	ClusterName       string
	BmcKeySetName     string
}

// NewBmcKeySetID returns a new BmcKeySetId struct
func NewBmcKeySetID(subscriptionId string, resourceGroupName string, clusterName string, bmcKeySetName string) BmcKeySetId {
	return BmcKeySetId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ClusterName:       clusterName,
		BmcKeySetName:     bmcKeySetName,
	}
}

// ParseBmcKeySetID parses 'input' into a BmcKeySetId
func ParseBmcKeySetID(input string) (*BmcKeySetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BmcKeySetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BmcKeySetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBmcKeySetIDInsensitively parses 'input' case-insensitively into a BmcKeySetId
// note: this method should only be used for API response data and not user input
func ParseBmcKeySetIDInsensitively(input string) (*BmcKeySetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BmcKeySetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BmcKeySetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BmcKeySetId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ClusterName, ok = input.Parsed["clusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterName", input)
	}

	if id.BmcKeySetName, ok = input.Parsed["bmcKeySetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bmcKeySetName", input)
	}

	return nil
}

// ValidateBmcKeySetID checks that 'input' can be parsed as a Bmc Key Set ID
func ValidateBmcKeySetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBmcKeySetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Bmc Key Set ID
func (id BmcKeySetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/clusters/%s/bmcKeySets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.BmcKeySetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Bmc Key Set ID
func (id BmcKeySetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterName"),
		resourceids.StaticSegment("staticBmcKeySets", "bmcKeySets", "bmcKeySets"),
		resourceids.UserSpecifiedSegment("bmcKeySetName", "bmcKeySetName"),
	}
}

// String returns a human-readable description of this Bmc Key Set ID
func (id BmcKeySetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Bmc Key Set Name: %q", id.BmcKeySetName),
	}
	return fmt.Sprintf("Bmc Key Set (%s)", strings.Join(components, "\n"))
}
