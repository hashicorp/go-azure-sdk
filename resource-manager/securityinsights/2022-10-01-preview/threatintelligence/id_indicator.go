package threatintelligence

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = IndicatorId{}

// IndicatorId is a struct representing the Resource ID for a Indicator
type IndicatorId struct {
	SubscriptionId    string
	ResourceGroupName string
	WorkspaceName     string
	IndicatorName     string
}

// NewIndicatorID returns a new IndicatorId struct
func NewIndicatorID(subscriptionId string, resourceGroupName string, workspaceName string, indicatorName string) IndicatorId {
	return IndicatorId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		WorkspaceName:     workspaceName,
		IndicatorName:     indicatorName,
	}
}

// ParseIndicatorID parses 'input' into a IndicatorId
func ParseIndicatorID(input string) (*IndicatorId, error) {
	parser := resourceids.NewParserFromResourceIdType(IndicatorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IndicatorId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.IndicatorName, ok = parsed.Parsed["indicatorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "indicatorName", *parsed)
	}

	return &id, nil
}

// ParseIndicatorIDInsensitively parses 'input' case-insensitively into a IndicatorId
// note: this method should only be used for API response data and not user input
func ParseIndicatorIDInsensitively(input string) (*IndicatorId, error) {
	parser := resourceids.NewParserFromResourceIdType(IndicatorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := IndicatorId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.WorkspaceName, ok = parsed.Parsed["workspaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", *parsed)
	}

	if id.IndicatorName, ok = parsed.Parsed["indicatorName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "indicatorName", *parsed)
	}

	return &id, nil
}

// ValidateIndicatorID checks that 'input' can be parsed as a Indicator ID
func ValidateIndicatorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIndicatorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Indicator ID
func (id IndicatorId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/threatIntelligence/main/indicators/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.IndicatorName)
}

// Segments returns a slice of Resource ID Segments which comprise this Indicator ID
func (id IndicatorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceValue"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticThreatIntelligence", "threatIntelligence", "threatIntelligence"),
		resourceids.StaticSegment("staticMain", "main", "main"),
		resourceids.StaticSegment("staticIndicators", "indicators", "indicators"),
		resourceids.UserSpecifiedSegment("indicatorName", "indicatorValue"),
	}
}

// String returns a human-readable description of this Indicator ID
func (id IndicatorId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Indicator Name: %q", id.IndicatorName),
	}
	return fmt.Sprintf("Indicator (%s)", strings.Join(components, "\n"))
}
