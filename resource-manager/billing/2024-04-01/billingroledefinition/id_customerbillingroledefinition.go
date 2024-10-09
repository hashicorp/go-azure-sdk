package billingroledefinition

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CustomerBillingRoleDefinitionId{})
}

var _ resourceids.ResourceId = &CustomerBillingRoleDefinitionId{}

// CustomerBillingRoleDefinitionId is a struct representing the Resource ID for a Customer Billing Role Definition
type CustomerBillingRoleDefinitionId struct {
	BillingAccountName        string
	BillingProfileName        string
	CustomerName              string
	BillingRoleDefinitionName string
}

// NewCustomerBillingRoleDefinitionID returns a new CustomerBillingRoleDefinitionId struct
func NewCustomerBillingRoleDefinitionID(billingAccountName string, billingProfileName string, customerName string, billingRoleDefinitionName string) CustomerBillingRoleDefinitionId {
	return CustomerBillingRoleDefinitionId{
		BillingAccountName:        billingAccountName,
		BillingProfileName:        billingProfileName,
		CustomerName:              customerName,
		BillingRoleDefinitionName: billingRoleDefinitionName,
	}
}

// ParseCustomerBillingRoleDefinitionID parses 'input' into a CustomerBillingRoleDefinitionId
func ParseCustomerBillingRoleDefinitionID(input string) (*CustomerBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCustomerBillingRoleDefinitionIDInsensitively parses 'input' case-insensitively into a CustomerBillingRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseCustomerBillingRoleDefinitionIDInsensitively(input string) (*CustomerBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CustomerBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CustomerBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CustomerBillingRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.BillingRoleDefinitionName, ok = input.Parsed["billingRoleDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleDefinitionName", input)
	}

	return nil
}

// ValidateCustomerBillingRoleDefinitionID checks that 'input' can be parsed as a Customer Billing Role Definition ID
func ValidateCustomerBillingRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCustomerBillingRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Customer Billing Role Definition ID
func (id CustomerBillingRoleDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/customers/%s/billingRoleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.CustomerName, id.BillingRoleDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Customer Billing Role Definition ID
func (id CustomerBillingRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountName"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileName"),
		resourceids.StaticSegment("staticCustomers", "customers", "customers"),
		resourceids.UserSpecifiedSegment("customerName", "customerName"),
		resourceids.StaticSegment("staticBillingRoleDefinitions", "billingRoleDefinitions", "billingRoleDefinitions"),
		resourceids.UserSpecifiedSegment("billingRoleDefinitionName", "billingRoleDefinitionName"),
	}
}

// String returns a human-readable description of this Customer Billing Role Definition ID
func (id CustomerBillingRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Customer Name: %q", id.CustomerName),
		fmt.Sprintf("Billing Role Definition Name: %q", id.BillingRoleDefinitionName),
	}
	return fmt.Sprintf("Customer Billing Role Definition (%s)", strings.Join(components, "\n"))
}
