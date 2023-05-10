package dataflows

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = DataflowId{}

// DataflowId is a struct representing the Resource ID for a Dataflow
type DataflowId struct {
	SubscriptionId    string
	ResourceGroupName string
	FactoryName       string
	DataflowName      string
}

// NewDataflowID returns a new DataflowId struct
func NewDataflowID(subscriptionId string, resourceGroupName string, factoryName string, dataflowName string) DataflowId {
	return DataflowId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		FactoryName:       factoryName,
		DataflowName:      dataflowName,
	}
}

// ParseDataflowID parses 'input' into a DataflowId
func ParseDataflowID(input string) (*DataflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataflowId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataflowId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "factoryName", *parsed)
	}

	if id.DataflowName, ok = parsed.Parsed["dataflowName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "dataflowName", *parsed)
	}

	return &id, nil
}

// ParseDataflowIDInsensitively parses 'input' case-insensitively into a DataflowId
// note: this method should only be used for API response data and not user input
func ParseDataflowIDInsensitively(input string) (*DataflowId, error) {
	parser := resourceids.NewParserFromResourceIdType(DataflowId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DataflowId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "factoryName", *parsed)
	}

	if id.DataflowName, ok = parsed.Parsed["dataflowName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "dataflowName", *parsed)
	}

	return &id, nil
}

// ValidateDataflowID checks that 'input' can be parsed as a Dataflow ID
func ValidateDataflowID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataflowID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dataflow ID
func (id DataflowId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/dataflows/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.DataflowName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dataflow ID
func (id DataflowId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticDataflows", "dataflows", "dataflows"),
		resourceids.UserSpecifiedSegment("dataflowName", "dataflowValue"),
	}
}

// String returns a human-readable description of this Dataflow ID
func (id DataflowId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Dataflow Name: %q", id.DataflowName),
	}
	return fmt.Sprintf("Dataflow (%s)", strings.Join(components, "\n"))
}
