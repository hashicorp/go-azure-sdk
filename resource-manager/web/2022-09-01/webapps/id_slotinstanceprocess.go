package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotInstanceProcessId{}

// SlotInstanceProcessId is a struct representing the Resource ID for a Slot Instance Process
type SlotInstanceProcessId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	InstanceId        string
	ProcessId         string
}

// NewSlotInstanceProcessID returns a new SlotInstanceProcessId struct
func NewSlotInstanceProcessID(subscriptionId string, resourceGroupName string, siteName string, slotName string, instanceId string, processId string) SlotInstanceProcessId {
	return SlotInstanceProcessId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		InstanceId:        instanceId,
		ProcessId:         processId,
	}
}

// ParseSlotInstanceProcessID parses 'input' into a SlotInstanceProcessId
func ParseSlotInstanceProcessID(input string) (*SlotInstanceProcessId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotInstanceProcessId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotInstanceProcessId{}

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

	if id.ProcessId, ok = parsed.Parsed["processId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "processId", *parsed)
	}

	return &id, nil
}

// ParseSlotInstanceProcessIDInsensitively parses 'input' case-insensitively into a SlotInstanceProcessId
// note: this method should only be used for API response data and not user input
func ParseSlotInstanceProcessIDInsensitively(input string) (*SlotInstanceProcessId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotInstanceProcessId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotInstanceProcessId{}

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

	if id.ProcessId, ok = parsed.Parsed["processId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "processId", *parsed)
	}

	return &id, nil
}

// ValidateSlotInstanceProcessID checks that 'input' can be parsed as a Slot Instance Process ID
func ValidateSlotInstanceProcessID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotInstanceProcessID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Instance Process ID
func (id SlotInstanceProcessId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/instances/%s/processes/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.InstanceId, id.ProcessId)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Instance Process ID
func (id SlotInstanceProcessId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticProcesses", "processes", "processes"),
		resourceids.UserSpecifiedSegment("processId", "processIdValue"),
	}
}

// String returns a human-readable description of this Slot Instance Process ID
func (id SlotInstanceProcessId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Instance: %q", id.InstanceId),
		fmt.Sprintf("Process: %q", id.ProcessId),
	}
	return fmt.Sprintf("Slot Instance Process (%s)", strings.Join(components, "\n"))
}
