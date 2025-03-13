package batchdeployment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BatchEndpointDeploymentId{})
}

var _ resourceids.ResourceId = &BatchEndpointDeploymentId{}

// BatchEndpointDeploymentId is a struct representing the Resource ID for a Batch Endpoint Deployment
type BatchEndpointDeploymentId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	BatchEndpointName string
	DeploymentName    string
}

// NewBatchEndpointDeploymentID returns a new BatchEndpointDeploymentId struct
func NewBatchEndpointDeploymentID(subscriptionId string, resourceGroupName string, workspaceName string, batchEndpointName string, deploymentName string) BatchEndpointDeploymentId {
	return BatchEndpointDeploymentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		BatchEndpointName: batchEndpointName,
		DeploymentName:    deploymentName,
	}
}

// ParseBatchEndpointDeploymentID parses 'input' into a BatchEndpointDeploymentId
func ParseBatchEndpointDeploymentID(input string) (*BatchEndpointDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BatchEndpointDeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BatchEndpointDeploymentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBatchEndpointDeploymentIDInsensitively parses 'input' case-insensitively into a BatchEndpointDeploymentId
// note: this method should only be used for API response data and not user input
func ParseBatchEndpointDeploymentIDInsensitively(input string) (*BatchEndpointDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BatchEndpointDeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BatchEndpointDeploymentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BatchEndpointDeploymentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DeploymentName, ok = input.Parsed["deploymentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", input)
	}

	return nil
}

// ValidateBatchEndpointDeploymentID checks that 'input' can be parsed as a Batch Endpoint Deployment ID
func ValidateBatchEndpointDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBatchEndpointDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Batch Endpoint Deployment ID
func (id BatchEndpointDeploymentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/batchEndpoints/%s/deployments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.BatchEndpointName, id.DeploymentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Batch Endpoint Deployment ID
func (id BatchEndpointDeploymentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticBatchEndpoints", "batchEndpoints", "batchEndpoints"),
		resourceids.UserSpecifiedSegment("batchEndpointName", "batchEndpointName"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentName"),
	}
}

// String returns a human-readable description of this Batch Endpoint Deployment ID
func (id BatchEndpointDeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Batch Endpoint Name: %q", id.BatchEndpointName),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
	}
	return fmt.Sprintf("Batch Endpoint Deployment (%s)", strings.Join(components, "\n"))
}
