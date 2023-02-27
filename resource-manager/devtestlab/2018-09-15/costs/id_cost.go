// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package costs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

var _ resourceids.ResourceId = CostId{}

// CostId is a struct representing the Resource ID for a Cost
type CostId struct {
	SubscriptionId    string
	ResourceGroupName string
	LabName           string
	CostName          string
}

// NewCostID returns a new CostId struct
func NewCostID(subscriptionId string, resourceGroupName string, labName string, costName string) CostId {
	return CostId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LabName:           labName,
		CostName:          costName,
	}
}

// ParseCostID parses 'input' into a CostId
func ParseCostID(input string) (*CostId, error) {
	parser := resourceids.NewParserFromResourceIdType(CostId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CostId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.CostName, ok = parsed.Parsed["costName"]; !ok {
		return nil, fmt.Errorf("the segment 'costName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseCostIDInsensitively parses 'input' case-insensitively into a CostId
// note: this method should only be used for API response data and not user input
func ParseCostIDInsensitively(input string) (*CostId, error) {
	parser := resourceids.NewParserFromResourceIdType(CostId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := CostId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, fmt.Errorf("the segment 'labName' was not found in the resource id %q", input)
	}

	if id.CostName, ok = parsed.Parsed["costName"]; !ok {
		return nil, fmt.Errorf("the segment 'costName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateCostID checks that 'input' can be parsed as a Cost ID
func ValidateCostID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCostID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cost ID
func (id CostId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/costs/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.CostName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cost ID
func (id CostId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticCosts", "costs", "costs"),
		resourceids.UserSpecifiedSegment("costName", "costValue"),
	}
}

// String returns a human-readable description of this Cost ID
func (id CostId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Cost Name: %q", id.CostName),
	}
	return fmt.Sprintf("Cost (%s)", strings.Join(components, "\n"))
}
