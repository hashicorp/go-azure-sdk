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
	recaser.RegisterResourceId(&DiagnosticDetectorId{})
}

var _ resourceids.ResourceId = &DiagnosticDetectorId{}

// DiagnosticDetectorId is a struct representing the Resource ID for a Diagnostic Detector
type DiagnosticDetectorId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	DiagnosticName    string
	DetectorName      string
}

// NewDiagnosticDetectorID returns a new DiagnosticDetectorId struct
func NewDiagnosticDetectorID(subscriptionId string, resourceGroupName string, siteName string, diagnosticName string, detectorName string) DiagnosticDetectorId {
	return DiagnosticDetectorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		DiagnosticName:    diagnosticName,
		DetectorName:      detectorName,
	}
}

// ParseDiagnosticDetectorID parses 'input' into a DiagnosticDetectorId
func ParseDiagnosticDetectorID(input string) (*DiagnosticDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiagnosticDetectorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiagnosticDetectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDiagnosticDetectorIDInsensitively parses 'input' case-insensitively into a DiagnosticDetectorId
// note: this method should only be used for API response data and not user input
func ParseDiagnosticDetectorIDInsensitively(input string) (*DiagnosticDetectorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiagnosticDetectorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiagnosticDetectorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DiagnosticDetectorId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.DiagnosticName, ok = input.Parsed["diagnosticName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "diagnosticName", input)
	}

	if id.DetectorName, ok = input.Parsed["detectorName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "detectorName", input)
	}

	return nil
}

// ValidateDiagnosticDetectorID checks that 'input' can be parsed as a Diagnostic Detector ID
func ValidateDiagnosticDetectorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiagnosticDetectorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Diagnostic Detector ID
func (id DiagnosticDetectorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/diagnostics/%s/detectors/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.DiagnosticName, id.DetectorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Diagnostic Detector ID
func (id DiagnosticDetectorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteName"),
		resourceids.StaticSegment("staticDiagnostics", "diagnostics", "diagnostics"),
		resourceids.UserSpecifiedSegment("diagnosticName", "diagnosticCategory"),
		resourceids.StaticSegment("staticDetectors", "detectors", "detectors"),
		resourceids.UserSpecifiedSegment("detectorName", "detectorName"),
	}
}

// String returns a human-readable description of this Diagnostic Detector ID
func (id DiagnosticDetectorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Diagnostic Name: %q", id.DiagnosticName),
		fmt.Sprintf("Detector Name: %q", id.DetectorName),
	}
	return fmt.Sprintf("Diagnostic Detector (%s)", strings.Join(components, "\n"))
}
