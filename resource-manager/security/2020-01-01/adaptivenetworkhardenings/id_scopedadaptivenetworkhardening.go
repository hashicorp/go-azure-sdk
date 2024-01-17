package adaptivenetworkhardenings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedAdaptiveNetworkHardeningId{}

// ScopedAdaptiveNetworkHardeningId is a struct representing the Resource ID for a Scoped Adaptive Network Hardening
type ScopedAdaptiveNetworkHardeningId struct {
	Scope                        string
	AdaptiveNetworkHardeningName string
}

// NewScopedAdaptiveNetworkHardeningID returns a new ScopedAdaptiveNetworkHardeningId struct
func NewScopedAdaptiveNetworkHardeningID(scope string, adaptiveNetworkHardeningName string) ScopedAdaptiveNetworkHardeningId {
	return ScopedAdaptiveNetworkHardeningId{
		Scope:                        scope,
		AdaptiveNetworkHardeningName: adaptiveNetworkHardeningName,
	}
}

// ParseScopedAdaptiveNetworkHardeningID parses 'input' into a ScopedAdaptiveNetworkHardeningId
func ParseScopedAdaptiveNetworkHardeningID(input string) (*ScopedAdaptiveNetworkHardeningId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedAdaptiveNetworkHardeningId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedAdaptiveNetworkHardeningId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedAdaptiveNetworkHardeningIDInsensitively parses 'input' case-insensitively into a ScopedAdaptiveNetworkHardeningId
// note: this method should only be used for API response data and not user input
func ParseScopedAdaptiveNetworkHardeningIDInsensitively(input string) (*ScopedAdaptiveNetworkHardeningId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedAdaptiveNetworkHardeningId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedAdaptiveNetworkHardeningId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedAdaptiveNetworkHardeningId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.AdaptiveNetworkHardeningName, ok = input.Parsed["adaptiveNetworkHardeningName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "adaptiveNetworkHardeningName", input)
	}

	return nil
}

// ValidateScopedAdaptiveNetworkHardeningID checks that 'input' can be parsed as a Scoped Adaptive Network Hardening ID
func ValidateScopedAdaptiveNetworkHardeningID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedAdaptiveNetworkHardeningID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Adaptive Network Hardening ID
func (id ScopedAdaptiveNetworkHardeningId) ID() string {
	fmtString := "/%s/providers/Microsoft.Security/adaptiveNetworkHardenings/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.AdaptiveNetworkHardeningName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Adaptive Network Hardening ID
func (id ScopedAdaptiveNetworkHardeningId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticAdaptiveNetworkHardenings", "adaptiveNetworkHardenings", "adaptiveNetworkHardenings"),
		resourceids.UserSpecifiedSegment("adaptiveNetworkHardeningName", "adaptiveNetworkHardeningValue"),
	}
}

// String returns a human-readable description of this Scoped Adaptive Network Hardening ID
func (id ScopedAdaptiveNetworkHardeningId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Adaptive Network Hardening Name: %q", id.AdaptiveNetworkHardeningName),
	}
	return fmt.Sprintf("Scoped Adaptive Network Hardening (%s)", strings.Join(components, "\n"))
}
