// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package placementpolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = VirtualMachineId{}

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
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineId, ok = parsed.Parsed["virtualMachineId"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineId' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseVirtualMachineIDInsensitively parses 'input' case-insensitively into a VirtualMachineId
// note: this method should only be used for API response data and not user input
func ParseVirtualMachineIDInsensitively(input string) (*VirtualMachineId, error) {
	parser := resourceids.NewParserFromResourceIdType(VirtualMachineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VirtualMachineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PrivateCloudName, ok = parsed.Parsed["privateCloudName"]; !ok {
		return nil, fmt.Errorf("the segment 'privateCloudName' was not found in the resource id %q", input)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, fmt.Errorf("the segment 'clusterName' was not found in the resource id %q", input)
	}

	if id.VirtualMachineId, ok = parsed.Parsed["virtualMachineId"]; !ok {
		return nil, fmt.Errorf("the segment 'virtualMachineId' was not found in the resource id %q", input)
	}

	return &id, nil
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
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineId", "virtualMachineIdValue"),
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
