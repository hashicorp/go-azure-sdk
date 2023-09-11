package namedvalue

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkspaceNamedValueId{}

// WorkspaceNamedValueId is a struct representing the Resource ID for a Workspace Named Value
type WorkspaceNamedValueId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	WorkspaceId       string
	NamedValueId      string
}

// NewWorkspaceNamedValueID returns a new WorkspaceNamedValueId struct
func NewWorkspaceNamedValueID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, namedValueId string) WorkspaceNamedValueId {
	return WorkspaceNamedValueId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		WorkspaceId:       workspaceId,
		NamedValueId:      namedValueId,
	}
}

// ParseWorkspaceNamedValueID parses 'input' into a WorkspaceNamedValueId
func ParseWorkspaceNamedValueID(input string) (*WorkspaceNamedValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspaceNamedValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspaceNamedValueId{}

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

	if id.NamedValueId, ok = parsed.Parsed["namedValueId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "namedValueId", *parsed)
	}

	return &id, nil
}

// ParseWorkspaceNamedValueIDInsensitively parses 'input' case-insensitively into a WorkspaceNamedValueId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceNamedValueIDInsensitively(input string) (*WorkspaceNamedValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspaceNamedValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspaceNamedValueId{}

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

	if id.NamedValueId, ok = parsed.Parsed["namedValueId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "namedValueId", *parsed)
	}

	return &id, nil
}

// ValidateWorkspaceNamedValueID checks that 'input' can be parsed as a Workspace Named Value ID
func ValidateWorkspaceNamedValueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceNamedValueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Named Value ID
func (id WorkspaceNamedValueId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/namedValues/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, id.NamedValueId)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Named Value ID
func (id WorkspaceNamedValueId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticNamedValues", "namedValues", "namedValues"),
		resourceids.UserSpecifiedSegment("namedValueId", "namedValueIdValue"),
	}
}

// String returns a human-readable description of this Workspace Named Value ID
func (id WorkspaceNamedValueId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Named Value: %q", id.NamedValueId),
	}
	return fmt.Sprintf("Workspace Named Value (%s)", strings.Join(components, "\n"))
}
