// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementassociation

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = ManagementAssociationId{}

// ManagementAssociationId is a struct representing the Resource ID for a Management Association
type ManagementAssociationId struct {
	SubscriptionId            string
	ResourceGroupName         string
	ProviderName              string
	ResourceType              string
	ResourceName              string
	ManagementAssociationName string
}

// NewManagementAssociationID returns a new ManagementAssociationId struct
func NewManagementAssociationID(subscriptionId string, resourceGroupName string, providerName string, resourceType string, resourceName string, managementAssociationName string) ManagementAssociationId {
	return ManagementAssociationId{
		SubscriptionId:            subscriptionId,
		ResourceGroupName:         resourceGroupName,
		ProviderName:              providerName,
		ResourceType:              resourceType,
		ResourceName:              resourceName,
		ManagementAssociationName: managementAssociationName,
	}
}

// ParseManagementAssociationID parses 'input' into a ManagementAssociationId
func ParseManagementAssociationID(input string) (*ManagementAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagementAssociationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagementAssociationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceType' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.ManagementAssociationName, ok = parsed.Parsed["managementAssociationName"]; !ok {
		return nil, fmt.Errorf("the segment 'managementAssociationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseManagementAssociationIDInsensitively parses 'input' case-insensitively into a ManagementAssociationId
// note: this method should only be used for API response data and not user input
func ParseManagementAssociationIDInsensitively(input string) (*ManagementAssociationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagementAssociationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagementAssociationId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, fmt.Errorf("the segment 'providerName' was not found in the resource id %q", input)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceType' was not found in the resource id %q", input)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceName' was not found in the resource id %q", input)
	}

	if id.ManagementAssociationName, ok = parsed.Parsed["managementAssociationName"]; !ok {
		return nil, fmt.Errorf("the segment 'managementAssociationName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateManagementAssociationID checks that 'input' can be parsed as a Management Association ID
func ValidateManagementAssociationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagementAssociationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Management Association ID
func (id ManagementAssociationId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.OperationsManagement/managementAssociations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ResourceType, id.ResourceName, id.ManagementAssociationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Management Association ID
func (id ManagementAssociationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("resourceType", "resourceTypeValue"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationsManagement", "Microsoft.OperationsManagement", "Microsoft.OperationsManagement"),
		resourceids.StaticSegment("staticManagementAssociations", "managementAssociations", "managementAssociations"),
		resourceids.UserSpecifiedSegment("managementAssociationName", "managementAssociationValue"),
	}
}

// String returns a human-readable description of this Management Association ID
func (id ManagementAssociationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Resource Type: %q", id.ResourceType),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Management Association Name: %q", id.ManagementAssociationName),
	}
	return fmt.Sprintf("Management Association (%s)", strings.Join(components, "\n"))
}
