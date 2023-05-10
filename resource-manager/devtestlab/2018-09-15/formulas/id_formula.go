package formulas

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = FormulaId{}

// FormulaId is a struct representing the Resource ID for a Formula
type FormulaId struct {
	SubscriptionId    string
	ResourceGroupName string
	LabName           string
	FormulaName       string
}

// NewFormulaID returns a new FormulaId struct
func NewFormulaID(subscriptionId string, resourceGroupName string, labName string, formulaName string) FormulaId {
	return FormulaId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		LabName:           labName,
		FormulaName:       formulaName,
	}
}

// ParseFormulaID parses 'input' into a FormulaId
func ParseFormulaID(input string) (*FormulaId, error) {
	parser := resourceids.NewParserFromResourceIdType(FormulaId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FormulaId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "labName", *parsed)
	}

	if id.FormulaName, ok = parsed.Parsed["formulaName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "formulaName", *parsed)
	}

	return &id, nil
}

// ParseFormulaIDInsensitively parses 'input' case-insensitively into a FormulaId
// note: this method should only be used for API response data and not user input
func ParseFormulaIDInsensitively(input string) (*FormulaId, error) {
	parser := resourceids.NewParserFromResourceIdType(FormulaId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := FormulaId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.LabName, ok = parsed.Parsed["labName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "labName", *parsed)
	}

	if id.FormulaName, ok = parsed.Parsed["formulaName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "formulaName", *parsed)
	}

	return &id, nil
}

// ValidateFormulaID checks that 'input' can be parsed as a Formula ID
func ValidateFormulaID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseFormulaID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Formula ID
func (id FormulaId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DevTestLab/labs/%s/formulas/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.LabName, id.FormulaName)
}

// Segments returns a slice of Resource ID Segments which comprise this Formula ID
func (id FormulaId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftDevTestLab", "Microsoft.DevTestLab", "Microsoft.DevTestLab"),
		resourceids.StaticSegment("staticLabs", "labs", "labs"),
		resourceids.UserSpecifiedSegment("labName", "labValue"),
		resourceids.StaticSegment("staticFormulas", "formulas", "formulas"),
		resourceids.UserSpecifiedSegment("formulaName", "formulaValue"),
	}
}

// String returns a human-readable description of this Formula ID
func (id FormulaId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Lab Name: %q", id.LabName),
		fmt.Sprintf("Formula Name: %q", id.FormulaName),
	}
	return fmt.Sprintf("Formula (%s)", strings.Join(components, "\n"))
}
