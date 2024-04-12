package contentpackages

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ContentPackageId{})
}

var _ resourceids.ResourceId = &ContentPackageId{}

// ContentPackageId is a struct representing the Resource ID for a Content Package
type ContentPackageId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	PackageId         string
}

// NewContentPackageID returns a new ContentPackageId struct
func NewContentPackageID(subscriptionId string, resourceGroupName string, workspaceName string, packageId string) ContentPackageId {
	return ContentPackageId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		PackageId:         packageId,
	}
}

// ParseContentPackageID parses 'input' into a ContentPackageId
func ParseContentPackageID(input string) (*ContentPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContentPackageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentPackageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseContentPackageIDInsensitively parses 'input' case-insensitively into a ContentPackageId
// note: this method should only be used for API response data and not user input
func ParseContentPackageIDInsensitively(input string) (*ContentPackageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ContentPackageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ContentPackageId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ContentPackageId) FromParseResult(input resourceids.ParseResult) error {
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

// ValidateContentPackageID checks that 'input' can be parsed as a Content Package ID
func ValidateContentPackageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseContentPackageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Content Package ID
func (id ContentPackageId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/contentPackages/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.PackageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Content Package ID
func (id ContentPackageId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticContentPackages", "contentPackages", "contentPackages"),
		resourceids.UserSpecifiedSegment("packageId", "packageIdValue"),
	}
}

// String returns a human-readable description of this Content Package ID
func (id ContentPackageId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Package: %q", id.PackageId),
	}
	return fmt.Sprintf("Content Package (%s)", strings.Join(components, "\n"))
}
