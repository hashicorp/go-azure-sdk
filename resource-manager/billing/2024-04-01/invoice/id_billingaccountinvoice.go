package invoice

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingAccountInvoiceId{})
}

var _ resourceids.ResourceId = &BillingAccountInvoiceId{}

// BillingAccountInvoiceId is a struct representing the Resource ID for a Billing Account Invoice
type BillingAccountInvoiceId struct {
	BillingAccountName string
	InvoiceName        string
}

// NewBillingAccountInvoiceID returns a new BillingAccountInvoiceId struct
func NewBillingAccountInvoiceID(billingAccountName string, invoiceName string) BillingAccountInvoiceId {
	return BillingAccountInvoiceId{
		BillingAccountName: billingAccountName,
		InvoiceName:        invoiceName,
	}
}

// ParseBillingAccountInvoiceID parses 'input' into a BillingAccountInvoiceId
func ParseBillingAccountInvoiceID(input string) (*BillingAccountInvoiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingAccountInvoiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingAccountInvoiceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingAccountInvoiceIDInsensitively parses 'input' case-insensitively into a BillingAccountInvoiceId
// note: this method should only be used for API response data and not user input
func ParseBillingAccountInvoiceIDInsensitively(input string) (*BillingAccountInvoiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingAccountInvoiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingAccountInvoiceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingAccountInvoiceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.InvoiceName, ok = input.Parsed["invoiceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "invoiceName", input)
	}

	return nil
}

// ValidateBillingAccountInvoiceID checks that 'input' can be parsed as a Billing Account Invoice ID
func ValidateBillingAccountInvoiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingAccountInvoiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Account Invoice ID
func (id BillingAccountInvoiceId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/invoices/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.InvoiceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Account Invoice ID
func (id BillingAccountInvoiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticInvoices", "invoices", "invoices"),
		resourceids.UserSpecifiedSegment("invoiceName", "invoiceValue"),
	}
}

// String returns a human-readable description of this Billing Account Invoice ID
func (id BillingAccountInvoiceId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Invoice Name: %q", id.InvoiceName),
	}
	return fmt.Sprintf("Billing Account Invoice (%s)", strings.Join(components, "\n"))
}
