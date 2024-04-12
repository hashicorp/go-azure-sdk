package iotsecuritysolutionsanalytics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AggregatedRecommendationId{})
}

var _ resourceids.ResourceId = &AggregatedRecommendationId{}

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
	parser := resourceids.NewParserFromResourceIdType(&AggregatedRecommendationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AggregatedRecommendationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAggregatedRecommendationIDInsensitively parses 'input' case-insensitively into a AggregatedRecommendationId
// note: this method should only be used for API response data and not user input
func ParseAggregatedRecommendationIDInsensitively(input string) (*AggregatedRecommendationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AggregatedRecommendationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AggregatedRecommendationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AggregatedRecommendationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.IotSecuritySolutionName, ok = input.Parsed["iotSecuritySolutionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "iotSecuritySolutionName", input)
	}

	if id.AggregatedRecommendationName, ok = input.Parsed["aggregatedRecommendationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "aggregatedRecommendationName", input)
	}

	return nil
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
