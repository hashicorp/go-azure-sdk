package policyevents

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PolicyAssignmentId{})
}

var _ resourceids.ResourceId = &PolicyAssignmentId{}

// PolicyAssignmentId is a struct representing the Resource ID for a Policy Assignment
type PolicyAssignmentId struct {
	SubscriptionId       string
	PolicyAssignmentName string
}

// NewPolicyAssignmentID returns a new PolicyAssignmentId struct
func NewPolicyAssignmentID(subscriptionId string, policyAssignmentName string) PolicyAssignmentId {
	return PolicyAssignmentId{
		SubscriptionId:       subscriptionId,
		PolicyAssignmentName: policyAssignmentName,
	}
}

// ParsePolicyAssignmentID parses 'input' into a PolicyAssignmentId
func ParsePolicyAssignmentID(input string) (*PolicyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicyAssignmentIDInsensitively parses 'input' case-insensitively into a PolicyAssignmentId
// note: this method should only be used for API response data and not user input
func ParsePolicyAssignmentIDInsensitively(input string) (*PolicyAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicyAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicyAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicyAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.PolicyAssignmentName, ok = input.Parsed["policyAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policyAssignmentName", input)
	}

	return nil
}

// ValidatePolicyAssignmentID checks that 'input' can be parsed as a Policy Assignment ID
func ValidatePolicyAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicyAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policy Assignment ID
func (id PolicyAssignmentId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Authorization/policyAssignments/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.PolicyAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Policy Assignment ID
func (id PolicyAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.StaticSegment("authorizationNamespace", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicyAssignments", "policyAssignments", "policyAssignments"),
		resourceids.UserSpecifiedSegment("policyAssignmentName", "policyAssignmentName"),
	}
}

// String returns a human-readable description of this Policy Assignment ID
func (id PolicyAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Policy Assignment Name: %q", id.PolicyAssignmentName),
	}
	return fmt.Sprintf("Policy Assignment (%s)", strings.Join(components, "\n"))
}
