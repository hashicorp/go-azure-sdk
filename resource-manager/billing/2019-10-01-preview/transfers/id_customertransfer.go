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
	recaser.RegisterResourceId(&CustomerTransferId{})
}

var _ resourceids.ResourceId = &CustomerTransferId{}

// CustomerTransferId is a struct representing the Resource ID for a Customer Transfer
type CustomerTransferId struct {
	BillingAccountName string
	BillingProfileName string
	CustomerName       string
	TransferName       string
}

// NewCustomerTransferID returns a new CustomerTransferId struct
func NewCustomerTransferID(billingAccountName string, billingProfileName string, customerName string, transferName string) CustomerTransferId {
	return CustomerTransferId{
		BillingAccountName: billingAccountName,
		BillingProfileName: billingProfileName,
		CustomerName:       customerName,
		TransferName:       transferName,
	}
}

// ParseCustomerTransferID parses 'input' into a CustomerTransferId
func ParseCustomerTransferID(input string) (*CustomerTransferId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerTransferId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerTransferId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCustomerTransferIDInsensitively parses 'input' case-insensitively into a CustomerTransferId
// note: this method should only be used for API response data and not user input
func ParseCustomerTransferIDInsensitively(input string) (*CustomerTransferId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerTransferId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerTransferId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CustomerTransferId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.TransferName, ok = input.Parsed["transferName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "transferName", input)
	}

	return nil
}

// ValidateCustomerTransferID checks that 'input' can be parsed as a Customer Transfer ID
func ValidateCustomerTransferID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomerTransferID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Customer Transfer ID
func (id CustomerTransferId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/customers/%s/transfers/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.CustomerName, id.TransferName)
}

// Segments returns a slice of Resource ID Segments which comprise this Customer Transfer ID
func (id CustomerTransferId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticCustomers", "customers", "customers"),
		resourceids.UserSpecifiedSegment("customerName", "customerValue"),
		resourceids.StaticSegment("staticTransfers", "transfers", "transfers"),
		resourceids.UserSpecifiedSegment("transferName", "transferValue"),
	}
}

// String returns a human-readable description of this Customer Transfer ID
func (id CustomerTransferId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Customer Name: %q", id.CustomerName),
		fmt.Sprintf("Transfer Name: %q", id.TransferName),
	}
	return fmt.Sprintf("Customer Transfer (%s)", strings.Join(components, "\n"))
}
