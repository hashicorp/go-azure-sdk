package lots

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &CustomerId{}

// CustomerId is a struct representing the Resource ID for a Customer
type CustomerId struct {
	BillingAccountId string
	CustomerId       string
}

// NewCustomerID returns a new CustomerId struct
func NewCustomerID(billingAccountId string, customerId string) CustomerId {
	return CustomerId{
		BillingAccountId: billingAccountId,
		CustomerId:       customerId,
	}
}

// ParseCustomerID parses 'input' into a CustomerId
func ParseCustomerID(input string) (*CustomerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCustomerIDInsensitively parses 'input' case-insensitively into a CustomerId
// note: this method should only be used for API response data and not user input
func ParseCustomerIDInsensitively(input string) (*CustomerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CustomerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountId, ok = input.Parsed["billingAccountId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountId", input)
	}

	if id.CustomerId, ok = input.Parsed["customerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customerId", input)
	}

	return nil
}

// ValidateCustomerID checks that 'input' can be parsed as a Customer ID
func ValidateCustomerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Customer ID
func (id CustomerId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/customers/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountId, id.CustomerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Customer ID
func (id CustomerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountId", "billingAccountIdValue"),
		resourceids.StaticSegment("staticCustomers", "customers", "customers"),
		resourceids.UserSpecifiedSegment("customerId", "customerIdValue"),
	}
}

// String returns a human-readable description of this Customer ID
func (id CustomerId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account: %q", id.BillingAccountId),
		fmt.Sprintf("Customer: %q", id.CustomerId),
	}
	return fmt.Sprintf("Customer (%s)", strings.Join(components, "\n"))
}
