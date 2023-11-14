package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ConfigReferenceConnectionStringId{}

// ConfigReferenceConnectionStringId is a struct representing the Resource ID for a Config Reference Connection String
type ConfigReferenceConnectionStringId struct {
	SubscriptionId      string
	ResourceGroupName   string
	SiteName            string
	SlotName            string
	ConnectionStringKey string
}

// NewConfigReferenceConnectionStringID returns a new ConfigReferenceConnectionStringId struct
func NewConfigReferenceConnectionStringID(subscriptionId string, resourceGroupName string, siteName string, slotName string, connectionStringKey string) ConfigReferenceConnectionStringId {
	return ConfigReferenceConnectionStringId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		SiteName:            siteName,
		SlotName:            slotName,
		ConnectionStringKey: connectionStringKey,
	}
}

// ParseConfigReferenceConnectionStringID parses 'input' into a ConfigReferenceConnectionStringId
func ParseConfigReferenceConnectionStringID(input string) (*ConfigReferenceConnectionStringId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigReferenceConnectionStringId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigReferenceConnectionStringId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.SlotName, ok = parsed.Parsed["slotName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "slotName", *parsed)
	}

	if id.ConnectionStringKey, ok = parsed.Parsed["connectionStringKey"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "connectionStringKey", *parsed)
	}

	return &id, nil
}

// ParseConfigReferenceConnectionStringIDInsensitively parses 'input' case-insensitively into a ConfigReferenceConnectionStringId
// note: this method should only be used for API response data and not user input
func ParseConfigReferenceConnectionStringIDInsensitively(input string) (*ConfigReferenceConnectionStringId, error) {
	parser := resourceids.NewParserFromResourceIdType(ConfigReferenceConnectionStringId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ConfigReferenceConnectionStringId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.SiteName, ok = parsed.Parsed["siteName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "siteName", *parsed)
	}

	if id.SlotName, ok = parsed.Parsed["slotName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "slotName", *parsed)
	}

	if id.ConnectionStringKey, ok = parsed.Parsed["connectionStringKey"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "connectionStringKey", *parsed)
	}

	return &id, nil
}

// ValidateConfigReferenceConnectionStringID checks that 'input' can be parsed as a Config Reference Connection String ID
func ValidateConfigReferenceConnectionStringID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseConfigReferenceConnectionStringID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Config Reference Connection String ID
func (id ConfigReferenceConnectionStringId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/config/configReferences/connectionStrings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.ConnectionStringKey)
}

// Segments returns a slice of Resource ID Segments which comprise this Config Reference Connection String ID
func (id ConfigReferenceConnectionStringId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteValue"),
		resourceids.StaticSegment("staticSlots", "slots", "slots"),
		resourceids.UserSpecifiedSegment("slotName", "slotValue"),
		resourceids.StaticSegment("staticConfig", "config", "config"),
		resourceids.StaticSegment("staticConfigReferences", "configReferences", "configReferences"),
		resourceids.StaticSegment("staticConnectionStrings", "connectionStrings", "connectionStrings"),
		resourceids.UserSpecifiedSegment("connectionStringKey", "connectionStringKeyValue"),
	}
}

// String returns a human-readable description of this Config Reference Connection String ID
func (id ConfigReferenceConnectionStringId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Connection String Key: %q", id.ConnectionStringKey),
	}
	return fmt.Sprintf("Config Reference Connection String (%s)", strings.Join(components, "\n"))
}
