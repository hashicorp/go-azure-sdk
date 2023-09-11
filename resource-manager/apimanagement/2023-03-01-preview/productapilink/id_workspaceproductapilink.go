package productapilink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WorkspaceProductApiLinkId{}

// WorkspaceProductApiLinkId is a struct representing the Resource ID for a Workspace Product Api Link
type WorkspaceProductApiLinkId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	WorkspaceId       string
	ProductId         string
	ApiLinkId         string
}

// NewWorkspaceProductApiLinkID returns a new WorkspaceProductApiLinkId struct
func NewWorkspaceProductApiLinkID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, productId string, apiLinkId string) WorkspaceProductApiLinkId {
	return WorkspaceProductApiLinkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		WorkspaceId:       workspaceId,
		ProductId:         productId,
		ApiLinkId:         apiLinkId,
	}
}

// ParseWorkspaceProductApiLinkID parses 'input' into a WorkspaceProductApiLinkId
func ParseWorkspaceProductApiLinkID(input string) (*WorkspaceProductApiLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspaceProductApiLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspaceProductApiLinkId{}

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

	if id.ProductId, ok = parsed.Parsed["productId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "productId", *parsed)
	}

	if id.ApiLinkId, ok = parsed.Parsed["apiLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "apiLinkId", *parsed)
	}

	return &id, nil
}

// ParseWorkspaceProductApiLinkIDInsensitively parses 'input' case-insensitively into a WorkspaceProductApiLinkId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceProductApiLinkIDInsensitively(input string) (*WorkspaceProductApiLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(WorkspaceProductApiLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WorkspaceProductApiLinkId{}

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

	if id.ProductId, ok = parsed.Parsed["productId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "productId", *parsed)
	}

	if id.ApiLinkId, ok = parsed.Parsed["apiLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "apiLinkId", *parsed)
	}

	return &id, nil
}

// ValidateWorkspaceProductApiLinkID checks that 'input' can be parsed as a Workspace Product Api Link ID
func ValidateWorkspaceProductApiLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceProductApiLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Product Api Link ID
func (id WorkspaceProductApiLinkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/products/%s/apiLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, id.ProductId, id.ApiLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Product Api Link ID
func (id WorkspaceProductApiLinkId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticProducts", "products", "products"),
		resourceids.UserSpecifiedSegment("productId", "productIdValue"),
		resourceids.StaticSegment("staticApiLinks", "apiLinks", "apiLinks"),
		resourceids.UserSpecifiedSegment("apiLinkId", "apiLinkIdValue"),
	}
}

// String returns a human-readable description of this Workspace Product Api Link ID
func (id WorkspaceProductApiLinkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Product: %q", id.ProductId),
		fmt.Sprintf("Api Link: %q", id.ApiLinkId),
	}
	return fmt.Sprintf("Workspace Product Api Link (%s)", strings.Join(components, "\n"))
}
