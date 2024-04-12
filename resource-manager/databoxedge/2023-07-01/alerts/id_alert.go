package alerts

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&AlertId{})
}

var _ resourceids.ResourceId = &AlertId{}

// AlertId is a struct representing the Resource ID for a Alert
type AlertId struct {
	SubscriptionId        string
	ResourceGroupName     string
	DataBoxEdgeDeviceName string
	AlertName             string
}

// NewAlertID returns a new AlertId struct
func NewAlertID(subscriptionId string, resourceGroupName string, dataBoxEdgeDeviceName string, alertName string) AlertId {
	return AlertId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		DataBoxEdgeDeviceName: dataBoxEdgeDeviceName,
		AlertName:             alertName,
	}
}

// ParseAlertID parses 'input' into a AlertId
func ParseAlertID(input string) (*AlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AlertId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AlertId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAlertIDInsensitively parses 'input' case-insensitively into a AlertId
// note: this method should only be used for API response data and not user input
func ParseAlertIDInsensitively(input string) (*AlertId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AlertId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AlertId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AlertId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DataBoxEdgeDeviceName, ok = input.Parsed["dataBoxEdgeDeviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataBoxEdgeDeviceName", input)
	}

	if id.AlertName, ok = input.Parsed["alertName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "alertName", input)
	}

	return nil
}

// ValidateAlertID checks that 'input' can be parsed as a Alert ID
func ValidateAlertID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAlertID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Alert ID
func (id AlertId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/%s/alerts/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DataBoxEdgeDeviceName, id.AlertName)
}

// Segments returns a slice of Resource ID Segments which comprise this Alert ID
func (id AlertId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataBoxEdge", "Microsoft.DataBoxEdge", "Microsoft.DataBoxEdge"),
		resourceids.StaticSegment("staticDataBoxEdgeDevices", "dataBoxEdgeDevices", "dataBoxEdgeDevices"),
		resourceids.UserSpecifiedSegment("dataBoxEdgeDeviceName", "dataBoxEdgeDeviceValue"),
		resourceids.StaticSegment("staticAlerts", "alerts", "alerts"),
		resourceids.UserSpecifiedSegment("alertName", "alertValue"),
	}
}

// String returns a human-readable description of this Alert ID
func (id AlertId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Data Box Edge Device Name: %q", id.DataBoxEdgeDeviceName),
		fmt.Sprintf("Alert Name: %q", id.AlertName),
	}
	return fmt.Sprintf("Alert (%s)", strings.Join(components, "\n"))
}
