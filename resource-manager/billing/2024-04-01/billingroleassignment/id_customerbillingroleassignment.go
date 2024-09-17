package billingroleassignment

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CustomerBillingRoleAssignmentId{})
}

var _ resourceids.ResourceId = &CustomerBillingRoleAssignmentId{}

// CustomerBillingRoleAssignmentId is a struct representing the Resource ID for a Customer Billing Role Assignment
type CustomerBillingRoleAssignmentId struct {
	BillingAccountName        string
	BillingProfileName        string
	CustomerName              string
	BillingRoleAssignmentName string
}

// NewCustomerBillingRoleAssignmentID returns a new CustomerBillingRoleAssignmentId struct
func NewCustomerBillingRoleAssignmentID(billingAccountName string, billingProfileName string, customerName string, billingRoleAssignmentName string) CustomerBillingRoleAssignmentId {
	return CustomerBillingRoleAssignmentId{
		BillingAccountName:        billingAccountName,
		BillingProfileName:        billingProfileName,
		CustomerName:              customerName,
		BillingRoleAssignmentName: billingRoleAssignmentName,
	}
}

// ParseCustomerBillingRoleAssignmentID parses 'input' into a CustomerBillingRoleAssignmentId
func ParseCustomerBillingRoleAssignmentID(input string) (*CustomerBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerBillingRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCustomerBillingRoleAssignmentIDInsensitively parses 'input' case-insensitively into a CustomerBillingRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseCustomerBillingRoleAssignmentIDInsensitively(input string) (*CustomerBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerBillingRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CustomerBillingRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.CustomerName, ok = input.Parsed["customerName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customerName", input)
	}

	if id.BillingRoleAssignmentName, ok = input.Parsed["billingRoleAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleAssignmentName", input)
	}

	return nil
}

// ValidateCustomerBillingRoleAssignmentID checks that 'input' can be parsed as a Customer Billing Role Assignment ID
func ValidateCustomerBillingRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomerBillingRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Customer Billing Role Assignment ID
func (id CustomerBillingRoleAssignmentId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/customers/%s/billingRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.CustomerName, id.BillingRoleAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Customer Billing Role Assignment ID
func (id CustomerBillingRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticCustomers", "customers", "customers"),
		resourceids.UserSpecifiedSegment("customerName", "customerValue"),
		resourceids.StaticSegment("staticBillingRoleAssignments", "billingRoleAssignments", "billingRoleAssignments"),
		resourceids.UserSpecifiedSegment("billingRoleAssignmentName", "billingRoleAssignmentValue"),
	}
}

// String returns a human-readable description of this Customer Billing Role Assignment ID
func (id CustomerBillingRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Customer Name: %q", id.CustomerName),
		fmt.Sprintf("Billing Role Assignment Name: %q", id.BillingRoleAssignmentName),
	}
	return fmt.Sprintf("Customer Billing Role Assignment (%s)", strings.Join(components, "\n"))
}
