package tagoperationlink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TagOperationLinkId{}

// TagOperationLinkId is a struct representing the Resource ID for a Tag Operation Link
type TagOperationLinkId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	WorkspaceId       string
	TagId             string
	OperationLinkId   string
}

// NewTagOperationLinkID returns a new TagOperationLinkId struct
func NewTagOperationLinkID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, tagId string, operationLinkId string) TagOperationLinkId {
	return TagOperationLinkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		WorkspaceId:       workspaceId,
		TagId:             tagId,
		OperationLinkId:   operationLinkId,
	}
}

// ParseTagOperationLinkID parses 'input' into a TagOperationLinkId
func ParseTagOperationLinkID(input string) (*TagOperationLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(TagOperationLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TagOperationLinkId{}

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

	if id.OperationLinkId, ok = parsed.Parsed["operationLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationLinkId", *parsed)
	}

	return &id, nil
}

// ParseTagOperationLinkIDInsensitively parses 'input' case-insensitively into a TagOperationLinkId
// note: this method should only be used for API response data and not user input
func ParseTagOperationLinkIDInsensitively(input string) (*TagOperationLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(TagOperationLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TagOperationLinkId{}

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

	if id.OperationLinkId, ok = parsed.Parsed["operationLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationLinkId", *parsed)
	}

	return &id, nil
}

// ValidateTagOperationLinkID checks that 'input' can be parsed as a Tag Operation Link ID
func ValidateTagOperationLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTagOperationLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Tag Operation Link ID
func (id TagOperationLinkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/tags/%s/operationLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, id.TagId, id.OperationLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Tag Operation Link ID
func (id TagOperationLinkId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticOperationLinks", "operationLinks", "operationLinks"),
		resourceids.UserSpecifiedSegment("operationLinkId", "operationLinkIdValue"),
	}
}

// String returns a human-readable description of this Tag Operation Link ID
func (id TagOperationLinkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Tag: %q", id.TagId),
		fmt.Sprintf("Operation Link: %q", id.OperationLinkId),
	}
	return fmt.Sprintf("Tag Operation Link (%s)", strings.Join(components, "\n"))
}
