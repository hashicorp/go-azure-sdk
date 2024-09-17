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
	recaser.RegisterResourceId(&InvoiceId{})
}

var _ resourceids.ResourceId = &InvoiceId{}

// InvoiceId is a struct representing the Resource ID for a Invoice
type InvoiceId struct {
	InvoiceName string
}

// NewInvoiceID returns a new InvoiceId struct
func NewInvoiceID(invoiceName string) InvoiceId {
	return InvoiceId{
		InvoiceName: invoiceName,
	}
}

// ParseInvoiceID parses 'input' into a InvoiceId
func ParseInvoiceID(input string) (*InvoiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvoiceIDInsensitively parses 'input' case-insensitively into a InvoiceId
// note: this method should only be used for API response data and not user input
func ParseInvoiceIDInsensitively(input string) (*InvoiceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvoiceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.InvoiceName, ok = input.Parsed["invoiceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "invoiceName", input)
	}

	return nil
}

// ValidateInvoiceID checks that 'input' can be parsed as a Invoice ID
func ValidateInvoiceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvoiceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invoice ID
func (id InvoiceId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/default/invoices/%s"
	return fmt.Sprintf(fmtString, id.InvoiceName)
}

// Segments returns a slice of Resource ID Segments which comprise this Invoice ID
func (id InvoiceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.StaticSegment("staticDefault", "default", "default"),
		resourceids.StaticSegment("staticInvoices", "invoices", "invoices"),
		resourceids.UserSpecifiedSegment("invoiceName", "invoiceValue"),
	}
}

// String returns a human-readable description of this Invoice ID
func (id InvoiceId) String() string {
	components := []string{
		fmt.Sprintf("Invoice Name: %q", id.InvoiceName),
	}
	return fmt.Sprintf("Invoice (%s)", strings.Join(components, "\n"))
}
