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
	recaser.RegisterResourceId(&PolicyStatePolicyStatesResourceId{})
}

var _ resourceids.ResourceId = &PolicyStatePolicyStatesResourceId{}

// PolicyStatePolicyStatesResourceId is a struct representing the Resource ID for a Policy State Policy States Resource
type PolicyStatePolicyStatesResourceId struct {
	SubscriptionId       string
	ResourceGroupName    string
	PolicyStatesResource PolicyStatesResource
}

// NewPolicyStatePolicyStatesResourceID returns a new PolicyStatePolicyStatesResourceId struct
func NewPolicyStatePolicyStatesResourceID(subscriptionId string, resourceGroupName string, policyStatesResource PolicyStatesResource) PolicyStatePolicyStatesResourceId {
	return PolicyStatePolicyStatesResourceId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		PolicyStatesResource: policyStatesResource,
	}
}

// ParsePolicyStatePolicyStatesResourceID parses 'input' into a PolicyStatePolicyStatesResourceId
func ParsePolicyStatePolicyStatesResourceID(input string) (*PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyStatePolicyStatesResourceIDInsensitively parses 'input' case-insensitively into a PolicyStatePolicyStatesResourceId
// note: this method should only be used for API response data and not user input
func ParsePolicyStatePolicyStatesResourceIDInsensitively(input string) (*PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyStatePolicyStatesResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
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

// ValidatePolicyStatePolicyStatesResourceID checks that 'input' can be parsed as a Policy State Policy States Resource ID
func ValidatePolicyStatePolicyStatesResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyStatePolicyStatesResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy State Policy States Resource ID
func (id PolicyStatePolicyStatesResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.PolicyInsights/policyStates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, string(id.PolicyStatesResource))
}

// Segments returns a slice of Resource ID Segments which comprise this Policy State Policy States Resource ID
func (id PolicyStatePolicyStatesResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPolicyInsights", "Microsoft.PolicyInsights", "Microsoft.PolicyInsights"),
		resourceids.StaticSegment("staticPolicyStates", "policyStates", "policyStates"),
		resourceids.ConstantSegment("policyStatesResource", PossibleValuesForPolicyStatesResource(), "default"),
	}
}

// String returns a human-readable description of this Policy State Policy States Resource ID
func (id PolicyStatePolicyStatesResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Policy States Resource: %q", string(id.PolicyStatesResource)),
	}
	return fmt.Sprintf("Policy State Policy States Resource (%s)", strings.Join(components, "\n"))
}
