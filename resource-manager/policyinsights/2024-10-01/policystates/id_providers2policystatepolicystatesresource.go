package policystates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&Providers2PolicyStatePolicyStatesResourceId{})
}

var _ resourceids.ResourceId = &Providers2PolicyStatePolicyStatesResourceId{}

// Providers2PolicyStatePolicyStatesResourceId is a struct representing the Resource ID for a Providers 2 Policy State Policy States Resource
type Providers2PolicyStatePolicyStatesResourceId struct {
	ManagementGroupName  string
	PolicyStatesResource PolicyStatesResource
}

// NewProviders2PolicyStatePolicyStatesResourceID returns a new Providers2PolicyStatePolicyStatesResourceId struct
func NewProviders2PolicyStatePolicyStatesResourceID(managementGroupName string, policyStatesResource PolicyStatesResource) Providers2PolicyStatePolicyStatesResourceId {
	return Providers2PolicyStatePolicyStatesResourceId{
		ManagementGroupName:  managementGroupName,
		PolicyStatesResource: policyStatesResource,
	}
}

// ParseProviders2PolicyStatePolicyStatesResourceID parses 'input' into a Providers2PolicyStatePolicyStatesResourceId
func ParseProviders2PolicyStatePolicyStatesResourceID(input string) (*Providers2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviders2PolicyStatePolicyStatesResourceIDInsensitively parses 'input' case-insensitively into a Providers2PolicyStatePolicyStatesResourceId
// note: this method should only be used for API response data and not user input
func ParseProviders2PolicyStatePolicyStatesResourceIDInsensitively(input string) (*Providers2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *Providers2PolicyStatePolicyStatesResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagementGroupName, ok = input.Parsed["managementGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementGroupName", input)
	}

	if v, ok := input.Parsed["policyStatesResource"]; true {
		if !ok {
			return resourceids.NewSegmentNotSpecifiedError(id, "policyStatesResource", input)
		}

		policyStatesResource, err := parsePolicyStatesResource(v)
		if err != nil {
			return fmt.Errorf("parsing %q: %+v", v, err)
		}
		id.PolicyStatesResource = *policyStatesResource
	}

	return nil
}

// ValidateProviders2PolicyStatePolicyStatesResourceID checks that 'input' can be parsed as a Providers 2 Policy State Policy States Resource ID
func ValidateProviders2PolicyStatePolicyStatesResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2PolicyStatePolicyStatesResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Policy State Policy States Resource ID
func (id Providers2PolicyStatePolicyStatesResourceId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.PolicyInsights/policyStates/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupName, string(id.PolicyStatesResource))
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Policy State Policy States Resource ID
func (id Providers2PolicyStatePolicyStatesResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.StaticSegment("managementGroupsNamespace", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupName", "managementGroupName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticPolicyStates", "policyStates", "policyStates"),
		resourceids.ConstantSegment("policyStatesResource", PossibleValuesForPolicyStatesResource(), "default"),
	}
}

// String returns a human-readable description of this Providers 2 Policy State Policy States Resource ID
func (id Providers2PolicyStatePolicyStatesResourceId) String() string {
	components := []string{
		fmt.Sprintf("Management Group Name: %q", id.ManagementGroupName),
		fmt.Sprintf("Policy States Resource: %q", string(id.PolicyStatesResource)),
	}
	return fmt.Sprintf("Providers 2 Policy State Policy States Resource (%s)", strings.Join(components, "\n"))
}
