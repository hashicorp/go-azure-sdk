package v2workspaceconnectionresource

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ConnectionRaiPolicyId{})
}

var _ resourceids.ResourceId = &ConnectionRaiPolicyId{}

// ConnectionRaiPolicyId is a struct representing the Resource ID for a Connection Rai Policy
type ConnectionRaiPolicyId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	ConnectionName    string
	RaiPolicyName     string
}

// NewConnectionRaiPolicyID returns a new ConnectionRaiPolicyId struct
func NewConnectionRaiPolicyID(subscriptionId string, resourceGroupName string, workspaceName string, connectionName string, raiPolicyName string) ConnectionRaiPolicyId {
	return ConnectionRaiPolicyId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		ConnectionName:    connectionName,
		RaiPolicyName:     raiPolicyName,
	}
}

// ParseConnectionRaiPolicyID parses 'input' into a ConnectionRaiPolicyId
func ParseConnectionRaiPolicyID(input string) (*ConnectionRaiPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConnectionRaiPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConnectionRaiPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseConnectionRaiPolicyIDInsensitively parses 'input' case-insensitively into a ConnectionRaiPolicyId
// note: this method should only be used for API response data and not user input
func ParseConnectionRaiPolicyIDInsensitively(input string) (*ConnectionRaiPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ConnectionRaiPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ConnectionRaiPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ConnectionRaiPolicyId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.ConnectionName, ok = input.Parsed["connectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "connectionName", input)
	}

	if id.RaiPolicyName, ok = input.Parsed["raiPolicyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "raiPolicyName", input)
	}

	return nil
}

// ValidateConnectionRaiPolicyID checks that 'input' can be parsed as a Connection Rai Policy ID
func ValidateConnectionRaiPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConnectionRaiPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Connection Rai Policy ID
func (id ConnectionRaiPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.MachineLearningServices/workspaces/%s/connections/%s/raiPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.ConnectionName, id.RaiPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Connection Rai Policy ID
func (id ConnectionRaiPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMachineLearningServices", "Microsoft.MachineLearningServices", "Microsoft.MachineLearningServices"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticConnections", "connections", "connections"),
		resourceids.UserSpecifiedSegment("connectionName", "connectionName"),
		resourceids.StaticSegment("staticRaiPolicies", "raiPolicies", "raiPolicies"),
		resourceids.UserSpecifiedSegment("raiPolicyName", "raiPolicyName"),
	}
}

// String returns a human-readable description of this Connection Rai Policy ID
func (id ConnectionRaiPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Connection Name: %q", id.ConnectionName),
		fmt.Sprintf("Rai Policy Name: %q", id.RaiPolicyName),
	}
	return fmt.Sprintf("Connection Rai Policy (%s)", strings.Join(components, "\n"))
}
