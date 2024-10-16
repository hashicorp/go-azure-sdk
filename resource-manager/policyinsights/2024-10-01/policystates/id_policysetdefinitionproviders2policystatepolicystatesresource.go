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
	recaser.RegisterResourceId(&PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{})
}

var _ resourceids.ResourceId = &PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{}

// PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId is a struct representing the Resource ID for a Policy Set Definition Providers 2 Policy State Policy States Resource
type PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId struct {
	SubscriptionId          string
	PolicySetDefinitionName string
	PolicyStatesResource    PolicyStatesResource
}

// NewPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID returns a new PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId struct
func NewPolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID(subscriptionId string, policySetDefinitionName string, policyStatesResource PolicyStatesResource) PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId {
	return PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{
		SubscriptionId:          subscriptionId,
		PolicySetDefinitionName: policySetDefinitionName,
		PolicyStatesResource:    policyStatesResource,
	}
}

// ParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID parses 'input' into a PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId
func ParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID(input string) (*PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceIDInsensitively parses 'input' case-insensitively into a PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId
// note: this method should only be used for API response data and not user input
func ParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceIDInsensitively(input string) (*PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.PolicySetDefinitionName, ok = input.Parsed["policySetDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policySetDefinitionName", input)
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

// ValidatePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID checks that 'input' can be parsed as a Policy Set Definition Providers 2 Policy State Policy States Resource ID
func ValidatePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicySetDefinitionProviders2PolicyStatePolicyStatesResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Set Definition Providers 2 Policy State Policy States Resource ID
func (id PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Authorization/policySetDefinitions/%s/providers/Microsoft.PolicyInsights/policyStates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.PolicySetDefinitionName, string(id.PolicyStatesResource))
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Set Definition Providers 2 Policy State Policy States Resource ID
func (id PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.StaticSegment("authorizationNamespace", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicySetDefinitions", "policySetDefinitions", "policySetDefinitions"),
		resourceids.UserSpecifiedSegment("policySetDefinitionName", "policySetDefinitionName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticPolicyStates", "policyStates", "policyStates"),
		resourceids.ConstantSegment("policyStatesResource", PossibleValuesForPolicyStatesResource(), "default"),
	}
}

// String returns a human-readable description of this Policy Set Definition Providers 2 Policy State Policy States Resource ID
func (id PolicySetDefinitionProviders2PolicyStatePolicyStatesResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Policy Set Definition Name: %q", id.PolicySetDefinitionName),
		fmt.Sprintf("Policy States Resource: %q", string(id.PolicyStatesResource)),
	}
	return fmt.Sprintf("Policy Set Definition Providers 2 Policy State Policy States Resource (%s)", strings.Join(components, "\n"))
}
