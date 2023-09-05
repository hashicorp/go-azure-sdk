package configurationprofilehciassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ConfigurationProfileAssignmentId{}

// ConfigurationProfileAssignmentId is a struct representing the Resource ID for a Configuration Profile Assignment
type ConfigurationProfileAssignmentId struct {
	SubscriptionId                     string
	ResourceGroupName                  string
	ClusterName                        string
	ConfigurationProfileAssignmentName string
}

// NewConfigurationProfileAssignmentID returns a new ConfigurationProfileAssignmentId struct
func NewConfigurationProfileAssignmentID(subscriptionId string, resourceGroupName string, clusterName string, configurationProfileAssignmentName string) ConfigurationProfileAssignmentId {
	return ConfigurationProfileAssignmentId{
		SubscriptionId:                     subscriptionId,
		ResourceGroupName:                  resourceGroupName,
		ClusterName:                        clusterName,
		ConfigurationProfileAssignmentName: configurationProfileAssignmentName,
	}
}

// ParseConfigurationProfileAssignmentID parses 'input' into a ConfigurationProfileAssignmentId
func ParseConfigurationProfileAssignmentID(input string) (*ConfigurationProfileAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigurationProfileAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigurationProfileAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterName", *parsed)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", *parsed)
	}

	return &id, nil
}

// ParseConfigurationProfileAssignmentIDInsensitively parses 'input' case-insensitively into a ConfigurationProfileAssignmentId
// note: this method should only be used for API response data and not user input
func ParseConfigurationProfileAssignmentIDInsensitively(input string) (*ConfigurationProfileAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigurationProfileAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigurationProfileAssignmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterName", *parsed)
	}

	if id.ConfigurationProfileAssignmentName, ok = parsed.Parsed["configurationProfileAssignmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "configurationProfileAssignmentName", *parsed)
	}

	return &id, nil
}

// ValidateConfigurationProfileAssignmentID checks that 'input' can be parsed as a Configuration Profile Assignment ID
func ValidateConfigurationProfileAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConfigurationProfileAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Configuration Profile Assignment ID
func (id ConfigurationProfileAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AzureStackHCI/clusters/%s/providers/Microsoft.AutoManage/configurationProfileAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.ConfigurationProfileAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Configuration Profile Assignment ID
func (id ConfigurationProfileAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAzureStackHCI", "Microsoft.AzureStackHCI", "Microsoft.AzureStackHCI"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAutoManage", "Microsoft.AutoManage", "Microsoft.AutoManage"),
		resourceids.StaticSegment("staticConfigurationProfileAssignments", "configurationProfileAssignments", "configurationProfileAssignments"),
		resourceids.UserSpecifiedSegment("configurationProfileAssignmentName", "configurationProfileAssignmentValue"),
	}
}

// String returns a human-readable description of this Configuration Profile Assignment ID
func (id ConfigurationProfileAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Configuration Profile Assignment Name: %q", id.ConfigurationProfileAssignmentName),
	}
	return fmt.Sprintf("Configuration Profile Assignment (%s)", strings.Join(components, "\n"))
}
