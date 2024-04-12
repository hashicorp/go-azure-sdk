package costs

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CostId{})
}

var _ resourceids.ResourceId = &CostId{}

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
	parser := resourceids.NewParserFromResourceIdType(&CostId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CostId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCostIDInsensitively parses 'input' case-insensitively into a CostId
// note: this method should only be used for API response data and not user input
func ParseCostIDInsensitively(input string) (*CostId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CostId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CostId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CostId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.LabName, ok = input.Parsed["labName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "labName", input)
	}

	if id.CostName, ok = input.Parsed["costName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "costName", input)
	}

	return nil
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
