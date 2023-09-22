package policydefinitionversions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = VersionId{}

// VersionId is a struct representing the Resource ID for a Version
type VersionId struct {
	PolicyDefinitionName string
	VersionName          string
}

// NewVersionID returns a new VersionId struct
func NewVersionID(policyDefinitionName string, versionName string) VersionId {
	return VersionId{
		PolicyDefinitionName: policyDefinitionName,
		VersionName:          versionName,
	}
}

// ParseVersionID parses 'input' into a VersionId
func ParseVersionID(input string) (*VersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VersionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VersionId{}

	if id.PolicyDefinitionName, ok = parsed.Parsed["policyDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policyDefinitionName", *parsed)
	}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionName", *parsed)
	}

	return &id, nil
}

// ParseVersionIDInsensitively parses 'input' case-insensitively into a VersionId
// note: this method should only be used for API response data and not user input
func ParseVersionIDInsensitively(input string) (*VersionId, error) {
	parser := resourceids.NewParserFromResourceIdType(VersionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := VersionId{}

	if id.PolicyDefinitionName, ok = parsed.Parsed["policyDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policyDefinitionName", *parsed)
	}

	if id.VersionName, ok = parsed.Parsed["versionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "versionName", *parsed)
	}

	return &id, nil
}

// ValidateVersionID checks that 'input' can be parsed as a Version ID
func ValidateVersionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseVersionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Version ID
func (id VersionId) ID() string {
	fmtString := "/providers/Microsoft.Authorization/policyDefinitions/%s/versions/%s"
	return fmt.Sprintf(fmtString, id.PolicyDefinitionName, id.VersionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Version ID
func (id VersionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyDefinitions", "policyDefinitions", "policyDefinitions"),
		resourceids.UserSpecifiedSegment("policyDefinitionName", "policyDefinitionValue"),
		resourceids.StaticSegment("staticVersions", "versions", "versions"),
		resourceids.UserSpecifiedSegment("versionName", "versionValue"),
	}
}

// String returns a human-readable description of this Version ID
func (id VersionId) String() string {
	components := []string{
		fmt.Sprintf("Policy Definition Name: %q", id.PolicyDefinitionName),
		fmt.Sprintf("Version Name: %q", id.VersionName),
	}
	return fmt.Sprintf("Version (%s)", strings.Join(components, "\n"))
}
