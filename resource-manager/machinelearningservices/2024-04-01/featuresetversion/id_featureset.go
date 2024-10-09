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
	recaser.RegisterResourceId(&FeatureSetId{})
}

var _ resourceids.ResourceId = &FeatureSetId{}

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
	parser := resourceids.NewParserFromResourceIdType(&FeatureSetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFeatureSetIDInsensitively parses 'input' case-insensitively into a FeatureSetId
// note: this method should only be used for API response data and not user input
func ParseFeatureSetIDInsensitively(input string) (*FeatureSetId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureSetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureSetId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FeatureSetId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
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
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticFeatureSets", "featureSets", "featureSets"),
		resourceids.UserSpecifiedSegment("featureSetName", "featureSetName"),
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
