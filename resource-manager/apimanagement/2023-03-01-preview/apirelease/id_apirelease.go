package apirelease

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApiReleaseId{}

// ApiReleaseId is a struct representing the Resource ID for a Api Release
type ApiReleaseId struct {
	SubscriptionId    string
	ResourceGroupName string
	ServiceName       string
	WorkspaceId       string
	ApiId             string
	ReleaseId         string
}

// NewApiReleaseID returns a new ApiReleaseId struct
func NewApiReleaseID(subscriptionId string, resourceGroupName string, serviceName string, workspaceId string, apiId string, releaseId string) ApiReleaseId {
	return ApiReleaseId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ServiceName:       serviceName,
		WorkspaceId:       workspaceId,
		ApiId:             apiId,
		ReleaseId:         releaseId,
	}
}

// ParseApiReleaseID parses 'input' into a ApiReleaseId
func ParseApiReleaseID(input string) (*ApiReleaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApiReleaseId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApiReleaseId{}

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

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "apiId", *parsed)
	}

	if id.ReleaseId, ok = parsed.Parsed["releaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "releaseId", *parsed)
	}

	return &id, nil
}

// ParseApiReleaseIDInsensitively parses 'input' case-insensitively into a ApiReleaseId
// note: this method should only be used for API response data and not user input
func ParseApiReleaseIDInsensitively(input string) (*ApiReleaseId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApiReleaseId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApiReleaseId{}

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

	if id.ApiId, ok = parsed.Parsed["apiId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "apiId", *parsed)
	}

	if id.ReleaseId, ok = parsed.Parsed["releaseId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "releaseId", *parsed)
	}

	return &id, nil
}

// ValidateApiReleaseID checks that 'input' can be parsed as a Api Release ID
func ValidateApiReleaseID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApiReleaseID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Api Release ID
func (id ApiReleaseId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ApiManagement/service/%s/workspaces/%s/apis/%s/releases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ServiceName, id.WorkspaceId, id.ApiId, id.ReleaseId)
}

// Segments returns a slice of Resource ID Segments which comprise this Api Release ID
func (id ApiReleaseId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticApis", "apis", "apis"),
		resourceids.UserSpecifiedSegment("apiId", "apiIdValue"),
		resourceids.StaticSegment("staticReleases", "releases", "releases"),
		resourceids.UserSpecifiedSegment("releaseId", "releaseIdValue"),
	}
}

// String returns a human-readable description of this Api Release ID
func (id ApiReleaseId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Service Name: %q", id.ServiceName),
		fmt.Sprintf("Workspace: %q", id.WorkspaceId),
		fmt.Sprintf("Api: %q", id.ApiId),
		fmt.Sprintf("Release: %q", id.ReleaseId),
	}
	return fmt.Sprintf("Api Release (%s)", strings.Join(components, "\n"))
}
