package iotsecuritysolutions

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IotSecuritySolutionId{}

// IotSecuritySolutionId is a struct representing the Resource ID for a Iot Security Solution
type IotSecuritySolutionId struct {
	SubscriptionId          string
	ResourceGroupName       string
	IotSecuritySolutionName string
}

// NewIotSecuritySolutionID returns a new IotSecuritySolutionId struct
func NewIotSecuritySolutionID(subscriptionId string, resourceGroupName string, iotSecuritySolutionName string) IotSecuritySolutionId {
	return IotSecuritySolutionId{
		SubscriptionId:          subscriptionId,
		ResourceGroupName:       resourceGroupName,
		IotSecuritySolutionName: iotSecuritySolutionName,
	}
}

// ParseIotSecuritySolutionID parses 'input' into a IotSecuritySolutionId
func ParseIotSecuritySolutionID(input string) (*IotSecuritySolutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(IotSecuritySolutionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IotSecuritySolutionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.IotSecuritySolutionName, ok = parsed.Parsed["iotSecuritySolutionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "iotSecuritySolutionName", *parsed)
	}

	return &id, nil
}

// ParseIotSecuritySolutionIDInsensitively parses 'input' case-insensitively into a IotSecuritySolutionId
// note: this method should only be used for API response data and not user input
func ParseIotSecuritySolutionIDInsensitively(input string) (*IotSecuritySolutionId, error) {
	parser := resourceids.NewParserFromResourceIdType(IotSecuritySolutionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IotSecuritySolutionId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.IotSecuritySolutionName, ok = parsed.Parsed["iotSecuritySolutionName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "iotSecuritySolutionName", *parsed)
	}

	return &id, nil
}

// ValidateIotSecuritySolutionID checks that 'input' can be parsed as a Iot Security Solution ID
func ValidateIotSecuritySolutionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIotSecuritySolutionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Iot Security Solution ID
func (id IotSecuritySolutionId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Security/iotSecuritySolutions/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.IotSecuritySolutionName)
}

// Segments returns a slice of Resource ID Segments which comprise this Iot Security Solution ID
func (id IotSecuritySolutionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurity", "Microsoft.Security", "Microsoft.Security"),
		resourceids.StaticSegment("staticIotSecuritySolutions", "iotSecuritySolutions", "iotSecuritySolutions"),
		resourceids.UserSpecifiedSegment("iotSecuritySolutionName", "iotSecuritySolutionValue"),
	}
}

// String returns a human-readable description of this Iot Security Solution ID
func (id IotSecuritySolutionId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Iot Security Solution Name: %q", id.IotSecuritySolutionName),
	}
	return fmt.Sprintf("Iot Security Solution (%s)", strings.Join(components, "\n"))
}
