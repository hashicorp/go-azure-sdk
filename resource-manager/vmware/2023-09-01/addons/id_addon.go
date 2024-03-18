package addons

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AddonId{}

// AddonId is a struct representing the Resource ID for a Addon
type AddonId struct {
	SubscriptionId    string
	ResourceGroupName string
	PrivateCloudName  string
	AddonName         string
}

// NewAddonID returns a new AddonId struct
func NewAddonID(subscriptionId string, resourceGroupName string, privateCloudName string, addonName string) AddonId {
	return AddonId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		PrivateCloudName:  privateCloudName,
		AddonName:         addonName,
	}
}

// ParseAddonID parses 'input' into a AddonId
func ParseAddonID(input string) (*AddonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AddonId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AddonId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAddonIDInsensitively parses 'input' case-insensitively into a AddonId
// note: this method should only be used for API response data and not user input
func ParseAddonIDInsensitively(input string) (*AddonId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AddonId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AddonId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AddonId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AddonName, ok = input.Parsed["addonName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "addonName", input)
	}

	return nil
}

// ValidateAddonID checks that 'input' can be parsed as a Addon ID
func ValidateAddonID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAddonID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Addon ID
func (id AddonId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AVS/privateClouds/%s/addons/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PrivateCloudName, id.AddonName)
}

// Segments returns a slice of Resource ID Segments which comprise this Addon ID
func (id AddonId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAVS", "Microsoft.AVS", "Microsoft.AVS"),
		resourceids.StaticSegment("staticPrivateClouds", "privateClouds", "privateClouds"),
		resourceids.UserSpecifiedSegment("privateCloudName", "privateCloudValue"),
		resourceids.StaticSegment("staticAddons", "addons", "addons"),
		resourceids.UserSpecifiedSegment("addonName", "addonValue"),
	}
}

// String returns a human-readable description of this Addon ID
func (id AddonId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Private Cloud Name: %q", id.PrivateCloudName),
		fmt.Sprintf("Addon Name: %q", id.AddonName),
	}
	return fmt.Sprintf("Addon (%s)", strings.Join(components, "\n"))
}
