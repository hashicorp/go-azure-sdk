package balances

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = BillingAccountBillingPeriodId{}

// BillingAccountBillingPeriodId is a struct representing the Resource ID for a Billing Account Billing Period
type BillingAccountBillingPeriodId struct {
	BillingAccountId  string
	BillingPeriodName string
}

// NewBillingAccountBillingPeriodID returns a new BillingAccountBillingPeriodId struct
func NewBillingAccountBillingPeriodID(billingAccountId string, billingPeriodName string) BillingAccountBillingPeriodId {
	return BillingAccountBillingPeriodId{
		BillingAccountId:  billingAccountId,
		BillingPeriodName: billingPeriodName,
	}
}

// ParseBillingAccountBillingPeriodID parses 'input' into a BillingAccountBillingPeriodId
func ParseBillingAccountBillingPeriodID(input string) (*BillingAccountBillingPeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(BillingAccountBillingPeriodId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BillingAccountBillingPeriodId{}

	if id.BillingAccountId, ok = parsed.Parsed["billingAccountId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "billingAccountId", *parsed)
	}

	if id.BillingPeriodName, ok = parsed.Parsed["billingPeriodName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "billingPeriodName", *parsed)
	}

	return &id, nil
}

// ParseBillingAccountBillingPeriodIDInsensitively parses 'input' case-insensitively into a BillingAccountBillingPeriodId
// note: this method should only be used for API response data and not user input
func ParseBillingAccountBillingPeriodIDInsensitively(input string) (*BillingAccountBillingPeriodId, error) {
	parser := resourceids.NewParserFromResourceIdType(BillingAccountBillingPeriodId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := BillingAccountBillingPeriodId{}

	if id.BillingAccountId, ok = parsed.Parsed["billingAccountId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "billingAccountId", *parsed)
	}

	if id.BillingPeriodName, ok = parsed.Parsed["billingPeriodName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "billingPeriodName", *parsed)
	}

	return &id, nil
}

// ValidateBillingAccountBillingPeriodID checks that 'input' can be parsed as a Billing Account Billing Period ID
func ValidateBillingAccountBillingPeriodID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingAccountBillingPeriodID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Account Billing Period ID
func (id BillingAccountBillingPeriodId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/billingPeriods/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountId, id.BillingPeriodName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Account Billing Period ID
func (id BillingAccountBillingPeriodId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountId", "billingAccountIdValue"),
		resourceids.StaticSegment("staticBillingPeriods", "billingPeriods", "billingPeriods"),
		resourceids.UserSpecifiedSegment("billingPeriodName", "billingPeriodValue"),
	}
}

// String returns a human-readable description of this Billing Account Billing Period ID
func (id BillingAccountBillingPeriodId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account: %q", id.BillingAccountId),
		fmt.Sprintf("Billing Period Name: %q", id.BillingPeriodName),
	}
	return fmt.Sprintf("Billing Account Billing Period (%s)", strings.Join(components, "\n"))
}
