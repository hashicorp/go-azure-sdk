package componentversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ComponentVersionId{})
}

var _ resourceids.ResourceId = &ComponentVersionId{}

// ComponentVersionId is a struct representing the Resource ID for a Component Version
type ComponentVersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	ComponentName     string
	VersionName       string
}

// NewComponentVersionID returns a new ComponentVersionId struct
func NewComponentVersionID(subscriptionId string, resourceGroupName string, workspaceName string, componentName string, versionName string) ComponentVersionId {
	return ComponentVersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		ComponentName:     componentName,
		VersionName:       versionName,
	}
}

// ParseComponentVersionID parses 'input' into a ComponentVersionId
func ParseComponentVersionID(input string) (*ComponentVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ComponentVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ComponentVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseComponentVersionIDInsensitively parses 'input' case-insensitively into a ComponentVersionId
// note: this method should only be used for API response data and not user input
func ParseComponentVersionIDInsensitively(input string) (*ComponentVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ComponentVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ComponentVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ComponentVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ComponentName, ok = input.Parsed["componentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "componentName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateComponentVersionID checks that 'input' can be parsed as a Component Version ID
func ValidateComponentVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseComponentVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Component Version ID
func (id ComponentVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/components/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.ComponentName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Component Version ID
func (id ComponentVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticComponents", "components", "components"),
		resourceids.UserSpecifiedSegment("componentName", "componentName"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Component Version ID
func (id ComponentVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Component Name: %q", id.ComponentName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Component Version (%s)", strings.Join(components, "\n"))
}
