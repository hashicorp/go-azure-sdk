package managementgroupdiagnosticsettings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&Providers2DiagnosticSettingId{})
}

var _ resourceids.ResourceId = &Providers2DiagnosticSettingId{}

// Providers2DiagnosticSettingId is a struct representing the Resource ID for a Providers 2 Diagnostic Setting
type Providers2DiagnosticSettingId struct {
	ManagementGroupId     string
	DiagnosticSettingName string
}

// NewProviders2DiagnosticSettingID returns a new Providers2DiagnosticSettingId struct
func NewProviders2DiagnosticSettingID(managementGroupId string, diagnosticSettingName string) Providers2DiagnosticSettingId {
	return Providers2DiagnosticSettingId{
		ManagementGroupId:     managementGroupId,
		DiagnosticSettingName: diagnosticSettingName,
	}
}

// ParseProviders2DiagnosticSettingID parses 'input' into a Providers2DiagnosticSettingId
func ParseProviders2DiagnosticSettingID(input string) (*Providers2DiagnosticSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2DiagnosticSettingId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2DiagnosticSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviders2DiagnosticSettingIDInsensitively parses 'input' case-insensitively into a Providers2DiagnosticSettingId
// note: this method should only be used for API response data and not user input
func ParseProviders2DiagnosticSettingIDInsensitively(input string) (*Providers2DiagnosticSettingId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2DiagnosticSettingId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2DiagnosticSettingId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *Providers2DiagnosticSettingId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagementGroupId, ok = input.Parsed["managementGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementGroupId", input)
	}

	if id.DiagnosticSettingName, ok = input.Parsed["diagnosticSettingName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "diagnosticSettingName", input)
	}

	return nil
}

// ValidateProviders2DiagnosticSettingID checks that 'input' can be parsed as a Providers 2 Diagnostic Setting ID
func ValidateProviders2DiagnosticSettingID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2DiagnosticSettingID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Diagnostic Setting ID
func (id Providers2DiagnosticSettingId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Insights/diagnosticSettings/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.DiagnosticSettingName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Diagnostic Setting ID
func (id Providers2DiagnosticSettingId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupId", "managementGroupId"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftInsights", "Microsoft.Insights", "Microsoft.Insights"),
		resourceids.StaticSegment("staticDiagnosticSettings", "diagnosticSettings", "diagnosticSettings"),
		resourceids.UserSpecifiedSegment("diagnosticSettingName", "name"),
	}
}

// String returns a human-readable description of this Providers 2 Diagnostic Setting ID
func (id Providers2DiagnosticSettingId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Diagnostic Setting Name: %q", id.DiagnosticSettingName),
	}
	return fmt.Sprintf("Providers 2 Diagnostic Setting (%s)", strings.Join(components, "\n"))
}
