package targets

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = TargetId{}

// TargetId is a struct representing the Resource ID for a Target
type TargetId struct {
	SubscriptionId     string
	ResourceGroupName  string
	ProviderName       string
	ParentResourceType string
	ParentResourceName string
	TargetName         string
}

// NewTargetID returns a new TargetId struct
func NewTargetID(subscriptionId string, resourceGroupName string, providerName string, parentResourceType string, parentResourceName string, targetName string) TargetId {
	return TargetId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		ProviderName:       providerName,
		ParentResourceType: parentResourceType,
		ParentResourceName: parentResourceName,
		TargetName:         targetName,
	}
}

// ParseTargetID parses 'input' into a TargetId
func ParseTargetID(input string) (*TargetId, error) {
	parser := resourceids.NewParserFromResourceIdType(TargetId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TargetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ParentResourceType, ok = parsed.Parsed["parentResourceType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "parentResourceType", *parsed)
	}

	if id.ParentResourceName, ok = parsed.Parsed["parentResourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "parentResourceName", *parsed)
	}

	if id.TargetName, ok = parsed.Parsed["targetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "targetName", *parsed)
	}

	return &id, nil
}

// ParseTargetIDInsensitively parses 'input' case-insensitively into a TargetId
// note: this method should only be used for API response data and not user input
func ParseTargetIDInsensitively(input string) (*TargetId, error) {
	parser := resourceids.NewParserFromResourceIdType(TargetId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := TargetId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.ProviderName, ok = parsed.Parsed["providerName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "providerName", *parsed)
	}

	if id.ParentResourceType, ok = parsed.Parsed["parentResourceType"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "parentResourceType", *parsed)
	}

	if id.ParentResourceName, ok = parsed.Parsed["parentResourceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "parentResourceName", *parsed)
	}

	if id.TargetName, ok = parsed.Parsed["targetName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "targetName", *parsed)
	}

	return &id, nil
}

// ValidateTargetID checks that 'input' can be parsed as a Target ID
func ValidateTargetID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseTargetID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Target ID
func (id TargetId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/%s/%s/%s/providers/Microsoft.Chaos/targets/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ProviderName, id.ParentResourceType, id.ParentResourceName, id.TargetName)
}

// Segments returns a slice of Resource ID Segments which comprise this Target ID
func (id TargetId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.UserSpecifiedSegment("providerName", "providerValue"),
		resourceids.UserSpecifiedSegment("parentResourceType", "parentResourceTypeValue"),
		resourceids.UserSpecifiedSegment("parentResourceName", "parentResourceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftChaos", "Microsoft.Chaos", "Microsoft.Chaos"),
		resourceids.StaticSegment("staticTargets", "targets", "targets"),
		resourceids.UserSpecifiedSegment("targetName", "targetValue"),
	}
}

// String returns a human-readable description of this Target ID
func (id TargetId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Provider Name: %q", id.ProviderName),
		fmt.Sprintf("Parent Resource Type: %q", id.ParentResourceType),
		fmt.Sprintf("Parent Resource Name: %q", id.ParentResourceName),
		fmt.Sprintf("Target Name: %q", id.TargetName),
	}
	return fmt.Sprintf("Target (%s)", strings.Join(components, "\n"))
}
