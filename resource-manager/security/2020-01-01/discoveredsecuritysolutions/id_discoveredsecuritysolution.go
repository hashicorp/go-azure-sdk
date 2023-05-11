package discoveredsecuritysolutions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DiscoveredSecuritySolutionId{}

// DiscoveredSecuritySolutionId is a struct representing the Resource ID for a Discovered Security Solution
type DiscoveredSecuritySolutionId struct {
	SubscriptionId                 string
	ResourceGroupName              string
	LocationName                   string
	DiscoveredSecuritySolutionName string
}

// NewDiscoveredSecuritySolutionID returns a new DiscoveredSecuritySolutionId struct
func NewDiscoveredSecuritySolutionID(subscriptionId string, resourceGroupName string, locationName string, discoveredSecuritySolutionName string) DiscoveredSecuritySolutionId {
	return DiscoveredSecuritySolutionId{
		SubscriptionId:                 subscriptionId,
		ResourceGroupName:              resourceGroupName,
		LocationName:                   locationName,
		DiscoveredSecuritySolutionName: discoveredSecuritySolutionName,
	}
}

// ParseDiscoveredSecuritySolutionID parses 'input' into a DiscoveredSecuritySolutionId
func ParseDiscoveredSecuritySolutionID(input string) (*DiscoveredSecuritySolutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(DiscoveredSecuritySolutionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DiscoveredSecuritySolutionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.DiscoveredSecuritySolutionName, ok = parsed.Parsed["discoveredSecuritySolutionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "discoveredSecuritySolutionName", *parsed)
	}

	return &id, nil
}

// ParseDiscoveredSecuritySolutionIDInsensitively parses 'input' case-insensitively into a DiscoveredSecuritySolutionId
// note: this method should only be used for API response data and not user input
func ParseDiscoveredSecuritySolutionIDInsensitively(input string) (*DiscoveredSecuritySolutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(DiscoveredSecuritySolutionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DiscoveredSecuritySolutionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.DiscoveredSecuritySolutionName, ok = parsed.Parsed["discoveredSecuritySolutionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "discoveredSecuritySolutionName", *parsed)
	}

	return &id, nil
}

// ValidateDiscoveredSecuritySolutionID checks that 'input' can be parsed as a Discovered Security Solution ID
func ValidateDiscoveredSecuritySolutionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiscoveredSecuritySolutionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Discovered Security Solution ID
func (id DiscoveredSecuritySolutionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/locations/%s/discoveredSecuritySolutions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.DiscoveredSecuritySolutionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Discovered Security Solution ID
func (id DiscoveredSecuritySolutionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticDiscoveredSecuritySolutions", "discoveredSecuritySolutions", "discoveredSecuritySolutions"),
		resourceids.UserSpecifiedSegment("discoveredSecuritySolutionName", "discoveredSecuritySolutionValue"),
	}
}

// String returns a human-readable description of this Discovered Security Solution ID
func (id DiscoveredSecuritySolutionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("Discovered Security Solution Name: %q", id.DiscoveredSecuritySolutionName),
	}
	return fmt.Sprintf("Discovered Security Solution (%s)", strings.Join(components, "\n"))
}
