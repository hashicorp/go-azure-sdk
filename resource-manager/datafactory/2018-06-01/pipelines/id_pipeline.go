package pipelines

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PipelineId{}

// PipelineId is a struct representing the Resource ID for a Pipeline
type PipelineId struct {
	SubscriptionId    string
	ResourceGroupName string
	FactoryName       string
	PipelineName      string
}

// NewPipelineID returns a new PipelineId struct
func NewPipelineID(subscriptionId string, resourceGroupName string, factoryName string, pipelineName string) PipelineId {
	return PipelineId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		FactoryName:       factoryName,
		PipelineName:      pipelineName,
	}
}

// ParsePipelineID parses 'input' into a PipelineId
func ParsePipelineID(input string) (*PipelineId, error) {
	parser := resourceids.NewParserFromResourceIdType(PipelineId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PipelineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, fmt.Errorf("the segment 'factoryName' was not found in the resource id %q", input)
	}

	if id.PipelineName, ok = parsed.Parsed["pipelineName"]; !ok {
		return nil, fmt.Errorf("the segment 'pipelineName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParsePipelineIDInsensitively parses 'input' case-insensitively into a PipelineId
// note: this method should only be used for API response data and not user input
func ParsePipelineIDInsensitively(input string) (*PipelineId, error) {
	parser := resourceids.NewParserFromResourceIdType(PipelineId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PipelineId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, fmt.Errorf("the segment 'factoryName' was not found in the resource id %q", input)
	}

	if id.PipelineName, ok = parsed.Parsed["pipelineName"]; !ok {
		return nil, fmt.Errorf("the segment 'pipelineName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidatePipelineID checks that 'input' can be parsed as a Pipeline ID
func ValidatePipelineID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePipelineID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Pipeline ID
func (id PipelineId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/pipelines/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.PipelineName)
}

// Segments returns a slice of Resource ID Segments which comprise this Pipeline ID
func (id PipelineId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticPipelines", "pipelines", "pipelines"),
		resourceids.UserSpecifiedSegment("pipelineName", "pipelineValue"),
	}
}

// String returns a human-readable description of this Pipeline ID
func (id PipelineId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Pipeline Name: %q", id.PipelineName),
	}
	return fmt.Sprintf("Pipeline (%s)", strings.Join(components, "\n"))
}
