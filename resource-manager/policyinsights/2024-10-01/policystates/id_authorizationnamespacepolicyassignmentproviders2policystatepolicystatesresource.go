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
	recaser.RegisterResourceId(&AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{})
}

var _ resourceids.ResourceId = &AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{}

// AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId is a struct representing the Resource ID for a Authorization Namespace Policy Assignment Providers 2 Policy State Policy States Resource
type AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId struct {
	SubscriptionId       string
	ResourceGroupName    string
	PolicyAssignmentName string
	PolicyStatesResource PolicyStatesResource
}

// NewAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID returns a new AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId struct
func NewAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(subscriptionId string, resourceGroupName string, policyAssignmentName string, policyStatesResource PolicyStatesResource) AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId {
	return AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		PolicyAssignmentName: policyAssignmentName,
		PolicyStatesResource: policyStatesResource,
	}
}

// ParseAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID parses 'input' into a AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId
func ParseAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(input string) (*AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceIDInsensitively parses 'input' case-insensitively into a AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId
// note: this method should only be used for API response data and not user input
func ParseAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceIDInsensitively(input string) (*AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
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

// ValidateAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID checks that 'input' can be parsed as a Authorization Namespace Policy Assignment Providers 2 Policy State Policy States Resource ID
func ValidateAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Authorization Namespace Policy Assignment Providers 2 Policy State Policy States Resource ID
func (id AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Authorization/policyAssignments/%s/providers/Microsoft.PolicyInsights/policyStates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PolicyAssignmentName, string(id.PolicyStatesResource))
}

// Segments returns a slice of Resource ID Segments which comprise this Authorization Namespace Policy Assignment Providers 2 Policy State Policy States Resource ID
func (id AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
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

// String returns a human-readable description of this Authorization Namespace Policy Assignment Providers 2 Policy State Policy States Resource ID
func (id AuthorizationNamespacePolicyAssignmentProviders2PolicyStatePolicyStatesResourceId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Policy Assignment Name: %q", id.PolicyAssignmentName),
		fmt.Sprintf("Policy States Resource: %q", string(id.PolicyStatesResource)),
	}
	return fmt.Sprintf("Authorization Namespace Policy Assignment Providers 2 Policy State Policy States Resource (%s)", strings.Join(components, "\n"))
}
