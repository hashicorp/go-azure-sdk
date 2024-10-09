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
	recaser.RegisterResourceId(&PaymentMethodLinkId{})
}

var _ resourceids.ResourceId = &PaymentMethodLinkId{}

// PaymentMethodLinkId is a struct representing the Resource ID for a Payment Method Link
type PaymentMethodLinkId struct {
	BillingAccountName    string
	BillingProfileName    string
	PaymentMethodLinkName string
}

// NewPaymentMethodLinkID returns a new PaymentMethodLinkId struct
func NewPaymentMethodLinkID(billingAccountName string, billingProfileName string, paymentMethodLinkName string) PaymentMethodLinkId {
	return PaymentMethodLinkId{
		BillingAccountName:    billingAccountName,
		BillingProfileName:    billingProfileName,
		PaymentMethodLinkName: paymentMethodLinkName,
	}
}

// ParsePaymentMethodLinkID parses 'input' into a PaymentMethodLinkId
func ParsePaymentMethodLinkID(input string) (*PaymentMethodLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PaymentMethodLinkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PaymentMethodLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParsePaymentMethodLinkIDInsensitively parses 'input' case-insensitively into a PaymentMethodLinkId
// note: this method should only be used for API response data and not user input
func ParsePaymentMethodLinkIDInsensitively(input string) (*PaymentMethodLinkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&PaymentMethodLinkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := PaymentMethodLinkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *PaymentMethodLinkId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.PaymentMethodLinkName, ok = input.Parsed["paymentMethodLinkName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "paymentMethodLinkName", input)
	}

	return nil
}

// ValidatePaymentMethodLinkID checks that 'input' can be parsed as a Payment Method Link ID
func ValidatePaymentMethodLinkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParsePaymentMethodLinkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Payment Method Link ID
func (id PaymentMethodLinkId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/paymentMethodLinks/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.PaymentMethodLinkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Payment Method Link ID
func (id PaymentMethodLinkId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountName"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileName"),
		resourceids.StaticSegment("staticPaymentMethodLinks", "paymentMethodLinks", "paymentMethodLinks"),
		resourceids.UserSpecifiedSegment("paymentMethodLinkName", "paymentMethodLinkName"),
	}
}

// String returns a human-readable description of this Payment Method Link ID
func (id PaymentMethodLinkId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Payment Method Link Name: %q", id.PaymentMethodLinkName),
	}
	return fmt.Sprintf("Payment Method Link (%s)", strings.Join(components, "\n"))
}
