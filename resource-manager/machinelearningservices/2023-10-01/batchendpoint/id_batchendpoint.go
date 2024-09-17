package batchendpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BatchEndpointId{})
}

var _ resourceids.ResourceId = &BatchEndpointId{}

// BatchEndpointId is a struct representing the Resource ID for a Batch Endpoint
type BatchEndpointId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	BatchEndpointName string
}

// NewBatchEndpointID returns a new BatchEndpointId struct
func NewBatchEndpointID(subscriptionId string, resourceGroupName string, workspaceName string, batchEndpointName string) BatchEndpointId {
	return BatchEndpointId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		BatchEndpointName: batchEndpointName,
	}
}

// ParseBatchEndpointID parses 'input' into a BatchEndpointId
func ParseBatchEndpointID(input string) (*BatchEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BatchEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BatchEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBatchEndpointIDInsensitively parses 'input' case-insensitively into a BatchEndpointId
// note: this method should only be used for API response data and not user input
func ParseBatchEndpointIDInsensitively(input string) (*BatchEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BatchEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BatchEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BatchEndpointId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.BatchEndpointName, ok = input.Parsed["batchEndpointName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "batchEndpointName", input)
	}

	return nil
}

// ValidateBatchEndpointID checks that 'input' can be parsed as a Batch Endpoint ID
func ValidateBatchEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBatchEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Batch Endpoint ID
func (id BatchEndpointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/batchEndpoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.BatchEndpointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Batch Endpoint ID
func (id BatchEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticBatchEndpoints", "batchEndpoints", "batchEndpoints"),
		resourceids.UserSpecifiedSegment("batchEndpointName", "batchEndpointValue"),
	}
}

// String returns a human-readable description of this Batch Endpoint ID
func (id BatchEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Batch Endpoint Name: %q", id.BatchEndpointName),
	}
	return fmt.Sprintf("Batch Endpoint (%s)", strings.Join(components, "\n"))
}
