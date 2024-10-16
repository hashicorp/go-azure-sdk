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
	recaser.RegisterResourceId(&ScopedPolicyStatesResourceId{})
}

var _ resourceids.ResourceId = &ScopedPolicyStatesResourceId{}

// ScopedPolicyStatesResourceId is a struct representing the Resource ID for a Scoped Policy States Resource
type ScopedPolicyStatesResourceId struct {
	ResourceId           string
	PolicyStatesResource PolicyStatesResource
}

// NewScopedPolicyStatesResourceID returns a new ScopedPolicyStatesResourceId struct
func NewScopedPolicyStatesResourceID(resourceId string, policyStatesResource PolicyStatesResource) ScopedPolicyStatesResourceId {
	return ScopedPolicyStatesResourceId{
		ResourceId:           resourceId,
		PolicyStatesResource: policyStatesResource,
	}
}

// ParseScopedPolicyStatesResourceID parses 'input' into a ScopedPolicyStatesResourceId
func ParseScopedPolicyStatesResourceID(input string) (*ScopedPolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedPolicyStatesResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedPolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseScopedPolicyStatesResourceIDInsensitively parses 'input' case-insensitively into a ScopedPolicyStatesResourceId
// note: this method should only be used for API response data and not user input
func ParseScopedPolicyStatesResourceIDInsensitively(input string) (*ScopedPolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ScopedPolicyStatesResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ScopedPolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ScopedPolicyStatesResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ResourceId, ok = input.Parsed["resourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceId", input)
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

// ValidateScopedPolicyStatesResourceID checks that 'input' can be parsed as a Scoped Policy States Resource ID
func ValidateScopedPolicyStatesResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseScopedPolicyStatesResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Scoped Policy States Resource ID
func (id ScopedPolicyStatesResourceId) ID() string {
	fmtString := "/%s/providers/Microsoft.PolicyInsights/policyStates/%s"
	return fmt.Sprintf(fmtString, strings.TrimPrefix(id.ResourceId, "/"), string(id.PolicyStatesResource))
}

// Segments returns a slice of Resource ID Segments which comprise this Scoped Policy States Resource ID
func (id ScopedPolicyStatesResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.ScopeSegment("resourceId", "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticPolicyStates", "policyStates", "policyStates"),
		resourceids.ConstantSegment("policyStatesResource", PossibleValuesForPolicyStatesResource(), "default"),
	}
}

// String returns a human-readable description of this Scoped Policy States Resource ID
func (id ScopedPolicyStatesResourceId) String() string {
	components := []string{
		fmt.Sprintf("Resource: %q", id.ResourceId),
		fmt.Sprintf("Policy States Resource: %q", string(id.PolicyStatesResource)),
	}
	return fmt.Sprintf("Scoped Policy States Resource (%s)", strings.Join(components, "\n"))
}
