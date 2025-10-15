package dataflowgraph

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DataflowGraphId{})
}

var _ resourceids.ResourceId = &DataflowGraphId{}

// DataflowGraphId is a struct representing the Resource ID for a Dataflow Graph
type DataflowGraphId struct {
	SubscriptionId      string
	ResourceGroupName   string
	InstanceName        string
	DataflowProfileName string
	DataflowGraphName   string
}

// NewDataflowGraphID returns a new DataflowGraphId struct
func NewDataflowGraphID(subscriptionId string, resourceGroupName string, instanceName string, dataflowProfileName string, dataflowGraphName string) DataflowGraphId {
	return DataflowGraphId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		InstanceName:        instanceName,
		DataflowProfileName: dataflowProfileName,
		DataflowGraphName:   dataflowGraphName,
	}
}

// ParseDataflowGraphID parses 'input' into a DataflowGraphId
func ParseDataflowGraphID(input string) (*DataflowGraphId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataflowGraphId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataflowGraphId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDataflowGraphIDInsensitively parses 'input' case-insensitively into a DataflowGraphId
// note: this method should only be used for API response data and not user input
func ParseDataflowGraphIDInsensitively(input string) (*DataflowGraphId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DataflowGraphId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DataflowGraphId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DataflowGraphId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.InstanceName, ok = input.Parsed["instanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "instanceName", input)
	}

	if id.DataflowProfileName, ok = input.Parsed["dataflowProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataflowProfileName", input)
	}

	if id.DataflowGraphName, ok = input.Parsed["dataflowGraphName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataflowGraphName", input)
	}

	return nil
}

// ValidateDataflowGraphID checks that 'input' can be parsed as a Dataflow Graph ID
func ValidateDataflowGraphID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDataflowGraphID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Dataflow Graph ID
func (id DataflowGraphId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.IoTOperations/instances/%s/dataflowProfiles/%s/dataflowGraphs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.InstanceName, id.DataflowProfileName, id.DataflowGraphName)
}

// Segments returns a slice of Resource ID Segments which comprise this Dataflow Graph ID
func (id DataflowGraphId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftIoTOperations", "Microsoft.IoTOperations", "Microsoft.IoTOperations"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("instanceName", "instanceName"),
		resourceids.StaticSegment("staticDataflowProfiles", "dataflowProfiles", "dataflowProfiles"),
		resourceids.UserSpecifiedSegment("dataflowProfileName", "dataflowProfileName"),
		resourceids.StaticSegment("staticDataflowGraphs", "dataflowGraphs", "dataflowGraphs"),
		resourceids.UserSpecifiedSegment("dataflowGraphName", "dataflowGraphName"),
	}
}

// String returns a human-readable description of this Dataflow Graph ID
func (id DataflowGraphId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Instance Name: %q", id.InstanceName),
		fmt.Sprintf("Dataflow Profile Name: %q", id.DataflowProfileName),
		fmt.Sprintf("Dataflow Graph Name: %q", id.DataflowGraphName),
	}
	return fmt.Sprintf("Dataflow Graph (%s)", strings.Join(components, "\n"))
}
