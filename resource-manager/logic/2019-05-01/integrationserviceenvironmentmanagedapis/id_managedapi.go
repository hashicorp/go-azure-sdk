package integrationserviceenvironmentmanagedapis

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ManagedApiId{}

// ManagedApiId is a struct representing the Resource ID for a Managed Api
type ManagedApiId struct {
	SubscriptionId                    string
	ResourceGroup                     string
	IntegrationServiceEnvironmentName string
	ManagedApiName                    string
}

// NewManagedApiID returns a new ManagedApiId struct
func NewManagedApiID(subscriptionId string, resourceGroup string, integrationServiceEnvironmentName string, managedApiName string) ManagedApiId {
	return ManagedApiId{
		SubscriptionId:                    subscriptionId,
		ResourceGroup:                     resourceGroup,
		IntegrationServiceEnvironmentName: integrationServiceEnvironmentName,
		ManagedApiName:                    managedApiName,
	}
}

// ParseManagedApiID parses 'input' into a ManagedApiId
func ParseManagedApiID(input string) (*ManagedApiId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedApiId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedApiId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroup, ok = parsed.Parsed["resourceGroup"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroup", *parsed)
	}

	if id.IntegrationServiceEnvironmentName, ok = parsed.Parsed["integrationServiceEnvironmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "integrationServiceEnvironmentName", *parsed)
	}

	if id.ManagedApiName, ok = parsed.Parsed["managedApiName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedApiName", *parsed)
	}

	return &id, nil
}

// ParseManagedApiIDInsensitively parses 'input' case-insensitively into a ManagedApiId
// note: this method should only be used for API response data and not user input
func ParseManagedApiIDInsensitively(input string) (*ManagedApiId, error) {
	parser := resourceids.NewParserFromResourceIdType(ManagedApiId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ManagedApiId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroup, ok = parsed.Parsed["resourceGroup"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroup", *parsed)
	}

	if id.IntegrationServiceEnvironmentName, ok = parsed.Parsed["integrationServiceEnvironmentName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "integrationServiceEnvironmentName", *parsed)
	}

	if id.ManagedApiName, ok = parsed.Parsed["managedApiName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managedApiName", *parsed)
	}

	return &id, nil
}

// ValidateManagedApiID checks that 'input' can be parsed as a Managed Api ID
func ValidateManagedApiID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedApiID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Api ID
func (id ManagedApiId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Logic/integrationServiceEnvironments/%s/managedApis/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.IntegrationServiceEnvironmentName, id.ManagedApiName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Api ID
func (id ManagedApiId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroup", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftLogic", "Microsoft.Logic", "Microsoft.Logic"),
		resourceids.StaticSegment("staticIntegrationServiceEnvironments", "integrationServiceEnvironments", "integrationServiceEnvironments"),
		resourceids.UserSpecifiedSegment("integrationServiceEnvironmentName", "integrationServiceEnvironmentValue"),
		resourceids.StaticSegment("staticManagedApis", "managedApis", "managedApis"),
		resourceids.UserSpecifiedSegment("managedApiName", "managedApiValue"),
	}
}

// String returns a human-readable description of this Managed Api ID
func (id ManagedApiId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group: %q", id.ResourceGroup),
		fmt.Sprintf("Integration Service Environment Name: %q", id.IntegrationServiceEnvironmentName),
		fmt.Sprintf("Managed Api Name: %q", id.ManagedApiName),
	}
	return fmt.Sprintf("Managed Api (%s)", strings.Join(components, "\n"))
}
