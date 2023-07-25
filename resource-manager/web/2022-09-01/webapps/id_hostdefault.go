package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = HostDefaultId{}

// HostDefaultId is a struct representing the Resource ID for a Host Default
type HostDefaultId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	DefaultName       string
	KeyName           string
}

// NewHostDefaultID returns a new HostDefaultId struct
func NewHostDefaultID(subscriptionId string, resourceGroupName string, siteName string, slotName string, defaultName string, keyName string) HostDefaultId {
	return HostDefaultId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		DefaultName:       defaultName,
		KeyName:           keyName,
	}
}

// ParseHostDefaultID parses 'input' into a HostDefaultId
func ParseHostDefaultID(input string) (*HostDefaultId, error) {
	parser := resourceids.NewParserFromResourceIdType(HostDefaultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HostDefaultId{}

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

	if id.DefaultName, ok = parsed.Parsed["defaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "defaultName", *parsed)
	}

	if id.KeyName, ok = parsed.Parsed["keyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "keyName", *parsed)
	}

	return &id, nil
}

// ParseHostDefaultIDInsensitively parses 'input' case-insensitively into a HostDefaultId
// note: this method should only be used for API response data and not user input
func ParseHostDefaultIDInsensitively(input string) (*HostDefaultId, error) {
	parser := resourceids.NewParserFromResourceIdType(HostDefaultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := HostDefaultId{}

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

	if id.DefaultName, ok = parsed.Parsed["defaultName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "defaultName", *parsed)
	}

	if id.KeyName, ok = parsed.Parsed["keyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "keyName", *parsed)
	}

	return &id, nil
}

// ValidateHostDefaultID checks that 'input' can be parsed as a Host Default ID
func ValidateHostDefaultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseHostDefaultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Host Default ID
func (id HostDefaultId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/host/default/%s/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.DefaultName, id.KeyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Host Default ID
func (id HostDefaultId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticHost", "host", "host"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.UserSpecifiedSegment("defaultName", "defaultValue"),
		resourceids.UserSpecifiedSegment("keyName", "keyValue"),
	}
}

// String returns a human-readable description of this Host Default ID
func (id HostDefaultId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Default Name: %q", id.DefaultName),
		fmt.Sprintf("Key Name: %q", id.KeyName),
	}
	return fmt.Sprintf("Host Default (%s)", strings.Join(components, "\n"))
}
