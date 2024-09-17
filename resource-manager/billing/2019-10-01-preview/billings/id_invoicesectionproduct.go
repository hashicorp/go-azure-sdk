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
	recaser.RegisterResourceId(&InvoiceSectionProductId{})
}

var _ resourceids.ResourceId = &InvoiceSectionProductId{}

// InvoiceSectionProductId is a struct representing the Resource ID for a Invoice Section Product
type InvoiceSectionProductId struct {
	BillingAccountName string
	BillingProfileName string
	InvoiceSectionName string
	ProductName        string
}

// NewInvoiceSectionProductID returns a new InvoiceSectionProductId struct
func NewInvoiceSectionProductID(billingAccountName string, billingProfileName string, invoiceSectionName string, productName string) InvoiceSectionProductId {
	return InvoiceSectionProductId{
		BillingAccountName: billingAccountName,
		BillingProfileName: billingProfileName,
		InvoiceSectionName: invoiceSectionName,
		ProductName:        productName,
	}
}

// ParseInvoiceSectionProductID parses 'input' into a InvoiceSectionProductId
func ParseInvoiceSectionProductID(input string) (*InvoiceSectionProductId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionProductId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionProductId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvoiceSectionProductIDInsensitively parses 'input' case-insensitively into a InvoiceSectionProductId
// note: this method should only be used for API response data and not user input
func ParseInvoiceSectionProductIDInsensitively(input string) (*InvoiceSectionProductId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionProductId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionProductId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvoiceSectionProductId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.InvoiceSectionName, ok = input.Parsed["invoiceSectionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "invoiceSectionName", input)
	}

	if id.ProductName, ok = input.Parsed["productName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "productName", input)
	}

	return nil
}

// ValidateInvoiceSectionProductID checks that 'input' can be parsed as a Invoice Section Product ID
func ValidateInvoiceSectionProductID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvoiceSectionProductID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invoice Section Product ID
func (id InvoiceSectionProductId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/invoiceSections/%s/products/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.InvoiceSectionName, id.ProductName)
}

// Segments returns a slice of Resource ID Segments which comprise this Invoice Section Product ID
func (id InvoiceSectionProductId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticInvoiceSections", "invoiceSections", "invoiceSections"),
		resourceids.UserSpecifiedSegment("invoiceSectionName", "invoiceSectionValue"),
		resourceids.StaticSegment("staticProducts", "products", "products"),
		resourceids.UserSpecifiedSegment("productName", "productValue"),
	}
}

// String returns a human-readable description of this Invoice Section Product ID
func (id InvoiceSectionProductId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Invoice Section Name: %q", id.InvoiceSectionName),
		fmt.Sprintf("Product Name: %q", id.ProductName),
	}
	return fmt.Sprintf("Invoice Section Product (%s)", strings.Join(components, "\n"))
}
