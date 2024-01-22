package products

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ProductId{}

// ProductId is a struct representing the Resource ID for a Product
type ProductId struct {
	BillingAccountName string
	ProductName        string
}

// NewProductID returns a new ProductId struct
func NewProductID(billingAccountName string, productName string) ProductId {
	return ProductId{
		BillingAccountName: billingAccountName,
		ProductName:        productName,
	}
}

// ParseProductID parses 'input' into a ProductId
func ParseProductID(input string) (*ProductId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProductId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProductId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProductIDInsensitively parses 'input' case-insensitively into a ProductId
// note: this method should only be used for API response data and not user input
func ParseProductIDInsensitively(input string) (*ProductId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ProductId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ProductId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ProductId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.ProductName, ok = input.Parsed["productName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "productName", input)
	}

	return nil
}

// ValidateProductID checks that 'input' can be parsed as a Product ID
func ValidateProductID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProductID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Product ID
func (id ProductId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/products/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.ProductName)
}

// Segments returns a slice of Resource ID Segments which comprise this Product ID
func (id ProductId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticProducts", "products", "products"),
		resourceids.UserSpecifiedSegment("productName", "productValue"),
	}
}

// String returns a human-readable description of this Product ID
func (id ProductId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Product Name: %q", id.ProductName),
	}
	return fmt.Sprintf("Product (%s)", strings.Join(components, "\n"))
}
