package policyfragment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkspacePolicyFragmentId{}

// WorkspacePolicyFragmentId is a struct representing the Resource ID for a Workspace Policy Fragment
type WorkspacePolicyFragmentId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ServiceName        string
	WorkspaceId        string
	PolicyFragmentName string
}

// NewWorkspacePolicyFragmentID returns a new WorkspacePolicyFragmentId struct
func NewWorkspacePolicyFragmentID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, policyFragmentName string) WorkspacePolicyFragmentId {
	return WorkspacePolicyFragmentId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ServiceName:        serviceName,
		WorkspaceId:        workspaceId,
		PolicyFragmentName: policyFragmentName,
	}
}

// ParseWorkspacePolicyFragmentID parses 'input' into a WorkspacePolicyFragmentId
func ParseWorkspacePolicyFragmentID(input string) (*WorkspacePolicyFragmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspacePolicyFragmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspacePolicyFragmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.WorkspaceId, ok = parsed.Parsed["workspaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceId", *parsed)
	}

	if id.PolicyFragmentName, ok = parsed.Parsed["policyFragmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policyFragmentName", *parsed)
	}

	return &id, nil
}

// ParseWorkspacePolicyFragmentIDInsensitively parses 'input' case-insensitively into a WorkspacePolicyFragmentId
// note: this method should only be used for API response data and not user input
func ParseWorkspacePolicyFragmentIDInsensitively(input string) (*WorkspacePolicyFragmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspacePolicyFragmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspacePolicyFragmentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ServiceName, ok = parsed.Parsed["serviceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "serviceName", *parsed)
	}

	if id.WorkspaceId, ok = parsed.Parsed["workspaceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceId", *parsed)
	}

	if id.PolicyFragmentName, ok = parsed.Parsed["policyFragmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policyFragmentName", *parsed)
	}

	return &id, nil
}

// ValidateWorkspacePolicyFragmentID checks that 'input' can be parsed as a Workspace Policy Fragment ID
func ValidateWorkspacePolicyFragmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspacePolicyFragmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Policy Fragment ID
func (id WorkspacePolicyFragmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/policyFragments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, id.PolicyFragmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Policy Fragment ID
func (id WorkspacePolicyFragmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftApiManagement", "Microsoft.ApiManagement", "Microsoft.ApiManagement"),
		resourceids.StaticSegment("staticService", "service", "service"),
		resourceids.UserSpecifiedSegment("serviceName", "serviceValue"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceId", "workspaceIdValue"),
		resourceids.StaticSegment("staticPolicyFragments", "policyFragments", "policyFragments"),
		resourceids.UserSpecifiedSegment("policyFragmentName", "policyFragmentValue"),
	}
}

// String returns a human-readable description of this Workspace Policy Fragment ID
func (id WorkspacePolicyFragmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Policy Fragment Name: %q", id.PolicyFragmentName),
	}
	return fmt.Sprintf("Workspace Policy Fragment (%s)", strings.Join(components, "\n"))
}
