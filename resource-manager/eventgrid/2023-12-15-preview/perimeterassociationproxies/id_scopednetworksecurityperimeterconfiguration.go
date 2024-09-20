package perimeterassociationproxies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedNetworkSecurityPerimeterConfigurationId{})
}

var _ resourceids.ResourceId = &ScopedNetworkSecurityPerimeterConfigurationId{}

// ScopedNetworkSecurityPerimeterConfigurationId is a struct representing the Resource ID for a Scoped Network Security Perimeter Configuration
type ScopedNetworkSecurityPerimeterConfigurationId struct {
	Scope                                     string
	NetworkSecurityPerimeterConfigurationName string
}

// NewScopedNetworkSecurityPerimeterConfigurationID returns a new ScopedNetworkSecurityPerimeterConfigurationId struct
func NewScopedNetworkSecurityPerimeterConfigurationID(scope string, networkSecurityPerimeterConfigurationName string) ScopedNetworkSecurityPerimeterConfigurationId {
	return ScopedNetworkSecurityPerimeterConfigurationId{
		Scope: scope,
		NetworkSecurityPerimeterConfigurationName: networkSecurityPerimeterConfigurationName,
	}
}

// ParseScopedNetworkSecurityPerimeterConfigurationID parses 'input' into a ScopedNetworkSecurityPerimeterConfigurationId
func ParseScopedNetworkSecurityPerimeterConfigurationID(input string) (*ScopedNetworkSecurityPerimeterConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedNetworkSecurityPerimeterConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedNetworkSecurityPerimeterConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedNetworkSecurityPerimeterConfigurationIDInsensitively parses 'input' case-insensitively into a ScopedNetworkSecurityPerimeterConfigurationId
// note: this method should only be used for API response data and not user input
func ParseScopedNetworkSecurityPerimeterConfigurationIDInsensitively(input string) (*ScopedNetworkSecurityPerimeterConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedNetworkSecurityPerimeterConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedNetworkSecurityPerimeterConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedNetworkSecurityPerimeterConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.NetworkSecurityPerimeterConfigurationName, ok = input.Parsed["networkSecurityPerimeterConfigurationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "networkSecurityPerimeterConfigurationName", input)
	}

	return nil
}

// ValidateScopedNetworkSecurityPerimeterConfigurationID checks that 'input' can be parsed as a Scoped Network Security Perimeter Configuration ID
func ValidateScopedNetworkSecurityPerimeterConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedNetworkSecurityPerimeterConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Network Security Perimeter Configuration ID
func (id ScopedNetworkSecurityPerimeterConfigurationId) ID() string {
	fmtString := "/%s/networkSecurityPerimeterConfigurations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.NetworkSecurityPerimeterConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Network Security Perimeter Configuration ID
func (id ScopedNetworkSecurityPerimeterConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticNetworkSecurityPerimeterConfigurations", "networkSecurityPerimeterConfigurations", "networkSecurityPerimeterConfigurations"),
		resourceids.UserSpecifiedSegment("networkSecurityPerimeterConfigurationName", "perimeterGuidassociationName"),
	}
}

// String returns a human-readable description of this Scoped Network Security Perimeter Configuration ID
func (id ScopedNetworkSecurityPerimeterConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Network Security Perimeter Configuration Name: %q", id.NetworkSecurityPerimeterConfigurationName),
	}
	return fmt.Sprintf("Scoped Network Security Perimeter Configuration (%s)", strings.Join(components, "\n"))
}
