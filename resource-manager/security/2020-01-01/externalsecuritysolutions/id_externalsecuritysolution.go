package externalsecuritysolutions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ExternalSecuritySolutionId{}

// ExternalSecuritySolutionId is a struct representing the Resource ID for a External Security Solution
type ExternalSecuritySolutionId struct {
	SubscriptionId               string
	ResourceGroupName            string
	LocationName                 string
	ExternalSecuritySolutionName string
}

// NewExternalSecuritySolutionID returns a new ExternalSecuritySolutionId struct
func NewExternalSecuritySolutionID(subscriptionId string, resourceGroupName string, locationName string, externalSecuritySolutionName string) ExternalSecuritySolutionId {
	return ExternalSecuritySolutionId{
		SubscriptionId:               subscriptionId,
		ResourceGroupName:            resourceGroupName,
		LocationName:                 locationName,
		ExternalSecuritySolutionName: externalSecuritySolutionName,
	}
}

// ParseExternalSecuritySolutionID parses 'input' into a ExternalSecuritySolutionId
func ParseExternalSecuritySolutionID(input string) (*ExternalSecuritySolutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExternalSecuritySolutionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ExternalSecuritySolutionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.ExternalSecuritySolutionName, ok = parsed.Parsed["externalSecuritySolutionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "externalSecuritySolutionName", *parsed)
	}

	return &id, nil
}

// ParseExternalSecuritySolutionIDInsensitively parses 'input' case-insensitively into a ExternalSecuritySolutionId
// note: this method should only be used for API response data and not user input
func ParseExternalSecuritySolutionIDInsensitively(input string) (*ExternalSecuritySolutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(ExternalSecuritySolutionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ExternalSecuritySolutionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LocationName, ok = parsed.Parsed["locationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "locationName", *parsed)
	}

	if id.ExternalSecuritySolutionName, ok = parsed.Parsed["externalSecuritySolutionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "externalSecuritySolutionName", *parsed)
	}

	return &id, nil
}

// ValidateExternalSecuritySolutionID checks that 'input' can be parsed as a External Security Solution ID
func ValidateExternalSecuritySolutionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseExternalSecuritySolutionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted External Security Solution ID
func (id ExternalSecuritySolutionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/locations/%s/externalSecuritySolutions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LocationName, id.ExternalSecuritySolutionName)
}

// Segments returns a slice of Resource ID Segments which comprise this External Security Solution ID
func (id ExternalSecuritySolutionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticLocations", "locations", "locations"),
		resourceids.UserSpecifiedSegment("locationName", "locationValue"),
		resourceids.StaticSegment("staticExternalSecuritySolutions", "externalSecuritySolutions", "externalSecuritySolutions"),
		resourceids.UserSpecifiedSegment("externalSecuritySolutionName", "externalSecuritySolutionValue"),
	}
}

// String returns a human-readable description of this External Security Solution ID
func (id ExternalSecuritySolutionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Location Name: %q", id.LocationName),
		fmt.Sprintf("External Security Solution Name: %q", id.ExternalSecuritySolutionName),
	}
	return fmt.Sprintf("External Security Solution (%s)", strings.Join(components, "\n"))
}
