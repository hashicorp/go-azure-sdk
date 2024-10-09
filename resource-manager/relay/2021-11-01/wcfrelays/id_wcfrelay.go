package wcfrelays

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&WcfRelayId{})
}

var _ resourceids.ResourceId = &WcfRelayId{}

// WcfRelayId is a struct representing the Resource ID for a Wcf Relay
type WcfRelayId struct {
	SubscriptionId    string
	ResourceGroupName string
	NamespaceName     string
	WcfRelayName      string
}

// NewWcfRelayID returns a new WcfRelayId struct
func NewWcfRelayID(subscriptionId string, resourceGroupName string, namespaceName string, wcfRelayName string) WcfRelayId {
	return WcfRelayId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NamespaceName:     namespaceName,
		WcfRelayName:      wcfRelayName,
	}
}

// ParseWcfRelayID parses 'input' into a WcfRelayId
func ParseWcfRelayID(input string) (*WcfRelayId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WcfRelayId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WcfRelayId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseWcfRelayIDInsensitively parses 'input' case-insensitively into a WcfRelayId
// note: this method should only be used for API response data and not user input
func ParseWcfRelayIDInsensitively(input string) (*WcfRelayId, error) {
	parser := resourceids.NewParserFromResourceIdType(&WcfRelayId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := WcfRelayId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *WcfRelayId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NamespaceName, ok = input.Parsed["namespaceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "namespaceName", input)
	}

	if id.WcfRelayName, ok = input.Parsed["wcfRelayName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "wcfRelayName", input)
	}

	return nil
}

// ValidateWcfRelayID checks that 'input' can be parsed as a Wcf Relay ID
func ValidateWcfRelayID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseWcfRelayID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Wcf Relay ID
func (id WcfRelayId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Relay/namespaces/%s/wcfRelays/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NamespaceName, id.WcfRelayName)
}

// Segments returns a slice of Resource ID Segments which comprise this Wcf Relay ID
func (id WcfRelayId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRelay", "Microsoft.Relay", "Microsoft.Relay"),
		resourceids.StaticSegment("staticNamespaces", "namespaces", "namespaces"),
		resourceids.UserSpecifiedSegment("namespaceName", "namespaceName"),
		resourceids.StaticSegment("staticWcfRelays", "wcfRelays", "wcfRelays"),
		resourceids.UserSpecifiedSegment("wcfRelayName", "wcfRelayName"),
	}
}

// String returns a human-readable description of this Wcf Relay ID
func (id WcfRelayId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Namespace Name: %q", id.NamespaceName),
		fmt.Sprintf("Wcf Relay Name: %q", id.WcfRelayName),
	}
	return fmt.Sprintf("Wcf Relay (%s)", strings.Join(components, "\n"))
}
