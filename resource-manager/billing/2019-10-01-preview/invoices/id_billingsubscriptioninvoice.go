package invoices

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingSubscriptionInvoiceId{})
}

var _ resourceids.ResourceId = &BillingSubscriptionInvoiceId{}

// BillingSubscriptionInvoiceId is a struct representing the Resource ID for a Billing Subscription Invoice
type BillingSubscriptionInvoiceId struct {
	BillingAccountName      string
	BillingSubscriptionName string
	InvoiceName             string
}

// NewBillingSubscriptionInvoiceID returns a new BillingSubscriptionInvoiceId struct
func NewBillingSubscriptionInvoiceID(billingAccountName string, billingSubscriptionName string, invoiceName string) BillingSubscriptionInvoiceId {
	return BillingSubscriptionInvoiceId{
		BillingAccountName:      billingAccountName,
		BillingSubscriptionName: billingSubscriptionName,
		InvoiceName:             invoiceName,
	}
}

// ParseBillingSubscriptionInvoiceID parses 'input' into a BillingSubscriptionInvoiceId
func ParseBillingSubscriptionInvoiceID(input string) (*BillingSubscriptionInvoiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingSubscriptionInvoiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingSubscriptionInvoiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingSubscriptionInvoiceIDInsensitively parses 'input' case-insensitively into a BillingSubscriptionInvoiceId
// note: this method should only be used for API response data and not user input
func ParseBillingSubscriptionInvoiceIDInsensitively(input string) (*BillingSubscriptionInvoiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingSubscriptionInvoiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingSubscriptionInvoiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingSubscriptionInvoiceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingSubscriptionName, ok = input.Parsed["billingSubscriptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingSubscriptionName", input)
	}

	if id.InvoiceName, ok = input.Parsed["invoiceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "invoiceName", input)
	}

	return nil
}

// ValidateBillingSubscriptionInvoiceID checks that 'input' can be parsed as a Billing Subscription Invoice ID
func ValidateBillingSubscriptionInvoiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingSubscriptionInvoiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Subscription Invoice ID
func (id BillingSubscriptionInvoiceId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingSubscriptions/%s/invoices/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingSubscriptionName, id.InvoiceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Subscription Invoice ID
func (id BillingSubscriptionInvoiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountName"),
		resourceids.StaticSegment("staticBillingSubscriptions", "billingSubscriptions", "billingSubscriptions"),
		resourceids.UserSpecifiedSegment("billingSubscriptionName", "billingSubscriptionName"),
		resourceids.StaticSegment("staticInvoices", "invoices", "invoices"),
		resourceids.UserSpecifiedSegment("invoiceName", "invoiceName"),
	}
}

// String returns a human-readable description of this Billing Subscription Invoice ID
func (id BillingSubscriptionInvoiceId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Subscription Name: %q", id.BillingSubscriptionName),
		fmt.Sprintf("Invoice Name: %q", id.InvoiceName),
	}
	return fmt.Sprintf("Billing Subscription Invoice (%s)", strings.Join(components, "\n"))
}
