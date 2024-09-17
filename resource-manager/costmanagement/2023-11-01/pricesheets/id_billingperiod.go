package pricesheets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingPeriodId{})
}

var _ resourceids.ResourceId = &BillingPeriodId{}

// BillingPeriodId is a struct representing the Resource ID for a Billing Period
type BillingPeriodId struct {
	BillingAccountId  string
	BillingPeriodName string
}

// NewBillingPeriodID returns a new BillingPeriodId struct
func NewBillingPeriodID(billingAccountId string, billingPeriodName string) BillingPeriodId {
	return BillingPeriodId{
		BillingAccountId:  billingAccountId,
		BillingPeriodName: billingPeriodName,
	}
}

// ParseBillingPeriodID parses 'input' into a BillingPeriodId
func ParseBillingPeriodID(input string) (*BillingPeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingPeriodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingPeriodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingPeriodIDInsensitively parses 'input' case-insensitively into a BillingPeriodId
// note: this method should only be used for API response data and not user input
func ParseBillingPeriodIDInsensitively(input string) (*BillingPeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingPeriodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingPeriodId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingPeriodId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountId, ok = input.Parsed["billingAccountId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountId", input)
	}

	if id.BillingPeriodName, ok = input.Parsed["billingPeriodName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingPeriodName", input)
	}

	return nil
}

// ValidateBillingPeriodID checks that 'input' can be parsed as a Billing Period ID
func ValidateBillingPeriodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingPeriodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Period ID
func (id BillingPeriodId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingPeriods/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountId, id.BillingPeriodName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Period ID
func (id BillingPeriodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountId", "billingAccountIdValue"),
		resourceids.StaticSegment("staticBillingPeriods", "billingPeriods", "billingPeriods"),
		resourceids.UserSpecifiedSegment("billingPeriodName", "billingPeriodValue"),
	}
}

// String returns a human-readable description of this Billing Period ID
func (id BillingPeriodId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account: %q", id.BillingAccountId),
		fmt.Sprintf("Billing Period Name: %q", id.BillingPeriodName),
	}
	return fmt.Sprintf("Billing Period (%s)", strings.Join(components, "\n"))
}
