package netappresourcequotalimitsaccount

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&NetAppAccountQuotaLimitId{})
}

var _ resourceids.ResourceId = &NetAppAccountQuotaLimitId{}

// NetAppAccountQuotaLimitId is a struct representing the Resource ID for a Net App Account Quota Limit
type NetAppAccountQuotaLimitId struct {
	SubscriptionId    string
	ResourceGroupName string
	NetAppAccountName string
	QuotaLimitName    string
}

// NewNetAppAccountQuotaLimitID returns a new NetAppAccountQuotaLimitId struct
func NewNetAppAccountQuotaLimitID(subscriptionId string, resourceGroupName string, netAppAccountName string, quotaLimitName string) NetAppAccountQuotaLimitId {
	return NetAppAccountQuotaLimitId{
		SubscriptionId:    subscriptionId,
		ResourceGroupName: resourceGroupName,
		NetAppAccountName: netAppAccountName,
		QuotaLimitName:    quotaLimitName,
	}
}

// ParseNetAppAccountQuotaLimitID parses 'input' into a NetAppAccountQuotaLimitId
func ParseNetAppAccountQuotaLimitID(input string) (*NetAppAccountQuotaLimitId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NetAppAccountQuotaLimitId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NetAppAccountQuotaLimitId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseNetAppAccountQuotaLimitIDInsensitively parses 'input' case-insensitively into a NetAppAccountQuotaLimitId
// note: this method should only be used for API response data and not user input
func ParseNetAppAccountQuotaLimitIDInsensitively(input string) (*NetAppAccountQuotaLimitId, error) {
	parser := resourceids.NewParserFromResourceIdType(&NetAppAccountQuotaLimitId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := NetAppAccountQuotaLimitId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *NetAppAccountQuotaLimitId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.NetAppAccountName, ok = input.Parsed["netAppAccountName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "netAppAccountName", input)
	}

	if id.QuotaLimitName, ok = input.Parsed["quotaLimitName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "quotaLimitName", input)
	}

	return nil
}

// ValidateNetAppAccountQuotaLimitID checks that 'input' can be parsed as a Net App Account Quota Limit ID
func ValidateNetAppAccountQuotaLimitID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseNetAppAccountQuotaLimitID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Net App Account Quota Limit ID
func (id NetAppAccountQuotaLimitId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.NetApp/netAppAccounts/%s/quotaLimits/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.NetAppAccountName, id.QuotaLimitName)
}

// Segments returns a slice of Resource ID Segments which comprise this Net App Account Quota Limit ID
func (id NetAppAccountQuotaLimitId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftNetApp", "Microsoft.NetApp", "Microsoft.NetApp"),
		resourceids.StaticSegment("staticNetAppAccounts", "netAppAccounts", "netAppAccounts"),
		resourceids.UserSpecifiedSegment("netAppAccountName", "netAppAccountName"),
		resourceids.StaticSegment("staticQuotaLimits", "quotaLimits", "quotaLimits"),
		resourceids.UserSpecifiedSegment("quotaLimitName", "quotaLimitName"),
	}
}

// String returns a human-readable description of this Net App Account Quota Limit ID
func (id NetAppAccountQuotaLimitId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Net App Account Name: %q", id.NetAppAccountName),
		fmt.Sprintf("Quota Limit Name: %q", id.QuotaLimitName),
	}
	return fmt.Sprintf("Net App Account Quota Limit (%s)", strings.Join(components, "\n"))
}
