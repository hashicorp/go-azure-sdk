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
	recaser.RegisterResourceId(&InvoiceSectionBillingRoleAssignmentId{})
}

var _ resourceids.ResourceId = &InvoiceSectionBillingRoleAssignmentId{}

// InvoiceSectionBillingRoleAssignmentId is a struct representing the Resource ID for a Invoice Section Billing Role Assignment
type InvoiceSectionBillingRoleAssignmentId struct {
	BillingAccountName        string
	BillingProfileName        string
	InvoiceSectionName        string
	BillingRoleAssignmentName string
}

// NewInvoiceSectionBillingRoleAssignmentID returns a new InvoiceSectionBillingRoleAssignmentId struct
func NewInvoiceSectionBillingRoleAssignmentID(billingAccountName string, billingProfileName string, invoiceSectionName string, billingRoleAssignmentName string) InvoiceSectionBillingRoleAssignmentId {
	return InvoiceSectionBillingRoleAssignmentId{
		BillingAccountName:        billingAccountName,
		BillingProfileName:        billingProfileName,
		InvoiceSectionName:        invoiceSectionName,
		BillingRoleAssignmentName: billingRoleAssignmentName,
	}
}

// ParseInvoiceSectionBillingRoleAssignmentID parses 'input' into a InvoiceSectionBillingRoleAssignmentId
func ParseInvoiceSectionBillingRoleAssignmentID(input string) (*InvoiceSectionBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionBillingRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseInvoiceSectionBillingRoleAssignmentIDInsensitively parses 'input' case-insensitively into a InvoiceSectionBillingRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseInvoiceSectionBillingRoleAssignmentIDInsensitively(input string) (*InvoiceSectionBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&InvoiceSectionBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := InvoiceSectionBillingRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *InvoiceSectionBillingRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.BillingRoleAssignmentName, ok = input.Parsed["billingRoleAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleAssignmentName", input)
	}

	return nil
}

// ValidateInvoiceSectionBillingRoleAssignmentID checks that 'input' can be parsed as a Invoice Section Billing Role Assignment ID
func ValidateInvoiceSectionBillingRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseInvoiceSectionBillingRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Invoice Section Billing Role Assignment ID
func (id InvoiceSectionBillingRoleAssignmentId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/invoiceSections/%s/billingRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.InvoiceSectionName, id.BillingRoleAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Invoice Section Billing Role Assignment ID
func (id InvoiceSectionBillingRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountName"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileName"),
		resourceids.StaticSegment("staticInvoiceSections", "invoiceSections", "invoiceSections"),
		resourceids.UserSpecifiedSegment("invoiceSectionName", "invoiceSectionName"),
		resourceids.StaticSegment("staticBillingRoleAssignments", "billingRoleAssignments", "billingRoleAssignments"),
		resourceids.UserSpecifiedSegment("billingRoleAssignmentName", "billingRoleAssignmentName"),
	}
}

// String returns a human-readable description of this Invoice Section Billing Role Assignment ID
func (id InvoiceSectionBillingRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Invoice Section Name: %q", id.InvoiceSectionName),
		fmt.Sprintf("Billing Role Assignment Name: %q", id.BillingRoleAssignmentName),
	}
	return fmt.Sprintf("Invoice Section Billing Role Assignment (%s)", strings.Join(components, "\n"))
}
