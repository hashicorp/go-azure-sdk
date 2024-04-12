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
	recaser.RegisterResourceId(&DepartmentBillingRoleAssignmentId{})
}

var _ resourceids.ResourceId = &DepartmentBillingRoleAssignmentId{}

// DepartmentBillingRoleAssignmentId is a struct representing the Resource ID for a Department Billing Role Assignment
type DepartmentBillingRoleAssignmentId struct {
	BillingAccountName        string
	DepartmentName            string
	BillingRoleAssignmentName string
}

// NewDepartmentBillingRoleAssignmentID returns a new DepartmentBillingRoleAssignmentId struct
func NewDepartmentBillingRoleAssignmentID(billingAccountName string, departmentName string, billingRoleAssignmentName string) DepartmentBillingRoleAssignmentId {
	return DepartmentBillingRoleAssignmentId{
		BillingAccountName:        billingAccountName,
		DepartmentName:            departmentName,
		BillingRoleAssignmentName: billingRoleAssignmentName,
	}
}

// ParseDepartmentBillingRoleAssignmentID parses 'input' into a DepartmentBillingRoleAssignmentId
func ParseDepartmentBillingRoleAssignmentID(input string) (*DepartmentBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DepartmentBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DepartmentBillingRoleAssignmentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDepartmentBillingRoleAssignmentIDInsensitively parses 'input' case-insensitively into a DepartmentBillingRoleAssignmentId
// note: this method should only be used for API response data and not user input
func ParseDepartmentBillingRoleAssignmentIDInsensitively(input string) (*DepartmentBillingRoleAssignmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DepartmentBillingRoleAssignmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DepartmentBillingRoleAssignmentId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DepartmentBillingRoleAssignmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.DepartmentName, ok = input.Parsed["departmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "departmentName", input)
	}

	if id.BillingRoleAssignmentName, ok = input.Parsed["billingRoleAssignmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleAssignmentName", input)
	}

	return nil
}

// ValidateDepartmentBillingRoleAssignmentID checks that 'input' can be parsed as a Department Billing Role Assignment ID
func ValidateDepartmentBillingRoleAssignmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDepartmentBillingRoleAssignmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Department Billing Role Assignment ID
func (id DepartmentBillingRoleAssignmentId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/departments/%s/billingRoleAssignments/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.DepartmentName, id.BillingRoleAssignmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Department Billing Role Assignment ID
func (id DepartmentBillingRoleAssignmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticDepartments", "departments", "departments"),
		resourceids.UserSpecifiedSegment("departmentName", "departmentValue"),
		resourceids.StaticSegment("staticBillingRoleAssignments", "billingRoleAssignments", "billingRoleAssignments"),
		resourceids.UserSpecifiedSegment("billingRoleAssignmentName", "billingRoleAssignmentValue"),
	}
}

// String returns a human-readable description of this Department Billing Role Assignment ID
func (id DepartmentBillingRoleAssignmentId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Department Name: %q", id.DepartmentName),
		fmt.Sprintf("Billing Role Assignment Name: %q", id.BillingRoleAssignmentName),
	}
	return fmt.Sprintf("Department Billing Role Assignment (%s)", strings.Join(components, "\n"))
}
