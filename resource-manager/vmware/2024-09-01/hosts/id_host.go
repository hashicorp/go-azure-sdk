package hosts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&HostId{})
}

var _ resourceids.ResourceId = &HostId{}

// HostId is a struct representing the Resource ID for a Host
type HostId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	ClusterName       string
	HostId            string
}

// NewHostID returns a new HostId struct
func NewHostID(subscriptionId string, resourceGroupName string, privateCloudName string, clusterName string, hostId string) HostId {
	return HostId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		ClusterName:       clusterName,
		HostId:            hostId,
	}
}

// ParseHostID parses 'input' into a HostId
func ParseHostID(input string) (*HostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HostId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HostId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseHostIDInsensitively parses 'input' case-insensitively into a HostId
// note: this method should only be used for API response data and not user input
func ParseHostIDInsensitively(input string) (*HostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&HostId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := HostId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *HostId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.HostId, ok = input.Parsed["hostId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "hostId", input)
	}

	return nil
}

// ValidateHostID checks that 'input' can be parsed as a Host ID
func ValidateHostID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHostID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Host ID
func (id HostId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/clusters/%s/hosts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.ClusterName, id.HostId)
}

// Segments returns a slice of Resource ID Segments which comprise this Host ID
func (id HostId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudName"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterName"),
		resourceids.StaticSegment("staticHosts", "hosts", "hosts"),
		resourceids.UserSpecifiedSegment("hostId", "hostId"),
	}
}

// String returns a human-readable description of this Host ID
func (id HostId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Host: %q", id.HostId),
	}
	return fmt.Sprintf("Host (%s)", strings.Join(components, "\n"))
}
