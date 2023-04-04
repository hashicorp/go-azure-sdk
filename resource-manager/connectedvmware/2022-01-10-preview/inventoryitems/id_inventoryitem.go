package inventoryitems

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = InventoryItemId{}

// InventoryItemId is a struct representing the Resource ID for a Inventory Item
type InventoryItemId struct {
	SubscriptionId    string
	ResourceGroupName string
	VCenterName       string
	InventoryItemName string
}

// NewInventoryItemID returns a new InventoryItemId struct
func NewInventoryItemID(subscriptionId string, resourceGroupName string, vCenterName string, inventoryItemName string) InventoryItemId {
	return InventoryItemId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		VCenterName:       vCenterName,
		InventoryItemName: inventoryItemName,
	}
}

// ParseInventoryItemID parses 'input' into a InventoryItemId
func ParseInventoryItemID(input string) (*InventoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(InventoryItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := InventoryItemId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VCenterName, ok = parsed.Parsed["vCenterName"]; !ok {
		return nil, fmt.Errorf("the segment 'vCenterName' was not found in the resource id %q", input)
	}

	if id.InventoryItemName, ok = parsed.Parsed["inventoryItemName"]; !ok {
		return nil, fmt.Errorf("the segment 'inventoryItemName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseInventoryItemIDInsensitively parses 'input' case-insensitively into a InventoryItemId
// note: this method should only be used for API response data and not user input
func ParseInventoryItemIDInsensitively(input string) (*InventoryItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(InventoryItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := InventoryItemId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VCenterName, ok = parsed.Parsed["vCenterName"]; !ok {
		return nil, fmt.Errorf("the segment 'vCenterName' was not found in the resource id %q", input)
	}

	if id.InventoryItemName, ok = parsed.Parsed["inventoryItemName"]; !ok {
		return nil, fmt.Errorf("the segment 'inventoryItemName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateInventoryItemID checks that 'input' can be parsed as a Inventory Item ID
func ValidateInventoryItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInventoryItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Inventory Item ID
func (id InventoryItemId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ConnectedVMwarevSphere/vCenters/%s/inventoryItems/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VCenterName, id.InventoryItemName)
}

// Segments returns a slice of Resource ID Segments which comprise this Inventory Item ID
func (id InventoryItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere", "Microsoft.ConnectedVMwarevSphere"),
		resourceids.StaticSegment("staticVCenters", "vCenters", "vCenters"),
		resourceids.UserSpecifiedSegment("vCenterName", "vCenterValue"),
		resourceids.StaticSegment("staticInventoryItems", "inventoryItems", "inventoryItems"),
		resourceids.UserSpecifiedSegment("inventoryItemName", "inventoryItemValue"),
	}
}

// String returns a human-readable description of this Inventory Item ID
func (id InventoryItemId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("V Center Name: %q", id.VCenterName),
		fmt.Sprintf("Inventory Item Name: %q", id.InventoryItemName),
	}
	return fmt.Sprintf("Inventory Item (%s)", strings.Join(components, "\n"))
}
