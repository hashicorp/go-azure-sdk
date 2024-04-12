package compliances

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedComplianceId{})
}

var _ resourceids.ResourceId = &ScopedComplianceId{}

// ScopedComplianceId is a struct representing the Resource ID for a Scoped Compliance
type ScopedComplianceId struct {
	Scope          string
	ComplianceName string
}

// NewScopedComplianceID returns a new ScopedComplianceId struct
func NewScopedComplianceID(scope string, complianceName string) ScopedComplianceId {
	return ScopedComplianceId{
		Scope:          scope,
		ComplianceName: complianceName,
	}
}

// ParseScopedComplianceID parses 'input' into a ScopedComplianceId
func ParseScopedComplianceID(input string) (*ScopedComplianceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedComplianceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedComplianceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedComplianceIDInsensitively parses 'input' case-insensitively into a ScopedComplianceId
// note: this method should only be used for API response data and not user input
func ParseScopedComplianceIDInsensitively(input string) (*ScopedComplianceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedComplianceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedComplianceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedComplianceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.ComplianceName, ok = input.Parsed["complianceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "complianceName", input)
	}

	return nil
}

// ValidateScopedComplianceID checks that 'input' can be parsed as a Scoped Compliance ID
func ValidateScopedComplianceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedComplianceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Compliance ID
func (id ScopedComplianceId) ID() string {
	fmtString := "/%s/providers/Microsoft.Security/compliances/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.ComplianceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Compliance ID
func (id ScopedComplianceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticCompliances", "compliances", "compliances"),
		resourceids.UserSpecifiedSegment("complianceName", "complianceValue"),
	}
}

// String returns a human-readable description of this Scoped Compliance ID
func (id ScopedComplianceId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Compliance Name: %q", id.ComplianceName),
	}
	return fmt.Sprintf("Scoped Compliance (%s)", strings.Join(components, "\n"))
}
