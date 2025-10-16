package registryendpoint

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&RegistryEndpointId{})
}

var _ resourceids.ResourceId = &RegistryEndpointId{}

// RegistryEndpointId is a struct representing the Resource ID for a Registry Endpoint
type RegistryEndpointId struct {
	SubscriptionId       string
	ResourceGroupName    string
	InstanceName         string
	RegistryEndpointName string
}

// NewRegistryEndpointID returns a new RegistryEndpointId struct
func NewRegistryEndpointID(subscriptionId string, resourceGroupName string, instanceName string, registryEndpointName string) RegistryEndpointId {
	return RegistryEndpointId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		InstanceName:         instanceName,
		RegistryEndpointName: registryEndpointName,
	}
}

// ParseRegistryEndpointID parses 'input' into a RegistryEndpointId
func ParseRegistryEndpointID(input string) (*RegistryEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegistryEndpointId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegistryEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRegistryEndpointIDInsensitively parses 'input' case-insensitively into a RegistryEndpointId
// note: this method should only be used for API response data and not user input
func ParseRegistryEndpointIDInsensitively(input string) (*RegistryEndpointId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RegistryEndpointId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RegistryEndpointId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RegistryEndpointId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.InstanceName, ok = input.Parsed["instanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "instanceName", input)
	}

	if id.RegistryEndpointName, ok = input.Parsed["registryEndpointName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "registryEndpointName", input)
	}

	return nil
}

// ValidateRegistryEndpointID checks that 'input' can be parsed as a Registry Endpoint ID
func ValidateRegistryEndpointID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegistryEndpointID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Registry Endpoint ID
func (id RegistryEndpointId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.IoTOperations/instances/%s/registryEndpoints/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.InstanceName, id.RegistryEndpointName)
}

// Segments returns a slice of Resource ID Segments which comprise this Registry Endpoint ID
func (id RegistryEndpointId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftIoTOperations", "Microsoft.IoTOperations", "Microsoft.IoTOperations"),
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("instanceName", "instanceName"),
		resourceids.StaticSegment("staticRegistryEndpoints", "registryEndpoints", "registryEndpoints"),
		resourceids.UserSpecifiedSegment("registryEndpointName", "registryEndpointName"),
	}
}

// String returns a human-readable description of this Registry Endpoint ID
func (id RegistryEndpointId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Instance Name: %q", id.InstanceName),
		fmt.Sprintf("Registry Endpoint Name: %q", id.RegistryEndpointName),
	}
	return fmt.Sprintf("Registry Endpoint (%s)", strings.Join(components, "\n"))
}
