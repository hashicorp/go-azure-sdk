package integrationruntimeobjectmetadata

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IntegrationRuntimeId{}

// IntegrationRuntimeId is a struct representing the Resource ID for a Integration Runtime
type IntegrationRuntimeId struct {
	SubscriptionId         string
	ResourceGroupName      string
	FactoryName            string
	IntegrationRuntimeName string
}

// NewIntegrationRuntimeID returns a new IntegrationRuntimeId struct
func NewIntegrationRuntimeID(subscriptionId string, resourceGroupName string, factoryName string, integrationRuntimeName string) IntegrationRuntimeId {
	return IntegrationRuntimeId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		FactoryName:            factoryName,
		IntegrationRuntimeName: integrationRuntimeName,
	}
}

// ParseIntegrationRuntimeID parses 'input' into a IntegrationRuntimeId
func ParseIntegrationRuntimeID(input string) (*IntegrationRuntimeId, error) {
	parser := resourceids.NewParserFromResourceIdType(IntegrationRuntimeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IntegrationRuntimeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, fmt.Errorf("the segment 'factoryName' was not found in the resource id %q", input)
	}

	if id.IntegrationRuntimeName, ok = parsed.Parsed["integrationRuntimeName"]; !ok {
		return nil, fmt.Errorf("the segment 'integrationRuntimeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseIntegrationRuntimeIDInsensitively parses 'input' case-insensitively into a IntegrationRuntimeId
// note: this method should only be used for API response data and not user input
func ParseIntegrationRuntimeIDInsensitively(input string) (*IntegrationRuntimeId, error) {
	parser := resourceids.NewParserFromResourceIdType(IntegrationRuntimeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IntegrationRuntimeId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.FactoryName, ok = parsed.Parsed["factoryName"]; !ok {
		return nil, fmt.Errorf("the segment 'factoryName' was not found in the resource id %q", input)
	}

	if id.IntegrationRuntimeName, ok = parsed.Parsed["integrationRuntimeName"]; !ok {
		return nil, fmt.Errorf("the segment 'integrationRuntimeName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateIntegrationRuntimeID checks that 'input' can be parsed as a Integration Runtime ID
func ValidateIntegrationRuntimeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIntegrationRuntimeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Integration Runtime ID
func (id IntegrationRuntimeId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DataFactory/factories/%s/integrationRuntimes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.FactoryName, id.IntegrationRuntimeName)
}

// Segments returns a slice of Resource ID Segments which comprise this Integration Runtime ID
func (id IntegrationRuntimeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDataFactory", "Microsoft.DataFactory", "Microsoft.DataFactory"),
		resourceids.StaticSegment("staticFactories", "factories", "factories"),
		resourceids.UserSpecifiedSegment("factoryName", "factoryValue"),
		resourceids.StaticSegment("staticIntegrationRuntimes", "integrationRuntimes", "integrationRuntimes"),
		resourceids.UserSpecifiedSegment("integrationRuntimeName", "integrationRuntimeValue"),
	}
}

// String returns a human-readable description of this Integration Runtime ID
func (id IntegrationRuntimeId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Factory Name: %q", id.FactoryName),
		fmt.Sprintf("Integration Runtime Name: %q", id.IntegrationRuntimeName),
	}
	return fmt.Sprintf("Integration Runtime (%s)", strings.Join(components, "\n"))
}
