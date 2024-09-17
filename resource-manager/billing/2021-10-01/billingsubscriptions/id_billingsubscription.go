package billingsubscriptions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingSubscriptionId{})
}

var _ resourceids.ResourceId = &BillingSubscriptionId{}

// BillingSubscriptionId is a struct representing the Resource ID for a Billing Subscription
type BillingSubscriptionId struct {
	BillingAccountName      string
	BillingSubscriptionName string
}

// NewBillingSubscriptionID returns a new BillingSubscriptionId struct
func NewBillingSubscriptionID(billingAccountName string, billingSubscriptionName string) BillingSubscriptionId {
	return BillingSubscriptionId{
		BillingAccountName:      billingAccountName,
		BillingSubscriptionName: billingSubscriptionName,
	}
}

// ParseBillingSubscriptionID parses 'input' into a BillingSubscriptionId
func ParseBillingSubscriptionID(input string) (*BillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingSubscriptionIDInsensitively parses 'input' case-insensitively into a BillingSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseBillingSubscriptionIDInsensitively(input string) (*BillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingSubscriptionName, ok = input.Parsed["billingSubscriptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingSubscriptionName", input)
	}

	return nil
}

// ValidateBillingSubscriptionID checks that 'input' can be parsed as a Billing Subscription ID
func ValidateBillingSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Subscription ID
func (id BillingSubscriptionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Subscription ID
func (id BillingSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingSubscriptions", "billingSubscriptions", "billingSubscriptions"),
		resourceids.UserSpecifiedSegment("billingSubscriptionName", "billingSubscriptionValue"),
	}
}

// String returns a human-readable description of this Billing Subscription ID
func (id BillingSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Subscription Name: %q", id.BillingSubscriptionName),
	}
	return fmt.Sprintf("Billing Subscription (%s)", strings.Join(components, "\n"))
}
