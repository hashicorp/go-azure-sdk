package policysetdefinitionversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2PolicySetDefinitionVersionId{}

// Providers2PolicySetDefinitionVersionId is a struct representing the Resource ID for a Providers 2 Policy Set Definition Version
type Providers2PolicySetDefinitionVersionId struct {
	ManagementGroupName     string
	PolicySetDefinitionName string
	VersionName             string
}

// NewProviders2PolicySetDefinitionVersionID returns a new Providers2PolicySetDefinitionVersionId struct
func NewProviders2PolicySetDefinitionVersionID(managementGroupName string, policySetDefinitionName string, versionName string) Providers2PolicySetDefinitionVersionId {
	return Providers2PolicySetDefinitionVersionId{
		ManagementGroupName:     managementGroupName,
		PolicySetDefinitionName: policySetDefinitionName,
		VersionName:             versionName,
	}
}

// ParseProviders2PolicySetDefinitionVersionID parses 'input' into a Providers2PolicySetDefinitionVersionId
func ParseProviders2PolicySetDefinitionVersionID(input string) (*Providers2PolicySetDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2PolicySetDefinitionVersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2PolicySetDefinitionVersionId{}

	if id.ManagementGroupName, ok = parsed.Parsed["managementGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managementGroupName", *parsed)
	}

	if id.PolicySetDefinitionName, ok = parsed.Parsed["policySetDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policySetDefinitionName", *parsed)
	}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionName", *parsed)
	}

	return &id, nil
}

// ParseProviders2PolicySetDefinitionVersionIDInsensitively parses 'input' case-insensitively into a Providers2PolicySetDefinitionVersionId
// note: this method should only be used for API response data and not user input
func ParseProviders2PolicySetDefinitionVersionIDInsensitively(input string) (*Providers2PolicySetDefinitionVersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2PolicySetDefinitionVersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2PolicySetDefinitionVersionId{}

	if id.ManagementGroupName, ok = parsed.Parsed["managementGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managementGroupName", *parsed)
	}

	if id.PolicySetDefinitionName, ok = parsed.Parsed["policySetDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policySetDefinitionName", *parsed)
	}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionName", *parsed)
	}

	return &id, nil
}

// ValidateProviders2PolicySetDefinitionVersionID checks that 'input' can be parsed as a Providers 2 Policy Set Definition Version ID
func ValidateProviders2PolicySetDefinitionVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2PolicySetDefinitionVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Policy Set Definition Version ID
func (id Providers2PolicySetDefinitionVersionId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Authorization/policySetDefinitions/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupName, id.PolicySetDefinitionName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Policy Set Definition Version ID
func (id Providers2PolicySetDefinitionVersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupName", "managementGroupValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicySetDefinitions", "policySetDefinitions", "policySetDefinitions"),
		resourceids.UserSpecifiedSegment("policySetDefinitionName", "policySetDefinitionValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionValue"),
	}
}

// String returns a human-readable description of this Providers 2 Policy Set Definition Version ID
func (id Providers2PolicySetDefinitionVersionId) String() string {
	components := []string{
		fmt.Sprintf("Management Group Name: %q", id.ManagementGroupName),
		fmt.Sprintf("Policy Set Definition Name: %q", id.PolicySetDefinitionName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Providers 2 Policy Set Definition Version (%s)", strings.Join(components, "\n"))
}
