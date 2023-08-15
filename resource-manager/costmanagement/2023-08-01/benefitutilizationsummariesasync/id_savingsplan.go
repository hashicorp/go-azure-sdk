package benefitutilizationsummariesasync

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SavingsPlanId{}

// SavingsPlanId is a struct representing the Resource ID for a Savings Plan
type SavingsPlanId struct {
	SavingsPlanOrderId string
	SavingsPlanId      string
}

// NewSavingsPlanID returns a new SavingsPlanId struct
func NewSavingsPlanID(savingsPlanOrderId string, savingsPlanId string) SavingsPlanId {
	return SavingsPlanId{
		SavingsPlanOrderId: savingsPlanOrderId,
		SavingsPlanId:      savingsPlanId,
	}
}

// ParseSavingsPlanID parses 'input' into a SavingsPlanId
func ParseSavingsPlanID(input string) (*SavingsPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(SavingsPlanId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SavingsPlanId{}

	if id.SavingsPlanOrderId, ok = parsed.Parsed["savingsPlanOrderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "savingsPlanOrderId", *parsed)
	}

	if id.SavingsPlanId, ok = parsed.Parsed["savingsPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "savingsPlanId", *parsed)
	}

	return &id, nil
}

// ParseSavingsPlanIDInsensitively parses 'input' case-insensitively into a SavingsPlanId
// note: this method should only be used for API response data and not user input
func ParseSavingsPlanIDInsensitively(input string) (*SavingsPlanId, error) {
	parser := resourceids.NewParserFromResourceIdType(SavingsPlanId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SavingsPlanId{}

	if id.SavingsPlanOrderId, ok = parsed.Parsed["savingsPlanOrderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "savingsPlanOrderId", *parsed)
	}

	if id.SavingsPlanId, ok = parsed.Parsed["savingsPlanId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "savingsPlanId", *parsed)
	}

	return &id, nil
}

// ValidateSavingsPlanID checks that 'input' can be parsed as a Savings Plan ID
func ValidateSavingsPlanID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSavingsPlanID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Savings Plan ID
func (id SavingsPlanId) ID() string {
	fmtString := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/%s/savingsPlans/%s"
	return fmt.Sprintf(fmtString, id.SavingsPlanOrderId, id.SavingsPlanId)
}

// Segments returns a slice of Resource ID Segments which comprise this Savings Plan ID
func (id SavingsPlanId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBillingBenefits", "Microsoft.BillingBenefits", "Microsoft.BillingBenefits"),
		resourceids.StaticSegment("staticSavingsPlanOrders", "savingsPlanOrders", "savingsPlanOrders"),
		resourceids.UserSpecifiedSegment("savingsPlanOrderId", "savingsPlanOrderIdValue"),
		resourceids.StaticSegment("staticSavingsPlans", "savingsPlans", "savingsPlans"),
		resourceids.UserSpecifiedSegment("savingsPlanId", "savingsPlanIdValue"),
	}
}

// String returns a human-readable description of this Savings Plan ID
func (id SavingsPlanId) String() string {
	components := []string{
		fmt.Sprintf("Savings Plan Order: %q", id.SavingsPlanOrderId),
		fmt.Sprintf("Savings Plan: %q", id.SavingsPlanId),
	}
	return fmt.Sprintf("Savings Plan (%s)", strings.Join(components, "\n"))
}
