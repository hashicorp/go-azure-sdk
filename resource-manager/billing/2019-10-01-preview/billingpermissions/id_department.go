package billingpermissions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&DepartmentId{})
}

var _ resourceids.ResourceId = &DepartmentId{}

// DepartmentId is a struct representing the Resource ID for a Department
type DepartmentId struct {
	BillingAccountName string
	DepartmentName     string
}

// NewDepartmentID returns a new DepartmentId struct
func NewDepartmentID(billingAccountName string, departmentName string) DepartmentId {
	return DepartmentId{
		BillingAccountName: billingAccountName,
		DepartmentName:     departmentName,
	}
}

// ParseDepartmentID parses 'input' into a DepartmentId
func ParseDepartmentID(input string) (*DepartmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DepartmentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DepartmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDepartmentIDInsensitively parses 'input' case-insensitively into a DepartmentId
// note: this method should only be used for API response data and not user input
func ParseDepartmentIDInsensitively(input string) (*DepartmentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DepartmentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DepartmentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DepartmentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.DepartmentName, ok = input.Parsed["departmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "departmentName", input)
	}

	return nil
}

// ValidateDepartmentID checks that 'input' can be parsed as a Department ID
func ValidateDepartmentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDepartmentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Department ID
func (id DepartmentId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/departments/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.DepartmentName)
}

// Segments returns a slice of Resource ID Segments which comprise this Department ID
func (id DepartmentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticDepartments", "departments", "departments"),
		resourceids.UserSpecifiedSegment("departmentName", "departmentValue"),
	}
}

// String returns a human-readable description of this Department ID
func (id DepartmentId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Department Name: %q", id.DepartmentName),
	}
	return fmt.Sprintf("Department (%s)", strings.Join(components, "\n"))
}
