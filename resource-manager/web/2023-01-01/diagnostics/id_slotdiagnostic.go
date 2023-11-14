package diagnostics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SlotDiagnosticId{}

// SlotDiagnosticId is a struct representing the Resource ID for a Slot Diagnostic
type SlotDiagnosticId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	DiagnosticName    string
}

// NewSlotDiagnosticID returns a new SlotDiagnosticId struct
func NewSlotDiagnosticID(subscriptionId string, resourceGroupName string, siteName string, slotName string, diagnosticName string) SlotDiagnosticId {
	return SlotDiagnosticId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		DiagnosticName:    diagnosticName,
	}
}

// ParseSlotDiagnosticID parses 'input' into a SlotDiagnosticId
func ParseSlotDiagnosticID(input string) (*SlotDiagnosticId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotDiagnosticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotDiagnosticId{}

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

	return &id, nil
}

// ParseSlotDiagnosticIDInsensitively parses 'input' case-insensitively into a SlotDiagnosticId
// note: this method should only be used for API response data and not user input
func ParseSlotDiagnosticIDInsensitively(input string) (*SlotDiagnosticId, error) {
	parser := resourceids.NewParserFromResourceIdType(SlotDiagnosticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SlotDiagnosticId{}

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

	return &id, nil
}

// ValidateSlotDiagnosticID checks that 'input' can be parsed as a Slot Diagnostic ID
func ValidateSlotDiagnosticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSlotDiagnosticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Slot Diagnostic ID
func (id SlotDiagnosticId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/diagnostics/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.DiagnosticName)
}

// Segments returns a slice of Resource ID Segments which comprise this Slot Diagnostic ID
func (id SlotDiagnosticId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Slot Diagnostic ID
func (id SlotDiagnosticId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Diagnostic Name: %q", id.DiagnosticName),
	}
	return fmt.Sprintf("Slot Diagnostic (%s)", strings.Join(components, "\n"))
}
