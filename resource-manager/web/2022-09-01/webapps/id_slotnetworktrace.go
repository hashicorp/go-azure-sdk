package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotNetworkTraceId{}

// SlotNetworkTraceId is a struct representing the Resource ID for a Slot Network Trace
type SlotNetworkTraceId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	OperationId       string
}

// NewSlotNetworkTraceID returns a new SlotNetworkTraceId struct
func NewSlotNetworkTraceID(subscriptionId string, resourceGroupName string, siteName string, slotName string, operationId string) SlotNetworkTraceId {
	return SlotNetworkTraceId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		OperationId:       operationId,
	}
}

// ParseSlotNetworkTraceID parses 'input' into a SlotNetworkTraceId
func ParseSlotNetworkTraceID(input string) (*SlotNetworkTraceId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotNetworkTraceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotNetworkTraceId{}

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

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ParseSlotNetworkTraceIDInsensitively parses 'input' case-insensitively into a SlotNetworkTraceId
// note: this method should only be used for API response data and not user input
func ParseSlotNetworkTraceIDInsensitively(input string) (*SlotNetworkTraceId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotNetworkTraceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotNetworkTraceId{}

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

	if id.OperationId, ok = parsed.Parsed["operationId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "operationId", *parsed)
	}

	return &id, nil
}

// ValidateSlotNetworkTraceID checks that 'input' can be parsed as a Slot Network Trace ID
func ValidateSlotNetworkTraceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotNetworkTraceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Network Trace ID
func (id SlotNetworkTraceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/networkTrace/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.OperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Network Trace ID
func (id SlotNetworkTraceId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticNetworkTrace", "networkTrace", "networkTrace"),
		resourceids.UserSpecifiedSegment("operationId", "operationIdValue"),
	}
}

// String returns a human-readable description of this Slot Network Trace ID
func (id SlotNetworkTraceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Operation: %q", id.OperationId),
	}
	return fmt.Sprintf("Slot Network Trace (%s)", strings.Join(components, "\n"))
}
