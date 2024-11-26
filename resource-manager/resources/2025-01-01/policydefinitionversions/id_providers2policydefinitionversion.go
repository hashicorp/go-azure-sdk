package policydefinitionversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&Providers2PolicyDefinitionVersionId{})
}

var _ resourceids.ResourceId = &Providers2PolicyDefinitionVersionId{}

// Providers2PolicyDefinitionVersionId is a struct representing the Resource ID for a Providers 2 Policy Definition Version
type Providers2PolicyDefinitionVersionId struct {
	ManagementGroupName  string
	PolicyDefinitionName string
	VersionName          string
}

// NewProviders2PolicyDefinitionVersionID returns a new Providers2PolicyDefinitionVersionId struct
func NewProviders2PolicyDefinitionVersionID(managementGroupName string, policyDefinitionName string, versionName string) Providers2PolicyDefinitionVersionId {
	return Providers2PolicyDefinitionVersionId{
		ManagementGroupName:  managementGroupName,
		PolicyDefinitionName: policyDefinitionName,
		VersionName:          versionName,
	}
}

// ParseProviders2PolicyDefinitionVersionID parses 'input' into a Providers2PolicyDefinitionVersionId
func ParseProviders2PolicyDefinitionVersionID(input string) (*Providers2PolicyDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2PolicyDefinitionVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2PolicyDefinitionVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviders2PolicyDefinitionVersionIDInsensitively parses 'input' case-insensitively into a Providers2PolicyDefinitionVersionId
// note: this method should only be used for API response data and not user input
func ParseProviders2PolicyDefinitionVersionIDInsensitively(input string) (*Providers2PolicyDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2PolicyDefinitionVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2PolicyDefinitionVersionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *Providers2PolicyDefinitionVersionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagementGroupName, ok = input.Parsed["managementGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementGroupName", input)
	}

	if id.PolicyDefinitionName, ok = input.Parsed["policyDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policyDefinitionName", input)
	}

	if id.VersionName, ok = input.Parsed["versionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "versionName", input)
	}

	return nil
}

// ValidateProviders2PolicyDefinitionVersionID checks that 'input' can be parsed as a Providers 2 Policy Definition Version ID
func ValidateProviders2PolicyDefinitionVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2PolicyDefinitionVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Policy Definition Version ID
func (id Providers2PolicyDefinitionVersionId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Authorization/policyDefinitions/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupName, id.PolicyDefinitionName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Policy Definition Version ID
func (id Providers2PolicyDefinitionVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupName", "managementGroupName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyDefinitions", "policyDefinitions", "policyDefinitions"),
		resourceids.UserSpecifiedSegment("policyDefinitionName", "policyDefinitionName"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionName"),
	}
}

// String returns a human-readable description of this Providers 2 Policy Definition Version ID
func (id Providers2PolicyDefinitionVersionId) String() string {
	components := []string{
		fmt.Sprintf("Management Group Name: %q", id.ManagementGroupName),
		fmt.Sprintf("Policy Definition Name: %q", id.PolicyDefinitionName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Providers 2 Policy Definition Version (%s)", strings.Join(components, "\n"))
}
