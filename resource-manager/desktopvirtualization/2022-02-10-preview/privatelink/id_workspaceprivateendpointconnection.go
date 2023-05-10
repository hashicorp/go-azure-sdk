package privatelink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkspacePrivateEndpointConnectionId{}

// WorkspacePrivateEndpointConnectionId is a struct representing the Resource ID for a Workspace Private Endpoint Connection
type WorkspacePrivateEndpointConnectionId struct {
	SubscriptionId                string
	ResourceGroupName             string
	WorkspaceName                 string
	PrivateEndpointConnectionName string
}

// NewWorkspacePrivateEndpointConnectionID returns a new WorkspacePrivateEndpointConnectionId struct
func NewWorkspacePrivateEndpointConnectionID(subscriptionId string, resourceGroupName string, workspaceName string, privateEndpointConnectionName string) WorkspacePrivateEndpointConnectionId {
	return WorkspacePrivateEndpointConnectionId{
		SubscriptionId:                subscriptionId,
		ResourceGroupName:             resourceGroupName,
		WorkspaceName:                 workspaceName,
		PrivateEndpointConnectionName: privateEndpointConnectionName,
	}
}

// ParseWorkspacePrivateEndpointConnectionID parses 'input' into a WorkspacePrivateEndpointConnectionId
func ParseWorkspacePrivateEndpointConnectionID(input string) (*WorkspacePrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspacePrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspacePrivateEndpointConnectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.PrivateEndpointConnectionName, ok = parsed.Parsed["privateEndpointConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "privateEndpointConnectionName", *parsed)
	}

	return &id, nil
}

// ParseWorkspacePrivateEndpointConnectionIDInsensitively parses 'input' case-insensitively into a WorkspacePrivateEndpointConnectionId
// note: this method should only be used for API response data and not user input
func ParseWorkspacePrivateEndpointConnectionIDInsensitively(input string) (*WorkspacePrivateEndpointConnectionId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspacePrivateEndpointConnectionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspacePrivateEndpointConnectionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.PrivateEndpointConnectionName, ok = parsed.Parsed["privateEndpointConnectionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "privateEndpointConnectionName", *parsed)
	}

	return &id, nil
}

// ValidateWorkspacePrivateEndpointConnectionID checks that 'input' can be parsed as a Workspace Private Endpoint Connection ID
func ValidateWorkspacePrivateEndpointConnectionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspacePrivateEndpointConnectionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Private Endpoint Connection ID
func (id WorkspacePrivateEndpointConnectionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DesktopVirtualization/workspaces/%s/privateEndpointConnections/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.PrivateEndpointConnectionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Private Endpoint Connection ID
func (id WorkspacePrivateEndpointConnectionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDesktopVirtualization", "Microsoft.DesktopVirtualization", "Microsoft.DesktopVirtualization"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticPrivateEndpointConnections", "privateEndpointConnections", "privateEndpointConnections"),
		resourceids.UserSpecifiedSegment("privateEndpointConnectionName", "privateEndpointConnectionValue"),
	}
}

// String returns a human-readable description of this Workspace Private Endpoint Connection ID
func (id WorkspacePrivateEndpointConnectionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Private Endpoint Connection Name: %q", id.PrivateEndpointConnectionName),
	}
	return fmt.Sprintf("Workspace Private Endpoint Connection (%s)", strings.Join(components, "\n"))
}
