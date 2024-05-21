package policydefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PolicydefinitionId{})
}

var _ resourceids.ResourceId = &PolicydefinitionId{}

// PolicydefinitionId is a struct representing the Resource ID for a Policydefinition
type PolicydefinitionId struct {
	SubscriptionId       string
	PolicydefinitionName string
}

// NewPolicydefinitionID returns a new PolicydefinitionId struct
func NewPolicydefinitionID(subscriptionId string, policydefinitionName string) PolicydefinitionId {
	return PolicydefinitionId{
		SubscriptionId:       subscriptionId,
		PolicydefinitionName: policydefinitionName,
	}
}

// ParsePolicydefinitionID parses 'input' into a PolicydefinitionId
func ParsePolicydefinitionID(input string) (*PolicydefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicydefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicydefinitionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePolicydefinitionIDInsensitively parses 'input' case-insensitively into a PolicydefinitionId
// note: this method should only be used for API response data and not user input
func ParsePolicydefinitionIDInsensitively(input string) (*PolicydefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PolicydefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PolicydefinitionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PolicydefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.PolicydefinitionName, ok = input.Parsed["policydefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "policydefinitionName", input)
	}

	return nil
}

// ValidatePolicydefinitionID checks that 'input' can be parsed as a Policydefinition ID
func ValidatePolicydefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePolicydefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Policydefinition ID
func (id PolicydefinitionId) ID() string {
	fmtString := "/subscriptions/%s/providers/Microsoft.Authorization/policydefinitions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.PolicydefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Policydefinition ID
func (id PolicydefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftAuthorization", "Microsoft.Authorization", "Microsoft.Authorization"),
		resourceids.StaticSegment("staticPolicydefinitions", "policydefinitions", "policydefinitions"),
		resourceids.UserSpecifiedSegment("policydefinitionName", "policydefinitionValue"),
	}
}

// String returns a human-readable description of this Policydefinition ID
func (id PolicydefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Policydefinition Name: %q", id.PolicydefinitionName),
	}
	return fmt.Sprintf("Policydefinition (%s)", strings.Join(components, "\n"))
}
