package webapps

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TriggeredWebJobHistoryId{}

// TriggeredWebJobHistoryId is a struct representing the Resource ID for a Triggered Web Job History
type TriggeredWebJobHistoryId struct {
	SubscriptionId      string
	ResourceGroupName   string
	SiteName            string
	SlotName            string
	TriggeredWebJobName string
	HistoryName         string
}

// NewTriggeredWebJobHistoryID returns a new TriggeredWebJobHistoryId struct
func NewTriggeredWebJobHistoryID(subscriptionId string, resourceGroupName string, siteName string, slotName string, triggeredWebJobName string, historyName string) TriggeredWebJobHistoryId {
	return TriggeredWebJobHistoryId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		SiteName:            siteName,
		SlotName:            slotName,
		TriggeredWebJobName: triggeredWebJobName,
		HistoryName:         historyName,
	}
}

// ParseTriggeredWebJobHistoryID parses 'input' into a TriggeredWebJobHistoryId
func ParseTriggeredWebJobHistoryID(input string) (*TriggeredWebJobHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(TriggeredWebJobHistoryId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TriggeredWebJobHistoryId{}

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

	if id.HistoryName, ok = parsed.Parsed["historyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "historyName", *parsed)
	}

	return &id, nil
}

// ParseTriggeredWebJobHistoryIDInsensitively parses 'input' case-insensitively into a TriggeredWebJobHistoryId
// note: this method should only be used for API response data and not user input
func ParseTriggeredWebJobHistoryIDInsensitively(input string) (*TriggeredWebJobHistoryId, error) {
	parser := resourceids.NewParserFromResourceIdType(TriggeredWebJobHistoryId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TriggeredWebJobHistoryId{}

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

	if id.HistoryName, ok = parsed.Parsed["historyName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "historyName", *parsed)
	}

	return &id, nil
}

// ValidateTriggeredWebJobHistoryID checks that 'input' can be parsed as a Triggered Web Job History ID
func ValidateTriggeredWebJobHistoryID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTriggeredWebJobHistoryID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Triggered Web Job History ID
func (id TriggeredWebJobHistoryId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/triggeredWebJobs/%s/history/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.TriggeredWebJobName, id.HistoryName)
}

// Segments returns a slice of Resource ID Segments which comprise this Triggered Web Job History ID
func (id TriggeredWebJobHistoryId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticHistory", "history", "history"),
		resourceids.UserSpecifiedSegment("historyName", "historyValue"),
	}
}

// String returns a human-readable description of this Triggered Web Job History ID
func (id TriggeredWebJobHistoryId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Triggered Web Job Name: %q", id.TriggeredWebJobName),
		fmt.Sprintf("History Name: %q", id.HistoryName),
	}
	return fmt.Sprintf("Triggered Web Job History (%s)", strings.Join(components, "\n"))
}
