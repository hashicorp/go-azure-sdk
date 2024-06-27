package enrollmentaccount

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DepartmentEnrollmentAccountId{})
}

var _ resourceids.ResourceId = &DepartmentEnrollmentAccountId{}

// DepartmentEnrollmentAccountId is a struct representing the Resource ID for a Department Enrollment Account
type DepartmentEnrollmentAccountId struct {
	BillingAccountName    string
	DepartmentName        string
	EnrollmentAccountName string
}

// NewDepartmentEnrollmentAccountID returns a new DepartmentEnrollmentAccountId struct
func NewDepartmentEnrollmentAccountID(billingAccountName string, departmentName string, enrollmentAccountName string) DepartmentEnrollmentAccountId {
	return DepartmentEnrollmentAccountId{
		BillingAccountName:    billingAccountName,
		DepartmentName:        departmentName,
		EnrollmentAccountName: enrollmentAccountName,
	}
}

// ParseDepartmentEnrollmentAccountID parses 'input' into a DepartmentEnrollmentAccountId
func ParseDepartmentEnrollmentAccountID(input string) (*DepartmentEnrollmentAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DepartmentEnrollmentAccountId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DepartmentEnrollmentAccountId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDepartmentEnrollmentAccountIDInsensitively parses 'input' case-insensitively into a DepartmentEnrollmentAccountId
// note: this method should only be used for API response data and not user input
func ParseDepartmentEnrollmentAccountIDInsensitively(input string) (*DepartmentEnrollmentAccountId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DepartmentEnrollmentAccountId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DepartmentEnrollmentAccountId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DepartmentEnrollmentAccountId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.DepartmentName, ok = input.Parsed["departmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "departmentName", input)
	}

	if id.EnrollmentAccountName, ok = input.Parsed["enrollmentAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentAccountName", input)
	}

	return nil
}

// ValidateDepartmentEnrollmentAccountID checks that 'input' can be parsed as a Department Enrollment Account ID
func ValidateDepartmentEnrollmentAccountID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDepartmentEnrollmentAccountID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Department Enrollment Account ID
func (id DepartmentEnrollmentAccountId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/departments/%s/enrollmentAccounts/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.DepartmentName, id.EnrollmentAccountName)
}

// Segments returns a slice of Resource ID Segments which comprise this Department Enrollment Account ID
func (id DepartmentEnrollmentAccountId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticDepartments", "departments", "departments"),
		resourceids.UserSpecifiedSegment("departmentName", "departmentValue"),
		resourceids.StaticSegment("staticEnrollmentAccounts", "enrollmentAccounts", "enrollmentAccounts"),
		resourceids.UserSpecifiedSegment("enrollmentAccountName", "enrollmentAccountValue"),
	}
}

// String returns a human-readable description of this Department Enrollment Account ID
func (id DepartmentEnrollmentAccountId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Department Name: %q", id.DepartmentName),
		fmt.Sprintf("Enrollment Account Name: %q", id.EnrollmentAccountName),
	}
	return fmt.Sprintf("Department Enrollment Account (%s)", strings.Join(components, "\n"))
}
