package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotTriggeredWebJobId{}

// SlotTriggeredWebJobId is a struct representing the Resource ID for a Slot Triggered Web Job
type SlotTriggeredWebJobId struct {
	SubscriptionId      string
	ResourceGroupName   string
	SiteName            string
	SlotName            string
	TriggeredWebJobName string
}

// NewSlotTriggeredWebJobID returns a new SlotTriggeredWebJobId struct
func NewSlotTriggeredWebJobID(subscriptionId string, resourceGroupName string, siteName string, slotName string, triggeredWebJobName string) SlotTriggeredWebJobId {
	return SlotTriggeredWebJobId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		SiteName:            siteName,
		SlotName:            slotName,
		TriggeredWebJobName: triggeredWebJobName,
	}
}

// ParseSlotTriggeredWebJobID parses 'input' into a SlotTriggeredWebJobId
func ParseSlotTriggeredWebJobID(input string) (*SlotTriggeredWebJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotTriggeredWebJobId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotTriggeredWebJobId{}

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

	if id.TriggeredWebJobName, ok = parsed.Parsed["triggeredWebJobName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "triggeredWebJobName", *parsed)
	}

	return &id, nil
}

// ParseSlotTriggeredWebJobIDInsensitively parses 'input' case-insensitively into a SlotTriggeredWebJobId
// note: this method should only be used for API response data and not user input
func ParseSlotTriggeredWebJobIDInsensitively(input string) (*SlotTriggeredWebJobId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotTriggeredWebJobId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotTriggeredWebJobId{}

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

	if id.TriggeredWebJobName, ok = parsed.Parsed["triggeredWebJobName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "triggeredWebJobName", *parsed)
	}

	return &id, nil
}

// ValidateSlotTriggeredWebJobID checks that 'input' can be parsed as a Slot Triggered Web Job ID
func ValidateSlotTriggeredWebJobID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotTriggeredWebJobID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Triggered Web Job ID
func (id SlotTriggeredWebJobId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/triggeredWebJobs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.TriggeredWebJobName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Triggered Web Job ID
func (id SlotTriggeredWebJobId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticTriggeredWebJobs", "triggeredWebJobs", "triggeredWebJobs"),
		resourceids.UserSpecifiedSegment("triggeredWebJobName", "triggeredWebJobValue"),
	}
}

// String returns a human-readable description of this Slot Triggered Web Job ID
func (id SlotTriggeredWebJobId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Triggered Web Job Name: %q", id.TriggeredWebJobName),
	}
	return fmt.Sprintf("Slot Triggered Web Job (%s)", strings.Join(components, "\n"))
}
