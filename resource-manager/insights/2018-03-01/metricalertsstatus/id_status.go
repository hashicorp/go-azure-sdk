package metricalertsstatus

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&StatusId{})
}

var _ resourceids.ResourceId = &StatusId{}

// StatusId is a struct representing the Resource ID for a Status
type StatusId struct {
	SubscriptionId    string
	ResourceGroupName string
	MetricAlertName   string
	StatusName        string
}

// NewStatusID returns a new StatusId struct
func NewStatusID(subscriptionId string, resourceGroupName string, metricAlertName string, statusName string) StatusId {
	return StatusId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		MetricAlertName:   metricAlertName,
		StatusName:        statusName,
	}
}

// ParseStatusID parses 'input' into a StatusId
func ParseStatusID(input string) (*StatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StatusId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseStatusIDInsensitively parses 'input' case-insensitively into a StatusId
// note: this method should only be used for API response data and not user input
func ParseStatusIDInsensitively(input string) (*StatusId, error) {
	parser := resourceids.NewParserFromResourceIdType(&StatusId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := StatusId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *StatusId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.MetricAlertName, ok = input.Parsed["metricAlertName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "metricAlertName", input)
	}

	if id.StatusName, ok = input.Parsed["statusName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "statusName", input)
	}

	return nil
}

// ValidateStatusID checks that 'input' can be parsed as a Status ID
func ValidateStatusID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseStatusID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Status ID
func (id StatusId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Insights/metricAlerts/%s/status/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MetricAlertName, id.StatusName)
}

// Segments returns a slice of Resource ID Segments which comprise this Status ID
func (id StatusId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticMetricAlerts", "metricAlerts", "metricAlerts"),
		resourceids.UserSpecifiedSegment("metricAlertName", "ruleName"),
		resourceids.StaticSegment("staticStatus", "status", "status"),
		resourceids.UserSpecifiedSegment("statusName", "statusName"),
	}
}

// String returns a human-readable description of this Status ID
func (id StatusId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Metric Alert Name: %q", id.MetricAlertName),
		fmt.Sprintf("Status Name: %q", id.StatusName),
	}
	return fmt.Sprintf("Status (%s)", strings.Join(components, "\n"))
}
