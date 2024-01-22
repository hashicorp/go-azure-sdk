package paymentmethods

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &BillingAccountPaymentMethodId{}

// BillingAccountPaymentMethodId is a struct representing the Resource ID for a Billing Account Payment Method
type BillingAccountPaymentMethodId struct {
	BillingAccountName string
	PaymentMethodName  string
}

// NewBillingAccountPaymentMethodID returns a new BillingAccountPaymentMethodId struct
func NewBillingAccountPaymentMethodID(billingAccountName string, paymentMethodName string) BillingAccountPaymentMethodId {
	return BillingAccountPaymentMethodId{
		BillingAccountName: billingAccountName,
		PaymentMethodName:  paymentMethodName,
	}
}

// ParseBillingAccountPaymentMethodID parses 'input' into a BillingAccountPaymentMethodId
func ParseBillingAccountPaymentMethodID(input string) (*BillingAccountPaymentMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingAccountPaymentMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingAccountPaymentMethodId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingAccountPaymentMethodIDInsensitively parses 'input' case-insensitively into a BillingAccountPaymentMethodId
// note: this method should only be used for API response data and not user input
func ParseBillingAccountPaymentMethodIDInsensitively(input string) (*BillingAccountPaymentMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingAccountPaymentMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingAccountPaymentMethodId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingAccountPaymentMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.PaymentMethodName, ok = input.Parsed["paymentMethodName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "paymentMethodName", input)
	}

	return nil
}

// ValidateBillingAccountPaymentMethodID checks that 'input' can be parsed as a Billing Account Payment Method ID
func ValidateBillingAccountPaymentMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingAccountPaymentMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Account Payment Method ID
func (id BillingAccountPaymentMethodId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/paymentMethods/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.PaymentMethodName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Account Payment Method ID
func (id BillingAccountPaymentMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticPaymentMethods", "paymentMethods", "paymentMethods"),
		resourceids.UserSpecifiedSegment("paymentMethodName", "paymentMethodValue"),
	}
}

// String returns a human-readable description of this Billing Account Payment Method ID
func (id BillingAccountPaymentMethodId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Payment Method Name: %q", id.PaymentMethodName),
	}
	return fmt.Sprintf("Billing Account Payment Method (%s)", strings.Join(components, "\n"))
}
