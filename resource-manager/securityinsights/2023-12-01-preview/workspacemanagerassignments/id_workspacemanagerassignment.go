package workspacemanagerassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&WorkspaceManagerAssignmentId{})
}

var _ resourceids.ResourceId = &WorkspaceManagerAssignmentId{}

// WorkspaceManagerAssignmentId is a struct representing the Resource ID for a Workspace Manager Assignment
type WorkspaceManagerAssignmentId struct {
	SubscriptionId                 string
	ResourceGroupName              string
	WorkspaceName                  string
	WorkspaceManagerAssignmentName string
}

// NewWorkspaceManagerAssignmentID returns a new WorkspaceManagerAssignmentId struct
func NewWorkspaceManagerAssignmentID(subscriptionId string, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string) WorkspaceManagerAssignmentId {
	return WorkspaceManagerAssignmentId{
		SubscriptionId:                 subscriptionId,
		ResourceGroupName:              resourceGroupName,
		WorkspaceName:                  workspaceName,
		WorkspaceManagerAssignmentName: workspaceManagerAssignmentName,
	}
}

// ParseWorkspaceManagerAssignmentID parses 'input' into a WorkspaceManagerAssignmentId
func ParseWorkspaceManagerAssignmentID(input string) (*WorkspaceManagerAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceManagerAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceManagerAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWorkspaceManagerAssignmentIDInsensitively parses 'input' case-insensitively into a WorkspaceManagerAssignmentId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceManagerAssignmentIDInsensitively(input string) (*WorkspaceManagerAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceManagerAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceManagerAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WorkspaceManagerAssignmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.WorkspaceManagerAssignmentName, ok = input.Parsed["workspaceManagerAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceManagerAssignmentName", input)
	}

	return nil
}

// ValidateWorkspaceManagerAssignmentID checks that 'input' can be parsed as a Workspace Manager Assignment ID
func ValidateWorkspaceManagerAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceManagerAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Manager Assignment ID
func (id WorkspaceManagerAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.WorkspaceManagerAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Manager Assignment ID
func (id WorkspaceManagerAssignmentId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticWorkspaceManagerAssignments", "workspaceManagerAssignments", "workspaceManagerAssignments"),
		resourceids.UserSpecifiedSegment("workspaceManagerAssignmentName", "workspaceManagerAssignmentName"),
	}
}

// String returns a human-readable description of this Workspace Manager Assignment ID
func (id WorkspaceManagerAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Workspace Manager Assignment Name: %q", id.WorkspaceManagerAssignmentName),
	}
	return fmt.Sprintf("Workspace Manager Assignment (%s)", strings.Join(components, "\n"))
}
