package transfers

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&InvoiceSectionTransferId{})
}

var _ resourceids.ResourceId = &InvoiceSectionTransferId{}

// InvoiceSectionTransferId is a struct representing the Resource ID for a Invoice Section Transfer
type InvoiceSectionTransferId struct {
	BillingAccountName string
	BillingProfileName string
	InvoiceSectionName string
	TransferName       string
}

// NewInvoiceSectionTransferID returns a new InvoiceSectionTransferId struct
func NewInvoiceSectionTransferID(billingAccountName string, billingProfileName string, invoiceSectionName string, transferName string) InvoiceSectionTransferId {
	return InvoiceSectionTransferId{
		BillingAccountName: billingAccountName,
		BillingProfileName: billingProfileName,
		InvoiceSectionName: invoiceSectionName,
		TransferName:       transferName,
	}
}

// ParseInvoiceSectionTransferID parses 'input' into a InvoiceSectionTransferId
func ParseInvoiceSectionTransferID(input string) (*InvoiceSectionTransferId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionTransferId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionTransferId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvoiceSectionTransferIDInsensitively parses 'input' case-insensitively into a InvoiceSectionTransferId
// note: this method should only be used for API response data and not user input
func ParseInvoiceSectionTransferIDInsensitively(input string) (*InvoiceSectionTransferId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionTransferId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionTransferId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvoiceSectionTransferId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.TransferName, ok = input.Parsed["transferName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "transferName", input)
	}

	return nil
}

// ValidateInvoiceSectionTransferID checks that 'input' can be parsed as a Invoice Section Transfer ID
func ValidateInvoiceSectionTransferID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvoiceSectionTransferID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invoice Section Transfer ID
func (id InvoiceSectionTransferId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/invoiceSections/%s/transfers/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.InvoiceSectionName, id.TransferName)
}

// Segments returns a slice of Resource ID Segments which comprise this Invoice Section Transfer ID
func (id InvoiceSectionTransferId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticInvoiceSections", "invoiceSections", "invoiceSections"),
		resourceids.UserSpecifiedSegment("invoiceSectionName", "invoiceSectionValue"),
		resourceids.StaticSegment("staticTransfers", "transfers", "transfers"),
		resourceids.UserSpecifiedSegment("transferName", "transferValue"),
	}
}

// String returns a human-readable description of this Invoice Section Transfer ID
func (id InvoiceSectionTransferId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Invoice Section Name: %q", id.InvoiceSectionName),
		fmt.Sprintf("Transfer Name: %q", id.TransferName),
	}
	return fmt.Sprintf("Invoice Section Transfer (%s)", strings.Join(components, "\n"))
}
