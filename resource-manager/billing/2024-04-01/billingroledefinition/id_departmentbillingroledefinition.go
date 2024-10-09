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
	recaser.RegisterResourceId(&DepartmentBillingRoleDefinitionId{})
}

var _ resourceids.ResourceId = &DepartmentBillingRoleDefinitionId{}

// DepartmentBillingRoleDefinitionId is a struct representing the Resource ID for a Department Billing Role Definition
type DepartmentBillingRoleDefinitionId struct {
	BillingAccountName        string
	DepartmentName            string
	BillingRoleDefinitionName string
}

// NewDepartmentBillingRoleDefinitionID returns a new DepartmentBillingRoleDefinitionId struct
func NewDepartmentBillingRoleDefinitionID(billingAccountName string, departmentName string, billingRoleDefinitionName string) DepartmentBillingRoleDefinitionId {
	return DepartmentBillingRoleDefinitionId{
		BillingAccountName:        billingAccountName,
		DepartmentName:            departmentName,
		BillingRoleDefinitionName: billingRoleDefinitionName,
	}
}

// ParseDepartmentBillingRoleDefinitionID parses 'input' into a DepartmentBillingRoleDefinitionId
func ParseDepartmentBillingRoleDefinitionID(input string) (*DepartmentBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DepartmentBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DepartmentBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDepartmentBillingRoleDefinitionIDInsensitively parses 'input' case-insensitively into a DepartmentBillingRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseDepartmentBillingRoleDefinitionIDInsensitively(input string) (*DepartmentBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DepartmentBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DepartmentBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DepartmentBillingRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.DepartmentName, ok = input.Parsed["departmentName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "departmentName", input)
	}

	if id.BillingRoleDefinitionName, ok = input.Parsed["billingRoleDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleDefinitionName", input)
	}

	return nil
}

// ValidateDepartmentBillingRoleDefinitionID checks that 'input' can be parsed as a Department Billing Role Definition ID
func ValidateDepartmentBillingRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDepartmentBillingRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Department Billing Role Definition ID
func (id DepartmentBillingRoleDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/departments/%s/billingRoleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.DepartmentName, id.BillingRoleDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Department Billing Role Definition ID
func (id DepartmentBillingRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountName"),
		resourceids.StaticSegment("staticDepartments", "departments", "departments"),
		resourceids.UserSpecifiedSegment("departmentName", "departmentName"),
		resourceids.StaticSegment("staticBillingRoleDefinitions", "billingRoleDefinitions", "billingRoleDefinitions"),
		resourceids.UserSpecifiedSegment("billingRoleDefinitionName", "billingRoleDefinitionName"),
	}
}

// String returns a human-readable description of this Department Billing Role Definition ID
func (id DepartmentBillingRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Department Name: %q", id.DepartmentName),
		fmt.Sprintf("Billing Role Definition Name: %q", id.BillingRoleDefinitionName),
	}
	return fmt.Sprintf("Department Billing Role Definition (%s)", strings.Join(components, "\n"))
}
