package policy

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AttestationTypeId{}

// AttestationTypeId is a struct representing the Resource ID for a Attestation Type
type AttestationTypeId struct {
	BaseURI         string
	AttestationType AttestationType
}

// NewAttestationTypeID returns a new AttestationTypeId struct
func NewAttestationTypeID(baseURI string, attestationType AttestationType) AttestationTypeId {
	return AttestationTypeId{
		BaseURI:         strings.TrimSuffix(baseURI, "/"),
		AttestationType: attestationType,
	}
}

// ParseAttestationTypeID parses 'input' into a AttestationTypeId
func ParseAttestationTypeID(input string) (*AttestationTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AttestationTypeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AttestationTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAttestationTypeIDInsensitively parses 'input' case-insensitively into a AttestationTypeId
// note: this method should only be used for API response data and not user input
func ParseAttestationTypeIDInsensitively(input string) (*AttestationTypeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AttestationTypeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AttestationTypeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AttestationTypeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BaseURI, ok = input.Parsed["baseURI"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "baseURI", input)
	}

	if v, ok := input.Parsed["attestationType"]; true {
		if !ok {
			return resourceids.NewSegmentNotSpecifiedError(id, "attestationType", input)
		}

		attestationType, err := parseAttestationType(v)
		if err != nil {
			return fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.AttestationType = *attestationType
	}

	return nil
}

// ValidateAttestationTypeID checks that 'input' can be parsed as a Attestation Type ID
func ValidateAttestationTypeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAttestationTypeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Attestation Type ID
func (id AttestationTypeId) ID() string {
	fmtString := "%s/policies/%s"
	return fmt.Sprintf(fmtString, strings.TrimSuffix(id.BaseURI, "/"), string(id.AttestationType))
}

// Path returns the formatted Attestation Type ID without the BaseURI
func (id AttestationTypeId) Path() string {
	fmtString := "/policies/%s"
	return fmt.Sprintf(fmtString, string(id.AttestationType))
}

// PathElements returns the values of Attestation Type ID Segments without the BaseURI
func (id AttestationTypeId) PathElements() []any {
	return []any{string(id.AttestationType)}
}

// Segments returns a slice of Resource ID Segments which comprise this Attestation Type ID
func (id AttestationTypeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.DataPlaneBaseURISegment("baseURI", "https://endpoint-url.example.com"),
		resourceids.StaticSegment("staticPolicies", "policies", "policies"),
		resourceids.ConstantSegment("attestationType", PossibleValuesForAttestationType(), "OpenEnclave"),
	}
}

// String returns a human-readable description of this Attestation Type ID
func (id AttestationTypeId) String() string {
	components := []string{
		fmt.Sprintf("Base URI: %q", id.BaseURI),
		fmt.Sprintf("Attestation Type: %q", string(id.AttestationType)),
	}
	return fmt.Sprintf("Attestation Type (%s)", strings.Join(components, "\n"))
}
