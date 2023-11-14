package diagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotDetectorId{}

// SlotDetectorId is a struct representing the Resource ID for a Slot Detector
type SlotDetectorId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	DetectorName      string
}

// NewSlotDetectorID returns a new SlotDetectorId struct
func NewSlotDetectorID(subscriptionId string, resourceGroupName string, siteName string, slotName string, detectorName string) SlotDetectorId {
	return SlotDetectorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		DetectorName:      detectorName,
	}
}

// ParseSlotDetectorID parses 'input' into a SlotDetectorId
func ParseSlotDetectorID(input string) (*SlotDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotDetectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotDetectorId{}

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

	if id.DetectorName, ok = parsed.Parsed["detectorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "detectorName", *parsed)
	}

	return &id, nil
}

// ParseSlotDetectorIDInsensitively parses 'input' case-insensitively into a SlotDetectorId
// note: this method should only be used for API response data and not user input
func ParseSlotDetectorIDInsensitively(input string) (*SlotDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotDetectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotDetectorId{}

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

	if id.DetectorName, ok = parsed.Parsed["detectorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "detectorName", *parsed)
	}

	return &id, nil
}

// ValidateSlotDetectorID checks that 'input' can be parsed as a Slot Detector ID
func ValidateSlotDetectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotDetectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Detector ID
func (id SlotDetectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/detectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.DetectorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Detector ID
func (id SlotDetectorId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticDetectors", "detectors", "detectors"),
		resourceids.UserSpecifiedSegment("detectorName", "detectorValue"),
	}
}

// String returns a human-readable description of this Slot Detector ID
func (id SlotDetectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Detector Name: %q", id.DetectorName),
	}
	return fmt.Sprintf("Slot Detector (%s)", strings.Join(components, "\n"))
}
