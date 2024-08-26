package deploymentoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DeploymentOperationId{})
}

var _ resourceids.ResourceId = &DeploymentOperationId{}

// DeploymentOperationId is a struct representing the Resource ID for a Deployment Operation
type DeploymentOperationId struct {
	SubscriptionId string
	DeploymentName string
	OperationId    string
}

// NewDeploymentOperationID returns a new DeploymentOperationId struct
func NewDeploymentOperationID(subscriptionId string, deploymentName string, operationId string) DeploymentOperationId {
	return DeploymentOperationId{
		SubscriptionId: subscriptionId,
		DeploymentName: deploymentName,
		OperationId:    operationId,
	}
}

// ParseDeploymentOperationID parses 'input' into a DeploymentOperationId
func ParseDeploymentOperationID(input string) (*DeploymentOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeploymentOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeploymentOperationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDeploymentOperationIDInsensitively parses 'input' case-insensitively into a DeploymentOperationId
// note: this method should only be used for API response data and not user input
func ParseDeploymentOperationIDInsensitively(input string) (*DeploymentOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DeploymentOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DeploymentOperationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DeploymentOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.DeploymentName, ok = input.Parsed["deploymentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "deploymentName", input)
	}

	if id.OperationId, ok = input.Parsed["operationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "operationId", input)
	}

	return nil
}

// ValidateDeploymentOperationID checks that 'input' can be parsed as a Deployment Operation ID
func ValidateDeploymentOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDeploymentOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Deployment Operation ID
func (id DeploymentOperationId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Resources/deployments/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.DeploymentName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Deployment Operation ID
func (id DeploymentOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftResources", "Microsoft.Resources", "Microsoft.Resources"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Deployment Operation ID
func (id DeploymentOperationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Deployment Operation (%s)", strings.Join(components, "\n"))
}
