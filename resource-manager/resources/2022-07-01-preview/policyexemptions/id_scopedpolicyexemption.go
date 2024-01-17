package policyexemptions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedPolicyExemptionId{}

// ScopedPolicyExemptionId is a struct representing the Resource ID for a Scoped Policy Exemption
type ScopedPolicyExemptionId struct {
	Scope               string
	PolicyExemptionName string
}

// NewScopedPolicyExemptionID returns a new ScopedPolicyExemptionId struct
func NewScopedPolicyExemptionID(scope string, policyExemptionName string) ScopedPolicyExemptionId {
	return ScopedPolicyExemptionId{
		Scope:               scope,
		PolicyExemptionName: policyExemptionName,
	}
}

// ParseScopedPolicyExemptionID parses 'input' into a ScopedPolicyExemptionId
func ParseScopedPolicyExemptionID(input string) (*ScopedPolicyExemptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedPolicyExemptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedPolicyExemptionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedPolicyExemptionIDInsensitively parses 'input' case-insensitively into a ScopedPolicyExemptionId
// note: this method should only be used for API response data and not user input
func ParseScopedPolicyExemptionIDInsensitively(input string) (*ScopedPolicyExemptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedPolicyExemptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedPolicyExemptionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedPolicyExemptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.Scope, ok = input.Parsed["scope"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "scope", input)
	}

	if id.PolicyExemptionName, ok = input.Parsed["policyExemptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policyExemptionName", input)
	}

	return nil
}

// ValidateScopedPolicyExemptionID checks that 'input' can be parsed as a Scoped Policy Exemption ID
func ValidateScopedPolicyExemptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedPolicyExemptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Policy Exemption ID
func (id ScopedPolicyExemptionId) ID() string {
	fmtString := "/%s/providers/Microsoft.Authorization/policyExemptions/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.Scope, "/"), id.PolicyExemptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Policy Exemption ID
func (id ScopedPolicyExemptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("scope", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyExemptions", "policyExemptions", "policyExemptions"),
		resourceids.UserSpecifiedSegment("policyExemptionName", "policyExemptionValue"),
	}
}

// String returns a human-readable description of this Scoped Policy Exemption ID
func (id ScopedPolicyExemptionId) String() string {
	components := []string{
		fmt.Sprintf("Scope: %q", id.Scope),
		fmt.Sprintf("Policy Exemption Name: %q", id.PolicyExemptionName),
	}
	return fmt.Sprintf("Scoped Policy Exemption (%s)", strings.Join(components, "\n"))
}
