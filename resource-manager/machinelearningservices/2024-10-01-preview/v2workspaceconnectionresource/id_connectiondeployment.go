package v2workspaceconnectionresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ConnectionDeploymentId{})
}

var _ resourceids.ResourceId = &ConnectionDeploymentId{}

// ConnectionDeploymentId is a struct representing the Resource ID for a Connection Deployment
type ConnectionDeploymentId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	ConnectionName    string
	DeploymentName    string
}

// NewConnectionDeploymentID returns a new ConnectionDeploymentId struct
func NewConnectionDeploymentID(subscriptionId string, resourceGroupName string, workspaceName string, connectionName string, deploymentName string) ConnectionDeploymentId {
	return ConnectionDeploymentId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		ConnectionName:    connectionName,
		DeploymentName:    deploymentName,
	}
}

// ParseConnectionDeploymentID parses 'input' into a ConnectionDeploymentId
func ParseConnectionDeploymentID(input string) (*ConnectionDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConnectionDeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConnectionDeploymentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseConnectionDeploymentIDInsensitively parses 'input' case-insensitively into a ConnectionDeploymentId
// note: this method should only be used for API response data and not user input
func ParseConnectionDeploymentIDInsensitively(input string) (*ConnectionDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConnectionDeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConnectionDeploymentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ConnectionDeploymentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ConnectionName, ok = input.Parsed["connectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "connectionName", input)
	}

	if id.DeploymentName, ok = input.Parsed["deploymentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", input)
	}

	return nil
}

// ValidateConnectionDeploymentID checks that 'input' can be parsed as a Connection Deployment ID
func ValidateConnectionDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConnectionDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Connection Deployment ID
func (id ConnectionDeploymentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/connections/%s/deployments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.ConnectionName, id.DeploymentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Connection Deployment ID
func (id ConnectionDeploymentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticConnections", "connections", "connections"),
		resourceids.UserSpecifiedSegment("connectionName", "connectionName"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentName"),
	}
}

// String returns a human-readable description of this Connection Deployment ID
func (id ConnectionDeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Connection Name: %q", id.ConnectionName),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
	}
	return fmt.Sprintf("Connection Deployment (%s)", strings.Join(components, "\n"))
}
