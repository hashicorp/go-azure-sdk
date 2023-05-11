package applyupdates

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = ApplyUpdateId{}

// ApplyUpdateId is a struct representing the Resource ID for a Apply Update
type ApplyUpdateId struct {
	SubscriptionId    string
	ResourceGroupName string
	ProviderName      string
	ResourceType      string
	ResourceName      string
	ApplyUpdateName   string
}

// NewApplyUpdateID returns a new ApplyUpdateId struct
func NewApplyUpdateID(subscriptionId string, resourceGroupName string, providerName string, resourceType string, resourceName string, applyUpdateName string) ApplyUpdateId {
	return ApplyUpdateId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		ProviderName:      providerName,
		ResourceType:      resourceType,
		ResourceName:      resourceName,
		ApplyUpdateName:   applyUpdateName,
	}
}

// ParseApplyUpdateID parses 'input' into a ApplyUpdateId
func ParseApplyUpdateID(input string) (*ApplyUpdateId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplyUpdateId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplyUpdateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceType", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	if id.ApplyUpdateName, ok = parsed.Parsed["applyUpdateName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applyUpdateName", *parsed)
	}

	return &id, nil
}

// ParseApplyUpdateIDInsensitively parses 'input' case-insensitively into a ApplyUpdateId
// note: this method should only be used for API response data and not user input
func ParseApplyUpdateIDInsensitively(input string) (*ApplyUpdateId, error) {
	parser := resourceids.NewParserFromResourceIdType(ApplyUpdateId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := ApplyUpdateId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ResourceType, ok = parsed.Parsed["resourceType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceType", *parsed)
	}

	if id.ResourceName, ok = parsed.Parsed["resourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceName", *parsed)
	}

	if id.ApplyUpdateName, ok = parsed.Parsed["applyUpdateName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "applyUpdateName", *parsed)
	}

	return &id, nil
}

// ValidateApplyUpdateID checks that 'input' can be parsed as a Apply Update ID
func ValidateApplyUpdateID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseApplyUpdateID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Apply Update ID
func (id ApplyUpdateId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.Maintenance/applyUpdates/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ResourceType, id.ResourceName, id.ApplyUpdateName)
}

// Segments returns a slice of Resource ID Segments which comprise this Apply Update ID
func (id ApplyUpdateId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("resourceType", "resourceTypeValue"),
		resourceids.UserSpecifiedSegment("resourceName", "resourceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftMaintenance", "Microsoft.Maintenance", "Microsoft.Maintenance"),
		resourceids.StaticSegment("staticApplyUpdates", "applyUpdates", "applyUpdates"),
		resourceids.UserSpecifiedSegment("applyUpdateName", "applyUpdateValue"),
	}
}

// String returns a human-readable description of this Apply Update ID
func (id ApplyUpdateId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Resource Type: %q", id.ResourceType),
		fmt.Sprintf("Resource Name: %q", id.ResourceName),
		fmt.Sprintf("Apply Update Name: %q", id.ApplyUpdateName),
	}
	return fmt.Sprintf("Apply Update (%s)", strings.Join(components, "\n"))
}
