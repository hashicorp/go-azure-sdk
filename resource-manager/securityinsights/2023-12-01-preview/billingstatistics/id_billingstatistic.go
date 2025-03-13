package billingstatistics

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&BillingStatisticId{})
}

var _ resourceids.ResourceId = &BillingStatisticId{}

// BillingStatisticId is a struct representing the Resource ID for a Billing Statistic
type BillingStatisticId struct {
	SubscriptionId       string
	ResourceGroupName    string
	WorkspaceName        string
	BillingStatisticName string
}

// NewBillingStatisticID returns a new BillingStatisticId struct
func NewBillingStatisticID(subscriptionId string, resourceGroupName string, workspaceName string, billingStatisticName string) BillingStatisticId {
	return BillingStatisticId{
		SubscriptionId:       subscriptionId,
		ResourceGroupName:    resourceGroupName,
		WorkspaceName:        workspaceName,
		BillingStatisticName: billingStatisticName,
	}
}

// ParseBillingStatisticID parses 'input' into a BillingStatisticId
func ParseBillingStatisticID(input string) (*BillingStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingStatisticId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseBillingStatisticIDInsensitively parses 'input' case-insensitively into a BillingStatisticId
// note: this method should only be used for API response data and not user input
func ParseBillingStatisticIDInsensitively(input string) (*BillingStatisticId, error) {
	parser := resourceids.NewParserFromResourceIdType(&BillingStatisticId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := BillingStatisticId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *BillingStatisticId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.WorkspaceName, ok = input.Parsed["workspaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "workspaceName", input)
	}

	if id.BillingStatisticName, ok = input.Parsed["billingStatisticName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "billingStatisticName", input)
	}

	return nil
}

// ValidateBillingStatisticID checks that 'input' can be parsed as a Billing Statistic ID
func ValidateBillingStatisticID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseBillingStatisticID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Billing Statistic ID
func (id BillingStatisticId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.OperationalInsights/workspaces/%s/providers/Microsoft.SecurityInsights/billingStatistics/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.WorkspaceName, id.BillingStatisticName)
}

// Segments returns a slice of Resource ID Segments which comprise this Billing Statistic ID
func (id BillingStatisticId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftOperationalInsights", "Microsoft.OperationalInsights", "Microsoft.OperationalInsights"),
		resourceids.StaticSegment("staticWorkspaces", "workspaces", "workspaces"),
		resourceids.UserSpecifiedSegment("workspaceName", "workspaceName"),
		resourceids.StaticSegment("staticProviders2", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSecurityInsights", "Microsoft.SecurityInsights", "Microsoft.SecurityInsights"),
		resourceids.StaticSegment("staticBillingStatistics", "billingStatistics", "billingStatistics"),
		resourceids.UserSpecifiedSegment("billingStatisticName", "billingStatisticName"),
	}
}

// String returns a human-readable description of this Billing Statistic ID
func (id BillingStatisticId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Workspace Name: %q", id.WorkspaceName),
		fmt.Sprintf("Billing Statistic Name: %q", id.BillingStatisticName),
	}
	return fmt.Sprintf("Billing Statistic (%s)", strings.Join(components, "\n"))
}
