package batchdeployment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = DeploymentId{}

// DeploymentId is a struct representing the Resource ID for a Deployment
type DeploymentId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	EndpointName      string
	DeploymentName    string
}

// NewDeploymentID returns a new DeploymentId struct
func NewDeploymentID(subscriptionId string, resourceGroupName string, workspaceName string, endpointName string, deploymentName string) DeploymentId {
	return DeploymentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		EndpointName:      endpointName,
		DeploymentName:    deploymentName,
	}
}

// ParseDeploymentID parses 'input' into a DeploymentId
func ParseDeploymentID(input string) (*DeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeploymentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EndpointName, ok = parsed.Parsed["endpointName"]; !ok {
		return nil, fmt.Errorf("the segment 'endpointName' was not found in the resource id %q", input)
	}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, fmt.Errorf("the segment 'deploymentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseDeploymentIDInsensitively parses 'input' case-insensitively into a DeploymentId
// note: this method should only be used for API response data and not user input
func ParseDeploymentIDInsensitively(input string) (*DeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(DeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := DeploymentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, fmt.Errorf("the segment 'workspaceName' was not found in the resource id %q", input)
	}

	if id.EndpointName, ok = parsed.Parsed["endpointName"]; !ok {
		return nil, fmt.Errorf("the segment 'endpointName' was not found in the resource id %q", input)
	}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, fmt.Errorf("the segment 'deploymentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateDeploymentID checks that 'input' can be parsed as a Deployment ID
func ValidateDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deployment ID
func (id DeploymentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/batchEndpoints/%s/deployments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EndpointName, id.DeploymentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Deployment ID
func (id DeploymentId) Segments() []resourceids.Segment {
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
		resourceids.UserSpecifiedSegment("endpointName", "endpointValue"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
	}
}

// String returns a human-readable description of this Deployment ID
func (id DeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Endpoint Name: %q", id.EndpointName),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
	}
	return fmt.Sprintf("Deployment (%s)", strings.Join(components, "\n"))
}
