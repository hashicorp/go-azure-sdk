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
	recaser.RegisterResourceId(&InvoiceSectionBillingRoleDefinitionId{})
}

var _ resourceids.ResourceId = &InvoiceSectionBillingRoleDefinitionId{}

// InvoiceSectionBillingRoleDefinitionId is a struct representing the Resource ID for a Invoice Section Billing Role Definition
type InvoiceSectionBillingRoleDefinitionId struct {
	BillingAccountName        string
	BillingProfileName        string
	InvoiceSectionName        string
	BillingRoleDefinitionName string
}

// NewInvoiceSectionBillingRoleDefinitionID returns a new InvoiceSectionBillingRoleDefinitionId struct
func NewInvoiceSectionBillingRoleDefinitionID(billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleDefinitionName string) InvoiceSectionBillingRoleDefinitionId {
	return InvoiceSectionBillingRoleDefinitionId{
		BillingAccountName:        billingAccountName,
		BillingProfileName:        billingProfileName,
		InvoiceSectionName:        invoiceSectionName,
		BillingRoleDefinitionName: billingRoleDefinitionName,
	}
}

// ParseInvoiceSectionBillingRoleDefinitionID parses 'input' into a InvoiceSectionBillingRoleDefinitionId
func ParseInvoiceSectionBillingRoleDefinitionID(input string) (*InvoiceSectionBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvoiceSectionBillingRoleDefinitionIDInsensitively parses 'input' case-insensitively into a InvoiceSectionBillingRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseInvoiceSectionBillingRoleDefinitionIDInsensitively(input string) (*InvoiceSectionBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvoiceSectionBillingRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.BillingRoleDefinitionName, ok = input.Parsed["billingRoleDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleDefinitionName", input)
	}

	return nil
}

// ValidateInvoiceSectionBillingRoleDefinitionID checks that 'input' can be parsed as a Invoice Section Billing Role Definition ID
func ValidateInvoiceSectionBillingRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvoiceSectionBillingRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invoice Section Billing Role Definition ID
func (id InvoiceSectionBillingRoleDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/invoiceSections/%s/billingRoleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.InvoiceSectionName, id.BillingRoleDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Invoice Section Billing Role Definition ID
func (id InvoiceSectionBillingRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountName"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileName"),
		resourceids.StaticSegment("staticInvoiceSections", "invoiceSections", "invoiceSections"),
		resourceids.UserSpecifiedSegment("invoiceSectionName", "invoiceSectionName"),
		resourceids.StaticSegment("staticBillingRoleDefinitions", "billingRoleDefinitions", "billingRoleDefinitions"),
		resourceids.UserSpecifiedSegment("billingRoleDefinitionName", "roleDefinitionName"),
	}
}

// String returns a human-readable description of this Invoice Section Billing Role Definition ID
func (id InvoiceSectionBillingRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Invoice Section Name: %q", id.InvoiceSectionName),
		fmt.Sprintf("Billing Role Definition Name: %q", id.BillingRoleDefinitionName),
	}
	return fmt.Sprintf("Invoice Section Billing Role Definition (%s)", strings.Join(components, "\n"))
}
