package pricesheet

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingProfileInvoiceId{})
}

var _ resourceids.ResourceId = &BillingProfileInvoiceId{}

// BillingProfileInvoiceId is a struct representing the Resource ID for a Billing Profile Invoice
type BillingProfileInvoiceId struct {
	BillingAccountName string
	BillingProfileName string
	InvoiceName        string
}

// NewBillingProfileInvoiceID returns a new BillingProfileInvoiceId struct
func NewBillingProfileInvoiceID(billingAccountName string, billingProfileName string, invoiceName string) BillingProfileInvoiceId {
	return BillingProfileInvoiceId{
		BillingAccountName: billingAccountName,
		BillingProfileName: billingProfileName,
		InvoiceName:        invoiceName,
	}
}

// ParseBillingProfileInvoiceID parses 'input' into a BillingProfileInvoiceId
func ParseBillingProfileInvoiceID(input string) (*BillingProfileInvoiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileInvoiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileInvoiceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingProfileInvoiceIDInsensitively parses 'input' case-insensitively into a BillingProfileInvoiceId
// note: this method should only be used for API response data and not user input
func ParseBillingProfileInvoiceIDInsensitively(input string) (*BillingProfileInvoiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileInvoiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileInvoiceId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingProfileInvoiceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.InvoiceName, ok = input.Parsed["invoiceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "invoiceName", input)
	}

	return nil
}

// ValidateBillingProfileInvoiceID checks that 'input' can be parsed as a Billing Profile Invoice ID
func ValidateBillingProfileInvoiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingProfileInvoiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Profile Invoice ID
func (id BillingProfileInvoiceId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/invoices/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.InvoiceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Profile Invoice ID
func (id BillingProfileInvoiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticInvoices", "invoices", "invoices"),
		resourceids.UserSpecifiedSegment("invoiceName", "invoiceValue"),
	}
}

// String returns a human-readable description of this Billing Profile Invoice ID
func (id BillingProfileInvoiceId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Invoice Name: %q", id.InvoiceName),
	}
	return fmt.Sprintf("Billing Profile Invoice (%s)", strings.Join(components, "\n"))
}
