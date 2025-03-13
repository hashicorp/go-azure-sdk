package workspacemanagergroups

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&WorkspaceManagerGroupId{})
}

var _ resourceids.ResourceId = &WorkspaceManagerGroupId{}

// WorkspaceManagerGroupId is a struct representing the Resource ID for a Workspace Manager Group
type WorkspaceManagerGroupId struct {
	SubscriptionId            string
	ResourceGroupName         string
	WorkspaceName             string
	WorkspaceManagerGroupName string
}

// NewWorkspaceManagerGroupID returns a new WorkspaceManagerGroupId struct
func NewWorkspaceManagerGroupID(subscriptionId string, resourceGroupName string, workspaceName string, workspaceManagerGroupName string) WorkspaceManagerGroupId {
	return WorkspaceManagerGroupId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		WorkspaceName:             workspaceName,
		WorkspaceManagerGroupName: workspaceManagerGroupName,
	}
}

// ParseWorkspaceManagerGroupID parses 'input' into a WorkspaceManagerGroupId
func ParseWorkspaceManagerGroupID(input string) (*WorkspaceManagerGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceManagerGroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceManagerGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWorkspaceManagerGroupIDInsensitively parses 'input' case-insensitively into a WorkspaceManagerGroupId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceManagerGroupIDInsensitively(input string) (*WorkspaceManagerGroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceManagerGroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceManagerGroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WorkspaceManagerGroupId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.WorkspaceManagerGroupName, ok = input.Parsed["workspaceManagerGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceManagerGroupName", input)
	}

	return nil
}

// ValidateWorkspaceManagerGroupID checks that 'input' can be parsed as a Workspace Manager Group ID
func ValidateWorkspaceManagerGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceManagerGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Manager Group ID
func (id WorkspaceManagerGroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/workspaceManagerGroups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.WorkspaceManagerGroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Manager Group ID
func (id WorkspaceManagerGroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticWorkspaceManagerGroups", "workspaceManagerGroups", "workspaceManagerGroups"),
		resourceids.UserSpecifiedSegment("workspaceManagerGroupName", "workspaceManagerGroupName"),
	}
}

// String returns a human-readable description of this Workspace Manager Group ID
func (id WorkspaceManagerGroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Workspace Manager Group Name: %q", id.WorkspaceManagerGroupName),
	}
	return fmt.Sprintf("Workspace Manager Group (%s)", strings.Join(components, "\n"))
}
