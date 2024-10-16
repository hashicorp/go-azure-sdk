package discoveredassets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DiscoveredAssetId{})
}

var _ resourceids.ResourceId = &DiscoveredAssetId{}

// DiscoveredAssetId is a struct representing the Resource ID for a Discovered Asset
type DiscoveredAssetId struct {
	SubscriptionId      string
	ResourceGroupName   string
	DiscoveredAssetName string
}

// NewDiscoveredAssetID returns a new DiscoveredAssetId struct
func NewDiscoveredAssetID(subscriptionId string, resourceGroupName string, discoveredAssetName string) DiscoveredAssetId {
	return DiscoveredAssetId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		DiscoveredAssetName: discoveredAssetName,
	}
}

// ParseDiscoveredAssetID parses 'input' into a DiscoveredAssetId
func ParseDiscoveredAssetID(input string) (*DiscoveredAssetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiscoveredAssetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiscoveredAssetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDiscoveredAssetIDInsensitively parses 'input' case-insensitively into a DiscoveredAssetId
// note: this method should only be used for API response data and not user input
func ParseDiscoveredAssetIDInsensitively(input string) (*DiscoveredAssetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiscoveredAssetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiscoveredAssetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DiscoveredAssetId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DiscoveredAssetName, ok = input.Parsed["discoveredAssetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "discoveredAssetName", input)
	}

	return nil
}

// ValidateDiscoveredAssetID checks that 'input' can be parsed as a Discovered Asset ID
func ValidateDiscoveredAssetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiscoveredAssetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Discovered Asset ID
func (id DiscoveredAssetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DeviceRegistry/discoveredAssets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DiscoveredAssetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Discovered Asset ID
func (id DiscoveredAssetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDeviceRegistry", "Microsoft.DeviceRegistry", "Microsoft.DeviceRegistry"),
		resourceids.StaticSegment("staticDiscoveredAssets", "discoveredAssets", "discoveredAssets"),
		resourceids.UserSpecifiedSegment("discoveredAssetName", "discoveredAssetName"),
	}
}

// String returns a human-readable description of this Discovered Asset ID
func (id DiscoveredAssetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Discovered Asset Name: %q", id.DiscoveredAssetName),
	}
	return fmt.Sprintf("Discovered Asset (%s)", strings.Join(components, "\n"))
}
