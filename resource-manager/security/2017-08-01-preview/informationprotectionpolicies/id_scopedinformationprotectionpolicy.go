package informationprotectionpolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&ScopedInformationProtectionPolicyId{})
}

var _ resourceids.ResourceId = &ScopedInformationProtectionPolicyId{}

// ScopedInformationProtectionPolicyId is a struct representing the Resource ID for a Scoped Information Protection Policy
type ScopedInformationProtectionPolicyId struct {
	Scope                           string
	InformationProtectionPolicyName InformationProtectionPolicyName
}

// NewScopedInformationProtectionPolicyID returns a new ScopedInformationProtectionPolicyId struct
func NewScopedInformationProtectionPolicyID(scope string, informationProtectionPolicyName InformationProtectionPolicyName) ScopedInformationProtectionPolicyId {
	return ScopedInformationProtectionPolicyId{
		Scope:                           scope,
		InformationProtectionPolicyName: informationProtectionPolicyName,
	}
}

// ParseScopedInformationProtectionPolicyID parses 'input' into a ScopedInformationProtectionPolicyId
func ParseScopedInformationProtectionPolicyID(input string) (*ScopedInformationProtectionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedInformationProtectionPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedInformationProtectionPolicyId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedInformationProtectionPolicyIDInsensitively parses 'input' case-insensitively into a ScopedInformationProtectionPolicyId
// note: this method should only be used for API response data and not user input
func ParseScopedInformationProtectionPolicyIDInsensitively(input string) (*ScopedInformationProtectionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedInformationProtectionPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedInformationProtectionPolicyId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedInformationProtectionPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if v, ok := input.Parsed["informationProtectionPolicyName"]; true {
		if !ok {
			return resourceids.NewSegmentNotSpecifiedError(id, "informationProtectionPolicyName", input)
		}

		informationProtectionPolicyName, err := parseInformationProtectionPolicyName(v)
		if err != nil {
			return fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.InformationProtectionPolicyName = *informationProtectionPolicyName
	}

	return nil
}

// ValidateScopedInformationProtectionPolicyID checks that 'input' can be parsed as a Scoped Information Protection Policy ID
func ValidateScopedInformationProtectionPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedInformationProtectionPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Information Protection Policy ID
func (id ScopedInformationProtectionPolicyId) ID() string {
	fmtString := "/%s/providers/Microsoft.Security/informationProtectionPolicies/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), string(id.InformationProtectionPolicyName))
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Information Protection Policy ID
func (id ScopedInformationProtectionPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticInformationProtectionPolicies", "informationProtectionPolicies", "informationProtectionPolicies"),
		resourceids.ConstantSegment("informationProtectionPolicyName", PossibleValuesForInformationProtectionPolicyName(), "custom"),
	}
}

// String returns a human-readable description of this Scoped Information Protection Policy ID
func (id ScopedInformationProtectionPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Information Protection Policy Name: %q", string(id.InformationProtectionPolicyName)),
	}
	return fmt.Sprintf("Scoped Information Protection Policy (%s)", strings.Join(components, "\n"))
}
