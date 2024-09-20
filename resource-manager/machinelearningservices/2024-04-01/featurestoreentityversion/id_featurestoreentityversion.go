package featurestoreentityversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&FeatureStoreEntityVersionId{})
}

var _ resourceids.ResourceId = &FeatureStoreEntityVersionId{}

// FeatureStoreEntityVersionId is a struct representing the Resource ID for a Feature Store Entity Version
type FeatureStoreEntityVersionId struct {
	SubscriptionId         string
	ResourceGroupName      string
	WorkspaceName          string
	FeatureStoreEntityName string
	VersionName            string
}

// NewFeatureStoreEntityVersionID returns a new FeatureStoreEntityVersionId struct
func NewFeatureStoreEntityVersionID(subscriptionId string, resourceGroupName string, workspaceName string, featureStoreEntityName string, versionName string) FeatureStoreEntityVersionId {
	return FeatureStoreEntityVersionId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		WorkspaceName:          workspaceName,
		FeatureStoreEntityName: featureStoreEntityName,
		VersionName:            versionName,
	}
}

// ParseFeatureStoreEntityVersionID parses 'input' into a FeatureStoreEntityVersionId
func ParseFeatureStoreEntityVersionID(input string) (*FeatureStoreEntityVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureStoreEntityVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureStoreEntityVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFeatureStoreEntityVersionIDInsensitively parses 'input' case-insensitively into a FeatureStoreEntityVersionId
// note: this method should only be used for API response data and not user input
func ParseFeatureStoreEntityVersionIDInsensitively(input string) (*FeatureStoreEntityVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureStoreEntityVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureStoreEntityVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FeatureStoreEntityVersionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.FeatureStoreEntityName, ok = input.Parsed["featureStoreEntityName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "featureStoreEntityName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateFeatureStoreEntityVersionID checks that 'input' can be parsed as a Feature Store Entity Version ID
func ValidateFeatureStoreEntityVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFeatureStoreEntityVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Feature Store Entity Version ID
func (id FeatureStoreEntityVersionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/featureStoreEntities/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.FeatureStoreEntityName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Feature Store Entity Version ID
func (id FeatureStoreEntityVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticFeatureStoreEntities", "featureStoreEntities", "featureStoreEntities"),
		resourceids.UserSpecifiedSegment("featureStoreEntityName", "name"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "version"),
	}
}

// String returns a human-readable description of this Feature Store Entity Version ID
func (id FeatureStoreEntityVersionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Feature Store Entity Name: %q", id.FeatureStoreEntityName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Feature Store Entity Version (%s)", strings.Join(components, "\n"))
}
