package deploymentoperations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ProviderDeploymentId{}

// ProviderDeploymentId is a struct representing the Resource ID for a Provider Deployment
type ProviderDeploymentId struct {
	SubscriptionId string
	DeploymentName string
}

// NewProviderDeploymentID returns a new ProviderDeploymentId struct
func NewProviderDeploymentID(subscriptionId string, deploymentName string) ProviderDeploymentId {
	return ProviderDeploymentId{
		SubscriptionId: subscriptionId,
		DeploymentName: deploymentName,
	}
}

// ParseProviderDeploymentID parses 'input' into a ProviderDeploymentId
func ParseProviderDeploymentID(input string) (*ProviderDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderDeploymentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderDeploymentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, fmt.Errorf("the segment 'deploymentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseProviderDeploymentIDInsensitively parses 'input' case-insensitively into a ProviderDeploymentId
// note: this method should only be used for API response data and not user input
func ParseProviderDeploymentIDInsensitively(input string) (*ProviderDeploymentId, error) {
	parser := resourceids.NewParserFromResourceIdType(ProviderDeploymentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ProviderDeploymentId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.DeploymentName, ok = parsed.Parsed["deploymentName"]; !ok {
		return nil, fmt.Errorf("the segment 'deploymentName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateProviderDeploymentID checks that 'input' can be parsed as a Provider Deployment ID
func ValidateProviderDeploymentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviderDeploymentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Provider Deployment ID
func (id ProviderDeploymentId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Resources/deployments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.DeploymentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Provider Deployment ID
func (id ProviderDeploymentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftResources", "Microsoft.Resources", "Microsoft.Resources"),
		resourceids.StaticSegment("staticDeployments", "deployments", "deployments"),
		resourceids.UserSpecifiedSegment("deploymentName", "deploymentValue"),
	}
}

// String returns a human-readable description of this Provider Deployment ID
func (id ProviderDeploymentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Deployment Name: %q", id.DeploymentName),
	}
	return fmt.Sprintf("Provider Deployment (%s)", strings.Join(components, "\n"))
}
