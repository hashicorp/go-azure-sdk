package autoscalevcores

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &AutoScaleVCoreId{}

// AutoScaleVCoreId is a struct representing the Resource ID for a Auto Scale V Core
type AutoScaleVCoreId struct {
	SubscriptionId     string
	ResourceGroupName  string
	AutoScaleVCoreName string
}

// NewAutoScaleVCoreID returns a new AutoScaleVCoreId struct
func NewAutoScaleVCoreID(subscriptionId string, resourceGroupName string, autoScaleVCoreName string) AutoScaleVCoreId {
	return AutoScaleVCoreId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		AutoScaleVCoreName: autoScaleVCoreName,
	}
}

// ParseAutoScaleVCoreID parses 'input' into a AutoScaleVCoreId
func ParseAutoScaleVCoreID(input string) (*AutoScaleVCoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AutoScaleVCoreId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AutoScaleVCoreId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseAutoScaleVCoreIDInsensitively parses 'input' case-insensitively into a AutoScaleVCoreId
// note: this method should only be used for API response data and not user input
func ParseAutoScaleVCoreIDInsensitively(input string) (*AutoScaleVCoreId, error) {
	parser := resourceids.NewParserFromResourceIdType(&AutoScaleVCoreId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := AutoScaleVCoreId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *AutoScaleVCoreId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AutoScaleVCoreName, ok = input.Parsed["autoScaleVCoreName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "autoScaleVCoreName", input)
	}

	return nil
}

// ValidateAutoScaleVCoreID checks that 'input' can be parsed as a Auto Scale V Core ID
func ValidateAutoScaleVCoreID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseAutoScaleVCoreID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Auto Scale V Core ID
func (id AutoScaleVCoreId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.PowerBIDedicated/autoScaleVCores/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AutoScaleVCoreName)
}

// Segments returns a slice of Resource ID Segments which comprise this Auto Scale V Core ID
func (id AutoScaleVCoreId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftPowerBIDedicated", "Microsoft.PowerBIDedicated", "Microsoft.PowerBIDedicated"),
		resourceids.StaticSegment("staticAutoScaleVCores", "autoScaleVCores", "autoScaleVCores"),
		resourceids.UserSpecifiedSegment("autoScaleVCoreName", "autoScaleVCoreValue"),
	}
}

// String returns a human-readable description of this Auto Scale V Core ID
func (id AutoScaleVCoreId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Auto Scale V Core Name: %q", id.AutoScaleVCoreName),
	}
	return fmt.Sprintf("Auto Scale V Core (%s)", strings.Join(components, "\n"))
}
