package featuresetversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&FeatureSetVersionId{})
}

var _ resourceids.ResourceId = &FeatureSetVersionId{}

// FeatureSetVersionId is a struct representing the Resource ID for a Feature Set Version
type FeatureSetVersionId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	FeatureSetName    string
	VersionName       string
}

// NewFeatureSetVersionID returns a new FeatureSetVersionId struct
func NewFeatureSetVersionID(subscriptionId string, resourceGroupName string, workspaceName string, featureSetName string, versionName string) FeatureSetVersionId {
	return FeatureSetVersionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		FeatureSetName:    featureSetName,
		VersionName:       versionName,
	}
}

// ParseFeatureSetVersionID parses 'input' into a FeatureSetVersionId
func ParseFeatureSetVersionID(input string) (*FeatureSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureSetVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFeatureSetVersionIDInsensitively parses 'input' case-insensitively into a FeatureSetVersionId
// note: this method should only be used for API response data and not user input
func ParseFeatureSetVersionIDInsensitively(input string) (*FeatureSetVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureSetVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureSetVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FeatureSetVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.FeatureSetName, ok = input.Parsed["featureSetName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "featureSetName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateFeatureSetVersionID checks that 'input' can be parsed as a Feature Set Version ID
func ValidateFeatureSetVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFeatureSetVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Feature Set Version ID
func (id FeatureSetVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/featureSets/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.FeatureSetName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Feature Set Version ID
func (id FeatureSetVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticFeatureSets", "featureSets", "featureSets"),
		resourceids.UserSpecifiedSegment("featureSetName", "featureSetName"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Feature Set Version ID
func (id FeatureSetVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Feature Set Name: %q", id.FeatureSetName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Feature Set Version (%s)", strings.Join(components, "\n"))
}
