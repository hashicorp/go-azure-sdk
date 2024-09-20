package virtualmachines

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&VirtualMachineId{})
}

var _ resourceids.ResourceId = &VirtualMachineId{}

// VirtualMachineId is a struct representing the Resource ID for a Virtual Machine
type VirtualMachineId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	ClusterName       string
	VirtualMachineId  string
}

// NewVirtualMachineID returns a new VirtualMachineId struct
func NewVirtualMachineID(subscriptionId string, resourceGroupName string, privateCloudName string, clusterName string, virtualMachineId string) VirtualMachineId {
	return VirtualMachineId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		ClusterName:       clusterName,
		VirtualMachineId:  virtualMachineId,
	}
}

// ParseVirtualMachineID parses 'input' into a VirtualMachineId
func ParseVirtualMachineID(input string) (*VirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseVirtualMachineIDInsensitively parses 'input' case-insensitively into a VirtualMachineId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineIDInsensitively(input string) (*VirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(&VirtualMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := VirtualMachineId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *VirtualMachineId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.VirtualMachineId, ok = input.Parsed["virtualMachineId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineId", input)
	}

	return nil
}

// ValidateVirtualMachineID checks that 'input' can be parsed as a Virtual Machine ID
func ValidateVirtualMachineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVirtualMachineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Virtual Machine ID
func (id VirtualMachineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/clusters/%s/virtualMachines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.ClusterName, id.VirtualMachineId)
}

// Segments returns a slice of Resource ID Segments which comprise this Virtual Machine ID
func (id VirtualMachineId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineId", "virtualMachineId"),
	}
}

// String returns a human-readable description of this Virtual Machine ID
func (id VirtualMachineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Virtual Machine: %q", id.VirtualMachineId),
	}
	return fmt.Sprintf("Virtual Machine (%s)", strings.Join(components, "\n"))
}
