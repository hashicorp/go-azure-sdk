package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotHostNameBindingId{}

// SlotHostNameBindingId is a struct representing the Resource ID for a Slot Host Name Binding
type SlotHostNameBindingId struct {
	SubscriptionId      string
	ResourceGroupName   string
	SiteName            string
	SlotName            string
	HostNameBindingName string
}

// NewSlotHostNameBindingID returns a new SlotHostNameBindingId struct
func NewSlotHostNameBindingID(subscriptionId string, resourceGroupName string, siteName string, slotName string, hostNameBindingName string) SlotHostNameBindingId {
	return SlotHostNameBindingId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		SiteName:            siteName,
		SlotName:            slotName,
		HostNameBindingName: hostNameBindingName,
	}
}

// ParseSlotHostNameBindingID parses 'input' into a SlotHostNameBindingId
func ParseSlotHostNameBindingID(input string) (*SlotHostNameBindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotHostNameBindingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotHostNameBindingId{}

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

	if id.HostNameBindingName, ok = parsed.Parsed["hostNameBindingName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hostNameBindingName", *parsed)
	}

	return &id, nil
}

// ParseSlotHostNameBindingIDInsensitively parses 'input' case-insensitively into a SlotHostNameBindingId
// note: this method should only be used for API response data and not user input
func ParseSlotHostNameBindingIDInsensitively(input string) (*SlotHostNameBindingId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotHostNameBindingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotHostNameBindingId{}

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

	if id.HostNameBindingName, ok = parsed.Parsed["hostNameBindingName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "hostNameBindingName", *parsed)
	}

	return &id, nil
}

// ValidateSlotHostNameBindingID checks that 'input' can be parsed as a Slot Host Name Binding ID
func ValidateSlotHostNameBindingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotHostNameBindingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Host Name Binding ID
func (id SlotHostNameBindingId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/hostNameBindings/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.HostNameBindingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Host Name Binding ID
func (id SlotHostNameBindingId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticHostNameBindings", "hostNameBindings", "hostNameBindings"),
		resourceids.UserSpecifiedSegment("hostNameBindingName", "hostNameBindingValue"),
	}
}

// String returns a human-readable description of this Slot Host Name Binding ID
func (id SlotHostNameBindingId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Host Name Binding Name: %q", id.HostNameBindingName),
	}
	return fmt.Sprintf("Slot Host Name Binding (%s)", strings.Join(components, "\n"))
}
