package productgrouplink

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProductGroupLinkId{}

// ProductGroupLinkId is a struct representing the Resource ID for a Product Group Link
type ProductGroupLinkId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	WorkspaceId       string
	ProductId         string
	GroupLinkId       string
}

// NewProductGroupLinkID returns a new ProductGroupLinkId struct
func NewProductGroupLinkID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, productId string, groupLinkId string) ProductGroupLinkId {
	return ProductGroupLinkId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		WorkspaceId:       workspaceId,
		ProductId:         productId,
		GroupLinkId:       groupLinkId,
	}
}

// ParseProductGroupLinkID parses 'input' into a ProductGroupLinkId
func ParseProductGroupLinkID(input string) (*ProductGroupLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProductGroupLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProductGroupLinkId{}

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

	if id.GroupLinkId, ok = parsed.Parsed["groupLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupLinkId", *parsed)
	}

	return &id, nil
}

// ParseProductGroupLinkIDInsensitively parses 'input' case-insensitively into a ProductGroupLinkId
// note: this method should only be used for API response data and not user input
func ParseProductGroupLinkIDInsensitively(input string) (*ProductGroupLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProductGroupLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProductGroupLinkId{}

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

	if id.GroupLinkId, ok = parsed.Parsed["groupLinkId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "groupLinkId", *parsed)
	}

	return &id, nil
}

// ValidateProductGroupLinkID checks that 'input' can be parsed as a Product Group Link ID
func ValidateProductGroupLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProductGroupLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Product Group Link ID
func (id ProductGroupLinkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/products/%s/groupLinks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, id.ProductId, id.GroupLinkId)
}

// Segments returns a slice of Resource ID Segments which comprise this Product Group Link ID
func (id ProductGroupLinkId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticGroupLinks", "groupLinks", "groupLinks"),
		resourceids.UserSpecifiedSegment("groupLinkId", "groupLinkIdValue"),
	}
}

// String returns a human-readable description of this Product Group Link ID
func (id ProductGroupLinkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Product: %q", id.ProductId),
		fmt.Sprintf("Group Link: %q", id.GroupLinkId),
	}
	return fmt.Sprintf("Product Group Link (%s)", strings.Join(components, "\n"))
}
