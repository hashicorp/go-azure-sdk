package publicipprefixes

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PublicIPPrefixId{}

// PublicIPPrefixId is a struct representing the Resource ID for a Public I P Prefix
type PublicIPPrefixId struct {
	SubscriptionId     string
	ResourceGroupName  string
	PublicIPPrefixName string
}

// NewPublicIPPrefixID returns a new PublicIPPrefixId struct
func NewPublicIPPrefixID(subscriptionId string, resourceGroupName string, publicIPPrefixName string) PublicIPPrefixId {
	return PublicIPPrefixId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		PublicIPPrefixName: publicIPPrefixName,
	}
}

// ParsePublicIPPrefixID parses 'input' into a PublicIPPrefixId
func ParsePublicIPPrefixID(input string) (*PublicIPPrefixId, error) {
	parser := resourceids.NewParserFromResourceIdType(PublicIPPrefixId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PublicIPPrefixId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PublicIPPrefixName, ok = parsed.Parsed["publicIPPrefixName"]; !ok {
		return nil, fmt.Errorf("the segment 'publicIPPrefixName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePublicIPPrefixIDInsensitively parses 'input' case-insensitively into a PublicIPPrefixId
// note: this method should only be used for API response data and not user input
func ParsePublicIPPrefixIDInsensitively(input string) (*PublicIPPrefixId, error) {
	parser := resourceids.NewParserFromResourceIdType(PublicIPPrefixId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PublicIPPrefixId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.PublicIPPrefixName, ok = parsed.Parsed["publicIPPrefixName"]; !ok {
		return nil, fmt.Errorf("the segment 'publicIPPrefixName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePublicIPPrefixID checks that 'input' can be parsed as a Public I P Prefix ID
func ValidatePublicIPPrefixID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePublicIPPrefixID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Public I P Prefix ID
func (id PublicIPPrefixId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/publicIPPrefixes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PublicIPPrefixName)
}

// Segments returns a slice of Resource ID Segments which comprise this Public I P Prefix ID
func (id PublicIPPrefixId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetwork", "Microsoft.Network", "Microsoft.Network"),
		resourceids.StaticSegment("staticPublicIPPrefixes", "publicIPPrefixes", "publicIPPrefixes"),
		resourceids.UserSpecifiedSegment("publicIPPrefixName", "publicIPPrefixValue"),
	}
}

// String returns a human-readable description of this Public I P Prefix ID
func (id PublicIPPrefixId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Public I P Prefix Name: %q", id.PublicIPPrefixName),
	}
	return fmt.Sprintf("Public I P Prefix (%s)", strings.Join(components, "\n"))
}
