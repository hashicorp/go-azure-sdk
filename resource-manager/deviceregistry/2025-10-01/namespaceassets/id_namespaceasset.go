package namespaceassets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&NamespaceAssetId{})
}

var _ resourceids.ResourceId = &NamespaceAssetId{}

// NamespaceAssetId is a struct representing the Resource ID for a Namespace Asset
type NamespaceAssetId struct {
	SubscriptionId    string
	ResourceGroupName string
	NamespaceName     string
	AssetName         string
}

// NewNamespaceAssetID returns a new NamespaceAssetId struct
func NewNamespaceAssetID(subscriptionId string, resourceGroupName string, namespaceName string, assetName string) NamespaceAssetId {
	return NamespaceAssetId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NamespaceName:     namespaceName,
		AssetName:         assetName,
	}
}

// ParseNamespaceAssetID parses 'input' into a NamespaceAssetId
func ParseNamespaceAssetID(input string) (*NamespaceAssetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NamespaceAssetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NamespaceAssetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseNamespaceAssetIDInsensitively parses 'input' case-insensitively into a NamespaceAssetId
// note: this method should only be used for API response data and not user input
func ParseNamespaceAssetIDInsensitively(input string) (*NamespaceAssetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NamespaceAssetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NamespaceAssetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *NamespaceAssetId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NamespaceName, ok = input.Parsed["namespaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "namespaceName", input)
	}

	if id.AssetName, ok = input.Parsed["assetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "assetName", input)
	}

	return nil
}

// ValidateNamespaceAssetID checks that 'input' can be parsed as a Namespace Asset ID
func ValidateNamespaceAssetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNamespaceAssetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Namespace Asset ID
func (id NamespaceAssetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DeviceRegistry/namespaces/%s/assets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NamespaceName, id.AssetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Namespace Asset ID
func (id NamespaceAssetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDeviceRegistry", "Microsoft.DeviceRegistry", "Microsoft.DeviceRegistry"),
		resourceids.StaticSegment("staticNamespaces", "namespaces", "namespaces"),
		resourceids.UserSpecifiedSegment("namespaceName", "namespaceName"),
		resourceids.StaticSegment("staticAssets", "assets", "assets"),
		resourceids.UserSpecifiedSegment("assetName", "assetName"),
	}
}

// String returns a human-readable description of this Namespace Asset ID
func (id NamespaceAssetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Namespace Name: %q", id.NamespaceName),
		fmt.Sprintf("Asset Name: %q", id.AssetName),
	}
	return fmt.Sprintf("Namespace Asset (%s)", strings.Join(components, "\n"))
}
