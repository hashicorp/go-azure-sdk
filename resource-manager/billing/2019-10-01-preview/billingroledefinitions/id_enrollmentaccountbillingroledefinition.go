package billingroledefinitions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&EnrollmentAccountBillingRoleDefinitionId{})
}

var _ resourceids.ResourceId = &EnrollmentAccountBillingRoleDefinitionId{}

// EnrollmentAccountBillingRoleDefinitionId is a struct representing the Resource ID for a Enrollment Account Billing Role Definition
type EnrollmentAccountBillingRoleDefinitionId struct {
	BillingAccountName        string
	EnrollmentAccountName     string
	BillingRoleDefinitionName string
}

// NewEnrollmentAccountBillingRoleDefinitionID returns a new EnrollmentAccountBillingRoleDefinitionId struct
func NewEnrollmentAccountBillingRoleDefinitionID(billingAccountName string, enrollmentAccountName string, billingRoleDefinitionName string) EnrollmentAccountBillingRoleDefinitionId {
	return EnrollmentAccountBillingRoleDefinitionId{
		BillingAccountName:        billingAccountName,
		EnrollmentAccountName:     enrollmentAccountName,
		BillingRoleDefinitionName: billingRoleDefinitionName,
	}
}

// ParseEnrollmentAccountBillingRoleDefinitionID parses 'input' into a EnrollmentAccountBillingRoleDefinitionId
func ParseEnrollmentAccountBillingRoleDefinitionID(input string) (*EnrollmentAccountBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnrollmentAccountBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnrollmentAccountBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseEnrollmentAccountBillingRoleDefinitionIDInsensitively parses 'input' case-insensitively into a EnrollmentAccountBillingRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseEnrollmentAccountBillingRoleDefinitionIDInsensitively(input string) (*EnrollmentAccountBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&EnrollmentAccountBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := EnrollmentAccountBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *EnrollmentAccountBillingRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.EnrollmentAccountName, ok = input.Parsed["enrollmentAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "enrollmentAccountName", input)
	}

	if id.BillingRoleDefinitionName, ok = input.Parsed["billingRoleDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleDefinitionName", input)
	}

	return nil
}

// ValidateEnrollmentAccountBillingRoleDefinitionID checks that 'input' can be parsed as a Enrollment Account Billing Role Definition ID
func ValidateEnrollmentAccountBillingRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseEnrollmentAccountBillingRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Enrollment Account Billing Role Definition ID
func (id EnrollmentAccountBillingRoleDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/enrollmentAccounts/%s/billingRoleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.EnrollmentAccountName, id.BillingRoleDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Enrollment Account Billing Role Definition ID
func (id EnrollmentAccountBillingRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticEnrollmentAccounts", "enrollmentAccounts", "enrollmentAccounts"),
		resourceids.UserSpecifiedSegment("enrollmentAccountName", "enrollmentAccountValue"),
		resourceids.StaticSegment("staticBillingRoleDefinitions", "billingRoleDefinitions", "billingRoleDefinitions"),
		resourceids.UserSpecifiedSegment("billingRoleDefinitionName", "billingRoleDefinitionValue"),
	}
}

// String returns a human-readable description of this Enrollment Account Billing Role Definition ID
func (id EnrollmentAccountBillingRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Enrollment Account Name: %q", id.EnrollmentAccountName),
		fmt.Sprintf("Billing Role Definition Name: %q", id.BillingRoleDefinitionName),
	}
	return fmt.Sprintf("Enrollment Account Billing Role Definition (%s)", strings.Join(components, "\n"))
}
