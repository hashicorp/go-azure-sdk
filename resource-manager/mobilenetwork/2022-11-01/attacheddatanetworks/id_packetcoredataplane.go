// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package attacheddatanetworks

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = PacketCoreDataPlaneId{}

// PacketCoreDataPlaneId is a struct representing the Resource ID for a Packet Core Data Plane
type PacketCoreDataPlaneId struct {
	SubscriptionId             string
	ResourceGroupName          string
	PacketCoreControlPlaneName string
	PacketCoreDataPlaneName    string
}

// NewPacketCoreDataPlaneID returns a new PacketCoreDataPlaneId struct
func NewPacketCoreDataPlaneID(subscriptionId string, resourceGroupName string, packetCoreControlPlaneName string, packetCoreDataPlaneName string) PacketCoreDataPlaneId {
	return PacketCoreDataPlaneId{
		SubscriptionId:             subscriptionId,
		ResourceGroupName:          resourceGroupName,
		PacketCoreControlPlaneName: packetCoreControlPlaneName,
		PacketCoreDataPlaneName:    packetCoreDataPlaneName,
	}
}

// ParsePacketCoreDataPlaneID parses 'input' into a PacketCoreDataPlaneId
func ParsePacketCoreDataPlaneID(input string) (*PacketCoreDataPlaneId, error) {
	parser := resourceids.NewParserFromResourceIdType(PacketCoreDataPlaneId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PacketCoreDataPlaneId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PacketCoreControlPlaneName, ok = parsed.Parsed["packetCoreControlPlaneName"]; !ok {
		return nil, fmt.Errorf("the segment 'packetCoreControlPlaneName' was not found in the resource id %q", input)
	}

	if id.PacketCoreDataPlaneName, ok = parsed.Parsed["packetCoreDataPlaneName"]; !ok {
		return nil, fmt.Errorf("the segment 'packetCoreDataPlaneName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePacketCoreDataPlaneIDInsensitively parses 'input' case-insensitively into a PacketCoreDataPlaneId
// note: this method should only be used for API response data and not user input
func ParsePacketCoreDataPlaneIDInsensitively(input string) (*PacketCoreDataPlaneId, error) {
	parser := resourceids.NewParserFromResourceIdType(PacketCoreDataPlaneId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PacketCoreDataPlaneId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PacketCoreControlPlaneName, ok = parsed.Parsed["packetCoreControlPlaneName"]; !ok {
		return nil, fmt.Errorf("the segment 'packetCoreControlPlaneName' was not found in the resource id %q", input)
	}

	if id.PacketCoreDataPlaneName, ok = parsed.Parsed["packetCoreDataPlaneName"]; !ok {
		return nil, fmt.Errorf("the segment 'packetCoreDataPlaneName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePacketCoreDataPlaneID checks that 'input' can be parsed as a Packet Core Data Plane ID
func ValidatePacketCoreDataPlaneID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePacketCoreDataPlaneID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Packet Core Data Plane ID
func (id PacketCoreDataPlaneId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MobileNetwork/packetCoreControlPlanes/%s/packetCoreDataPlanes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PacketCoreControlPlaneName, id.PacketCoreDataPlaneName)
}

// Segments returns a slice of Resource ID Segments which comprise this Packet Core Data Plane ID
func (id PacketCoreDataPlaneId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMobileNetwork", "Microsoft.MobileNetwork", "Microsoft.MobileNetwork"),
		resourceids.StaticSegment("staticPacketCoreControlPlanes", "packetCoreControlPlanes", "packetCoreControlPlanes"),
		resourceids.UserSpecifiedSegment("packetCoreControlPlaneName", "packetCoreControlPlaneValue"),
		resourceids.StaticSegment("staticPacketCoreDataPlanes", "packetCoreDataPlanes", "packetCoreDataPlanes"),
		resourceids.UserSpecifiedSegment("packetCoreDataPlaneName", "packetCoreDataPlaneValue"),
	}
}

// String returns a human-readable description of this Packet Core Data Plane ID
func (id PacketCoreDataPlaneId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Packet Core Control Plane Name: %q", id.PacketCoreControlPlaneName),
		fmt.Sprintf("Packet Core Data Plane Name: %q", id.PacketCoreDataPlaneName),
	}
	return fmt.Sprintf("Packet Core Data Plane (%s)", strings.Join(components, "\n"))
}
