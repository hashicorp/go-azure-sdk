package attestations

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AttestationId{}

// AttestationId is a struct representing the Resource ID for a Attestation
type AttestationId struct {
	SubscriptionId  string
	AttestationName string
}

// NewAttestationID returns a new AttestationId struct
func NewAttestationID(subscriptionId string, attestationName string) AttestationId {
	return AttestationId{
		SubscriptionId:  subscriptionId,
		AttestationName: attestationName,
	}
}

// ParseAttestationID parses 'input' into a AttestationId
func ParseAttestationID(input string) (*AttestationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AttestationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AttestationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAttestationIDInsensitively parses 'input' case-insensitively into a AttestationId
// note: this method should only be used for API response data and not user input
func ParseAttestationIDInsensitively(input string) (*AttestationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AttestationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AttestationId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AttestationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.AttestationName, ok = input.Parsed["attestationName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "attestationName", input)
	}

	return nil
}

// ValidateAttestationID checks that 'input' can be parsed as a Attestation ID
func ValidateAttestationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAttestationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Attestation ID
func (id AttestationId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.PolicyInsights/attestations/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.AttestationName)
}

// Segments returns a slice of Resource ID Segments which comprise this Attestation ID
func (id AttestationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticAttestations", "attestations", "attestations"),
		resourceids.UserSpecifiedSegment("attestationName", "attestationValue"),
	}
}

// String returns a human-readable description of this Attestation ID
func (id AttestationId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Attestation Name: %q", id.AttestationName),
	}
	return fmt.Sprintf("Attestation (%s)", strings.Join(components, "\n"))
}
