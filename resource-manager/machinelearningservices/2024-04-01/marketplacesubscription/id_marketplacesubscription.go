package marketplacesubscription

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&MarketplaceSubscriptionId{})
}

var _ resourceids.ResourceId = &MarketplaceSubscriptionId{}

// MarketplaceSubscriptionId is a struct representing the Resource ID for a Marketplace Subscription
type MarketplaceSubscriptionId struct {
	SubscriptionId              string
	ResourceGroupName           string
	WorkspaceName               string
	MarketplaceSubscriptionName string
}

// NewMarketplaceSubscriptionID returns a new MarketplaceSubscriptionId struct
func NewMarketplaceSubscriptionID(subscriptionId string, resourceGroupName string, workspaceName string, marketplaceSubscriptionName string) MarketplaceSubscriptionId {
	return MarketplaceSubscriptionId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		WorkspaceName:               workspaceName,
		MarketplaceSubscriptionName: marketplaceSubscriptionName,
	}
}

// ParseMarketplaceSubscriptionID parses 'input' into a MarketplaceSubscriptionId
func ParseMarketplaceSubscriptionID(input string) (*MarketplaceSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MarketplaceSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MarketplaceSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMarketplaceSubscriptionIDInsensitively parses 'input' case-insensitively into a MarketplaceSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseMarketplaceSubscriptionIDInsensitively(input string) (*MarketplaceSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MarketplaceSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MarketplaceSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MarketplaceSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.MarketplaceSubscriptionName, ok = input.Parsed["marketplaceSubscriptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "marketplaceSubscriptionName", input)
	}

	return nil
}

// ValidateMarketplaceSubscriptionID checks that 'input' can be parsed as a Marketplace Subscription ID
func ValidateMarketplaceSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMarketplaceSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Marketplace Subscription ID
func (id MarketplaceSubscriptionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/marketplaceSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.MarketplaceSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Marketplace Subscription ID
func (id MarketplaceSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticMarketplaceSubscriptions", "marketplaceSubscriptions", "marketplaceSubscriptions"),
		resourceids.UserSpecifiedSegment("marketplaceSubscriptionName", "name"),
	}
}

// String returns a human-readable description of this Marketplace Subscription ID
func (id MarketplaceSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Marketplace Subscription Name: %q", id.MarketplaceSubscriptionName),
	}
	return fmt.Sprintf("Marketplace Subscription (%s)", strings.Join(components, "\n"))
}
