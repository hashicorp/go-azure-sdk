package costallocationrules

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CostAllocationRuleId{})
}

var _ resourceids.ResourceId = &CostAllocationRuleId{}

// CostAllocationRuleId is a struct representing the Resource ID for a Cost Allocation Rule
type CostAllocationRuleId struct {
	BillingAccountId       string
	CostAllocationRuleName string
}

// NewCostAllocationRuleID returns a new CostAllocationRuleId struct
func NewCostAllocationRuleID(billingAccountId string, costAllocationRuleName string) CostAllocationRuleId {
	return CostAllocationRuleId{
		BillingAccountId:       billingAccountId,
		CostAllocationRuleName: costAllocationRuleName,
	}
}

// ParseCostAllocationRuleID parses 'input' into a CostAllocationRuleId
func ParseCostAllocationRuleID(input string) (*CostAllocationRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CostAllocationRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CostAllocationRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCostAllocationRuleIDInsensitively parses 'input' case-insensitively into a CostAllocationRuleId
// note: this method should only be used for API response data and not user input
func ParseCostAllocationRuleIDInsensitively(input string) (*CostAllocationRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CostAllocationRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CostAllocationRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CostAllocationRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BillingAccountId, ok = input.Parsed["billingAccountId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingAccountId", input)
	}

	if id.CostAllocationRuleName, ok = input.Parsed["costAllocationRuleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "costAllocationRuleName", input)
	}

	return nil
}

// ValidateCostAllocationRuleID checks that 'input' can be parsed as a Cost Allocation Rule ID
func ValidateCostAllocationRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCostAllocationRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cost Allocation Rule ID
func (id CostAllocationRuleId) ID() string {
	fmtString := "/providers/Microsoft.Billing/billingAccounts/%s/providers/Microsoft.CostManagement/costAllocationRules/%s"
	return fmt.Sprintf(fmtString, id.BillingAccountId, id.CostAllocationRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cost Allocation Rule ID
func (id CostAllocationRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftBilling", "Microsoft.Billing", "Microsoft.Billing"),
		resourceids.StaticSegment("staticBillingAccounts", "billingAccounts", "billingAccounts"),
		resourceids.UserSpecifiedSegment("billingAccountId", "billingAccountId"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCostManagement", "Microsoft.CostManagement", "Microsoft.CostManagement"),
		resourceids.StaticSegment("staticCostAllocationRules", "costAllocationRules", "costAllocationRules"),
		resourceids.UserSpecifiedSegment("costAllocationRuleName", "costAllocationRuleName"),
	}
}

// String returns a human-readable description of this Cost Allocation Rule ID
func (id CostAllocationRuleId) String() string {
	components := []string{
		fmt.Sprintf("Billing Account: %q", id.BillingAccountId),
		fmt.Sprintf("Cost Allocation Rule Name: %q", id.CostAllocationRuleName),
	}
	return fmt.Sprintf("Cost Allocation Rule (%s)", strings.Join(components, "\n"))
}
