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
	recaser.RegisterResourceId(&ManagedNetworkId{})
}

var _ resourceids.ResourceId = &ManagedNetworkId{}

// ManagedNetworkId is a struct representing the Resource ID for a Managed Network
type ManagedNetworkId struct {
	SubscriptionId     string
	ResourceGroupName  string
	AccountName        string
	ManagedNetworkName string
}

// NewManagedNetworkID returns a new ManagedNetworkId struct
func NewManagedNetworkID(subscriptionId string, resourceGroupName string, accountName string, managedNetworkName string) ManagedNetworkId {
	return ManagedNetworkId{
		SubscriptionId:     subscriptionId,
		ResourceGroupName:  resourceGroupName,
		AccountName:        accountName,
		ManagedNetworkName: managedNetworkName,
	}
}

// ParseManagedNetworkID parses 'input' into a ManagedNetworkId
func ParseManagedNetworkID(input string) (*ManagedNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedNetworkId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedNetworkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseManagedNetworkIDInsensitively parses 'input' case-insensitively into a ManagedNetworkId
// note: this method should only be used for API response data and not user input
func ParseManagedNetworkIDInsensitively(input string) (*ManagedNetworkId, error) {
	parser := resourceids.NewParserFromResourceIdType(&ManagedNetworkId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := ManagedNetworkId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *ManagedNetworkId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateManagedNetworkID checks that 'input' can be parsed as a Managed Network ID
func ValidateManagedNetworkID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseManagedNetworkID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Managed Network ID
func (id ManagedNetworkId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.CognitiveServices/accounts/%s/managedNetworks/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.AccountName, id.ManagedNetworkName)
}

// Segments returns a slice of Resource ID Segments which comprise this Managed Network ID
func (id ManagedNetworkId) Segments() []resourceids.Segment {
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
	}
}

// String returns a human-readable description of this Managed Network ID
func (id ManagedNetworkId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Account Name: %q", id.AccountName),
		fmt.Sprintf("Managed Network Name: %q", id.ManagedNetworkName),
	}
	return fmt.Sprintf("Managed Network (%s)", strings.Join(components, "\n"))
}
