package tagproductlink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkspaceTagId{}

// WorkspaceTagId is a struct representing the Resource ID for a Workspace Tag
type WorkspaceTagId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	WorkspaceId       string
	TagId             string
}

// NewWorkspaceTagID returns a new WorkspaceTagId struct
func NewWorkspaceTagID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, tagId string) WorkspaceTagId {
	return WorkspaceTagId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		WorkspaceId:       workspaceId,
		TagId:             tagId,
	}
}

// ParseWorkspaceTagID parses 'input' into a WorkspaceTagId
func ParseWorkspaceTagID(input string) (*WorkspaceTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspaceTagId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspaceTagId{}

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

	if id.TagId, ok = parsed.Parsed["tagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tagId", *parsed)
	}

	return &id, nil
}

// ParseWorkspaceTagIDInsensitively parses 'input' case-insensitively into a WorkspaceTagId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceTagIDInsensitively(input string) (*WorkspaceTagId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspaceTagId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspaceTagId{}

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

	if id.TagId, ok = parsed.Parsed["tagId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "tagId", *parsed)
	}

	return &id, nil
}

// ValidateWorkspaceTagID checks that 'input' can be parsed as a Workspace Tag ID
func ValidateWorkspaceTagID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceTagID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Tag ID
func (id WorkspaceTagId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/tags/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, id.TagId)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Tag ID
func (id WorkspaceTagId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticTags", "tags", "tags"),
		resourceids.UserSpecifiedSegment("tagId", "tagIdValue"),
	}
}

// String returns a human-readable description of this Workspace Tag ID
func (id WorkspaceTagId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Tag: %q", id.TagId),
	}
	return fmt.Sprintf("Workspace Tag (%s)", strings.Join(components, "\n"))
}
