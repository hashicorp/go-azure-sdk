package keyvalues

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&KeyValueId{})
}

var _ resourceids.ResourceId = &KeyValueId{}

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
	parser := resourceids.NewParserFromResourceIdType(&KeyValueId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KeyValueId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseKeyValueIDInsensitively parses 'input' case-insensitively into a KeyValueId
// note: this method should only be used for API response data and not user input
func ParseKeyValueIDInsensitively(input string) (*KeyValueId, error) {
	parser := resourceids.NewParserFromResourceIdType(&KeyValueId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := KeyValueId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *KeyValueId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ConfigurationStoreName, ok = input.Parsed["configurationStoreName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "configurationStoreName", input)
	}

	if id.KeyValueName, ok = input.Parsed["keyValueName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "keyValueName", input)
	}

	return nil
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
