package fluxconfigurationoperationstatus

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FluxConfigurationOperationId{}

// FluxConfigurationOperationId is a struct representing the Resource ID for a Flux Configuration Operation
type FluxConfigurationOperationId struct {
	SubscriptionId        string
	ResourceGroupName     string
	ProviderName          string
	ClusterResourceName   string
	ClusterName           string
	FluxConfigurationName string
	OperationId           string
}

// NewFluxConfigurationOperationID returns a new FluxConfigurationOperationId struct
func NewFluxConfigurationOperationID(subscriptionId string, resourceGroupName string, providerName string, clusterResourceName string, clusterName string, fluxConfigurationName string, operationId string) FluxConfigurationOperationId {
	return FluxConfigurationOperationId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		ProviderName:          providerName,
		ClusterResourceName:   clusterResourceName,
		ClusterName:           clusterName,
		FluxConfigurationName: fluxConfigurationName,
		OperationId:           operationId,
	}
}

// ParseFluxConfigurationOperationID parses 'input' into a FluxConfigurationOperationId
func ParseFluxConfigurationOperationID(input string) (*FluxConfigurationOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(FluxConfigurationOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FluxConfigurationOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ClusterResourceName, ok = parsed.Parsed["clusterResourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterResourceName", *parsed)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterName", *parsed)
	}

	if id.FluxConfigurationName, ok = parsed.Parsed["fluxConfigurationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fluxConfigurationName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseFluxConfigurationOperationIDInsensitively parses 'input' case-insensitively into a FluxConfigurationOperationId
// note: this method should only be used for API response data and not user input
func ParseFluxConfigurationOperationIDInsensitively(input string) (*FluxConfigurationOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(FluxConfigurationOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FluxConfigurationOperationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ClusterResourceName, ok = parsed.Parsed["clusterResourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterResourceName", *parsed)
	}

	if id.ClusterName, ok = parsed.Parsed["clusterName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "clusterName", *parsed)
	}

	if id.FluxConfigurationName, ok = parsed.Parsed["fluxConfigurationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "fluxConfigurationName", *parsed)
	}

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateFluxConfigurationOperationID checks that 'input' can be parsed as a Flux Configuration Operation ID
func ValidateFluxConfigurationOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFluxConfigurationOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Flux Configuration Operation ID
func (id FluxConfigurationOperationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.KubernetesConfiguration/fluxConfigurations/%s/operations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ClusterResourceName, id.ClusterName, id.FluxConfigurationName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Flux Configuration Operation ID
func (id FluxConfigurationOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("clusterResourceName", "clusterResourceValue"),
		resourceids.UserSpecifiedSegment("clusterName", "clusterValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftKubernetesConfiguration", "Microsoft.KubernetesConfiguration", "Microsoft.KubernetesConfiguration"),
		resourceids.StaticSegment("staticFluxConfigurations", "fluxConfigurations", "fluxConfigurations"),
		resourceids.UserSpecifiedSegment("fluxConfigurationName", "fluxConfigurationValue"),
		resourceids.StaticSegment("staticOperations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Flux Configuration Operation ID
func (id FluxConfigurationOperationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Cluster Resource Name: %q", id.ClusterResourceName),
		fmt.Sprintf("Cluster Name: %q", id.ClusterName),
		fmt.Sprintf("Flux Configuration Name: %q", id.FluxConfigurationName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Flux Configuration Operation (%s)", strings.Join(components, "\n"))
}
