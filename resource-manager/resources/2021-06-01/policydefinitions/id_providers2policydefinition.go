package policydefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = Providers2PolicyDefinitionId{}

// Providers2PolicyDefinitionId is a struct representing the Resource ID for a Providers 2 Policy Definition
type Providers2PolicyDefinitionId struct {
	ManagementGroupId    string
	PolicyDefinitionName string
}

// NewProviders2PolicyDefinitionID returns a new Providers2PolicyDefinitionId struct
func NewProviders2PolicyDefinitionID(managementGroupId string, policyDefinitionName string) Providers2PolicyDefinitionId {
	return Providers2PolicyDefinitionId{
		ManagementGroupId:    managementGroupId,
		PolicyDefinitionName: policyDefinitionName,
	}
}

// ParseProviders2PolicyDefinitionID parses 'input' into a Providers2PolicyDefinitionId
func ParseProviders2PolicyDefinitionID(input string) (*Providers2PolicyDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2PolicyDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2PolicyDefinitionId{}

	if id.ManagementGroupId, ok = parsed.Parsed["managementGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managementGroupId", *parsed)
	}

	if id.PolicyDefinitionName, ok = parsed.Parsed["policyDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policyDefinitionName", *parsed)
	}

	return &id, nil
}

// ParseProviders2PolicyDefinitionIDInsensitively parses 'input' case-insensitively into a Providers2PolicyDefinitionId
// note: this method should only be used for API response data and not user input
func ParseProviders2PolicyDefinitionIDInsensitively(input string) (*Providers2PolicyDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(Providers2PolicyDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := Providers2PolicyDefinitionId{}

	if id.ManagementGroupId, ok = parsed.Parsed["managementGroupId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "managementGroupId", *parsed)
	}

	if id.PolicyDefinitionName, ok = parsed.Parsed["policyDefinitionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "policyDefinitionName", *parsed)
	}

	return &id, nil
}

// ValidateProviders2PolicyDefinitionID checks that 'input' can be parsed as a Providers 2 Policy Definition ID
func ValidateProviders2PolicyDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2PolicyDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Policy Definition ID
func (id Providers2PolicyDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Authorization/policyDefinitions/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.PolicyDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Policy Definition ID
func (id Providers2PolicyDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupId", "managementGroupIdValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyDefinitions", "policyDefinitions", "policyDefinitions"),
		resourceids.UserSpecifiedSegment("policyDefinitionName", "policyDefinitionValue"),
	}
}

// String returns a human-readable description of this Providers 2 Policy Definition ID
func (id Providers2PolicyDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Policy Definition Name: %q", id.PolicyDefinitionName),
	}
	return fmt.Sprintf("Providers 2 Policy Definition (%s)", strings.Join(components, "\n"))
}
