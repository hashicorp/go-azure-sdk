package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotFunctionId{}

// SlotFunctionId is a struct representing the Resource ID for a Slot Function
type SlotFunctionId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	FunctionName      string
}

// NewSlotFunctionID returns a new SlotFunctionId struct
func NewSlotFunctionID(subscriptionId string, resourceGroupName string, siteName string, slotName string, functionName string) SlotFunctionId {
	return SlotFunctionId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		FunctionName:      functionName,
	}
}

// ParseSlotFunctionID parses 'input' into a SlotFunctionId
func ParseSlotFunctionID(input string) (*SlotFunctionId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotFunctionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotFunctionId{}

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

	if id.FunctionName, ok = parsed.Parsed["functionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "functionName", *parsed)
	}

	return &id, nil
}

// ParseSlotFunctionIDInsensitively parses 'input' case-insensitively into a SlotFunctionId
// note: this method should only be used for API response data and not user input
func ParseSlotFunctionIDInsensitively(input string) (*SlotFunctionId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotFunctionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotFunctionId{}

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

	if id.FunctionName, ok = parsed.Parsed["functionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "functionName", *parsed)
	}

	return &id, nil
}

// ValidateSlotFunctionID checks that 'input' can be parsed as a Slot Function ID
func ValidateSlotFunctionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotFunctionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Function ID
func (id SlotFunctionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/functions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.FunctionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Function ID
func (id SlotFunctionId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticFunctions", "functions", "functions"),
		resourceids.UserSpecifiedSegment("functionName", "functionValue"),
	}
}

// String returns a human-readable description of this Slot Function ID
func (id SlotFunctionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Function Name: %q", id.FunctionName),
	}
	return fmt.Sprintf("Slot Function (%s)", strings.Join(components, "\n"))
}
