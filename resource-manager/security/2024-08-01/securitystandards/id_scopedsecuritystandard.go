package securitystandards

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedSecurityStandardId{})
}

var _ resourceids.ResourceId = &ScopedSecurityStandardId{}

// ScopedSecurityStandardId is a struct representing the Resource ID for a Scoped Security Standard
type ScopedSecurityStandardId struct {
	Scope      string
	StandardId string
}

// NewScopedSecurityStandardID returns a new ScopedSecurityStandardId struct
func NewScopedSecurityStandardID(scope string, standardId string) ScopedSecurityStandardId {
	return ScopedSecurityStandardId{
		Scope:      scope,
		StandardId: standardId,
	}
}

// ParseScopedSecurityStandardID parses 'input' into a ScopedSecurityStandardId
func ParseScopedSecurityStandardID(input string) (*ScopedSecurityStandardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedSecurityStandardId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedSecurityStandardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedSecurityStandardIDInsensitively parses 'input' case-insensitively into a ScopedSecurityStandardId
// note: this method should only be used for API response data and not user input
func ParseScopedSecurityStandardIDInsensitively(input string) (*ScopedSecurityStandardId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedSecurityStandardId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedSecurityStandardId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedSecurityStandardId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.StandardId, ok = input.Parsed["standardId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "standardId", input)
	}

	return nil
}

// ValidateScopedSecurityStandardID checks that 'input' can be parsed as a Scoped Security Standard ID
func ValidateScopedSecurityStandardID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedSecurityStandardID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Security Standard ID
func (id ScopedSecurityStandardId) ID() string {
	fmtString := "/%s/providers/Microsoft.Security/securityStandards/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.StandardId)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Security Standard ID
func (id ScopedSecurityStandardId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticSecurityStandards", "securityStandards", "securityStandards"),
		resourceids.UserSpecifiedSegment("standardId", "standardId"),
	}
}

// String returns a human-readable description of this Scoped Security Standard ID
func (id ScopedSecurityStandardId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Standard: %q", id.StandardId),
	}
	return fmt.Sprintf("Scoped Security Standard (%s)", strings.Join(components, "\n"))
}
