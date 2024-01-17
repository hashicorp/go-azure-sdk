package modelcontainer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ModelId{}

// ModelId is a struct representing the Resource ID for a Model
type ModelId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	ModelName         string
}

// NewModelID returns a new ModelId struct
func NewModelID(subscriptionId string, resourceGroupName string, workspaceName string, modelName string) ModelId {
	return ModelId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		ModelName:         modelName,
	}
}

// ParseModelID parses 'input' into a ModelId
func ParseModelID(input string) (*ModelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModelId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseModelIDInsensitively parses 'input' case-insensitively into a ModelId
// note: this method should only be used for API response data and not user input
func ParseModelIDInsensitively(input string) (*ModelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ModelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ModelId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ModelId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ModelName, ok = input.Parsed["modelName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "modelName", input)
	}

	return nil
}

// ValidateModelID checks that 'input' can be parsed as a Model ID
func ValidateModelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseModelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Model ID
func (id ModelId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/models/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.ModelName)
}

// Segments returns a slice of Resource ID Segments which comprise this Model ID
func (id ModelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticModels", "models", "models"),
		resourceids.UserSpecifiedSegment("modelName", "modelValue"),
	}
}

// String returns a human-readable description of this Model ID
func (id ModelId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Model Name: %q", id.ModelName),
	}
	return fmt.Sprintf("Model (%s)", strings.Join(components, "\n"))
}
