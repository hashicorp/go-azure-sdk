package iotsecuritysolutionsanalytics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = AggregatedRecommendationId{}

// AggregatedRecommendationId is a struct representing the Resource ID for a Aggregated Recommendation
type AggregatedRecommendationId struct {
	SubscriptionId               string
	ResourceGroupName            string
	IotSecuritySolutionName      string
	AggregatedRecommendationName string
}

// NewAggregatedRecommendationID returns a new AggregatedRecommendationId struct
func NewAggregatedRecommendationID(subscriptionId string, resourceGroupName string, iotSecuritySolutionName string, aggregatedRecommendationName string) AggregatedRecommendationId {
	return AggregatedRecommendationId{
		SubscriptionId:               subscriptionId,
		ResourceGroupName:            resourceGroupName,
		IotSecuritySolutionName:      iotSecuritySolutionName,
		AggregatedRecommendationName: aggregatedRecommendationName,
	}
}

// ParseAggregatedRecommendationID parses 'input' into a AggregatedRecommendationId
func ParseAggregatedRecommendationID(input string) (*AggregatedRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(AggregatedRecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AggregatedRecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.IotSecuritySolutionName, ok = parsed.Parsed["iotSecuritySolutionName"]; !ok {
		return nil, fmt.Errorf("the segment 'iotSecuritySolutionName' was not found in the resource id %q", input)
	}

	if id.AggregatedRecommendationName, ok = parsed.Parsed["aggregatedRecommendationName"]; !ok {
		return nil, fmt.Errorf("the segment 'aggregatedRecommendationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseAggregatedRecommendationIDInsensitively parses 'input' case-insensitively into a AggregatedRecommendationId
// note: this method should only be used for API response data and not user input
func ParseAggregatedRecommendationIDInsensitively(input string) (*AggregatedRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(AggregatedRecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AggregatedRecommendationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.IotSecuritySolutionName, ok = parsed.Parsed["iotSecuritySolutionName"]; !ok {
		return nil, fmt.Errorf("the segment 'iotSecuritySolutionName' was not found in the resource id %q", input)
	}

	if id.AggregatedRecommendationName, ok = parsed.Parsed["aggregatedRecommendationName"]; !ok {
		return nil, fmt.Errorf("the segment 'aggregatedRecommendationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateAggregatedRecommendationID checks that 'input' can be parsed as a Aggregated Recommendation ID
func ValidateAggregatedRecommendationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAggregatedRecommendationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Aggregated Recommendation ID
func (id AggregatedRecommendationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/iotSecuritySolutions/%s/analyticsModels/default/aggregatedRecommendations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.IotSecuritySolutionName, id.AggregatedRecommendationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Aggregated Recommendation ID
func (id AggregatedRecommendationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticIotSecuritySolutions", "iotSecuritySolutions", "iotSecuritySolutions"),
		resourceids.UserSpecifiedSegment("iotSecuritySolutionName", "iotSecuritySolutionValue"),
		resourceids.StaticSegment("staticAnalyticsModels", "analyticsModels", "analyticsModels"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticAggregatedRecommendations", "aggregatedRecommendations", "aggregatedRecommendations"),
		resourceids.UserSpecifiedSegment("aggregatedRecommendationName", "aggregatedRecommendationValue"),
	}
}

// String returns a human-readable description of this Aggregated Recommendation ID
func (id AggregatedRecommendationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Iot Security Solution Name: %q", id.IotSecuritySolutionName),
		fmt.Sprintf("Aggregated Recommendation Name: %q", id.AggregatedRecommendationName),
	}
	return fmt.Sprintf("Aggregated Recommendation (%s)", strings.Join(components, "\n"))
}
