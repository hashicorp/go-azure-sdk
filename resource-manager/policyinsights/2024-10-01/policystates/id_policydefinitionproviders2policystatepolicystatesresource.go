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
	recaser.RegisterResourceId(&PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId{})
}

var _ resourceids.ResourceId = &PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId{}

// PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId is a struct representing the Resource ID for a Policy Definition Providers 2 Policy State Policy States Resource
type PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId struct {
	SubscriptionId       string
	PolicyDefinitionName string
	PolicyStatesResource PolicyStatesResource
}

// NewPolicyDefinitionProviders2PolicyStatePolicyStatesResourceID returns a new PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId struct
func NewPolicyDefinitionProviders2PolicyStatePolicyStatesResourceID(subscriptionId string, policyDefinitionName string, policyStatesResource PolicyStatesResource) PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId {
	return PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId{
		SubscriptionId:       subscriptionId,
		PolicyDefinitionName: policyDefinitionName,
		PolicyStatesResource: policyStatesResource,
	}
}

// ParsePolicyDefinitionProviders2PolicyStatePolicyStatesResourceID parses 'input' into a PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId
func ParsePolicyDefinitionProviders2PolicyStatePolicyStatesResourceID(input string) (*PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyDefinitionProviders2PolicyStatePolicyStatesResourceIDInsensitively parses 'input' case-insensitively into a PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId
// note: this method should only be used for API response data and not user input
func ParsePolicyDefinitionProviders2PolicyStatePolicyStatesResourceIDInsensitively(input string) (*PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.PolicyDefinitionName, ok = input.Parsed["policyDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policyDefinitionName", input)
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

// ValidatePolicyDefinitionProviders2PolicyStatePolicyStatesResourceID checks that 'input' can be parsed as a Policy Definition Providers 2 Policy State Policy States Resource ID
func ValidatePolicyDefinitionProviders2PolicyStatePolicyStatesResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyDefinitionProviders2PolicyStatePolicyStatesResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Definition Providers 2 Policy State Policy States Resource ID
func (id PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Authorization/policyDefinitions/%s/providers/Microsoft.PolicyInsights/policyStates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.PolicyDefinitionName, string(id.PolicyStatesResource))
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Definition Providers 2 Policy State Policy States Resource ID
func (id PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.StaticSegment("authorizationNamespace", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyDefinitions", "policyDefinitions", "policyDefinitions"),
		resourceids.UserSpecifiedSegment("policyDefinitionName", "policyDefinitionName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticPolicyStates", "policyStates", "policyStates"),
		resourceids.ConstantSegment("policyStatesResource", PossibleValuesForPolicyStatesResource(), "default"),
	}
}

// String returns a human-readable description of this Policy Definition Providers 2 Policy State Policy States Resource ID
func (id PolicyDefinitionProviders2PolicyStatePolicyStatesResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Policy Definition Name: %q", id.PolicyDefinitionName),
		fmt.Sprintf("Policy States Resource: %q", string(id.PolicyStatesResource)),
	}
	return fmt.Sprintf("Policy Definition Providers 2 Policy State Policy States Resource (%s)", strings.Join(components, "\n"))
}
