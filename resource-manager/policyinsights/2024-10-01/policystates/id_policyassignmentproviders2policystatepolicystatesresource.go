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
	recaser.RegisterResourceId(&PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{})
}

var _ resourceids.ResourceId = &PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{}

// PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId is a struct representing the Resource ID for a Policy Assignment Providers 2 Policy State Policy States Resource
type PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId struct {
	SubscriptionId       string
	PolicyAssignmentName string
	PolicyStatesResource PolicyStatesResource
}

// NewPolicyAssignmentProviders2PolicyStatePolicyStatesResourceID returns a new PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId struct
func NewPolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(subscriptionId string, policyAssignmentName string, policyStatesResource PolicyStatesResource) PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId {
	return PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{
		SubscriptionId:       subscriptionId,
		PolicyAssignmentName: policyAssignmentName,
		PolicyStatesResource: policyStatesResource,
	}
}

// ParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID parses 'input' into a PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId
func ParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(input string) (*PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceIDInsensitively parses 'input' case-insensitively into a PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId
// note: this method should only be used for API response data and not user input
func ParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceIDInsensitively(input string) (*PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.PolicyAssignmentName, ok = input.Parsed["policyAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policyAssignmentName", input)
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

// ValidatePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID checks that 'input' can be parsed as a Policy Assignment Providers 2 Policy State Policy States Resource ID
func ValidatePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Assignment Providers 2 Policy State Policy States Resource ID
func (id PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Authorization/policyAssignments/%s/providers/Microsoft.PolicyInsights/policyStates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.PolicyAssignmentName, string(id.PolicyStatesResource))
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Assignment Providers 2 Policy State Policy States Resource ID
func (id PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.StaticSegment("authorizationNamespace", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyAssignments", "policyAssignments", "policyAssignments"),
		resourceids.UserSpecifiedSegment("policyAssignmentName", "policyAssignmentName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticPolicyStates", "policyStates", "policyStates"),
		resourceids.ConstantSegment("policyStatesResource", PossibleValuesForPolicyStatesResource(), "default"),
	}
}

// String returns a human-readable description of this Policy Assignment Providers 2 Policy State Policy States Resource ID
func (id PolicyAssignmentProviders2PolicyStatePolicyStatesResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Policy Assignment Name: %q", id.PolicyAssignmentName),
		fmt.Sprintf("Policy States Resource: %q", string(id.PolicyStatesResource)),
	}
	return fmt.Sprintf("Policy Assignment Providers 2 Policy State Policy States Resource (%s)", strings.Join(components, "\n"))
}
