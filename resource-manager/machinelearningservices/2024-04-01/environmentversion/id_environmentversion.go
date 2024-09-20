package environmentversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&EnvironmentVersionId{})
}

var _ resourceids.ResourceId = &EnvironmentVersionId{}

// EnvironmentVersionId is a struct representing the Resource ID for a Environment Version
type EnvironmentVersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	EnvironmentName   string
	VersionName       string
}

// NewEnvironmentVersionID returns a new EnvironmentVersionId struct
func NewEnvironmentVersionID(subscriptionId string, resourceGroupName string, workspaceName string, environmentName string, versionName string) EnvironmentVersionId {
	return EnvironmentVersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		EnvironmentName:   environmentName,
		VersionName:       versionName,
	}
}

// ParseEnvironmentVersionID parses 'input' into a EnvironmentVersionId
func ParseEnvironmentVersionID(input string) (*EnvironmentVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnvironmentVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnvironmentVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEnvironmentVersionIDInsensitively parses 'input' case-insensitively into a EnvironmentVersionId
// note: this method should only be used for API response data and not user input
func ParseEnvironmentVersionIDInsensitively(input string) (*EnvironmentVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnvironmentVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnvironmentVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EnvironmentVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EnvironmentName, ok = input.Parsed["environmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "environmentName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateEnvironmentVersionID checks that 'input' can be parsed as a Environment Version ID
func ValidateEnvironmentVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnvironmentVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Environment Version ID
func (id EnvironmentVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/environments/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EnvironmentName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Environment Version ID
func (id EnvironmentVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticEnvironments", "environments", "environments"),
		resourceids.UserSpecifiedSegment("environmentName", "name"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "version"),
	}
}

// String returns a human-readable description of this Environment Version ID
func (id EnvironmentVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Environment Name: %q", id.EnvironmentName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Environment Version (%s)", strings.Join(components, "\n"))
}
