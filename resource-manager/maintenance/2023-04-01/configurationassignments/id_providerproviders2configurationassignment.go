package configurationassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProviderProviders2ConfigurationAssignmentId{}

// ProviderProviders2ConfigurationAssignmentId is a struct representing the Resource ID for a Provider Providers 2 Configuration Assignment
type ProviderProviders2ConfigurationAssignmentId struct {
	SubscriptionId              string
	ResourceGroupName           string
	ProviderName                string
	ResourceParentType          string
	ResourceParentName          string
	ResourceType                string
	ResourceName                string
	ConfigurationAssignmentName string
}

// NewProviderProviders2ConfigurationAssignmentID returns a new ProviderProviders2ConfigurationAssignmentId struct
func NewProviderProviders2ConfigurationAssignmentID(subscriptionId string, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string) ProviderProviders2ConfigurationAssignmentId {
	return ProviderProviders2ConfigurationAssignmentId{
		SubscriptionId:              subscriptionId,
		ResourceGroupName:           resourceGroupName,
		ProviderName:                providerName,
		ResourceParentType:          resourceParentType,
		ResourceParentName:          resourceParentName,
		ResourceType:                resourceType,
		ResourceName:                resourceName,
		ConfigurationAssignmentName: configurationAssignmentName,
	}
}

// ParseProviderProviders2ConfigurationAssignmentID parses 'input' into a ProviderProviders2ConfigurationAssignmentId
func ParseProviderProviders2ConfigurationAssignmentID(input string) (*ProviderProviders2ConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderProviders2ConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderProviders2ConfigurationAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ResourceParentType, ok = parsed.Parsed["resourceParentType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceParentType", *parsed)
	}

	if id.ResourceParentName, ok = parsed.Parsed["resourceParentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceParentName", *parsed)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceType", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	if id.ConfigurationAssignmentName, ok = parsed.Parsed["configurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationAssignmentName", *parsed)
	}

	return &id, nil
}

// ParseProviderProviders2ConfigurationAssignmentIDInsensitively parses 'input' case-insensitively into a ProviderProviders2ConfigurationAssignmentId
// note: this method should only be used for API response data and not user input
func ParseProviderProviders2ConfigurationAssignmentIDInsensitively(input string) (*ProviderProviders2ConfigurationAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderProviders2ConfigurationAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderProviders2ConfigurationAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ResourceParentType, ok = parsed.Parsed["resourceParentType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceParentType", *parsed)
	}

	if id.ResourceParentName, ok = parsed.Parsed["resourceParentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceParentName", *parsed)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceType", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	if id.ConfigurationAssignmentName, ok = parsed.Parsed["configurationAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationAssignmentName", *parsed)
	}

	return &id, nil
}

// ValidateProviderProviders2ConfigurationAssignmentID checks that 'input' can be parsed as a Provider Providers 2 Configuration Assignment ID
func ValidateProviderProviders2ConfigurationAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderProviders2ConfigurationAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Providers 2 Configuration Assignment ID
func (id ProviderProviders2ConfigurationAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/%s/%s/providers/Microsoft.Maintenance/configurationAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ResourceParentType, id.ResourceParentName, id.ResourceType, id.ResourceName, id.ConfigurationAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Providers 2 Configuration Assignment ID
func (id ProviderProviders2ConfigurationAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("resourceParentType", "resourceParentTypeValue"),
		resourceids.UserSpecifiedSegment("resourceParentName", "resourceParentValue"),
		resourceids.UserSpecifiedSegment("resourceType", "resourceTypeValue"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMaintenance", "Microsoft.Maintenance", "Microsoft.Maintenance"),
		resourceids.StaticSegment("staticConfigurationAssignments", "configurationAssignments", "configurationAssignments"),
		resourceids.UserSpecifiedSegment("configurationAssignmentName", "configurationAssignmentValue"),
	}
}

// String returns a human-readable description of this Provider Providers 2 Configuration Assignment ID
func (id ProviderProviders2ConfigurationAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Resource Parent Type: %q", id.ResourceParentType),
		fmt.Sprintf("Resource Parent Name: %q", id.ResourceParentName),
		fmt.Sprintf("Resource Type: %q", id.ResourceType),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Configuration Assignment Name: %q", id.ConfigurationAssignmentName),
	}
	return fmt.Sprintf("Provider Providers 2 Configuration Assignment (%s)", strings.Join(components, "\n"))
}
