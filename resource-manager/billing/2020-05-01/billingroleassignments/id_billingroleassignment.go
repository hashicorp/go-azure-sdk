package billingroleassignments

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingRoleAssignmentId{})
}

var _ resourceids.ResourceId = &BillingRoleAssignmentId{}

// BillingRoleAssignmentId is a struct representing the Resource ID for a Billing Role Assignment
type BillingRoleAssignmentId struct {
	BillingAccountName        string
	BillingRoleAssignmentName string
}

// NewBillingRoleAssignmentID returns a new BillingRoleAssignmentId struct
func NewBillingRoleAssignmentID(billingAccountName string, billingRoleAssignmentName string) BillingRoleAssignmentId {
	return BillingRoleAssignmentId{
		BillingAccountName:        billingAccountName,
		BillingRoleAssignmentName: billingRoleAssignmentName,
	}
}

// ParseBillingRoleAssignmentID parses 'input' into a BillingRoleAssignmentId
func ParseBillingRoleAssignmentID(input string) (*BillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingRoleAssignmentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingRoleAssignmentIDInsensitively parses 'input' case-insensitively into a BillingRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseBillingRoleAssignmentIDInsensitively(input string) (*BillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingRoleAssignmentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingRoleAssignmentName, ok = input.Parsed["billingRoleAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleAssignmentName", input)
	}

	return nil
}

// ValidateBillingRoleAssignmentID checks that 'input' can be parsed as a Billing Role Assignment ID
func ValidateBillingRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Role Assignment ID
func (id BillingRoleAssignmentId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingRoleAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Role Assignment ID
func (id BillingRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingRoleAssignments", "billingRoleAssignments", "billingRoleAssignments"),
		resourceids.UserSpecifiedSegment("billingRoleAssignmentName", "billingRoleAssignmentValue"),
	}
}

// String returns a human-readable description of this Billing Role Assignment ID
func (id BillingRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Role Assignment Name: %q", id.BillingRoleAssignmentName),
	}
	return fmt.Sprintf("Billing Role Assignment (%s)", strings.Join(components, "\n"))
}
