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
	recaser.RegisterResourceId(&EnrollmentAccountBillingRoleAssignmentId{})
}

var _ resourceids.ResourceId = &EnrollmentAccountBillingRoleAssignmentId{}

// EnrollmentAccountBillingRoleAssignmentId is a struct representing the Resource ID for a Enrollment Account Billing Role Assignment
type EnrollmentAccountBillingRoleAssignmentId struct {
	BillingAccountName        string
	EnrollmentAccountName     string
	BillingRoleAssignmentName string
}

// NewEnrollmentAccountBillingRoleAssignmentID returns a new EnrollmentAccountBillingRoleAssignmentId struct
func NewEnrollmentAccountBillingRoleAssignmentID(billingAccountName string, enrollmentAccountName string, billingRoleAssignmentName string) EnrollmentAccountBillingRoleAssignmentId {
	return EnrollmentAccountBillingRoleAssignmentId{
		BillingAccountName:        billingAccountName,
		EnrollmentAccountName:     enrollmentAccountName,
		BillingRoleAssignmentName: billingRoleAssignmentName,
	}
}

// ParseEnrollmentAccountBillingRoleAssignmentID parses 'input' into a EnrollmentAccountBillingRoleAssignmentId
func ParseEnrollmentAccountBillingRoleAssignmentID(input string) (*EnrollmentAccountBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnrollmentAccountBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnrollmentAccountBillingRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEnrollmentAccountBillingRoleAssignmentIDInsensitively parses 'input' case-insensitively into a EnrollmentAccountBillingRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseEnrollmentAccountBillingRoleAssignmentIDInsensitively(input string) (*EnrollmentAccountBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnrollmentAccountBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnrollmentAccountBillingRoleAssignmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EnrollmentAccountBillingRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.EnrollmentAccountName, ok = input.Parsed["enrollmentAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentAccountName", input)
	}

	if id.BillingRoleAssignmentName, ok = input.Parsed["billingRoleAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleAssignmentName", input)
	}

	return nil
}

// ValidateEnrollmentAccountBillingRoleAssignmentID checks that 'input' can be parsed as a Enrollment Account Billing Role Assignment ID
func ValidateEnrollmentAccountBillingRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnrollmentAccountBillingRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Enrollment Account Billing Role Assignment ID
func (id EnrollmentAccountBillingRoleAssignmentId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/enrollmentAccounts/%s/billingRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.EnrollmentAccountName, id.BillingRoleAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Enrollment Account Billing Role Assignment ID
func (id EnrollmentAccountBillingRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticEnrollmentAccounts", "enrollmentAccounts", "enrollmentAccounts"),
		resourceids.UserSpecifiedSegment("enrollmentAccountName", "enrollmentAccountValue"),
		resourceids.StaticSegment("staticBillingRoleAssignments", "billingRoleAssignments", "billingRoleAssignments"),
		resourceids.UserSpecifiedSegment("billingRoleAssignmentName", "billingRoleAssignmentValue"),
	}
}

// String returns a human-readable description of this Enrollment Account Billing Role Assignment ID
func (id EnrollmentAccountBillingRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Enrollment Account Name: %q", id.EnrollmentAccountName),
		fmt.Sprintf("Billing Role Assignment Name: %q", id.BillingRoleAssignmentName),
	}
	return fmt.Sprintf("Enrollment Account Billing Role Assignment (%s)", strings.Join(components, "\n"))
}
