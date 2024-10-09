package serverlessendpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ServerlessEndpointId{})
}

var _ resourceids.ResourceId = &ServerlessEndpointId{}

// ServerlessEndpointId is a struct representing the Resource ID for a Serverless Endpoint
type ServerlessEndpointId struct {
	SubscriptionId         string
	ResourceGroupName      string
	WorkspaceName          string
	ServerlessEndpointName string
}

// NewServerlessEndpointID returns a new ServerlessEndpointId struct
func NewServerlessEndpointID(subscriptionId string, resourceGroupName string, workspaceName string, serverlessEndpointName string) ServerlessEndpointId {
	return ServerlessEndpointId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		WorkspaceName:          workspaceName,
		ServerlessEndpointName: serverlessEndpointName,
	}
}

// ParseServerlessEndpointID parses 'input' into a ServerlessEndpointId
func ParseServerlessEndpointID(input string) (*ServerlessEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServerlessEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServerlessEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseServerlessEndpointIDInsensitively parses 'input' case-insensitively into a ServerlessEndpointId
// note: this method should only be used for API response data and not user input
func ParseServerlessEndpointIDInsensitively(input string) (*ServerlessEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ServerlessEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ServerlessEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ServerlessEndpointId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ServerlessEndpointName, ok = input.Parsed["serverlessEndpointName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "serverlessEndpointName", input)
	}

	return nil
}

// ValidateServerlessEndpointID checks that 'input' can be parsed as a Serverless Endpoint ID
func ValidateServerlessEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseServerlessEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Serverless Endpoint ID
func (id ServerlessEndpointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/serverlessEndpoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.ServerlessEndpointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Serverless Endpoint ID
func (id ServerlessEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticServerlessEndpoints", "serverlessEndpoints", "serverlessEndpoints"),
		resourceids.UserSpecifiedSegment("serverlessEndpointName", "serverlessEndpointName"),
	}
}

// String returns a human-readable description of this Serverless Endpoint ID
func (id ServerlessEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Serverless Endpoint Name: %q", id.ServerlessEndpointName),
	}
	return fmt.Sprintf("Serverless Endpoint (%s)", strings.Join(components, "\n"))
}
