package sourcecontrolconfiguration

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ScopedSourceControlConfigurationId{}

// ScopedSourceControlConfigurationId is a struct representing the Resource ID for a Scoped Source Control Configuration
type ScopedSourceControlConfigurationId struct {
	Scope                          string
	SourceControlConfigurationName string
}

// NewScopedSourceControlConfigurationID returns a new ScopedSourceControlConfigurationId struct
func NewScopedSourceControlConfigurationID(scope string, sourceControlConfigurationName string) ScopedSourceControlConfigurationId {
	return ScopedSourceControlConfigurationId{
		Scope:                          scope,
		SourceControlConfigurationName: sourceControlConfigurationName,
	}
}

// ParseScopedSourceControlConfigurationID parses 'input' into a ScopedSourceControlConfigurationId
func ParseScopedSourceControlConfigurationID(input string) (*ScopedSourceControlConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedSourceControlConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedSourceControlConfigurationId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.SourceControlConfigurationName, ok = parsed.Parsed["sourceControlConfigurationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sourceControlConfigurationName", *parsed)
	}

	return &id, nil
}

// ParseScopedSourceControlConfigurationIDInsensitively parses 'input' case-insensitively into a ScopedSourceControlConfigurationId
// note: this method should only be used for API response data and not user input
func ParseScopedSourceControlConfigurationIDInsensitively(input string) (*ScopedSourceControlConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(ScopedSourceControlConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ScopedSourceControlConfigurationId{}

	if id.Scope, ok = parsed.Parsed["scope"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "scope", *parsed)
	}

	if id.SourceControlConfigurationName, ok = parsed.Parsed["sourceControlConfigurationName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "sourceControlConfigurationName", *parsed)
	}

	return &id, nil
}

// ValidateScopedSourceControlConfigurationID checks that 'input' can be parsed as a Scoped Source Control Configuration ID
func ValidateScopedSourceControlConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedSourceControlConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Source Control Configuration ID
func (id ScopedSourceControlConfigurationId) ID() string {
	fmtString := "/%s/providers/Microsoft.KubernetesConfiguration/sourceControlConfigurations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.SourceControlConfigurationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Source Control Configuration ID
func (id ScopedSourceControlConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftKubernetesConfiguration", "Microsoft.KubernetesConfiguration", "Microsoft.KubernetesConfiguration"),
		resourceids.StaticSegment("staticSourceControlConfigurations", "sourceControlConfigurations", "sourceControlConfigurations"),
		resourceids.UserSpecifiedSegment("sourceControlConfigurationName", "sourceControlConfigurationValue"),
	}
}

// String returns a human-readable description of this Scoped Source Control Configuration ID
func (id ScopedSourceControlConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Source Control Configuration Name: %q", id.SourceControlConfigurationName),
	}
	return fmt.Sprintf("Scoped Source Control Configuration (%s)", strings.Join(components, "\n"))
}
