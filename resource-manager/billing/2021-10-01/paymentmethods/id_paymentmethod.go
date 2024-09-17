package paymentmethods

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&PaymentMethodId{})
}

var _ resourceids.ResourceId = &PaymentMethodId{}

// PaymentMethodId is a struct representing the Resource ID for a Payment Method
type PaymentMethodId struct {
	PaymentMethodName string
}

// NewPaymentMethodID returns a new PaymentMethodId struct
func NewPaymentMethodID(paymentMethodName string) PaymentMethodId {
	return PaymentMethodId{
		PaymentMethodName: paymentMethodName,
	}
}

// ParsePaymentMethodID parses 'input' into a PaymentMethodId
func ParsePaymentMethodID(input string) (*PaymentMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PaymentMethodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PaymentMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePaymentMethodIDInsensitively parses 'input' case-insensitively into a PaymentMethodId
// note: this method should only be used for API response data and not user input
func ParsePaymentMethodIDInsensitively(input string) (*PaymentMethodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PaymentMethodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PaymentMethodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PaymentMethodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.PaymentMethodName, ok = input.Parsed["paymentMethodName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "paymentMethodName", input)
	}

	return nil
}

// ValidatePaymentMethodID checks that 'input' can be parsed as a Payment Method ID
func ValidatePaymentMethodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePaymentMethodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Payment Method ID
func (id PaymentMethodId) ID() string {
	fmtString := "/providers/Microsoft.Billing/paymentMethods/%s"
	return fmt.Sprintf(fmtString, id.PaymentMethodName)
}

// Segments returns a slice of Resource ID Segments which comprise this Payment Method ID
func (id PaymentMethodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticPaymentMethods", "paymentMethods", "paymentMethods"),
		resourceids.UserSpecifiedSegment("paymentMethodName", "paymentMethodValue"),
	}
}

// String returns a human-readable description of this Payment Method ID
func (id PaymentMethodId) String() string {
	components := []string{
		fmt.Sprintf("Payment Method Name: %q", id.PaymentMethodName),
	}
	return fmt.Sprintf("Payment Method (%s)", strings.Join(components, "\n"))
}
