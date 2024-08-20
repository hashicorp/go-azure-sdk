package deploymentoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ResourceGroupDeploymentOperationId{})
}

var _ resourceids.ResourceId = &ResourceGroupDeploymentOperationId{}

// ResourceGroupDeploymentOperationId is a struct representing the Resource ID for a Resource Group Deployment Operation
type ResourceGroupDeploymentOperationId struct {
	SubscriptionId    string
	ResourceGroupName string
	DeploymentName    string
	OperationId       string
}

// NewResourceGroupDeploymentOperationID returns a new ResourceGroupDeploymentOperationId struct
func NewResourceGroupDeploymentOperationID(subscriptionId string, resourceGroupName string, deploymentName string, operationId string) ResourceGroupDeploymentOperationId {
	return ResourceGroupDeploymentOperationId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		DeploymentName:    deploymentName,
		OperationId:       operationId,
	}
}

// ParseResourceGroupDeploymentOperationID parses 'input' into a ResourceGroupDeploymentOperationId
func ParseResourceGroupDeploymentOperationID(input string) (*ResourceGroupDeploymentOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResourceGroupDeploymentOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourceGroupDeploymentOperationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseResourceGroupDeploymentOperationIDInsensitively parses 'input' case-insensitively into a ResourceGroupDeploymentOperationId
// note: this method should only be used for API response data and not user input
func ParseResourceGroupDeploymentOperationIDInsensitively(input string) (*ResourceGroupDeploymentOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResourceGroupDeploymentOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourceGroupDeploymentOperationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ResourceGroupDeploymentOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.DeploymentName, ok = input.Parsed["deploymentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", input)
	}

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
}

// ValidateResourceGroupDeploymentOperationID checks that 'input' can be parsed as a Resource Group Deployment Operation ID
func ValidateResourceGroupDeploymentOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResourceGroupDeploymentOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resource Group Deployment Operation ID
func (id ResourceGroupDeploymentOperationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/deployments/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DeploymentName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Resource Group Deployment Operation ID
func (id ResourceGroupDeploymentOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Resource Group Deployment Operation ID
func (id ResourceGroupDeploymentOperationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Resource Group Deployment Operation (%s)", strings.Join(components, "\n"))
}
