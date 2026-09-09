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
	recaser.RegisterResourceId(&DiagnosticAnalysisId{})
}

var _ resourceids.ResourceId = &DiagnosticAnalysisId{}

// DiagnosticAnalysisId is a struct representing the Resource ID for a Diagnostic Analysis
type DiagnosticAnalysisId struct {
	SubscriptionId    string
	ResourceGroupName string
	SiteName          string
	SlotName          string
	DiagnosticName    string
	AnalysisName      string
}

// NewDiagnosticAnalysisID returns a new DiagnosticAnalysisId struct
func NewDiagnosticAnalysisID(subscriptionId string, resourceGroupName string, siteName string, slotName string, diagnosticName string, analysisName string) DiagnosticAnalysisId {
	return DiagnosticAnalysisId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		SiteName:          siteName,
		SlotName:          slotName,
		DiagnosticName:    diagnosticName,
		AnalysisName:      analysisName,
	}
}

// ParseDiagnosticAnalysisID parses 'input' into a DiagnosticAnalysisId
func ParseDiagnosticAnalysisID(input string) (*DiagnosticAnalysisId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiagnosticAnalysisId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiagnosticAnalysisId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDiagnosticAnalysisIDInsensitively parses 'input' case-insensitively into a DiagnosticAnalysisId
// note: this method should only be used for API response data and not user input
func ParseDiagnosticAnalysisIDInsensitively(input string) (*DiagnosticAnalysisId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DiagnosticAnalysisId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DiagnosticAnalysisId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DiagnosticAnalysisId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AnalysisName, ok = input.Parsed["analysisName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "analysisName", input)
	}

	return nil
}

// ValidateDiagnosticAnalysisID checks that 'input' can be parsed as a Diagnostic Analysis ID
func ValidateDiagnosticAnalysisID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDiagnosticAnalysisID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Diagnostic Analysis ID
func (id DiagnosticAnalysisId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Web/sites/%s/slots/%s/diagnostics/%s/analyses/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.SiteName, id.SlotName, id.DiagnosticName, id.AnalysisName)
}

// Segments returns a slice of Resource ID Segments which comprise this Diagnostic Analysis ID
func (id DiagnosticAnalysisId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftWeb", "Microsoft.Web", "Microsoft.Web"),
		resourceids.StaticSegment("staticSites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteName", "siteName"),
		resourceids.StaticSegment("staticSlots", "slots", "slots"),
		resourceids.UserSpecifiedSegment("slotName", "slotName"),
		resourceids.StaticSegment("staticDiagnostics", "diagnostics", "diagnostics"),
		resourceids.UserSpecifiedSegment("diagnosticName", "diagnosticName"),
		resourceids.StaticSegment("staticAnalyses", "analyses", "analyses"),
		resourceids.UserSpecifiedSegment("analysisName", "analysisName"),
	}
}

// String returns a human-readable description of this Diagnostic Analysis ID
func (id DiagnosticAnalysisId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Site Name: %q", id.SiteName),
		fmt.Sprintf("Slot Name: %q", id.SlotName),
		fmt.Sprintf("Diagnostic Name: %q", id.DiagnosticName),
		fmt.Sprintf("Analysis Name: %q", id.AnalysisName),
	}
	return fmt.Sprintf("Diagnostic Analysis (%s)", strings.Join(components, "\n"))
}
