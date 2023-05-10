package hybrididentitymetadata

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = HybridIdentityMetadataId{}

// HybridIdentityMetadataId is a struct representing the Resource ID for a Hybrid Identity Metadata
type HybridIdentityMetadataId struct {
	SubscriptionId             string
	ResourceGroupName          string
	VirtualMachineName         string
	HybridIdentityMetadataName string
}

// NewHybridIdentityMetadataID returns a new HybridIdentityMetadataId struct
func NewHybridIdentityMetadataID(subscriptionId string, resourceGroupName string, virtualMachineName string, hybridIdentityMetadataName string) HybridIdentityMetadataId {
	return HybridIdentityMetadataId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		VirtualMachineName:         virtualMachineName,
		HybridIdentityMetadataName: hybridIdentityMetadataName,
	}
}

// ParseHybridIdentityMetadataID parses 'input' into a HybridIdentityMetadataId
func ParseHybridIdentityMetadataID(input string) (*HybridIdentityMetadataId, error) {
	parser := resourceids.NewParserFromResourceIdType(HybridIdentityMetadataId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HybridIdentityMetadataId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.HybridIdentityMetadataName, ok = parsed.Parsed["hybridIdentityMetadataName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hybridIdentityMetadataName", *parsed)
	}

	return &id, nil
}

// ParseHybridIdentityMetadataIDInsensitively parses 'input' case-insensitively into a HybridIdentityMetadataId
// note: this method should only be used for API response data and not user input
func ParseHybridIdentityMetadataIDInsensitively(input string) (*HybridIdentityMetadataId, error) {
	parser := resourceids.NewParserFromResourceIdType(HybridIdentityMetadataId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HybridIdentityMetadataId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.VirtualMachineName, ok = parsed.Parsed["virtualMachineName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "virtualMachineName", *parsed)
	}

	if id.HybridIdentityMetadataName, ok = parsed.Parsed["hybridIdentityMetadataName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hybridIdentityMetadataName", *parsed)
	}

	return &id, nil
}

// ValidateHybridIdentityMetadataID checks that 'input' can be parsed as a Hybrid Identity Metadata ID
func ValidateHybridIdentityMetadataID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHybridIdentityMetadataID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hybrid Identity Metadata ID
func (id HybridIdentityMetadataId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/virtualMachines/%s/hybridIdentityMetadata/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VirtualMachineName, id.HybridIdentityMetadataName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hybrid Identity Metadata ID
func (id HybridIdentityMetadataId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticVirtualMachines", "virtualMachines", "virtualMachines"),
		resourceids.UserSpecifiedSegment("virtualMachineName", "virtualMachineValue"),
		resourceids.StaticSegment("staticHybridIdentityMetadata", "hybridIdentityMetadata", "hybridIdentityMetadata"),
		resourceids.UserSpecifiedSegment("hybridIdentityMetadataName", "hybridIdentityMetadataValue"),
	}
}

// String returns a human-readable description of this Hybrid Identity Metadata ID
func (id HybridIdentityMetadataId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Virtual Machine Name: %q", id.VirtualMachineName),
		fmt.Sprintf("Hybrid Identity Metadata Name: %q", id.HybridIdentityMetadataName),
	}
	return fmt.Sprintf("Hybrid Identity Metadata (%s)", strings.Join(components, "\n"))
}
