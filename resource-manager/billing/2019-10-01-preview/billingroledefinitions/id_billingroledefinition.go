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
	recaser.RegisterResourceId(&BillingRoleDefinitionId{})
}

var _ resourceids.ResourceId = &BillingRoleDefinitionId{}

// BillingRoleDefinitionId is a struct representing the Resource ID for a Billing Role Definition
type BillingRoleDefinitionId struct {
	BillingAccountName        string
	BillingRoleDefinitionName string
}

// NewBillingRoleDefinitionID returns a new BillingRoleDefinitionId struct
func NewBillingRoleDefinitionID(billingAccountName string, billingRoleDefinitionName string) BillingRoleDefinitionId {
	return BillingRoleDefinitionId{
		BillingAccountName:        billingAccountName,
		BillingRoleDefinitionName: billingRoleDefinitionName,
	}
}

// ParseBillingRoleDefinitionID parses 'input' into a BillingRoleDefinitionId
func ParseBillingRoleDefinitionID(input string) (*BillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingRoleDefinitionIDInsensitively parses 'input' case-insensitively into a BillingRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseBillingRoleDefinitionIDInsensitively(input string) (*BillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingRoleDefinitionName, ok = input.Parsed["billingRoleDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleDefinitionName", input)
	}

	return nil
}

// ValidateBillingRoleDefinitionID checks that 'input' can be parsed as a Billing Role Definition ID
func ValidateBillingRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Role Definition ID
func (id BillingRoleDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingRoleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingRoleDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Role Definition ID
func (id BillingRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountValue"),
		resourceids.StaticSegment("staticBillingRoleDefinitions", "billingRoleDefinitions", "billingRoleDefinitions"),
		resourceids.UserSpecifiedSegment("billingRoleDefinitionName", "billingRoleDefinitionValue"),
	}
}

// String returns a human-readable description of this Billing Role Definition ID
func (id BillingRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Role Definition Name: %q", id.BillingRoleDefinitionName),
	}
	return fmt.Sprintf("Billing Role Definition (%s)", strings.Join(components, "\n"))
}
