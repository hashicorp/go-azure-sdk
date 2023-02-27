// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package commitmentplans

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = AccountCommitmentPlanId{}

// AccountCommitmentPlanId is a struct representing the Resource ID for a Account Commitment Plan
type AccountCommitmentPlanId struct {
	SubscriptionId     string
	ResourceGroupName  string
	AccountName        string
	CommitmentPlanName string
}

// NewAccountCommitmentPlanID returns a new AccountCommitmentPlanId struct
func NewAccountCommitmentPlanID(subscriptionId string, resourceGroupName string, accountName string, commitmentPlanName string) AccountCommitmentPlanId {
	return AccountCommitmentPlanId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		AccountName:        accountName,
		CommitmentPlanName: commitmentPlanName,
	}
}

// ParseAccountCommitmentPlanID parses 'input' into a AccountCommitmentPlanId
func ParseAccountCommitmentPlanID(input string) (*AccountCommitmentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountCommitmentPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccountCommitmentPlanId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.CommitmentPlanName, ok = parsed.Parsed["commitmentPlanName"]; !ok {
		return nil, fmt.Errorf("the segment 'commitmentPlanName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseAccountCommitmentPlanIDInsensitively parses 'input' case-insensitively into a AccountCommitmentPlanId
// note: this method should only be used for API response data and not user input
func ParseAccountCommitmentPlanIDInsensitively(input string) (*AccountCommitmentPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(AccountCommitmentPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := AccountCommitmentPlanId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.AccountName, ok = parsed.Parsed["accountName"]; !ok {
		return nil, fmt.Errorf("the segment 'accountName' was not found in the resource id %q", input)
	}

	if id.CommitmentPlanName, ok = parsed.Parsed["commitmentPlanName"]; !ok {
		return nil, fmt.Errorf("the segment 'commitmentPlanName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateAccountCommitmentPlanID checks that 'input' can be parsed as a Account Commitment Plan ID
func ValidateAccountCommitmentPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAccountCommitmentPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Account Commitment Plan ID
func (id AccountCommitmentPlanId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/commitmentPlans/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.CommitmentPlanName)
}

// Segments returns a slice of Resource ID Segments which comprise this Account Commitment Plan ID
func (id AccountCommitmentPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountValue"),
		resourceids.StaticSegment("staticCommitmentPlans", "commitmentPlans", "commitmentPlans"),
		resourceids.UserSpecifiedSegment("commitmentPlanName", "commitmentPlanValue"),
	}
}

// String returns a human-readable description of this Account Commitment Plan ID
func (id AccountCommitmentPlanId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Commitment Plan Name: %q", id.CommitmentPlanName),
	}
	return fmt.Sprintf("Account Commitment Plan (%s)", strings.Join(components, "\n"))
}
