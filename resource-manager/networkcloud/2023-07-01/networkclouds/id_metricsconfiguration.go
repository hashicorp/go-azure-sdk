package networkclouds

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&MetricsConfigurationId{})
}

var _ resourceids.ResourceId = &MetricsConfigurationId{}

// MetricsConfigurationId is a struct representing the Resource ID for a Metrics Configuration
type MetricsConfigurationId struct {
	SubscriptionId           string
	ResourceGroupName        string
	ClusterName              string
	MetricsConfigurationName string
}

// NewMetricsConfigurationID returns a new MetricsConfigurationId struct
func NewMetricsConfigurationID(subscriptionId string, resourceGroupName string, clusterName string, metricsConfigurationName string) MetricsConfigurationId {
	return MetricsConfigurationId{
		SubscriptionId:           subscriptionId,
		ResourceGroupName:        resourceGroupName,
		ClusterName:              clusterName,
		MetricsConfigurationName: metricsConfigurationName,
	}
}

// ParseMetricsConfigurationID parses 'input' into a MetricsConfigurationId
func ParseMetricsConfigurationID(input string) (*MetricsConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MetricsConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MetricsConfigurationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMetricsConfigurationIDInsensitively parses 'input' case-insensitively into a MetricsConfigurationId
// note: this method should only be used for API response data and not user input
func ParseMetricsConfigurationIDInsensitively(input string) (*MetricsConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MetricsConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MetricsConfigurationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MetricsConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ClusterName, ok = input.Parsed["clusterName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "clusterName", input)
	}

	if id.MetricsConfigurationName, ok = input.Parsed["metricsConfigurationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "metricsConfigurationName", input)
	}

	return nil
}

// ValidateMetricsConfigurationID checks that 'input' can be parsed as a Metrics Configuration ID
func ValidateMetricsConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMetricsConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Metrics Configuration ID
func (id MetricsConfigurationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetworkCloud/clusters/%s/metricsConfigurations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ClusterName, id.MetricsConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Metrics Configuration ID
func (id MetricsConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetworkCloud", "Microsoft.NetworkCloud", "Microsoft.NetworkCloud"),
		resourceids.StaticSegment("staticClusters", "clusters", "clusters"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticMetricsConfigurations", "metricsConfigurations", "metricsConfigurations"),
		resourceids.UserSpecifiedSegment("metricsConfigurationName", "metricsConfigurationValue"),
	}
}

// String returns a human-readable description of this Metrics Configuration ID
func (id MetricsConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Metrics Configuration Name: %q", id.MetricsConfigurationName),
	}
	return fmt.Sprintf("Metrics Configuration (%s)", strings.Join(components, "\n"))
}
