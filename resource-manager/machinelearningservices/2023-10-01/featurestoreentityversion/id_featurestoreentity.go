package featurestoreentityversion

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &FeatureStoreEntityId{}

// FeatureStoreEntityId is a struct representing the Resource ID for a Feature Store Entity
type FeatureStoreEntityId struct {
	SubscriptionId         string
	ResourceGroupName      string
	WorkspaceName          string
	FeatureStoreEntityName string
}

// NewFeatureStoreEntityID returns a new FeatureStoreEntityId struct
func NewFeatureStoreEntityID(subscriptionId string, resourceGroupName string, workspaceName string, featureStoreEntityName string) FeatureStoreEntityId {
	return FeatureStoreEntityId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		WorkspaceName:          workspaceName,
		FeatureStoreEntityName: featureStoreEntityName,
	}
}

// ParseFeatureStoreEntityID parses 'input' into a FeatureStoreEntityId
func ParseFeatureStoreEntityID(input string) (*FeatureStoreEntityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureStoreEntityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureStoreEntityId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseFeatureStoreEntityIDInsensitively parses 'input' case-insensitively into a FeatureStoreEntityId
// note: this method should only be used for API response data and not user input
func ParseFeatureStoreEntityIDInsensitively(input string) (*FeatureStoreEntityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&FeatureStoreEntityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := FeatureStoreEntityId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *FeatureStoreEntityId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateFeatureStoreEntityID checks that 'input' can be parsed as a Feature Store Entity ID
func ValidateFeatureStoreEntityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFeatureStoreEntityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Feature Store Entity ID
func (id FeatureStoreEntityId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/featureStoreEntities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.FeatureStoreEntityName)
}

// Segments returns a slice of Resource ID Segments which comprise this Feature Store Entity ID
func (id FeatureStoreEntityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticFeatureStoreEntities", "featureStoreEntities", "featureStoreEntities"),
		resourceids.UserSpecifiedSegment("featureStoreEntityName", "featureStoreEntityValue"),
	}
}

// String returns a human-readable description of this Feature Store Entity ID
func (id FeatureStoreEntityId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Feature Store Entity Name: %q", id.FeatureStoreEntityName),
	}
	return fmt.Sprintf("Feature Store Entity (%s)", strings.Join(components, "\n"))
}
