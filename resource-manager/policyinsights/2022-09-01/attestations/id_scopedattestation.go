package attestations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ScopedAttestationId{}

// ScopedAttestationId is a struct representing the Resource ID for a Scoped Attestation
type ScopedAttestationId struct {
	ResourceId      string
	AttestationName string
}

// NewScopedAttestationID returns a new ScopedAttestationId struct
func NewScopedAttestationID(resourceId string, attestationName string) ScopedAttestationId {
	return ScopedAttestationId{
		ResourceId:      resourceId,
		AttestationName: attestationName,
	}
}

// ParseScopedAttestationID parses 'input' into a ScopedAttestationId
func ParseScopedAttestationID(input string) (*ScopedAttestationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedAttestationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedAttestationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedAttestationIDInsensitively parses 'input' case-insensitively into a ScopedAttestationId
// note: this method should only be used for API response data and not user input
func ParseScopedAttestationIDInsensitively(input string) (*ScopedAttestationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedAttestationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedAttestationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedAttestationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceId, ok = input.Parsed["resourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceId", input)
	}

	if id.AttestationName, ok = input.Parsed["attestationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attestationName", input)
	}

	return nil
}

// ValidateScopedAttestationID checks that 'input' can be parsed as a Scoped Attestation ID
func ValidateScopedAttestationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedAttestationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Attestation ID
func (id ScopedAttestationId) ID() string {
	fmtString := "/%s/providers/Microsoft.PolicyInsights/attestations/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceId, "/"), id.AttestationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Attestation ID
func (id ScopedAttestationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceId", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticAttestations", "attestations", "attestations"),
		resourceids.UserSpecifiedSegment("attestationName", "attestationValue"),
	}
}

// String returns a human-readable description of this Scoped Attestation ID
func (id ScopedAttestationId) String() string {
	components := []string{
		fmt.Sprintf("Resource: %q", id.ResourceId),
		fmt.Sprintf("Attestation Name: %q", id.AttestationName),
	}
	return fmt.Sprintf("Scoped Attestation (%s)", strings.Join(components, "\n"))
}
