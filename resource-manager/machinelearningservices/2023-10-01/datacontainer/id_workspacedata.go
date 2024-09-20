package datacontainer

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&WorkspaceDataId{})
}

var _ resourceids.ResourceId = &WorkspaceDataId{}

// WorkspaceDataId is a struct representing the Resource ID for a Workspace Data
type WorkspaceDataId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	DataName          string
}

// NewWorkspaceDataID returns a new WorkspaceDataId struct
func NewWorkspaceDataID(subscriptionId string, resourceGroupName string, workspaceName string, dataName string) WorkspaceDataId {
	return WorkspaceDataId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		DataName:          dataName,
	}
}

// ParseWorkspaceDataID parses 'input' into a WorkspaceDataId
func ParseWorkspaceDataID(input string) (*WorkspaceDataId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceDataId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceDataId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWorkspaceDataIDInsensitively parses 'input' case-insensitively into a WorkspaceDataId
// note: this method should only be used for API response data and not user input
func ParseWorkspaceDataIDInsensitively(input string) (*WorkspaceDataId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WorkspaceDataId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WorkspaceDataId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WorkspaceDataId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DataName, ok = input.Parsed["dataName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataName", input)
	}

	return nil
}

// ValidateWorkspaceDataID checks that 'input' can be parsed as a Workspace Data ID
func ValidateWorkspaceDataID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWorkspaceDataID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Workspace Data ID
func (id WorkspaceDataId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/data/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.DataName)
}

// Segments returns a slice of Resource ID Segments which comprise this Workspace Data ID
func (id WorkspaceDataId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticData", "data", "data"),
		resourceids.UserSpecifiedSegment("dataName", "name"),
	}
}

// String returns a human-readable description of this Workspace Data ID
func (id WorkspaceDataId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Data Name: %q", id.DataName),
	}
	return fmt.Sprintf("Workspace Data (%s)", strings.Join(components, "\n"))
}
