package transactions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&TransactionId{})
}

var _ resourceids.ResourceId = &TransactionId{}

// TransactionId is a struct representing the Resource ID for a Transaction
type TransactionId struct {
	BillingAccountName string
	BillingProfileName string
	TransactionName    string
}

// NewTransactionID returns a new TransactionId struct
func NewTransactionID(billingAccountName string, billingProfileName string, transactionName string) TransactionId {
	return TransactionId{
		BillingAccountName: billingAccountName,
		BillingProfileName: billingProfileName,
		TransactionName:    transactionName,
	}
}

// ParseTransactionID parses 'input' into a TransactionId
func ParseTransactionID(input string) (*TransactionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TransactionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TransactionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseTransactionIDInsensitively parses 'input' case-insensitively into a TransactionId
// note: this method should only be used for API response data and not user input
func ParseTransactionIDInsensitively(input string) (*TransactionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&TransactionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := TransactionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *TransactionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.TransactionName, ok = input.Parsed["transactionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "transactionName", input)
	}

	return nil
}

// ValidateTransactionID checks that 'input' can be parsed as a Transaction ID
func ValidateTransactionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTransactionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Transaction ID
func (id TransactionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/transactions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.TransactionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Transaction ID
func (id TransactionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountName"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileName"),
		resourceids.StaticSegment("staticTransactions", "transactions", "transactions"),
		resourceids.UserSpecifiedSegment("transactionName", "transactionName"),
	}
}

// String returns a human-readable description of this Transaction ID
func (id TransactionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Transaction Name: %q", id.TransactionName),
	}
	return fmt.Sprintf("Transaction (%s)", strings.Join(components, "\n"))
}
