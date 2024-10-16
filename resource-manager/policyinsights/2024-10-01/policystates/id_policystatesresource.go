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
	recaser.RegisterResourceId(&PolicyStatesResourceId{})
}

var _ resourceids.ResourceId = &PolicyStatesResourceId{}

// PolicyStatesResourceId is a struct representing the Resource ID for a Policy States Resource
type PolicyStatesResourceId struct {
	SubscriptionId       string
	PolicyStatesResource PolicyStatesResource
}

// NewPolicyStatesResourceID returns a new PolicyStatesResourceId struct
func NewPolicyStatesResourceID(subscriptionId string, policyStatesResource PolicyStatesResource) PolicyStatesResourceId {
	return PolicyStatesResourceId{
		SubscriptionId:       subscriptionId,
		PolicyStatesResource: policyStatesResource,
	}
}

// ParsePolicyStatesResourceID parses 'input' into a PolicyStatesResourceId
func ParsePolicyStatesResourceID(input string) (*PolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyStatesResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyStatesResourceIDInsensitively parses 'input' case-insensitively into a PolicyStatesResourceId
// note: this method should only be used for API response data and not user input
func ParsePolicyStatesResourceIDInsensitively(input string) (*PolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyStatesResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyStatesResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
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

// ValidatePolicyStatesResourceID checks that 'input' can be parsed as a Policy States Resource ID
func ValidatePolicyStatesResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyStatesResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy States Resource ID
func (id PolicyStatesResourceId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.PolicyInsights/policyStates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, string(id.PolicyStatesResource))
}

// Segments returns a slice of Resource ID Segments which comprise this Policy States Resource ID
func (id PolicyStatesResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticPolicyStates", "policyStates", "policyStates"),
		resourceids.ConstantSegment("policyStatesResource", PossibleValuesForPolicyStatesResource(), "default"),
	}
}

// String returns a human-readable description of this Policy States Resource ID
func (id PolicyStatesResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Policy States Resource: %q", string(id.PolicyStatesResource)),
	}
	return fmt.Sprintf("Policy States Resource (%s)", strings.Join(components, "\n"))
}
