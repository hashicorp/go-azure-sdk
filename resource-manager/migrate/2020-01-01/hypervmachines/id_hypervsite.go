package hypervmachines

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = HyperVSiteId{}

// HyperVSiteId is a struct representing the Resource ID for a Hyper V Site
type HyperVSiteId struct {
	SubscriptionId    string
	ResourceGroupName string
	HyperVSiteName    string
}

// NewHyperVSiteID returns a new HyperVSiteId struct
func NewHyperVSiteID(subscriptionId string, resourceGroupName string, hyperVSiteName string) HyperVSiteId {
	return HyperVSiteId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		HyperVSiteName:    hyperVSiteName,
	}
}

// ParseHyperVSiteID parses 'input' into a HyperVSiteId
func ParseHyperVSiteID(input string) (*HyperVSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(HyperVSiteId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HyperVSiteId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.HyperVSiteName, ok = parsed.Parsed["hyperVSiteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hyperVSiteName", *parsed)
	}

	return &id, nil
}

// ParseHyperVSiteIDInsensitively parses 'input' case-insensitively into a HyperVSiteId
// note: this method should only be used for API response data and not user input
func ParseHyperVSiteIDInsensitively(input string) (*HyperVSiteId, error) {
	parser := resourceids.NewParserFromResourceIdType(HyperVSiteId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HyperVSiteId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.HyperVSiteName, ok = parsed.Parsed["hyperVSiteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hyperVSiteName", *parsed)
	}

	return &id, nil
}

// ValidateHyperVSiteID checks that 'input' can be parsed as a Hyper V Site ID
func ValidateHyperVSiteID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHyperVSiteID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Hyper V Site ID
func (id HyperVSiteId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OffAzure/hyperVSites/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.HyperVSiteName)
}

// Segments returns a slice of Resource ID Segments which comprise this Hyper V Site ID
func (id HyperVSiteId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOffAzure", "Microsoft.OffAzure", "Microsoft.OffAzure"),
		resourceids.StaticSegment("staticHyperVSites", "hyperVSites", "hyperVSites"),
		resourceids.UserSpecifiedSegment("hyperVSiteName", "hyperVSiteValue"),
	}
}

// String returns a human-readable description of this Hyper V Site ID
func (id HyperVSiteId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Hyper V Site Name: %q", id.HyperVSiteName),
	}
	return fmt.Sprintf("Hyper V Site (%s)", strings.Join(components, "\n"))
}
