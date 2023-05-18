package benefitutilizationsummaries

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = SavingsPlanOrderId{}

// SavingsPlanOrderId is a struct representing the Resource ID for a Savings Plan Order
type SavingsPlanOrderId struct {
	SavingsPlanOrderId string
}

// NewSavingsPlanOrderID returns a new SavingsPlanOrderId struct
func NewSavingsPlanOrderID(savingsPlanOrderId string) SavingsPlanOrderId {
	return SavingsPlanOrderId{
		SavingsPlanOrderId: savingsPlanOrderId,
	}
}

// ParseSavingsPlanOrderID parses 'input' into a SavingsPlanOrderId
func ParseSavingsPlanOrderID(input string) (*SavingsPlanOrderId, error) {
	parser := resourceids.NewParserFromResourceIdType(SavingsPlanOrderId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SavingsPlanOrderId{}

	if id.SavingsPlanOrderId, ok = parsed.Parsed["savingsPlanOrderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "savingsPlanOrderId", *parsed)
	}

	return &id, nil
}

// ParseSavingsPlanOrderIDInsensitively parses 'input' case-insensitively into a SavingsPlanOrderId
// note: this method should only be used for API response data and not user input
func ParseSavingsPlanOrderIDInsensitively(input string) (*SavingsPlanOrderId, error) {
	parser := resourceids.NewParserFromResourceIdType(SavingsPlanOrderId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := SavingsPlanOrderId{}

	if id.SavingsPlanOrderId, ok = parsed.Parsed["savingsPlanOrderId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "savingsPlanOrderId", *parsed)
	}

	return &id, nil
}

// ValidateSavingsPlanOrderID checks that 'input' can be parsed as a Savings Plan Order ID
func ValidateSavingsPlanOrderID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseSavingsPlanOrderID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Savings Plan Order ID
func (id SavingsPlanOrderId) ID() string {
	fmtString := "/providers/Microsoft.BillingBenefits/savingsPlanOrders/%s"
	return fmt.Sprintf(fmtString, id.SavingsPlanOrderId)
}

// Segments returns a slice of Resource ID Segments which comprise this Savings Plan Order ID
func (id SavingsPlanOrderId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBillingBenefits", "Microsoft.BillingBenefits", "Microsoft.BillingBenefits"),
		resourceids.StaticSegment("staticSavingsPlanOrders", "savingsPlanOrders", "savingsPlanOrders"),
		resourceids.UserSpecifiedSegment("savingsPlanOrderId", "savingsPlanOrderIdValue"),
	}
}

// String returns a human-readable description of this Savings Plan Order ID
func (id SavingsPlanOrderId) String() string {
	components := []string{
		fmt.Sprintf("Savings Plan Order: %q", id.SavingsPlanOrderId),
	}
	return fmt.Sprintf("Savings Plan Order (%s)", strings.Join(components, "\n"))
}
