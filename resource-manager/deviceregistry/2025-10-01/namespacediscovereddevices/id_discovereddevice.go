package namespacediscovereddevices

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DiscoveredDeviceId{})
}

var _ resourceids.ResourceId = &DiscoveredDeviceId{}

// DiscoveredDeviceId is a struct representing the Resource ID for a Discovered Device
type DiscoveredDeviceId struct {
	SubscriptionId       string
	ResourceGroupName    string
	NamespaceName        string
	DiscoveredDeviceName string
}

// NewDiscoveredDeviceID returns a new DiscoveredDeviceId struct
func NewDiscoveredDeviceID(subscriptionId string, resourceGroupName string, namespaceName string, discoveredDeviceName string) DiscoveredDeviceId {
	return DiscoveredDeviceId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		NamespaceName:        namespaceName,
		DiscoveredDeviceName: discoveredDeviceName,
	}
}

// ParseDiscoveredDeviceID parses 'input' into a DiscoveredDeviceId
func ParseDiscoveredDeviceID(input string) (*DiscoveredDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiscoveredDeviceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiscoveredDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDiscoveredDeviceIDInsensitively parses 'input' case-insensitively into a DiscoveredDeviceId
// note: this method should only be used for API response data and not user input
func ParseDiscoveredDeviceIDInsensitively(input string) (*DiscoveredDeviceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiscoveredDeviceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiscoveredDeviceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DiscoveredDeviceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NamespaceName, ok = input.Parsed["namespaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "namespaceName", input)
	}

	if id.DiscoveredDeviceName, ok = input.Parsed["discoveredDeviceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "discoveredDeviceName", input)
	}

	return nil
}

// ValidateDiscoveredDeviceID checks that 'input' can be parsed as a Discovered Device ID
func ValidateDiscoveredDeviceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiscoveredDeviceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Discovered Device ID
func (id DiscoveredDeviceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DeviceRegistry/namespaces/%s/discoveredDevices/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NamespaceName, id.DiscoveredDeviceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Discovered Device ID
func (id DiscoveredDeviceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDeviceRegistry", "Microsoft.DeviceRegistry", "Microsoft.DeviceRegistry"),
		resourceids.StaticSegment("staticNamespaces", "namespaces", "namespaces"),
		resourceids.UserSpecifiedSegment("namespaceName", "namespaceName"),
		resourceids.StaticSegment("staticDiscoveredDevices", "discoveredDevices", "discoveredDevices"),
		resourceids.UserSpecifiedSegment("discoveredDeviceName", "discoveredDeviceName"),
	}
}

// String returns a human-readable description of this Discovered Device ID
func (id DiscoveredDeviceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Namespace Name: %q", id.NamespaceName),
		fmt.Sprintf("Discovered Device Name: %q", id.DiscoveredDeviceName),
	}
	return fmt.Sprintf("Discovered Device (%s)", strings.Join(components, "\n"))
}
