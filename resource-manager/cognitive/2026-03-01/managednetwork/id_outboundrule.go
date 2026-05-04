package managednetwork

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&OutboundRuleId{})
}

var _ resourceids.ResourceId = &OutboundRuleId{}

// OutboundRuleId is a struct representing the Resource ID for a Outbound Rule
type OutboundRuleId struct {
	SubscriptionId     string
	ResourceGroupName  string
	AccountName        string
	ManagedNetworkName string
	OutboundRuleName   string
}

// NewOutboundRuleID returns a new OutboundRuleId struct
func NewOutboundRuleID(subscriptionId string, resourceGroupName string, accountName string, managedNetworkName string, outboundRuleName string) OutboundRuleId {
	return OutboundRuleId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		AccountName:        accountName,
		ManagedNetworkName: managedNetworkName,
		OutboundRuleName:   outboundRuleName,
	}
}

// ParseOutboundRuleID parses 'input' into a OutboundRuleId
func ParseOutboundRuleID(input string) (*OutboundRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OutboundRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OutboundRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseOutboundRuleIDInsensitively parses 'input' case-insensitively into a OutboundRuleId
// note: this method should only be used for API response data and not user input
func ParseOutboundRuleIDInsensitively(input string) (*OutboundRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(&OutboundRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := OutboundRuleId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *OutboundRuleId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.AccountName, ok = input.Parsed["accountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accountName", input)
	}

	if id.ManagedNetworkName, ok = input.Parsed["managedNetworkName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedNetworkName", input)
	}

	if id.OutboundRuleName, ok = input.Parsed["outboundRuleName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "outboundRuleName", input)
	}

	return nil
}

// ValidateOutboundRuleID checks that 'input' can be parsed as a Outbound Rule ID
func ValidateOutboundRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseOutboundRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Outbound Rule ID
func (id OutboundRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/managedNetworks/%s/outboundRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ManagedNetworkName, id.OutboundRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Outbound Rule ID
func (id OutboundRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCognitiveServices", "Microsoft.CognitiveServices", "Microsoft.CognitiveServices"),
		resourceids.StaticSegment("staticAccounts", "accounts", "accounts"),
		resourceids.UserSpecifiedSegment("accountName", "accountName"),
		resourceids.StaticSegment("staticManagedNetworks", "managedNetworks", "managedNetworks"),
		resourceids.UserSpecifiedSegment("managedNetworkName", "managedNetworkName"),
		resourceids.StaticSegment("staticOutboundRules", "outboundRules", "outboundRules"),
		resourceids.UserSpecifiedSegment("outboundRuleName", "outboundRuleName"),
	}
}

// String returns a human-readable description of this Outbound Rule ID
func (id OutboundRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Managed Network Name: %q", id.ManagedNetworkName),
		fmt.Sprintf("Outbound Rule Name: %q", id.OutboundRuleName),
	}
	return fmt.Sprintf("Outbound Rule (%s)", strings.Join(components, "\n"))
}
