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
	recaser.RegisterResourceId(&AuthorizationNamespacePolicyAssignmentId{})
}

var _ resourceids.ResourceId = &AuthorizationNamespacePolicyAssignmentId{}

// AuthorizationNamespacePolicyAssignmentId is a struct representing the Resource ID for a Authorization Namespace Policy Assignment
type AuthorizationNamespacePolicyAssignmentId struct {
	SubscriptionId       string
	ResourceGroupName    string
	PolicyAssignmentName string
}

// NewAuthorizationNamespacePolicyAssignmentID returns a new AuthorizationNamespacePolicyAssignmentId struct
func NewAuthorizationNamespacePolicyAssignmentID(subscriptionId string, resourceGroupName string, policyAssignmentName string) AuthorizationNamespacePolicyAssignmentId {
	return AuthorizationNamespacePolicyAssignmentId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		PolicyAssignmentName: policyAssignmentName,
	}
}

// ParseAuthorizationNamespacePolicyAssignmentID parses 'input' into a AuthorizationNamespacePolicyAssignmentId
func ParseAuthorizationNamespacePolicyAssignmentID(input string) (*AuthorizationNamespacePolicyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuthorizationNamespacePolicyAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuthorizationNamespacePolicyAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAuthorizationNamespacePolicyAssignmentIDInsensitively parses 'input' case-insensitively into a AuthorizationNamespacePolicyAssignmentId
// note: this method should only be used for API response data and not user input
func ParseAuthorizationNamespacePolicyAssignmentIDInsensitively(input string) (*AuthorizationNamespacePolicyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AuthorizationNamespacePolicyAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AuthorizationNamespacePolicyAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AuthorizationNamespacePolicyAssignmentId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateAuthorizationNamespacePolicyAssignmentID checks that 'input' can be parsed as a Authorization Namespace Policy Assignment ID
func ValidateAuthorizationNamespacePolicyAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAuthorizationNamespacePolicyAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Authorization Namespace Policy Assignment ID
func (id AuthorizationNamespacePolicyAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Authorization/policyAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.PolicyAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Authorization Namespace Policy Assignment ID
func (id AuthorizationNamespacePolicyAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.StaticSegment("authorizationNamespace", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyAssignments", "policyAssignments", "policyAssignments"),
		resourceids.UserSpecifiedSegment("policyAssignmentName", "policyAssignmentName"),
	}
}

// String returns a human-readable description of this Authorization Namespace Policy Assignment ID
func (id AuthorizationNamespacePolicyAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Policy Assignment Name: %q", id.PolicyAssignmentName),
	}
	return fmt.Sprintf("Authorization Namespace Policy Assignment (%s)", strings.Join(components, "\n"))
}
