package diagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&SlotDiagnosticDetectorId{})
}

var _ resourceids.ResourceId = &SlotDiagnosticDetectorId{}

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
	parser := resourceids.NewParserFromResourceIdType(&SlotDiagnosticDetectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SlotDiagnosticDetectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseSlotDiagnosticDetectorIDInsensitively parses 'input' case-insensitively into a SlotDiagnosticDetectorId
// note: this method should only be used for API response data and not user input
func ParseSlotDiagnosticDetectorIDInsensitively(input string) (*SlotDiagnosticDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&SlotDiagnosticDetectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := SlotDiagnosticDetectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *SlotDiagnosticDetectorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.SiteName, ok = input.Parsed["siteName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteName", input)
	}

	if id.SlotName, ok = input.Parsed["slotName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "slotName", input)
	}

	if id.DiagnosticName, ok = input.Parsed["diagnosticName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "diagnosticName", input)
	}

	if id.DetectorName, ok = input.Parsed["detectorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectorName", input)
	}

	return nil
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
