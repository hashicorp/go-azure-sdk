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
	recaser.RegisterResourceId(&BillingProfileBillingRoleAssignmentId{})
}

var _ resourceids.ResourceId = &BillingProfileBillingRoleAssignmentId{}

// BillingProfileBillingRoleAssignmentId is a struct representing the Resource ID for a Billing Profile Billing Role Assignment
type BillingProfileBillingRoleAssignmentId struct {
	BillingAccountName        string
	BillingProfileName        string
	BillingRoleAssignmentName string
}

// NewBillingProfileBillingRoleAssignmentID returns a new BillingProfileBillingRoleAssignmentId struct
func NewBillingProfileBillingRoleAssignmentID(billingAccountName string, billingProfileName string, billingRoleAssignmentName string) BillingProfileBillingRoleAssignmentId {
	return BillingProfileBillingRoleAssignmentId{
		BillingAccountName:        billingAccountName,
		BillingProfileName:        billingProfileName,
		BillingRoleAssignmentName: billingRoleAssignmentName,
	}
}

// ParseBillingProfileBillingRoleAssignmentID parses 'input' into a BillingProfileBillingRoleAssignmentId
func ParseBillingProfileBillingRoleAssignmentID(input string) (*BillingProfileBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileBillingRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingProfileBillingRoleAssignmentIDInsensitively parses 'input' case-insensitively into a BillingProfileBillingRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseBillingProfileBillingRoleAssignmentIDInsensitively(input string) (*BillingProfileBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileBillingRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingProfileBillingRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.BillingRoleAssignmentName, ok = input.Parsed["billingRoleAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleAssignmentName", input)
	}

	return nil
}

// ValidateBillingProfileBillingRoleAssignmentID checks that 'input' can be parsed as a Billing Profile Billing Role Assignment ID
func ValidateBillingProfileBillingRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingProfileBillingRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Profile Billing Role Assignment ID
func (id BillingProfileBillingRoleAssignmentId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/billingRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.BillingRoleAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Profile Billing Role Assignment ID
func (id BillingProfileBillingRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileValue"),
		resourceids.StaticSegment("staticBillingRoleAssignments", "billingRoleAssignments", "billingRoleAssignments"),
		resourceids.UserSpecifiedSegment("billingRoleAssignmentName", "billingRoleAssignmentValue"),
	}
}

// String returns a human-readable description of this Billing Profile Billing Role Assignment ID
func (id BillingProfileBillingRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Billing Role Assignment Name: %q", id.BillingRoleAssignmentName),
	}
	return fmt.Sprintf("Billing Profile Billing Role Assignment (%s)", strings.Join(components, "\n"))
}
