package transfers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingProfileCustomerId{})
}

var _ resourceids.ResourceId = &BillingProfileCustomerId{}

// BillingProfileCustomerId is a struct representing the Resource ID for a Billing Profile Customer
type BillingProfileCustomerId struct {
	BillingAccountName string
	BillingProfileName string
	CustomerName       string
}

// NewBillingProfileCustomerID returns a new BillingProfileCustomerId struct
func NewBillingProfileCustomerID(billingAccountName string, billingProfileName string, customerName string) BillingProfileCustomerId {
	return BillingProfileCustomerId{
		BillingAccountName: billingAccountName,
		BillingProfileName: billingProfileName,
		CustomerName:       customerName,
	}
}

// ParseBillingProfileCustomerID parses 'input' into a BillingProfileCustomerId
func ParseBillingProfileCustomerID(input string) (*BillingProfileCustomerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileCustomerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileCustomerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingProfileCustomerIDInsensitively parses 'input' case-insensitively into a BillingProfileCustomerId
// note: this method should only be used for API response data and not user input
func ParseBillingProfileCustomerIDInsensitively(input string) (*BillingProfileCustomerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileCustomerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileCustomerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingProfileCustomerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.CustomerName, ok = input.Parsed["customerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customerName", input)
	}

	return nil
}

// ValidateBillingProfileCustomerID checks that 'input' can be parsed as a Billing Profile Customer ID
func ValidateBillingProfileCustomerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingProfileCustomerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Profile Customer ID
func (id BillingProfileCustomerId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/customers/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.CustomerName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Profile Customer ID
func (id BillingProfileCustomerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticCustomers", "customers", "customers"),
		resourceids.UserSpecifiedSegment("customerName", "customerValue"),
	}
}

// String returns a human-readable description of this Billing Profile Customer ID
func (id BillingProfileCustomerId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Customer Name: %q", id.CustomerName),
	}
	return fmt.Sprintf("Billing Profile Customer (%s)", strings.Join(components, "\n"))
}
