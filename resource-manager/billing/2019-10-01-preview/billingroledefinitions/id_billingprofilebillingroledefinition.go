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
	recaser.RegisterResourceId(&BillingProfileBillingRoleDefinitionId{})
}

var _ resourceids.ResourceId = &BillingProfileBillingRoleDefinitionId{}

// BillingProfileBillingRoleDefinitionId is a struct representing the Resource ID for a Billing Profile Billing Role Definition
type BillingProfileBillingRoleDefinitionId struct {
	BillingAccountName        string
	BillingProfileName        string
	BillingRoleDefinitionName string
}

// NewBillingProfileBillingRoleDefinitionID returns a new BillingProfileBillingRoleDefinitionId struct
func NewBillingProfileBillingRoleDefinitionID(billingAccountName string, billingProfileName string, billingRoleDefinitionName string) BillingProfileBillingRoleDefinitionId {
	return BillingProfileBillingRoleDefinitionId{
		BillingAccountName:        billingAccountName,
		BillingProfileName:        billingProfileName,
		BillingRoleDefinitionName: billingRoleDefinitionName,
	}
}

// ParseBillingProfileBillingRoleDefinitionID parses 'input' into a BillingProfileBillingRoleDefinitionId
func ParseBillingProfileBillingRoleDefinitionID(input string) (*BillingProfileBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingProfileBillingRoleDefinitionIDInsensitively parses 'input' case-insensitively into a BillingProfileBillingRoleDefinitionId
// note: this method should only be used for API response data and not user input
func ParseBillingProfileBillingRoleDefinitionIDInsensitively(input string) (*BillingProfileBillingRoleDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingProfileBillingRoleDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingProfileBillingRoleDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingProfileBillingRoleDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountName, ok = input.Parsed["billingAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountName", input)
	}

	if id.BillingProfileName, ok = input.Parsed["billingProfileName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingProfileName", input)
	}

	if id.BillingRoleDefinitionName, ok = input.Parsed["billingRoleDefinitionName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingRoleDefinitionName", input)
	}

	return nil
}

// ValidateBillingProfileBillingRoleDefinitionID checks that 'input' can be parsed as a Billing Profile Billing Role Definition ID
func ValidateBillingProfileBillingRoleDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingProfileBillingRoleDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Profile Billing Role Definition ID
func (id BillingProfileBillingRoleDefinitionId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingProfiles/%s/billingRoleDefinitions/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountName, id.BillingProfileName, id.BillingRoleDefinitionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Profile Billing Role Definition ID
func (id BillingProfileBillingRoleDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountName", "billingAccountName"),
		resourceids.StaticSegment("staticBillingProfiles", "billingProfiles", "billingProfiles"),
		resourceids.UserSpecifiedSegment("billingProfileName", "billingProfileName"),
		resourceids.StaticSegment("staticBillingRoleDefinitions", "billingRoleDefinitions", "billingRoleDefinitions"),
		resourceids.UserSpecifiedSegment("billingRoleDefinitionName", "billingRoleDefinitionName"),
	}
}

// String returns a human-readable description of this Billing Profile Billing Role Definition ID
func (id BillingProfileBillingRoleDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account Name: %q", id.BillingAccountName),
		fmt.Sprintf("Billing Profile Name: %q", id.BillingProfileName),
		fmt.Sprintf("Billing Role Definition Name: %q", id.BillingRoleDefinitionName),
	}
	return fmt.Sprintf("Billing Profile Billing Role Definition (%s)", strings.Join(components, "\n"))
}
