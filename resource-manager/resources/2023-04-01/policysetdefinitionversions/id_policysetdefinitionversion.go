package policysetdefinitionversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &PolicySetDefinitionVersionId{}

// PolicySetDefinitionVersionId is a struct representing the Resource ID for a Policy Set Definition Version
type PolicySetDefinitionVersionId struct {
	PolicySetDefinitionName string
	VersionName             string
}

// NewPolicySetDefinitionVersionID returns a new PolicySetDefinitionVersionId struct
func NewPolicySetDefinitionVersionID(policySetDefinitionName string, versionName string) PolicySetDefinitionVersionId {
	return PolicySetDefinitionVersionId{
		PolicySetDefinitionName: policySetDefinitionName,
		VersionName:             versionName,
	}
}

// ParsePolicySetDefinitionVersionID parses 'input' into a PolicySetDefinitionVersionId
func ParsePolicySetDefinitionVersionID(input string) (*PolicySetDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicySetDefinitionVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicySetDefinitionVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicySetDefinitionVersionIDInsensitively parses 'input' case-insensitively into a PolicySetDefinitionVersionId
// note: this method should only be used for API response data and not user input
func ParsePolicySetDefinitionVersionIDInsensitively(input string) (*PolicySetDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicySetDefinitionVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicySetDefinitionVersionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicySetDefinitionVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PolicySetDefinitionName, ok = input.Parsed["policySetDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policySetDefinitionName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidatePolicySetDefinitionVersionID checks that 'input' can be parsed as a Policy Set Definition Version ID
func ValidatePolicySetDefinitionVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicySetDefinitionVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Set Definition Version ID
func (id PolicySetDefinitionVersionId) ID() string {
	fmtString := "/providers/Microsoft.Authorization/policySetDefinitions/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.PolicySetDefinitionName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Set Definition Version ID
func (id PolicySetDefinitionVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicySetDefinitions", "policySetDefinitions", "policySetDefinitions"),
		resourceids.UserSpecifiedSegment("policySetDefinitionName", "policySetDefinitionValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionValue"),
	}
}

// String returns a human-readable description of this Policy Set Definition Version ID
func (id PolicySetDefinitionVersionId) String() string {
	components := []string{
		fmt.Sprintf("Policy Set Definition Name: %q", id.PolicySetDefinitionName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Policy Set Definition Version (%s)", strings.Join(components, "\n"))
}
