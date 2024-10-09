package workspaceprivatelinkresources

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&WorkspacePrivateLinkResourceId{})
}

var _ resourceids.ResourceId = &WorkspacePrivateLinkResourceId{}

// WorkspacePrivateLinkResourceId is a struct representing the Resource ID for a Workspace Private Link Resource
type WorkspacePrivateLinkResourceId struct {
	SubscriptionId          string
	ResourceGroupName       string
	WorkspaceName           string
	PrivateLinkResourceName string
}

// NewWorkspacePrivateLinkResourceID returns a new WorkspacePrivateLinkResourceId struct
func NewWorkspacePrivateLinkResourceID(subscriptionId string, resourceGroupName string, workspaceName string, privateLinkResourceName string) WorkspacePrivateLinkResourceId {
	return WorkspacePrivateLinkResourceId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		WorkspaceName:           workspaceName,
		PrivateLinkResourceName: privateLinkResourceName,
	}
}

// ParseWorkspacePrivateLinkResourceID parses 'input' into a WorkspacePrivateLinkResourceId
func ParseWorkspacePrivateLinkResourceID(input string) (*WorkspacePrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspacePrivateLinkResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspacePrivateLinkResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWorkspacePrivateLinkResourceIDInsensitively parses 'input' case-insensitively into a WorkspacePrivateLinkResourceId
// note: this method should only be used for API response data and not user input
func ParseWorkspacePrivateLinkResourceIDInsensitively(input string) (*WorkspacePrivateLinkResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspacePrivateLinkResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspacePrivateLinkResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WorkspacePrivateLinkResourceId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PrivateLinkResourceName, ok = input.Parsed["privateLinkResourceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "privateLinkResourceName", input)
	}

	return nil
}

// ValidateWorkspacePrivateLinkResourceID checks that 'input' can be parsed as a Workspace Private Link Resource ID
func ValidateWorkspacePrivateLinkResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspacePrivateLinkResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Private Link Resource ID
func (id WorkspacePrivateLinkResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.HealthcareApis/workspaces/%s/privateLinkResources/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.PrivateLinkResourceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Private Link Resource ID
func (id WorkspacePrivateLinkResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftHealthcareApis", "Microsoft.HealthcareApis", "Microsoft.HealthcareApis"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticPrivateLinkResources", "privateLinkResources", "privateLinkResources"),
		resourceids.UserSpecifiedSegment("privateLinkResourceName", "privateLinkResourceName"),
	}
}

// String returns a human-readable description of this Workspace Private Link Resource ID
func (id WorkspacePrivateLinkResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Private Link Resource Name: %q", id.PrivateLinkResourceName),
	}
	return fmt.Sprintf("Workspace Private Link Resource (%s)", strings.Join(components, "\n"))
}
