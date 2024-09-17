package billings

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingAccountBillingSubscriptionId{})
}

var _ resourceids.ResourceId = &BillingAccountBillingSubscriptionId{}

// BillingAccountBillingSubscriptionId is a struct representing the Resource ID for a Billing Account Billing Subscription
type BillingAccountBillingSubscriptionId struct {
	BillingAccountName string
	SubscriptionId     string
}

// NewBillingAccountBillingSubscriptionID returns a new BillingAccountBillingSubscriptionId struct
func NewBillingAccountBillingSubscriptionID(billingAccountName string, subscriptionId string) BillingAccountBillingSubscriptionId {
	return BillingAccountBillingSubscriptionId{
		BillingAccountName: billingAccountName,
		SubscriptionId:     subscriptionId,
	}
}

// ParseBillingAccountBillingSubscriptionID parses 'input' into a BillingAccountBillingSubscriptionId
func ParseBillingAccountBillingSubscriptionID(input string) (*BillingAccountBillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingAccountBillingSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingAccountBillingSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingAccountBillingSubscriptionIDInsensitively parses 'input' case-insensitively into a BillingAccountBillingSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseBillingAccountBillingSubscriptionIDInsensitively(input string) (*BillingAccountBillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingAccountBillingSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingAccountBillingSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingAccountBillingSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	return nil
}

// ValidateBillingAccountBillingSubscriptionID checks that 'input' can be parsed as a Billing Account Billing Subscription ID
func ValidateBillingAccountBillingSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingAccountBillingSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Account Billing Subscription ID
func (id BillingAccountBillingSubscriptionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.SubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Account Billing Subscription ID
func (id BillingAccountBillingSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingSubscriptions", "billingSubscriptions", "billingSubscriptions"),
		resourceids.UserSpecifiedSegment("subscriptionId", "subscriptionIdValue"),
	}
}

// String returns a human-readable description of this Billing Account Billing Subscription ID
func (id BillingAccountBillingSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
	}
	return fmt.Sprintf("Billing Account Billing Subscription (%s)", strings.Join(components, "\n"))
}
