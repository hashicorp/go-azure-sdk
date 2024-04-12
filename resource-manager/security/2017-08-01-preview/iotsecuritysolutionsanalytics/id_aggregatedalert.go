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
	recaser.RegisterResourceId(&AggregatedAlertId{})
}

var _ resourceids.ResourceId = &AggregatedAlertId{}

// AggregatedAlertId is a struct representing the Resource ID for a Aggregated Alert
type AggregatedAlertId struct {
	SubscriptionId          string
	ResourceGroupName       string
	IotSecuritySolutionName string
	AggregatedAlertName     string
}

// NewAggregatedAlertID returns a new AggregatedAlertId struct
func NewAggregatedAlertID(subscriptionId string, resourceGroupName string, iotSecuritySolutionName string, aggregatedAlertName string) AggregatedAlertId {
	return AggregatedAlertId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		IotSecuritySolutionName: iotSecuritySolutionName,
		AggregatedAlertName:     aggregatedAlertName,
	}
}

// ParseAggregatedAlertID parses 'input' into a AggregatedAlertId
func ParseAggregatedAlertID(input string) (*AggregatedAlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AggregatedAlertId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AggregatedAlertId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAggregatedAlertIDInsensitively parses 'input' case-insensitively into a AggregatedAlertId
// note: this method should only be used for API response data and not user input
func ParseAggregatedAlertIDInsensitively(input string) (*AggregatedAlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AggregatedAlertId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AggregatedAlertId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AggregatedAlertId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AggregatedAlertName, ok = input.Parsed["aggregatedAlertName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "aggregatedAlertName", input)
	}

	return nil
}

// ValidateAggregatedAlertID checks that 'input' can be parsed as a Aggregated Alert ID
func ValidateAggregatedAlertID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAggregatedAlertID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Aggregated Alert ID
func (id AggregatedAlertId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/iotSecuritySolutions/%s/analyticsModels/default/aggregatedAlerts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.IotSecuritySolutionName, id.AggregatedAlertName)
}

// Segments returns a slice of Resource ID Segments which comprise this Aggregated Alert ID
func (id AggregatedAlertId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticAggregatedAlerts", "aggregatedAlerts", "aggregatedAlerts"),
		resourceids.UserSpecifiedSegment("aggregatedAlertName", "aggregatedAlertValue"),
	}
}

// String returns a human-readable description of this Aggregated Alert ID
func (id AggregatedAlertId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Iot Security Solution Name: %q", id.IotSecuritySolutionName),
		fmt.Sprintf("Aggregated Alert Name: %q", id.AggregatedAlertName),
	}
	return fmt.Sprintf("Aggregated Alert (%s)", strings.Join(components, "\n"))
}
