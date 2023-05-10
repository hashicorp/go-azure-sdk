package policysetdefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = PolicySetDefinitionId{}

// PolicySetDefinitionId is a struct representing the Resource ID for a Policy Set Definition
type PolicySetDefinitionId struct {
	PolicySetDefinitionName string
}

// NewPolicySetDefinitionID returns a new PolicySetDefinitionId struct
func NewPolicySetDefinitionID(policySetDefinitionName string) PolicySetDefinitionId {
	return PolicySetDefinitionId{
		PolicySetDefinitionName: policySetDefinitionName,
	}
}

// ParsePolicySetDefinitionID parses 'input' into a PolicySetDefinitionId
func ParsePolicySetDefinitionID(input string) (*PolicySetDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicySetDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicySetDefinitionId{}

	if id.PolicySetDefinitionName, ok = parsed.Parsed["policySetDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policySetDefinitionName", *parsed)
	}

	return &id, nil
}

// ParsePolicySetDefinitionIDInsensitively parses 'input' case-insensitively into a PolicySetDefinitionId
// note: this method should only be used for API response data and not user input
func ParsePolicySetDefinitionIDInsensitively(input string) (*PolicySetDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(PolicySetDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := PolicySetDefinitionId{}

	if id.PolicySetDefinitionName, ok = parsed.Parsed["policySetDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policySetDefinitionName", *parsed)
	}

	return &id, nil
}

// ValidatePolicySetDefinitionID checks that 'input' can be parsed as a Policy Set Definition ID
func ValidatePolicySetDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicySetDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Set Definition ID
func (id PolicySetDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.Authorization/policySetDefinitions/%s"
	return fmt.Sprintf(fmtString, id.PolicySetDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Set Definition ID
func (id PolicySetDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicySetDefinitions", "policySetDefinitions", "policySetDefinitions"),
		resourceids.UserSpecifiedSegment("policySetDefinitionName", "policySetDefinitionValue"),
	}
}

// String returns a human-readable description of this Policy Set Definition ID
func (id PolicySetDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Policy Set Definition Name: %q", id.PolicySetDefinitionName),
	}
	return fmt.Sprintf("Policy Set Definition (%s)", strings.Join(components, "\n"))
}
