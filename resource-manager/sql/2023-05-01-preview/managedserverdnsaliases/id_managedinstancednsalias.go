package managedserverdnsaliases

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &ManagedInstanceDnsAliasId{}

// ManagedInstanceDnsAliasId is a struct representing the Resource ID for a Managed Instance Dns Alias
type ManagedInstanceDnsAliasId struct {
	SubscriptionId      string
	ResourceGroupName   string
	ManagedInstanceName string
	DnsAliasName        string
}

// NewManagedInstanceDnsAliasID returns a new ManagedInstanceDnsAliasId struct
func NewManagedInstanceDnsAliasID(subscriptionId string, resourceGroupName string, managedInstanceName string, dnsAliasName string) ManagedInstanceDnsAliasId {
	return ManagedInstanceDnsAliasId{
		SubscriptionId:      subscriptionId,
		ResourceGroupName:   resourceGroupName,
		ManagedInstanceName: managedInstanceName,
		DnsAliasName:        dnsAliasName,
	}
}

// ParseManagedInstanceDnsAliasID parses 'input' into a ManagedInstanceDnsAliasId
func ParseManagedInstanceDnsAliasID(input string) (*ManagedInstanceDnsAliasId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedInstanceDnsAliasId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedInstanceDnsAliasId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedInstanceDnsAliasIDInsensitively parses 'input' case-insensitively into a ManagedInstanceDnsAliasId
// note: this method should only be used for API response data and not user input
func ParseManagedInstanceDnsAliasIDInsensitively(input string) (*ManagedInstanceDnsAliasId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedInstanceDnsAliasId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedInstanceDnsAliasId{}
	if err := id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedInstanceDnsAliasId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.SubscriptionId, ok = input.Parsed["subscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "subscriptionId", input)
	}

	if id.ResourceGroupName, ok = input.Parsed["resourceGroupName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceGroupName", input)
	}

	if id.ManagedInstanceName, ok = input.Parsed["managedInstanceName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedInstanceName", input)
	}

	if id.DnsAliasName, ok = input.Parsed["dnsAliasName"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dnsAliasName", input)
	}

	return nil
}

// ValidateManagedInstanceDnsAliasID checks that 'input' can be parsed as a Managed Instance Dns Alias ID
func ValidateManagedInstanceDnsAliasID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedInstanceDnsAliasID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Instance Dns Alias ID
func (id ManagedInstanceDnsAliasId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Sql/managedInstances/%s/dnsAliases/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.ManagedInstanceName, id.DnsAliasName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Instance Dns Alias ID
func (id ManagedInstanceDnsAliasId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftSql", "Microsoft.Sql", "Microsoft.Sql"),
		resourceids.StaticSegment("staticManagedInstances", "managedInstances", "managedInstances"),
		resourceids.UserSpecifiedSegment("managedInstanceName", "managedInstanceValue"),
		resourceids.StaticSegment("staticDnsAliases", "dnsAliases", "dnsAliases"),
		resourceids.UserSpecifiedSegment("dnsAliasName", "dnsAliasValue"),
	}
}

// String returns a human-readable description of this Managed Instance Dns Alias ID
func (id ManagedInstanceDnsAliasId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Managed Instance Name: %q", id.ManagedInstanceName),
		fmt.Sprintf("Dns Alias Name: %q", id.DnsAliasName),
	}
	return fmt.Sprintf("Managed Instance Dns Alias (%s)", strings.Join(components, "\n"))
}
