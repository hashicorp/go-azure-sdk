package monitoredsubscriptions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = MonitoredSubscriptionId{}

// MonitoredSubscriptionId is a struct representing the Resource ID for a Monitored Subscription
type MonitoredSubscriptionId struct {
	SubscriptionId            string
	ResourceGroupName         string
	MonitorName               string
	MonitoredSubscriptionName string
}

// NewMonitoredSubscriptionID returns a new MonitoredSubscriptionId struct
func NewMonitoredSubscriptionID(subscriptionId string, resourceGroupName string, monitorName string, monitoredSubscriptionName string) MonitoredSubscriptionId {
	return MonitoredSubscriptionId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		MonitorName:               monitorName,
		MonitoredSubscriptionName: monitoredSubscriptionName,
	}
}

// ParseMonitoredSubscriptionID parses 'input' into a MonitoredSubscriptionId
func ParseMonitoredSubscriptionID(input string) (*MonitoredSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MonitoredSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MonitoredSubscriptionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.MonitorName, ok = parsed.Parsed["monitorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "monitorName", *parsed)
	}

	if id.MonitoredSubscriptionName, ok = parsed.Parsed["monitoredSubscriptionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "monitoredSubscriptionName", *parsed)
	}

	return &id, nil
}

// ParseMonitoredSubscriptionIDInsensitively parses 'input' case-insensitively into a MonitoredSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseMonitoredSubscriptionIDInsensitively(input string) (*MonitoredSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(MonitoredSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := MonitoredSubscriptionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.MonitorName, ok = parsed.Parsed["monitorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "monitorName", *parsed)
	}

	if id.MonitoredSubscriptionName, ok = parsed.Parsed["monitoredSubscriptionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "monitoredSubscriptionName", *parsed)
	}

	return &id, nil
}

// ValidateMonitoredSubscriptionID checks that 'input' can be parsed as a Monitored Subscription ID
func ValidateMonitoredSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMonitoredSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Monitored Subscription ID
func (id MonitoredSubscriptionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Datadog/monitors/%s/monitoredSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.MonitorName, id.MonitoredSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Monitored Subscription ID
func (id MonitoredSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDatadog", "Microsoft.Datadog", "Microsoft.Datadog"),
		resourceids.StaticSegment("staticMonitors", "monitors", "monitors"),
		resourceids.UserSpecifiedSegment("monitorName", "monitorValue"),
		resourceids.StaticSegment("staticMonitoredSubscriptions", "monitoredSubscriptions", "monitoredSubscriptions"),
		resourceids.UserSpecifiedSegment("monitoredSubscriptionName", "monitoredSubscriptionValue"),
	}
}

// String returns a human-readable description of this Monitored Subscription ID
func (id MonitoredSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Monitor Name: %q", id.MonitorName),
		fmt.Sprintf("Monitored Subscription Name: %q", id.MonitoredSubscriptionName),
	}
	return fmt.Sprintf("Monitored Subscription (%s)", strings.Join(components, "\n"))
}
