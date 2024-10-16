package policymetadata

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PolicyMetadataId{})
}

var _ resourceids.ResourceId = &PolicyMetadataId{}

// PolicyMetadataId is a struct representing the Resource ID for a Policy Metadata
type PolicyMetadataId struct {
	PolicyMetadataName string
}

// NewPolicyMetadataID returns a new PolicyMetadataId struct
func NewPolicyMetadataID(policyMetadataName string) PolicyMetadataId {
	return PolicyMetadataId{
		PolicyMetadataName: policyMetadataName,
	}
}

// ParsePolicyMetadataID parses 'input' into a PolicyMetadataId
func ParsePolicyMetadataID(input string) (*PolicyMetadataId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMetadataId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMetadataId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyMetadataIDInsensitively parses 'input' case-insensitively into a PolicyMetadataId
// note: this method should only be used for API response data and not user input
func ParsePolicyMetadataIDInsensitively(input string) (*PolicyMetadataId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyMetadataId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyMetadataId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyMetadataId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PolicyMetadataName, ok = input.Parsed["policyMetadataName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policyMetadataName", input)
	}

	return nil
}

// ValidatePolicyMetadataID checks that 'input' can be parsed as a Policy Metadata ID
func ValidatePolicyMetadataID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyMetadataID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Metadata ID
func (id PolicyMetadataId) ID() string {
	fmtString := "/providers/Microsoft.PolicyInsights/policyMetadata/%s"
	return fmt.Sprintf(fmtString, id.PolicyMetadataName)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Metadata ID
func (id PolicyMetadataId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticPolicyMetadata", "policyMetadata", "policyMetadata"),
		resourceids.UserSpecifiedSegment("policyMetadataName", "policyMetadataName"),
	}
}

// String returns a human-readable description of this Policy Metadata ID
func (id PolicyMetadataId) String() string {
	components := []string{
		fmt.Sprintf("Policy Metadata Name: %q", id.PolicyMetadataName),
	}
	return fmt.Sprintf("Policy Metadata (%s)", strings.Join(components, "\n"))
}
