package contentproductpackages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ContentProductPackageId{}

// ContentProductPackageId is a struct representing the Resource ID for a Content Product Package
type ContentProductPackageId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	PackageId         string
}

// NewContentProductPackageID returns a new ContentProductPackageId struct
func NewContentProductPackageID(subscriptionId string, resourceGroupName string, workspaceName string, packageId string) ContentProductPackageId {
	return ContentProductPackageId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		PackageId:         packageId,
	}
}

// ParseContentProductPackageID parses 'input' into a ContentProductPackageId
func ParseContentProductPackageID(input string) (*ContentProductPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContentProductPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentProductPackageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseContentProductPackageIDInsensitively parses 'input' case-insensitively into a ContentProductPackageId
// note: this method should only be used for API response data and not user input
func ParseContentProductPackageIDInsensitively(input string) (*ContentProductPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContentProductPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentProductPackageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ContentProductPackageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.PackageId, ok = input.Parsed["packageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "packageId", input)
	}

	return nil
}

// ValidateContentProductPackageID checks that 'input' can be parsed as a Content Product Package ID
func ValidateContentProductPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseContentProductPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Content Product Package ID
func (id ContentProductPackageId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/contentProductPackages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.PackageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Content Product Package ID
func (id ContentProductPackageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticContentProductPackages", "contentProductPackages", "contentProductPackages"),
		resourceids.UserSpecifiedSegment("packageId", "packageIdValue"),
	}
}

// String returns a human-readable description of this Content Product Package ID
func (id ContentProductPackageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Package: %q", id.PackageId),
	}
	return fmt.Sprintf("Content Product Package (%s)", strings.Join(components, "\n"))
}
