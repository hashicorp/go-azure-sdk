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
	recaser.RegisterResourceId(&ResourceGroupDeploymentId{})
}

var _ resourceids.ResourceId = &ResourceGroupDeploymentId{}

// ResourceGroupDeploymentId is a struct representing the Resource ID for a Resource Group Deployment
type ResourceGroupDeploymentId struct {
	SubscriptionId    string
	ResourceGroupName string
	DeploymentName    string
}

// NewResourceGroupDeploymentID returns a new ResourceGroupDeploymentId struct
func NewResourceGroupDeploymentID(subscriptionId string, resourceGroupName string, deploymentName string) ResourceGroupDeploymentId {
	return ResourceGroupDeploymentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		DeploymentName:    deploymentName,
	}
}

// ParseResourceGroupDeploymentID parses 'input' into a ResourceGroupDeploymentId
func ParseResourceGroupDeploymentID(input string) (*ResourceGroupDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResourceGroupDeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourceGroupDeploymentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseResourceGroupDeploymentIDInsensitively parses 'input' case-insensitively into a ResourceGroupDeploymentId
// note: this method should only be used for API response data and not user input
func ParseResourceGroupDeploymentIDInsensitively(input string) (*ResourceGroupDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ResourceGroupDeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ResourceGroupDeploymentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ResourceGroupDeploymentId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateResourceGroupDeploymentID checks that 'input' can be parsed as a Resource Group Deployment ID
func ValidateResourceGroupDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseResourceGroupDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Resource Group Deployment ID
func (id ResourceGroupDeploymentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/deployments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.DeploymentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Resource Group Deployment ID
func (id ResourceGroupDeploymentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
	}
}

// String returns a human-readable description of this Resource Group Deployment ID
func (id ResourceGroupDeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
	}
	return fmt.Sprintf("Resource Group Deployment (%s)", strings.Join(components, "\n"))
}
