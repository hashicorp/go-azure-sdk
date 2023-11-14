package diagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotDiagnosticDetectorId{}

// SlotDiagnosticDetectorId is a struct representing the Resource ID for a Slot Diagnostic Detector
type SlotDiagnosticDetectorId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	DiagnosticName    string
	DetectorName      string
}

// NewSlotDiagnosticDetectorID returns a new SlotDiagnosticDetectorId struct
func NewSlotDiagnosticDetectorID(subscriptionId string, resourceGroupName string, siteName string, slotName string, diagnosticName string, detectorName string) SlotDiagnosticDetectorId {
	return SlotDiagnosticDetectorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		DiagnosticName:    diagnosticName,
		DetectorName:      detectorName,
	}
}

// ParseSlotDiagnosticDetectorID parses 'input' into a SlotDiagnosticDetectorId
func ParseSlotDiagnosticDetectorID(input string) (*SlotDiagnosticDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotDiagnosticDetectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotDiagnosticDetectorId{}

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

	if id.DiagnosticName, ok = parsed.Parsed["diagnosticName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "diagnosticName", *parsed)
	}

	if id.DetectorName, ok = parsed.Parsed["detectorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "detectorName", *parsed)
	}

	return &id, nil
}

// ParseSlotDiagnosticDetectorIDInsensitively parses 'input' case-insensitively into a SlotDiagnosticDetectorId
// note: this method should only be used for API response data and not user input
func ParseSlotDiagnosticDetectorIDInsensitively(input string) (*SlotDiagnosticDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotDiagnosticDetectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotDiagnosticDetectorId{}

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

	if id.DiagnosticName, ok = parsed.Parsed["diagnosticName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "diagnosticName", *parsed)
	}

	if id.DetectorName, ok = parsed.Parsed["detectorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "detectorName", *parsed)
	}

	return &id, nil
}

// ValidateSlotDiagnosticDetectorID checks that 'input' can be parsed as a Slot Diagnostic Detector ID
func ValidateSlotDiagnosticDetectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotDiagnosticDetectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Diagnostic Detector ID
func (id SlotDiagnosticDetectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/diagnostics/%s/detectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.DiagnosticName, id.DetectorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Diagnostic Detector ID
func (id SlotDiagnosticDetectorId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("staticDiagnostics", "diagnostics", "diagnostics"),
		resourceids.UserSpecifiedSegment("diagnosticName", "diagnosticValue"),
		resourceids.StaticSegment("staticDetectors", "detectors", "detectors"),
		resourceids.UserSpecifiedSegment("detectorName", "detectorValue"),
	}
}

// String returns a human-readable description of this Slot Diagnostic Detector ID
func (id SlotDiagnosticDetectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Diagnostic Name: %q", id.DiagnosticName),
		fmt.Sprintf("Detector Name: %q", id.DetectorName),
	}
	return fmt.Sprintf("Slot Diagnostic Detector (%s)", strings.Join(components, "\n"))
}
