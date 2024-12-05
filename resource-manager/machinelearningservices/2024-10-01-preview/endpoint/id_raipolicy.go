package endpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RaiPolicyId{})
}

var _ resourceids.ResourceId = &RaiPolicyId{}

// RaiPolicyId is a struct representing the Resource ID for a Rai Policy
type RaiPolicyId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	EndpointName      string
	RaiPolicyName     string
}

// NewRaiPolicyID returns a new RaiPolicyId struct
func NewRaiPolicyID(subscriptionId string, resourceGroupName string, workspaceName string, endpointName string, raiPolicyName string) RaiPolicyId {
	return RaiPolicyId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		EndpointName:      endpointName,
		RaiPolicyName:     raiPolicyName,
	}
}

// ParseRaiPolicyID parses 'input' into a RaiPolicyId
func ParseRaiPolicyID(input string) (*RaiPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaiPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaiPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRaiPolicyIDInsensitively parses 'input' case-insensitively into a RaiPolicyId
// note: this method should only be used for API response data and not user input
func ParseRaiPolicyIDInsensitively(input string) (*RaiPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RaiPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RaiPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RaiPolicyId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.EndpointName, ok = input.Parsed["endpointName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "endpointName", input)
	}

	if id.RaiPolicyName, ok = input.Parsed["raiPolicyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "raiPolicyName", input)
	}

	return nil
}

// ValidateRaiPolicyID checks that 'input' can be parsed as a Rai Policy ID
func ValidateRaiPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRaiPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Rai Policy ID
func (id RaiPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/endpoints/%s/raiPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.EndpointName, id.RaiPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Rai Policy ID
func (id RaiPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticEndpoints", "endpoints", "endpoints"),
		resourceids.UserSpecifiedSegment("endpointName", "endpointName"),
		resourceids.StaticSegment("staticRaiPolicies", "raiPolicies", "raiPolicies"),
		resourceids.UserSpecifiedSegment("raiPolicyName", "raiPolicyName"),
	}
}

// String returns a human-readable description of this Rai Policy ID
func (id RaiPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Endpoint Name: %q", id.EndpointName),
		fmt.Sprintf("Rai Policy Name: %q", id.RaiPolicyName),
	}
	return fmt.Sprintf("Rai Policy (%s)", strings.Join(components, "\n"))
}
