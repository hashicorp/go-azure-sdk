package registeredidentities

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = RegisteredIdentityId{}

// RegisteredIdentityId is a struct representing the Resource ID for a Registered Identity
type RegisteredIdentityId struct {
	SubscriptionId         string
	ResourceGroupName      string
	VaultName              string
	RegisteredIdentityName string
}

// NewRegisteredIdentityID returns a new RegisteredIdentityId struct
func NewRegisteredIdentityID(subscriptionId string, resourceGroupName string, vaultName string, registeredIdentityName string) RegisteredIdentityId {
	return RegisteredIdentityId{
		SubscriptionId:         subscriptionId,
		ResourceGroupName:      resourceGroupName,
		VaultName:              vaultName,
		RegisteredIdentityName: registeredIdentityName,
	}
}

// ParseRegisteredIdentityID parses 'input' into a RegisteredIdentityId
func ParseRegisteredIdentityID(input string) (*RegisteredIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegisteredIdentityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegisteredIdentityId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.RegisteredIdentityName, ok = parsed.Parsed["registeredIdentityName"]; !ok {
		return nil, fmt.Errorf("the segment 'registeredIdentityName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ParseRegisteredIdentityIDInsensitively parses 'input' case-insensitively into a RegisteredIdentityId
// note: this method should only be used for API response data and not user input
func ParseRegisteredIdentityIDInsensitively(input string) (*RegisteredIdentityId, error) {
	parser := resourceids.NewParserFromResourceIdType(RegisteredIdentityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	var ok bool
	id := RegisteredIdentityId{}

	if id.SubscriptionId, ok = parsed.Parsed["subscriptionId"]; !ok {
		return nil, fmt.Errorf("the segment 'subscriptionId' was not found in the resource id %q", input)
	}

	if id.ResourceGroupName, ok = parsed.Parsed["resourceGroupName"]; !ok {
		return nil, fmt.Errorf("the segment 'resourceGroupName' was not found in the resource id %q", input)
	}

	if id.VaultName, ok = parsed.Parsed["vaultName"]; !ok {
		return nil, fmt.Errorf("the segment 'vaultName' was not found in the resource id %q", input)
	}

	if id.RegisteredIdentityName, ok = parsed.Parsed["registeredIdentityName"]; !ok {
		return nil, fmt.Errorf("the segment 'registeredIdentityName' was not found in the resource id %q", input)
	}

	return &id, nil
}

// ValidateRegisteredIdentityID checks that 'input' can be parsed as a Registered Identity ID
func ValidateRegisteredIdentityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRegisteredIdentityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Registered Identity ID
func (id RegisteredIdentityId) ID() string {
	fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.RecoveryServices/vaults/%s/registeredIdentities/%s"
	return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroupName, id.VaultName, id.RegisteredIdentityName)
}

// Segments returns a slice of Resource ID Segments which comprise this Registered Identity ID
func (id RegisteredIdentityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("staticSubscriptions", "subscriptions", "subscriptions"),
		resourceids.SubscriptionIdSegment("subscriptionId", "12345678-1234-9876-4563-123456789012"),
		resourceids.StaticSegment("staticResourceGroups", "resourceGroups", "resourceGroups"),
		resourceids.ResourceGroupSegment("resourceGroupName", "example-resource-group"),
		resourceids.StaticSegment("staticProviders", "providers", "providers"),
		resourceids.ResourceProviderSegment("staticMicrosoftRecoveryServices", "Microsoft.RecoveryServices", "Microsoft.RecoveryServices"),
		resourceids.StaticSegment("staticVaults", "vaults", "vaults"),
		resourceids.UserSpecifiedSegment("vaultName", "vaultValue"),
		resourceids.StaticSegment("staticRegisteredIdentities", "registeredIdentities", "registeredIdentities"),
		resourceids.UserSpecifiedSegment("registeredIdentityName", "registeredIdentityValue"),
	}
}

// String returns a human-readable description of this Registered Identity ID
func (id RegisteredIdentityId) String() string {
	components := []string{
		fmt.Sprintf("Subscription: %q", id.SubscriptionId),
		fmt.Sprintf("Resource Group Name: %q", id.ResourceGroupName),
		fmt.Sprintf("Vault Name: %q", id.VaultName),
		fmt.Sprintf("Registered Identity Name: %q", id.RegisteredIdentityName),
	}
	return fmt.Sprintf("Registered Identity (%s)", strings.Join(components, "\n"))
}
