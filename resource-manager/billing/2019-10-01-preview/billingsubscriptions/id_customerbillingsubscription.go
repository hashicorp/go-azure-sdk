package billingsubscriptions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CustomerBillingSubscriptionId{}

// CustomerBillingSubscriptionId is a struct representing the Resource ID for a Customer Billing Subscription
type CustomerBillingSubscriptionId struct {
	BillingAccountName      string
	CustomerName            string
	BillingSubscriptionName string
}

// NewCustomerBillingSubscriptionID returns a new CustomerBillingSubscriptionId struct
func NewCustomerBillingSubscriptionID(billingAccountName string, customerName string, billingSubscriptionName string) CustomerBillingSubscriptionId {
	return CustomerBillingSubscriptionId{
		BillingAccountName:      billingAccountName,
		CustomerName:            customerName,
		BillingSubscriptionName: billingSubscriptionName,
	}
}

// ParseCustomerBillingSubscriptionID parses 'input' into a CustomerBillingSubscriptionId
func ParseCustomerBillingSubscriptionID(input string) (*CustomerBillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerBillingSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerBillingSubscriptionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCustomerBillingSubscriptionIDInsensitively parses 'input' case-insensitively into a CustomerBillingSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseCustomerBillingSubscriptionIDInsensitively(input string) (*CustomerBillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerBillingSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerBillingSubscriptionId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CustomerBillingSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.CustomerName, ok = input.Parsed["customerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customerName", input)
	}

	if id.BillingSubscriptionName, ok = input.Parsed["billingSubscriptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingSubscriptionName", input)
	}

	return nil
}

// ValidateCustomerBillingSubscriptionID checks that 'input' can be parsed as a Customer Billing Subscription ID
func ValidateCustomerBillingSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomerBillingSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Customer Billing Subscription ID
func (id CustomerBillingSubscriptionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/customers/%s/billingSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.CustomerName, id.BillingSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Customer Billing Subscription ID
func (id CustomerBillingSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticCustomers", "customers", "customers"),
		resourceids.UserSpecifiedSegment("customerName", "customerValue"),
		resourceids.StaticSegment("staticBillingSubscriptions", "billingSubscriptions", "billingSubscriptions"),
		resourceids.UserSpecifiedSegment("billingSubscriptionName", "billingSubscriptionValue"),
	}
}

// String returns a human-readable description of this Customer Billing Subscription ID
func (id CustomerBillingSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Customer Name: %q", id.CustomerName),
		fmt.Sprintf("Billing Subscription Name: %q", id.BillingSubscriptionName),
	}
	return fmt.Sprintf("Customer Billing Subscription (%s)", strings.Join(components, "\n"))
}
