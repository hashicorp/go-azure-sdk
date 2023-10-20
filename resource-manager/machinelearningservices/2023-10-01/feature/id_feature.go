package feature

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FeatureId{}

// FeatureId is a struct representing the Resource ID for a Feature
type FeatureId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	FeatureSetName    string
	VersionName       string
	FeatureName       string
}

// NewFeatureID returns a new FeatureId struct
func NewFeatureID(subscriptionId string, resourceGroupName string, workspaceName string, featureSetName string, versionName string, featureName string) FeatureId {
	return FeatureId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		FeatureSetName:    featureSetName,
		VersionName:       versionName,
		FeatureName:       featureName,
	}
}

// ParseFeatureID parses 'input' into a FeatureId
func ParseFeatureID(input string) (*FeatureId, error) {
	parser := resourceids.NewParserFromResourceIdType(FeatureId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FeatureId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.FeatureSetName, ok = parsed.Parsed["featureSetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureSetName", *parsed)
	}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionName", *parsed)
	}

	if id.FeatureName, ok = parsed.Parsed["featureName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureName", *parsed)
	}

	return &id, nil
}

// ParseFeatureIDInsensitively parses 'input' case-insensitively into a FeatureId
// note: this method should only be used for API response data and not user input
func ParseFeatureIDInsensitively(input string) (*FeatureId, error) {
	parser := resourceids.NewParserFromResourceIdType(FeatureId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FeatureId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.FeatureSetName, ok = parsed.Parsed["featureSetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureSetName", *parsed)
	}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionName", *parsed)
	}

	if id.FeatureName, ok = parsed.Parsed["featureName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "featureName", *parsed)
	}

	return &id, nil
}

// ValidateFeatureID checks that 'input' can be parsed as a Feature ID
func ValidateFeatureID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFeatureID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Feature ID
func (id FeatureId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/featureSets/%s/versions/%s/features/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.FeatureSetName, id.VersionName, id.FeatureName)
}

// Segments returns a slice of Resource ID Segments which comprise this Feature ID
func (id FeatureId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticFeatureSets", "featureSets", "featureSets"),
		resourceids.UserSpecifiedSegment("featureSetName", "featureSetValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionValue"),
		resourceids.StaticSegment("staticFeatures", "features", "features"),
		resourceids.UserSpecifiedSegment("featureName", "featureValue"),
	}
}

// String returns a human-readable description of this Feature ID
func (id FeatureId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Feature Set Name: %q", id.FeatureSetName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
		fmt.Sprintf("Feature Name: %q", id.FeatureName),
	}
	return fmt.Sprintf("Feature (%s)", strings.Join(components, "\n"))
}
