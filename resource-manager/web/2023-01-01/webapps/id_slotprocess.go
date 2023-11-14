package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotProcessId{}

// SlotProcessId is a struct representing the Resource ID for a Slot Process
type SlotProcessId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	ProcessId         string
}

// NewSlotProcessID returns a new SlotProcessId struct
func NewSlotProcessID(subscriptionId string, resourceGroupName string, siteName string, slotName string, processId string) SlotProcessId {
	return SlotProcessId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		ProcessId:         processId,
	}
}

// ParseSlotProcessID parses 'input' into a SlotProcessId
func ParseSlotProcessID(input string) (*SlotProcessId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotProcessId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotProcessId{}

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

	if id.ProcessId, ok = parsed.Parsed["processId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "processId", *parsed)
	}

	return &id, nil
}

// ParseSlotProcessIDInsensitively parses 'input' case-insensitively into a SlotProcessId
// note: this method should only be used for API response data and not user input
func ParseSlotProcessIDInsensitively(input string) (*SlotProcessId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotProcessId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotProcessId{}

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

	if id.ProcessId, ok = parsed.Parsed["processId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "processId", *parsed)
	}

	return &id, nil
}

// ValidateSlotProcessID checks that 'input' can be parsed as a Slot Process ID
func ValidateSlotProcessID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotProcessID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Process ID
func (id SlotProcessId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/processes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.ProcessId)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Process ID
func (id SlotProcessId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticProcesses", "processes", "processes"),
		resourceids.UserSpecifiedSegment("processId", "processIdValue"),
	}
}

// String returns a human-readable description of this Slot Process ID
func (id SlotProcessId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Process: %q", id.ProcessId),
	}
	return fmt.Sprintf("Slot Process (%s)", strings.Join(components, "\n"))
}
