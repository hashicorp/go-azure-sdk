package wcfrelays

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = WcfRelayAuthorizationRuleId{}

// WcfRelayAuthorizationRuleId is a struct representing the Resource ID for a Wcf Relay Authorization Rule
type WcfRelayAuthorizationRuleId struct {
	SubscriptionId        string
	ResourceGroupName     string
	NamespaceName         string
	WcfRelayName          string
	AuthorizationRuleName string
}

// NewWcfRelayAuthorizationRuleID returns a new WcfRelayAuthorizationRuleId struct
func NewWcfRelayAuthorizationRuleID(subscriptionId string, resourceGroupName string, namespaceName string, wcfRelayName string, authorizationRuleName string) WcfRelayAuthorizationRuleId {
	return WcfRelayAuthorizationRuleId{
		SubscriptionId:        subscriptionId,
		ResourceGroupName:     resourceGroupName,
		NamespaceName:         namespaceName,
		WcfRelayName:          wcfRelayName,
		AuthorizationRuleName: authorizationRuleName,
	}
}

// ParseWcfRelayAuthorizationRuleID parses 'input' into a WcfRelayAuthorizationRuleId
func ParseWcfRelayAuthorizationRuleID(input string) (*WcfRelayAuthorizationRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(WcfRelayAuthorizationRuleId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WcfRelayAuthorizationRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.NamespaceName, ok = parsed.Parsed["namespaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "namespaceName", *parsed)
	}

	if id.WcfRelayName, ok = parsed.Parsed["wcfRelayName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "wcfRelayName", *parsed)
	}

	if id.AuthorizationRuleName, ok = parsed.Parsed["authorizationRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authorizationRuleName", *parsed)
	}

	return &id, nil
}

// ParseWcfRelayAuthorizationRuleIDInsensitively parses 'input' case-insensitively into a WcfRelayAuthorizationRuleId
// note: this method should only be used for API response data and not user input
func ParseWcfRelayAuthorizationRuleIDInsensitively(input string) (*WcfRelayAuthorizationRuleId, error) {
	parser := resourceids.NewParserFromResourceIdType(WcfRelayAuthorizationRuleId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := WcfRelayAuthorizationRuleId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", *parsed)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", *parsed)
	}

	if id.NamespaceName, ok = parsed.Parsed["namespaceName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "namespaceName", *parsed)
	}

	if id.WcfRelayName, ok = parsed.Parsed["wcfRelayName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "wcfRelayName", *parsed)
	}

	if id.AuthorizationRuleName, ok = parsed.Parsed["authorizationRuleName"]; !ok {
		return nil, resourceids.NewSegmentNotSpecifiedError(id, "authorizationRuleName", *parsed)
	}

	return &id, nil
}

// ValidateWcfRelayAuthorizationRuleID checks that 'input' can be parsed as a Wcf Relay Authorization Rule ID
func ValidateWcfRelayAuthorizationRuleID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWcfRelayAuthorizationRuleID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Wcf Relay Authorization Rule ID
func (id WcfRelayAuthorizationRuleId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Relay/namespaces/%s/wcfRelays/%s/authorizationRules/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NamespaceName, id.WcfRelayName, id.AuthorizationRuleName)
}

// Segments returns a slice of Resource ID Segments which comprise this Wcf Relay Authorization Rule ID
func (id WcfRelayAuthorizationRuleId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRelay", "Microsoft.Relay", "Microsoft.Relay"),
		resourceids.StaticSegment("staticNamespaces", "namespaces", "namespaces"),
		resourceids.UserSpecifiedSegment("namespaceName", "namespaceValue"),
		resourceids.StaticSegment("staticWcfRelays", "wcfRelays", "wcfRelays"),
		resourceids.UserSpecifiedSegment("wcfRelayName", "wcfRelayValue"),
		resourceids.StaticSegment("staticAuthorizationRules", "authorizationRules", "authorizationRules"),
		resourceids.UserSpecifiedSegment("authorizationRuleName", "authorizationRuleValue"),
	}
}

// String returns a human-readable description of this Wcf Relay Authorization Rule ID
func (id WcfRelayAuthorizationRuleId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Namespace Name: %q", id.NamespaceName),
		fmt.Sprintf("Wcf Relay Name: %q", id.WcfRelayName),
		fmt.Sprintf("Authorization Rule Name: %q", id.AuthorizationRuleName),
	}
	return fmt.Sprintf("Wcf Relay Authorization Rule (%s)", strings.Join(components, "\n"))
}
