package aggregatedcost

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&Providers2BillingPeriodId{})
}

var _ resourceids.ResourceId = &Providers2BillingPeriodId{}

// Providers2BillingPeriodId is a struct representing the Resource ID for a Providers 2 Billing Period
type Providers2BillingPeriodId struct {
	ManagementGroupId string
	BillingPeriodName string
}

// NewProviders2BillingPeriodID returns a new Providers2BillingPeriodId struct
func NewProviders2BillingPeriodID(managementGroupId string, billingPeriodName string) Providers2BillingPeriodId {
	return Providers2BillingPeriodId{
		ManagementGroupId: managementGroupId,
		BillingPeriodName: billingPeriodName,
	}
}

// ParseProviders2BillingPeriodID parses 'input' into a Providers2BillingPeriodId
func ParseProviders2BillingPeriodID(input string) (*Providers2BillingPeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2BillingPeriodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2BillingPeriodId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseProviders2BillingPeriodIDInsensitively parses 'input' case-insensitively into a Providers2BillingPeriodId
// note: this method should only be used for API response data and not user input
func ParseProviders2BillingPeriodIDInsensitively(input string) (*Providers2BillingPeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&Providers2BillingPeriodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := Providers2BillingPeriodId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *Providers2BillingPeriodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ManagementGroupId, ok = input.Parsed["managementGroupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managementGroupId", input)
	}

	if id.BillingPeriodName, ok = input.Parsed["billingPeriodName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingPeriodName", input)
	}

	return nil
}

// ValidateProviders2BillingPeriodID checks that 'input' can be parsed as a Providers 2 Billing Period ID
func ValidateProviders2BillingPeriodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseProviders2BillingPeriodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Providers 2 Billing Period ID
func (id Providers2BillingPeriodId) ID() string {
	fmtString := "/providers/Microsoft.Management/managementGroups/%s/providers/Microsoft.Billing/billingPeriods/%s"
	return fmt.Sprintf(fmtString, id.ManagementGroupId, id.BillingPeriodName)
}

// Segments returns a slice of Resource ID Segments which comprise this Providers 2 Billing Period ID
func (id Providers2BillingPeriodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftManagement", "Microsoft.Management", "Microsoft.Management"),
		resourceids.StaticSegment("staticManagementGroups", "managementGroups", "managementGroups"),
		resourceids.UserSpecifiedSegment("managementGroupId", "managementGroupIdValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingPeriods", "billingPeriods", "billingPeriods"),
		resourceids.UserSpecifiedSegment("billingPeriodName", "billingPeriodValue"),
	}
}

// String returns a human-readable description of this Providers 2 Billing Period ID
func (id Providers2BillingPeriodId) String() string {
	components := []string{
		fmt.Sprintf("Management Group: %q", id.ManagementGroupId),
		fmt.Sprintf("Billing Period Name: %q", id.BillingPeriodName),
	}
	return fmt.Sprintf("Providers 2 Billing Period (%s)", strings.Join(components, "\n"))
}
