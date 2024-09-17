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
	recaser.RegisterResourceId(&ClusterManagerId{})
}

var _ resourceids.ResourceId = &ClusterManagerId{}

// ClusterManagerId is a struct representing the Resource ID for a Cluster Manager
type ClusterManagerId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ClusterManagerName string
}

// NewClusterManagerID returns a new ClusterManagerId struct
func NewClusterManagerID(subscriptionId string, resourceGroupName string, clusterManagerName string) ClusterManagerId {
	return ClusterManagerId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ClusterManagerName: clusterManagerName,
	}
}

// ParseClusterManagerID parses 'input' into a ClusterManagerId
func ParseClusterManagerID(input string) (*ClusterManagerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ClusterManagerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ClusterManagerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseClusterManagerIDInsensitively parses 'input' case-insensitively into a ClusterManagerId
// note: this method should only be used for API response data and not user input
func ParseClusterManagerIDInsensitively(input string) (*ClusterManagerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ClusterManagerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ClusterManagerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ClusterManagerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ClusterManagerName, ok = input.Parsed["clusterManagerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterManagerName", input)
	}

	return nil
}

// ValidateClusterManagerID checks that 'input' can be parsed as a Cluster Manager ID
func ValidateClusterManagerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseClusterManagerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cluster Manager ID
func (id ClusterManagerId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/clusterManagers/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterManagerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cluster Manager ID
func (id ClusterManagerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticClusterManagers", "clusterManagers", "clusterManagers"),
		resourceids.UserSpecifiedSegment("clusterManagerName", "clusterManagerValue"),
	}
}

// String returns a human-readable description of this Cluster Manager ID
func (id ClusterManagerId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Manager Name: %q", id.ClusterManagerName),
	}
	return fmt.Sprintf("Cluster Manager (%s)", strings.Join(components, "\n"))
}
