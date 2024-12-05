package inferenceendpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&InferencePoolId{})
}

var _ resourceids.ResourceId = &InferencePoolId{}

// InferencePoolId is a struct representing the Resource ID for a Inference Pool
type InferencePoolId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	InferencePoolName string
}

// NewInferencePoolID returns a new InferencePoolId struct
func NewInferencePoolID(subscriptionId string, resourceGroupName string, workspaceName string, inferencePoolName string) InferencePoolId {
	return InferencePoolId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		InferencePoolName: inferencePoolName,
	}
}

// ParseInferencePoolID parses 'input' into a InferencePoolId
func ParseInferencePoolID(input string) (*InferencePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InferencePoolId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InferencePoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInferencePoolIDInsensitively parses 'input' case-insensitively into a InferencePoolId
// note: this method should only be used for API response data and not user input
func ParseInferencePoolIDInsensitively(input string) (*InferencePoolId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InferencePoolId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InferencePoolId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InferencePoolId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateInferencePoolID checks that 'input' can be parsed as a Inference Pool ID
func ValidateInferencePoolID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInferencePoolID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Inference Pool ID
func (id InferencePoolId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/inferencePools/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.InferencePoolName)
}

// Segments returns a slice of Resource ID Segments which comprise this Inference Pool ID
func (id InferencePoolId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Inference Pool ID
func (id InferencePoolId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Inference Pool Name: %q", id.InferencePoolName),
	}
	return fmt.Sprintf("Inference Pool (%s)", strings.Join(components, "\n"))
}
