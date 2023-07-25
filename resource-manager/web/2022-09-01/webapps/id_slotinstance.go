package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotInstanceId{}

// SlotInstanceId is a struct representing the Resource ID for a Slot Instance
type SlotInstanceId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	InstanceId        string
}

// NewSlotInstanceID returns a new SlotInstanceId struct
func NewSlotInstanceID(subscriptionId string, resourceGroupName string, siteName string, slotName string, instanceId string) SlotInstanceId {
	return SlotInstanceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		InstanceId:        instanceId,
	}
}

// ParseSlotInstanceID parses 'input' into a SlotInstanceId
func ParseSlotInstanceID(input string) (*SlotInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotInstanceId{}

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

	if id.InstanceId, ok = parsed.Parsed["instanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "instanceId", *parsed)
	}

	return &id, nil
}

// ParseSlotInstanceIDInsensitively parses 'input' case-insensitively into a SlotInstanceId
// note: this method should only be used for API response data and not user input
func ParseSlotInstanceIDInsensitively(input string) (*SlotInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotInstanceId{}

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

	if id.InstanceId, ok = parsed.Parsed["instanceId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "instanceId", *parsed)
	}

	return &id, nil
}

// ValidateSlotInstanceID checks that 'input' can be parsed as a Slot Instance ID
func ValidateSlotInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Instance ID
func (id SlotInstanceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/instances/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.InstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Instance ID
func (id SlotInstanceId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticInstances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("instanceId", "instanceIdValue"),
	}
}

// String returns a human-readable description of this Slot Instance ID
func (id SlotInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Instance: %q", id.InstanceId),
	}
	return fmt.Sprintf("Slot Instance (%s)", strings.Join(components, "\n"))
}
