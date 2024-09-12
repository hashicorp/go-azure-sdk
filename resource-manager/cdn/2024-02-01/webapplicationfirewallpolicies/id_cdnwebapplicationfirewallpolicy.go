package webapplicationfirewallpolicies

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/recaser"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

func init() {
	recaser.RegisterResourceId(&CdnWebApplicationFirewallPolicyId{})
}

var _ resourceids.ResourceId = &CdnWebApplicationFirewallPolicyId{}

// CdnWebApplicationFirewallPolicyId is a struct representing the Resource ID for a Cdn Web Application Firewall Policy
type CdnWebApplicationFirewallPolicyId struct {
	SubscriptionId                      string
	ResourceGroupName                   string
	CdnWebApplicationFirewallPolicyName string
}

// NewCdnWebApplicationFirewallPolicyID returns a new CdnWebApplicationFirewallPolicyId struct
func NewCdnWebApplicationFirewallPolicyID(subscriptionId string, resourceGroupName string, cdnWebApplicationFirewallPolicyName string) CdnWebApplicationFirewallPolicyId {
	return CdnWebApplicationFirewallPolicyId{
		SubscriptionId:                      subscriptionId,
		ResourceGroupName:                   resourceGroupName,
		CdnWebApplicationFirewallPolicyName: cdnWebApplicationFirewallPolicyName,
	}
}

// ParseCdnWebApplicationFirewallPolicyID parses 'input' into a CdnWebApplicationFirewallPolicyId
func ParseCdnWebApplicationFirewallPolicyID(input string) (*CdnWebApplicationFirewallPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CdnWebApplicationFirewallPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CdnWebApplicationFirewallPolicyId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseCdnWebApplicationFirewallPolicyIDInsensitively parses 'input' case-insensitively into a CdnWebApplicationFirewallPolicyId
// note: this method should only be used for API response data and not user input
func ParseCdnWebApplicationFirewallPolicyIDInsensitively(input string) (*CdnWebApplicationFirewallPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&CdnWebApplicationFirewallPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := CdnWebApplicationFirewallPolicyId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *CdnWebApplicationFirewallPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.CdnWebApplicationFirewallPolicyName, ok = input.Parsed["cdnWebApplicationFirewallPolicyName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cdnWebApplicationFirewallPolicyName", input)
	}

	return nil
}

// ValidateCdnWebApplicationFirewallPolicyID checks that 'input' can be parsed as a Cdn Web Application Firewall Policy ID
func ValidateCdnWebApplicationFirewallPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseCdnWebApplicationFirewallPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Cdn Web Application Firewall Policy ID
func (id CdnWebApplicationFirewallPolicyId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CDN/cdnWebApplicationFirewallPolicies/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.CdnWebApplicationFirewallPolicyName)
}

// Segments returns a slice of Resource ID Segments which comprise this Cdn Web Application Firewall Policy ID
func (id CdnWebApplicationFirewallPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftCDN", "Microsoft.CDN", "Microsoft.CDN"),
		resourceids.StaticSegment("staticCdnWebApplicationFirewallPolicies", "cdnWebApplicationFirewallPolicies", "cdnWebApplicationFirewallPolicies"),
		resourceids.UserSpecifiedSegment("cdnWebApplicationFirewallPolicyName", "cdnWebApplicationFirewallPolicyValue"),
	}
}

// String returns a human-readable description of this Cdn Web Application Firewall Policy ID
func (id CdnWebApplicationFirewallPolicyId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Cdn Web Application Firewall Policy Name: %q", id.CdnWebApplicationFirewallPolicyName),
	}
	return fmt.Sprintf("Cdn Web Application Firewall Policy (%s)", strings.Join(components, "\n"))
}
