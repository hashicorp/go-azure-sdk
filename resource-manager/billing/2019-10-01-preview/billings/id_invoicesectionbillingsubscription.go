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
	recaser.RegisterResourceId(&InvoiceSectionBillingSubscriptionId{})
}

var _ resourceids.ResourceId = &InvoiceSectionBillingSubscriptionId{}

// InvoiceSectionBillingSubscriptionId is a struct representing the Resource ID for a Invoice Section Billing Subscription
type InvoiceSectionBillingSubscriptionId struct {
	BillingAccountName      string
	BillingProfileName      string
	InvoiceSectionName      string
	BillingSubscriptionName string
}

// NewInvoiceSectionBillingSubscriptionID returns a new InvoiceSectionBillingSubscriptionId struct
func NewInvoiceSectionBillingSubscriptionID(billingAccountName string, billingProfileName string, invoiceSectionName string, billingSubscriptionName string) InvoiceSectionBillingSubscriptionId {
	return InvoiceSectionBillingSubscriptionId{
		BillingAccountName:      billingAccountName,
		BillingProfileName:      billingProfileName,
		InvoiceSectionName:      invoiceSectionName,
		BillingSubscriptionName: billingSubscriptionName,
	}
}

// ParseInvoiceSectionBillingSubscriptionID parses 'input' into a InvoiceSectionBillingSubscriptionId
func ParseInvoiceSectionBillingSubscriptionID(input string) (*InvoiceSectionBillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionBillingSubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionBillingSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvoiceSectionBillingSubscriptionIDInsensitively parses 'input' case-insensitively into a InvoiceSectionBillingSubscriptionId
// note: this method should only be used for API response data and not user input
func ParseInvoiceSectionBillingSubscriptionIDInsensitively(input string) (*InvoiceSectionBillingSubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionBillingSubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionBillingSubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvoiceSectionBillingSubscriptionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.BillingSubscriptionName, ok = input.Parsed["billingSubscriptionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingSubscriptionName", input)
	}

	return nil
}

// ValidateInvoiceSectionBillingSubscriptionID checks that 'input' can be parsed as a Invoice Section Billing Subscription ID
func ValidateInvoiceSectionBillingSubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvoiceSectionBillingSubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invoice Section Billing Subscription ID
func (id InvoiceSectionBillingSubscriptionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/invoiceSections/%s/billingSubscriptions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.InvoiceSectionName, id.BillingSubscriptionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Invoice Section Billing Subscription ID
func (id InvoiceSectionBillingSubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticInvoiceSections", "invoiceSections", "invoiceSections"),
		resourceids.UserSpecifiedSegment("invoiceSectionName", "invoiceSectionValue"),
		resourceids.StaticSegment("staticBillingSubscriptions", "billingSubscriptions", "billingSubscriptions"),
		resourceids.UserSpecifiedSegment("billingSubscriptionName", "billingSubscriptionValue"),
	}
}

// String returns a human-readable description of this Invoice Section Billing Subscription ID
func (id InvoiceSectionBillingSubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Invoice Section Name: %q", id.InvoiceSectionName),
		fmt.Sprintf("Billing Subscription Name: %q", id.BillingSubscriptionName),
	}
	return fmt.Sprintf("Invoice Section Billing Subscription (%s)", strings.Join(components, "\n"))
}
