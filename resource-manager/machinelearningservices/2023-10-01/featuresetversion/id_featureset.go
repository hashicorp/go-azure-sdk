package featuresetversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FeatureSetId{}

// FeatureSetId is a struct representing the Resource ID for a Feature Set
type FeatureSetId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	FeatureSetName    string
}

// NewFeatureSetID returns a new FeatureSetId struct
func NewFeatureSetID(subscriptionId string, resourceGroupName string, workspaceName string, featureSetName string) FeatureSetId {
	return FeatureSetId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		FeatureSetName:    featureSetName,
	}
}

// ParseFeatureSetID parses 'input' into a FeatureSetId
func ParseFeatureSetID(input string) (*FeatureSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(FeatureSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FeatureSetId{}

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

	return &id, nil
}

// ParseFeatureSetIDInsensitively parses 'input' case-insensitively into a FeatureSetId
// note: this method should only be used for API response data and not user input
func ParseFeatureSetIDInsensitively(input string) (*FeatureSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(FeatureSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FeatureSetId{}

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

	return &id, nil
}

// ValidateFeatureSetID checks that 'input' can be parsed as a Feature Set ID
func ValidateFeatureSetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFeatureSetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Feature Set ID
func (id FeatureSetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/featureSets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.FeatureSetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Feature Set ID
func (id FeatureSetId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Feature Set ID
func (id FeatureSetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Feature Set Name: %q", id.FeatureSetName),
	}
	return fmt.Sprintf("Feature Set (%s)", strings.Join(components, "\n"))
}
