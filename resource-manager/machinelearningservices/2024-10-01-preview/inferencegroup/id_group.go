package inferencegroup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&GroupId{})
}

var _ resourceids.ResourceId = &GroupId{}

// GroupId is a struct representing the Resource ID for a Group
type GroupId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	InferencePoolName string
	GroupName         string
}

// NewGroupID returns a new GroupId struct
func NewGroupID(subscriptionId string, resourceGroupName string, workspaceName string, inferencePoolName string, groupName string) GroupId {
	return GroupId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		InferencePoolName: inferencePoolName,
		GroupName:         groupName,
	}
}

// ParseGroupID parses 'input' into a GroupId
func ParseGroupID(input string) (*GroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIDInsensitively parses 'input' case-insensitively into a GroupId
// note: this method should only be used for API response data and not user input
func ParseGroupIDInsensitively(input string) (*GroupId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.InferencePoolName, ok = input.Parsed["inferencePoolName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "inferencePoolName", input)
	}

	if id.GroupName, ok = input.Parsed["groupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupName", input)
	}

	return nil
}

// ValidateGroupID checks that 'input' can be parsed as a Group ID
func ValidateGroupID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group ID
func (id GroupId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/inferencePools/%s/groups/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.InferencePoolName, id.GroupName)
}

// Segments returns a slice of Resource ID Segments which comprise this Group ID
func (id GroupId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticInferencePools", "inferencePools", "inferencePools"),
		resourceids.UserSpecifiedSegment("inferencePoolName", "inferencePoolName"),
		resourceids.StaticSegment("staticGroups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupName", "groupName"),
	}
}

// String returns a human-readable description of this Group ID
func (id GroupId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Inference Pool Name: %q", id.InferencePoolName),
		fmt.Sprintf("Group Name: %q", id.GroupName),
	}
	return fmt.Sprintf("Group (%s)", strings.Join(components, "\n"))
}
