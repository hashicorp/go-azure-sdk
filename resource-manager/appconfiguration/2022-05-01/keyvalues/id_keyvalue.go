// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvalues

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = KeyValueId{}

// KeyValueId is a struct representing the Resource ID for a Key Value
type KeyValueId struct {
	SubscriptionId         string
	ResourceGroupName      string
	ConfigurationStoreName string
	KeyValueName           string
}

// NewKeyValueID returns a new KeyValueId struct
func NewKeyValueID(subscriptionId string, resourceGroupName string, configurationStoreName string, keyValueName string) KeyValueId {
	return KeyValueId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		ConfigurationStoreName: configurationStoreName,
		KeyValueName:           keyValueName,
	}
}

// ParseKeyValueID parses 'input' into a KeyValueId
func ParseKeyValueID(input string) (*KeyValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(KeyValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := KeyValueId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ConfigurationStoreName, ok = parsed.Parsed["configurationStoreName"]; !ok {
		return nil, fmt.Errorf("the segment 'configurationStoreName' was not found in the resource id %q", input)
	}

	if id.KeyValueName, ok = parsed.Parsed["keyValueName"]; !ok {
		return nil, fmt.Errorf("the segment 'keyValueName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseKeyValueIDInsensitively parses 'input' case-insensitively into a KeyValueId
// note: this method should only be used for API response data and not user input
func ParseKeyValueIDInsensitively(input string) (*KeyValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(KeyValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := KeyValueId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.ConfigurationStoreName, ok = parsed.Parsed["configurationStoreName"]; !ok {
		return nil, fmt.Errorf("the segment 'configurationStoreName' was not found in the resource id %q", input)
	}

	if id.KeyValueName, ok = parsed.Parsed["keyValueName"]; !ok {
		return nil, fmt.Errorf("the segment 'keyValueName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateKeyValueID checks that 'input' can be parsed as a Key Value ID
func ValidateKeyValueID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseKeyValueID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Key Value ID
func (id KeyValueId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.AppConfiguration/configurationStores/%s/keyValues/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ConfigurationStoreName, id.KeyValueName)
}

// Segments returns a slice of Resource ID Segments which comprise this Key Value ID
func (id KeyValueId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAppConfiguration", "Microsoft.AppConfiguration", "Microsoft.AppConfiguration"),
		resourceids.StaticSegment("staticConfigurationStores", "configurationStores", "configurationStores"),
		resourceids.UserSpecifiedSegment("configurationStoreName", "configurationStoreValue"),
		resourceids.StaticSegment("staticKeyValues", "keyValues", "keyValues"),
		resourceids.UserSpecifiedSegment("keyValueName", "keyValueValue"),
	}
}

// String returns a human-readable description of this Key Value ID
func (id KeyValueId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Configuration Store Name: %q", id.ConfigurationStoreName),
		fmt.Sprintf("Key Value Name: %q", id.KeyValueName),
	}
	return fmt.Sprintf("Key Value (%s)", strings.Join(components, "\n"))
}
